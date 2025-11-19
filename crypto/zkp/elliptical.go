package zkp

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

type EllipticIdentity struct {
	Curve  elliptic.Curve `json:"-"`
	Nonce  string         `json:"nonce"`
	HashX  string         `json:"hash_x"`
	HashY  string         `json:"hash_y"`
	Proofs []Proof        `json:"proofs"`
}

func NewEllipticIdentity() *EllipticIdentity {
	return &EllipticIdentity{
		Curve:  elliptic.P256(),
		Proofs: []Proof{},
	}
}

// --- Utility functions ---

func (id *EllipticIdentity) ensureCurve() {
	if id.Curve == nil {
		id.Curve = elliptic.P256()
	}
}

// --- IdentityBackend implementation ---

func (id *EllipticIdentity) New() error {
	id.ensureCurve()
	nonceBytes := make([]byte, 32)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return err
	}
	id.Nonce = hex.EncodeToString(nonceBytes)
	x := HashToScalar(id.Curve, "nonce:", id.Nonce)
	hashX, hashY := id.Curve.ScalarBaseMult(x.Bytes())
	id.HashX = hashX.Text(16)
	id.HashY = hashY.Text(16)
	id.Proofs = []Proof{}
	return nil
}

func (id *EllipticIdentity) Save() (string,error) {
	configData,err:=os.ReadFile("config.json");if err!=nil{
		return "",err
	}
		var cfg struct {
		IdentityDir string `json:"identity_dir"`
	}
	if err := json.Unmarshal(configData, &cfg); err != nil {
		return "", err
	}

	// Step 2: Prepend home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, cfg.IdentityDir)

	// Step 3: Ensure directory exists
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	// Step 4: Generate filename and full path
	filename := fmt.Sprintf("%d.json", time.Now().UnixNano())
	path := filepath.Join(dir, filename)

	// Step 5: Marshal and save JSON
	data, err := json.MarshalIndent(id, "", "  ")
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return "", err
	}
	return path,nil
}

func (id *EllipticIdentity) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, id); err != nil {
		return err
	}
	id.ensureCurve()
	return nil
}

func (id *EllipticIdentity) PublicHash() (hashX, hashY *big.Int) {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(id.HashX, 16)
	y.SetString(id.HashY, 16)
	return x, y
}

func (id *EllipticIdentity) GenerateProof(attribute string) (*Proof, error) {
	id.ensureCurve()
	x := HashToScalar(id.Curve, "nonce:", id.Nonce)
	hashX, hashY := id.PublicHash()

	r, err := rand.Int(rand.Reader, id.Curve.Params().N)
	if err != nil {
		return nil, err
	}
	rx, ry := id.Curve.ScalarBaseMult(r.Bytes())

	eBytes := sha256.Sum256([]byte(fmt.Sprintf("%x%x%x%x%s", rx, ry, hashX, hashY, attribute)))
	e := new(big.Int).SetBytes(eBytes[:])
	e.Mod(e, id.Curve.Params().N)

	s := new(big.Int).Mul(e, x)
	s.Add(s, r)
	s.Mod(s, id.Curve.Params().N)

	proof := Proof{
		Name: attribute,
		Rx:   rx.Text(16),
		Ry:   ry.Text(16),
		S:    s.Text(16),
	}
	id.Proofs = append(id.Proofs, proof)
	return &proof, nil
}

func (id *EllipticIdentity) Verify(proof *Proof, attribute string) bool {
	id.ensureCurve()
	hashX, hashY := id.PublicHash()
	rx := new(big.Int)
	ry := new(big.Int)
	s := new(big.Int)
	rx.SetString(proof.Rx, 16)
	ry.SetString(proof.Ry, 16)
	s.SetString(proof.S, 16)

	eBytes := sha256.Sum256([]byte(fmt.Sprintf("%x%x%x%x%s", rx, ry, hashX, hashY, attribute)))
	e := new(big.Int).SetBytes(eBytes[:])
	e.Mod(e, id.Curve.Params().N)

	sGx, sGy := id.Curve.ScalarBaseMult(s.Bytes())
	eHx, eHy := id.Curve.ScalarMult(hashX, hashY, e.Bytes())
	negEHy := new(big.Int).Neg(eHy)
	negEHy.Mod(negEHy, id.Curve.Params().P)
	vx, vy := id.Curve.Add(rx, ry, eHx, negEHy)

	return vx.Cmp(sGx) == 0 && vy.Cmp(sGy) == 0
}

func (id *EllipticIdentity) Serialize() ([]byte, error) {
	return json.MarshalIndent(id, "", "  ")
}

func (id *EllipticIdentity) Deserialize(data []byte) error {
	if err := json.Unmarshal(data, id); err != nil {
		return err
	}
	id.ensureCurve()
	return nil
}
