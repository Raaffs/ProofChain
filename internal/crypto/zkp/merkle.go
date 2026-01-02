package zkp

import (
	"slices"
	"fmt"
	"sort"

	"github.com/Suy56/ProofChain/storage/models"
)

// MerkleProof implements the ZKProof interface and holds the committed data state.
type MerkleProof struct {
	RootHash   Hash
	FieldLeaves map[string]FieldLeaf // Map for O(1) lookup during Disclosure
	LeafHashes []Hash                // Ordered list for Merkle Tree construction
}

func NewMerkleProof() *MerkleProof {
	mp:=&MerkleProof{}
	mp.New()
	return mp
}

func (id *MerkleProof) New()  {
	id.RootHash = ""
	id.FieldLeaves = make(map[string]FieldLeaf)
	id.LeafHashes = make([]Hash, 0)
}

// GenerateRootProof (Issuer side)
func (id *MerkleProof) GenerateRootProof(c models.CertificateData) (Hash, SaltedCertificate, error) {
	// 1. Salt the fields
	saltedCert, err := SaltCertificate(c)
	if err != nil {
		return "", SaltedCertificate{}, err
	}

	// 2. Prepare state for Merkle Tree generation (keys must be sorted for deterministic ordering)
	var keys []string
	for key := range saltedCert.SaltedFields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	id.FieldLeaves = saltedCert.SaltedFields
	id.LeafHashes = make([]Hash, 0, len(keys))

	// 3. Build the ordered LeafHashes list
	for _, key := range keys {
		leaf := id.FieldLeaves[key]
		id.LeafHashes = append(id.LeafHashes, leaf.Hash)
	}

	// 4. Calculate Merkle Root
	id.RootHash = calculateMerkleRoot(id.LeafHashes)

	// The Issuer sends the Root Hash (to Blockchain) and the SaltedCertificate (to Requestor)
	return id.RootHash, saltedCert, nil
}

// LoadSaltedData (Requestor side)
func (id *MerkleProof) LoadSaltedData(sc SaltedCertificate) error {
	id.New() // Reset state

	// 1. Prepare state for Merkle Tree generation (keys must be sorted for deterministic ordering)
	var keys []string
	for key := range sc.SaltedFields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	id.FieldLeaves = sc.SaltedFields
	id.LeafHashes = make([]Hash, 0, len(keys))

	// 2. Rebuild the ordered LeafHashes list from the received data
	for _, key := range keys {
		leaf := id.FieldLeaves[key]
		id.LeafHashes = append(id.LeafHashes, leaf.Hash)
	}
	
	// 3. Re-calculate Merkle Root (must match the one published by the issuer)
	id.RootHash = calculateMerkleRoot(id.LeafHashes)

	return nil
}

// Disclose (Requestor side)
func (id *MerkleProof) Disclose(attribute string) (Proof, error) {
	// Requestor can access FieldLeaves because they loaded the SaltedCertificate
	leaf, exists := id.FieldLeaves[attribute]
	if !exists {
		return Proof{}, fmt.Errorf("attribute '%s' does not exist in the certificate", attribute)
	}

	// 1. Find index in the ordered leaf list
	index := -1
	for i, h := range id.LeafHashes {
		if h == leaf.Hash {
			index = i
			break
		}
	}
	if index == -1 {
		return Proof{}, fmt.Errorf("integrity error: leaf hash not found in tree")
	}

	// 2. Calculate Merkle Path
	merklePath := calculateMerkleProof(id.LeafHashes, index)

	// 3. Construct Proof
	return Proof{
		RootHash:    id.RootHash,
		Attribute:   attribute,
		Value:       leaf.Value,
		Salt:        leaf.Salt,
		MerkleProof: merklePath,
	}, nil
}

// --- Verification Helper (Verifier Logic) ---

// VerifyProof checks if a disclosed proof matches the expected root hash.
// This runs on the client/verifier side.
func VerifyProof(p Proof, expectedRoot Hash) bool {
    // 1. Re-calculate the leaf hash for the field we are checking
    disclosedLeafHash := HashData([]byte(p.Value), []byte(p.Salt))
    
    // 2. We need to find this hash in the list of all leaves provided in the proof.
    // This proves the disclosed value actually belongs to the set.
    found := slices.Contains(p.MerkleProof, disclosedLeafHash)
    
    if !found {
        // If the hash isn't in the list, the user is showing 
        // a valid value/salt for a DIFFERENT certificate.
        return false 
    }

    // 3. Re-calculate the root using YOUR recursive function
    // We pass the full list of leaves (p.MerkleProof)
    calculatedRoot := calculateMerkleRoot(p.MerkleProof)

    // 4. Verification
    return calculatedRoot == expectedRoot && calculatedRoot == p.RootHash
}

// calculateMerkleRoot calculates the root hash from an ordered list of leaf hashes.
func calculateMerkleRoot(leaves []Hash) Hash {
	if len(leaves) == 0 {
		return ""
	}
	if len(leaves) == 1 {
		return leaves[0]
	}

	// Pad if odd
	currentLeaves := make([]Hash, len(leaves))
	copy(currentLeaves, leaves)
	if len(currentLeaves)%2 != 0 {
		currentLeaves = append(currentLeaves, currentLeaves[len(currentLeaves)-1])
	}

	var nextLevel []Hash
	for i := 0; i < len(currentLeaves); i += 2 {
		h1 := currentLeaves[i]
		h2 := currentLeaves[i+1]
		// Sort hashes before concatenating to ensure canonical parent hash
		if h1 < h2 {
			nextLevel = append(nextLevel, HashData([]byte(h1), []byte(h2)))
		} else {
			nextLevel = append(nextLevel, HashData([]byte(h2), []byte(h1)))
		}
	}

	return calculateMerkleRoot(nextLevel)
}

// calculateMerkleProof generates the sibling hashes needed to verify a leaf at a specific index.
func calculateMerkleProof(leaves []Hash, index int) []Hash {
	proof := []Hash{}
	current := leaves

	// Work with a copy to avoid modifying the original
	workingLeaves := make([]Hash, len(leaves))
	copy(workingLeaves, leaves)

	if len(workingLeaves)%2 != 0 {
		workingLeaves = append(workingLeaves, workingLeaves[len(workingLeaves)-1])
	}
	current = workingLeaves

	for len(current) > 1 {
		var nextLevel []Hash
		for i := 0; i < len(current); i += 2 {
			h1 := current[i]
			h2 := current[i+1]
			
			// If our target is in this pair, append the OTHER one to the proof
			if i == index || i+1 == index {
				if i == index {
					proof = append(proof, h2)
				} else {
					proof = append(proof, h1)
				}
			}

			if h1 < h2 {
				nextLevel = append(nextLevel, HashData([]byte(h1), []byte(h2)))
			} else {
				nextLevel = append(nextLevel, HashData([]byte(h2), []byte(h1)))
			}
		}

		current = nextLevel
		index /= 2 
		
		if len(current)%2 != 0 && len(current) > 1 {
			current = append(current, current[len(current)-1])
		}
	}

	return proof
}
