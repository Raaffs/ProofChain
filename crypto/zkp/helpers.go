package zkp

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// HashToScalar derives a scalar from a domain-separated string
func HashToScalar(curve elliptic.Curve, tag string, data string) *big.Int {
	h := sha256.Sum256([]byte(tag + data))
	scalar := new(big.Int).SetBytes(h[:])
	scalar.Mod(scalar, curve.Params().N)
	return scalar
}

// PointToHex serializes an elliptic curve point as uncompressed 04||x||y hex
func PointToHex(curve elliptic.Curve, x, y *big.Int) string {
	b := elliptic.Marshal(curve, x, y)
	return hex.EncodeToString(b)
}

// ScalarBaseMult multiplies the curve base point by a scalar
func ScalarBaseMult(curve elliptic.Curve, k *big.Int) (x, y *big.Int) {
	return curve.ScalarBaseMult(k.Bytes())
}

// PointAdd adds two curve points
func PointAdd(curve elliptic.Curve, x1, y1, x2, y2 *big.Int) (x, y *big.Int) {
	return curve.Add(x1, y1, x2, y2)
}

func Verify(curve elliptic.Curve, hashX, hashY *big.Int, proof *Proof, attribute string) bool {
	rx := new(big.Int)
	ry := new(big.Int)
	s := new(big.Int)

	rx.SetString(proof.Rx, 16)
	ry.SetString(proof.Ry, 16)
	s.SetString(proof.S, 16)

	eBytes := sha256.Sum256([]byte(fmt.Sprintf("%x%x%x%x%s", rx, ry, hashX, hashY, attribute)))
	e := new(big.Int).SetBytes(eBytes[:])
	e.Mod(e, curve.Params().N)

	sGx, sGy := curve.ScalarBaseMult(s.Bytes())
	eHx, eHy := curve.ScalarMult(hashX, hashY, e.Bytes())
	negEHy := new(big.Int).Neg(eHy)
	negEHy.Mod(negEHy, curve.Params().P)
	vx, vy := curve.Add(rx, ry, eHx, negEHy)

	return vx.Cmp(sGx) == 0 && vy.Cmp(sGy) == 0
}
