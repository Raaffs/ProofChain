package zkp

import (
	"github.com/Suy56/ProofChain/storage/models"
)


// FieldLeaf stores the necessary data for a single attribute leaf node.
type FieldLeaf struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Salt  string `json:"salt"`
	Hash  Hash   `json:"hash"` // The salted hash of the value
}


// SaltedCertificate is the data structure that the Issuer sends to the Requestor.
// It contains all field data and salts, allowing the Requestor to reconstruct the tree.
type SaltedCertificate struct {
	SaltedFields map[string]FieldLeaf `json:"salted_fields"`
}


type ZKProof interface {
	New() 
	
	// For Issuer: Generates Merkle Root and Salted Certificate for transmission
	GenerateRootProof(c models.CertificateData) (Hash, SaltedCertificate, error)

	// For Requestor: Loads salted data to prepare for disclosure
	LoadSaltedData(sc SaltedCertificate) error

	// For Requestor: Retrieves the proof for a single attribute
	Disclose(attribute string) (Proof, error)
}

// Proof contains the components needed for a third-party verifier.
// This is returned by the Disclose function.
type Proof struct {
	RootHash    Hash     `json:"root_hash"`    // The committed hash the verifier checks against
	Attribute   string   `json:"attribute"`    // The name of the field being disclosed
	Value       string   `json:"value"`        // The disclosed field value
	Salt        string   `json:"salt"`         // The salt used to generate the leaf hash
	MerkleProof []Hash   `json:"merkle_proof"` // The ordered list of sibling hashes for verification
}
// FieldData stores the salted and hashed commitment for a single certificate field.
type FieldData struct {
	Value    string `json:"value"`
	Salt     string `json:"salt"` // Salt will be empty for the Nonce field
	LeafHash Hash   `json:"leafHash"` // Now serializes as a hex string
	Index    uint64 `json:"index"`
}
