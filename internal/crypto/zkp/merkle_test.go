package zkp

import (
	"testing"

	"github.com/Suy56/ProofChain/storage/models"
)
func TestVerifyProof(t *testing.T) {
    // Setup initial valid state
    input := models.CertificateData{
        Name:            "Alice Smith",
        CertificateName: "Bachelors of Science",
        BirthDate:       "1995-01-01",
        Address:         "123 Go Lane",
        Extra:           map[string]string{"MembershipID": "1225789"},
    }

    merkle := NewMerkleProof()
    expectedRoot, saltedCert, err := merkle.GenerateRootProof(input)
    if err != nil {
        t.Fatalf("Setup failed: %v", err)
    }

    err = merkle.LoadSaltedData(saltedCert)
    if err != nil {
        t.Fatalf("Load failed: %v", err)
    }

    // Generate a clean proof for the "Name" attribute
    originalProof, err := merkle.Disclose("Name")
    if err != nil {
        t.Fatalf("Disclose failed: %v", err)
    }

    t.Run("Valid Proof", func(t *testing.T) {
        if !VerifyProof(originalProof, expectedRoot) {
            t.Error("Expected valid proof to pass")
        }
    })

    t.Run("Tampered Value", func(t *testing.T) {
        p := originalProof // Copy
        p.Value = "Bob"    // Change the disclosed value
        if VerifyProof(p, expectedRoot) {
            t.Error("Expected failure for tampered value")
        }
    })

    t.Run("Tampered Salt", func(t *testing.T) {
        p := originalProof
        p.Salt = "wrong_salt_123" // Change the salt
        if VerifyProof(p, expectedRoot) {
            t.Error("Expected failure for tampered salt")
        }
    })

    t.Run("Tampered Merkle Path", func(t *testing.T) {
        p := originalProof
        if len(p.MerkleProof) > 0 {
            // Swap a hash in the proof with a fake one
            p.MerkleProof[0] = "df6f4d226a2608447d95a1656b6851b471861053"
            if VerifyProof(p, expectedRoot) {
                t.Error("Expected failure for tampered merkle path")
            }
        }
    })

    t.Run("Wrong Expected Root", func(t *testing.T) {
        // Provide the correct proof but check against a completely different root
        fakeRoot := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
        if VerifyProof(originalProof, Hash(fakeRoot)) {
            t.Error("Expected failure when checking against incorrect root")
        }
    })

    t.Run("Empty/Missing Field", func(t *testing.T) {
        p := originalProof
        p.Value = ""
        if VerifyProof(p, expectedRoot) {
            t.Error("Expected failure for empty value")
        }
    })
}