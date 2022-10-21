package main

import (
	"fmt"

	"github.com/fsabiu/did-methods-spec/chaincode/doc"
)

func main() {

	// Variable assignments (bottom-up)
	s := "did:orcl:123"

	property := doc.Properties{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/suites/ed25519-2020/v1"},
		Id: s,
		Method: []doc.VerMethod{{
			Id:          s + "#key-1",
			Typ:         "Ed25519VerificationKey2020",
			Controller:  s,
			PublicKeyMB: "cTx0CiPUqsTLr2xy53VAQUVfOn7dvFqHeeC1k"}},
		Auth: []string{s}}

	datamodel := doc.DataModel{
		Property: &property,
	}

	// Did document
	didDocument := make(doc.DidDoc)
	didDocument["id"] = s
	didDocument["didMethod"] = "did:orcl"
	didDocument["implementation"] = "DID Oracle Test Suite"
	didDocument["implementater"] = "Oracle"
	didDocument["supportedContentTypes"] = []string{"application/did+json",
		"application/did+ld+json"}
	didDocument["dids"] = []string{"did:orcl:QC5S8KGCFN37Z5VP"}
	didDocument["did:orcl:QC5S8KGCFN37Z5VP"] = datamodel

	// Print did document JSON
	//printJson(didDocument)

	/////////////
	verMeth := doc.CreateVerMethod(s+"#key-1", "Ed25519VerificationKey2020", s, "cTx0CiPUqsTLr2xy53VAQUVfOn7dvFqHeeC1k")

	contexts := []string{"https://www.w3.org/ns/did/v1", "https://w3id.org/security/suites/ed25519-2020/v1"}

	prop := doc.CreateProperty(contexts, s+"#key-1", []doc.VerMethod{verMeth}, []string{s})

	dm := doc.CreateDataModel(prop)

	didMap := make(map[string]doc.DataModel)
	didMap["did:orcl:QC5S8KGCFN37Z5VP"] = dm

	publicKey := "cTx0CiPUqsTLr2xy53VAQUVfOn7dvFqXXXXXX"
	mydid := doc.Create(publicKey)

	fmt.Println(mydid)

	//printJson(verMeth)
	//printJson(prop)
	//printJson(didDoc)

	/* dataModel2Add := make(map[string]doc.DataModel)
	dataModel2Add["did:orcl:XXXXXXXXXXX"] = dm

	didDoc.AddDataModel(dataModel2Add)

	//printJson(didDoc)

	testMap := make(map[string][]doc.VerMethod)
	testMap["did:orcl:XXXXXXXXXXX"] = prop.Method
	didDoc.AddAuthMethod(testMap)
	doc.PrintJson(didDoc) */

}
