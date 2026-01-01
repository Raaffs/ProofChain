package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"

	blockchain "github.com/Suy56/ProofChain/chaincore/core"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func FilterDocument(docs []blockchain.VerificationDocument, condition func(blockchain.VerificationDocument)bool)[]blockchain.VerificationDocument{
	var userDocs []blockchain.VerificationDocument
	for _,doc :=range docs{
		if(condition(doc)){	
			userDocs=append(userDocs,doc)
		}
	}
	return userDocs
}



func Walk[S any](s S) func(yield func(string, any) bool) {
	v := reflect.ValueOf(s)

	// Dereference pointer if needed
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return func(yield func(string, any) bool) {}
	}

	return func(yield func(string, any) bool) {
		t := v.Type()
		numFields := v.NumField()

		for i := range numFields {
			field := t.Field(i)
			value := v.Field(i)

			if !field.IsExported() {
				continue
			}

			switch value.Kind() {
			case reflect.Map:
				// Iterate map keys
				for _, key := range value.MapKeys() {
					val := value.MapIndex(key)
					if !yield(fmt.Sprint(key.Interface()), val.Interface()) {
						return
					}
				}
			default:
				// Use the struct field name as the attribute key
				if !yield(field.Name, value.Interface()) {
					return
				}
			}
		}
	}
}