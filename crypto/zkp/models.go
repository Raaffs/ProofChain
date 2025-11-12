package zkp

import "math/big"

// IdentityBackend defines the core functions any identity type must implement
type DigitalIdentity interface {
	// PublicHash returns the serialized public commitment
	PublicHash() (hashX, hashY *big.Int)

	// Prove generates a ZK proof for a given attribute (name)
	Prove(attribute string) (*Proof, error)

	// Verify checks that a proof matches the public commitment for a given attribute
	Verify(proof *Proof, attribute string) bool

	// Serialize/Deserialize (optional, if needed for storage)
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}


type Proof struct {
    Name string
    Rx, Ry, S string
}