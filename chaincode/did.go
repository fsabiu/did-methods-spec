package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// Structures definition
	type VerMethod struct {
		Id          string `json:"id"`
		Typ         string `json:"type"`
		Controller  string `json:"controller"`
		PublicKeyMB string `json:"publicKeyMultibase"`
	}

	type Properties struct {
		Context []string    `json:"@context"`
		Id      string      `json:"id"`
		Method  []VerMethod `json:"verificationMethod"`
		Auth    []string    `json:"authentication"`
	}

	type DataModel struct {
		Property Properties `json:"didDocumentDataModel"`
	}

	// Variable assignments (bottom-up)
	s := "did:orcl:123"

	property := Properties{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/suites/ed25519-2020/v1"},
		Id: s,
		Method: []VerMethod{{
			Id:          s + "#key-1",
			Typ:         "Ed25519VerificationKey2020",
			Controller:  s,
			PublicKeyMB: "cTx0CiPUqsTLr2xy53VAQUVfOn7dvFqHeeC1k"}},
		Auth: []string{s}}

	datamodel := DataModel{
		Property: property,
	}

	// Did document
	didDocument := make(map[string]interface{})
	didDocument["didMethod"] = "did:orcl"
	didDocument["implementation"] = "DID Oracle Test Suite"
	didDocument["implementater"] = "Oracle"
	didDocument["supportedContentTypes"] = []string{"application/did+json",
		"application/did+ld+json"}
	didDocument["dids"] = []string{"did:orcl:QC5S8KGCFN37Z5VP"}
	didDocument["did:orcl:QC5S8KGCFN37Z5VP"] = datamodel

	// Print did document JSON
	a1_json, err := json.Marshal(didDocument)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
	}

}
