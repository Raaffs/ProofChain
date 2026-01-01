package download

import (
	"encoding/json"
	"log/slog"
	"os"
	"testing"
)
func TestExtractProofValues_WithExtra(t *testing.T) {
	// 1. Setup sample data with both Fixed and Extra fields
	input := DownloadProof{
		Name:      hashedField{Hash: "h1", Key: "Name", Salt: "s1", Value: "Maria"},
		CertificateName:      hashedField{Hash: "hn", Key: "CertificateName", Salt: "sn", Value: "CERTI"},
		BirthDate: hashedField{Hash: "h2", Key: "BirthDate", Salt: "s2", Value: "1995"},
		Address: hashedField{Hash: "h3",Key: "Address", Salt: "s3", Value: "Tokyo, Japan"},
		Extra: map[string]hashedField{
			"MembershipID": {
				Hash:  "m_hash",
				Key:   "MID",
				Salt:  "m_salt",
				Value: "GOLD_99",
			},
		},
	}

	inputBytes,err:=json.Marshal(input);if err!=nil{
		t.Fatal("error marshalling json: ",err)
	}
	d,err:=NewDownloader(inputBytes, slog.New(slog.NewJSONHandler(os.Stdout,nil)));if err!=nil{
		t.Fatal("error initializing downloader %w",err)
	}
	if err:=d.Exec();err!=nil{
		t.Fatal("Error downloading: ",err)
	}
}