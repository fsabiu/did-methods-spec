package main

import (
	"encoding/json"
	"fmt"
)

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

type DidDoc map[string]interface{}

func main() {

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
	didDocument := make(DidDoc)
	didDocument["id"] = s
	didDocument["didMethod"] = "did:orcl"
	didDocument["implementation"] = "DID Oracle Test Suite"
	didDocument["implementater"] = "Oracle"
	didDocument["supportedContentTypes"] = []string{"application/did+json",
		"application/did+ld+json"}
	didDocument["dids"] = []string{"did:orcl:QC5S8KGCFN37Z5VP"}
	didDocument["did:orcl:QC5S8KGCFN37Z5VP"] = datamodel

	// Print did document JSON
	printJson(didDocument)

	/////////////
	verMeth := createVerMethod(s+"#key-1", "Ed25519VerificationKey2020", s, "cTx0CiPUqsTLr2xy53VAQUVfOn7dvFqHeeC1k")

	contexts := []string{"https://www.w3.org/ns/did/v1", "https://w3id.org/security/suites/ed25519-2020/v1"}

	prop := createProperty(contexts, s+"#key-1", []VerMethod{verMeth}, []string{s})

	dm := createDataModel(prop)

	didMap := make(map[string]DataModel)
	didMap["did:orcl:QC5S8KGCFN37Z5VP"] = dm

	didDoc := createDidDocument(s, "did:orcl", "DID Oracle Test Suite", "Oracle", []string{"application/did+json",
		"application/did+ld+json"}, didMap)
	//printJson(verMeth)
	//printJson(prop)
	printJson(didDoc)

	dataModel2Add := make(map[string]DataModel)
	dataModel2Add["did:orcl:XXXXXXXXXXX"] = dm

	didDoc.addDataModel(dataModel2Add)

	printJson(didDoc)

}

func (didDocument DidDoc) addDataModel(datamodels map[string]DataModel) {

	for key, value := range datamodels {
		didDocument[key] = value
	}

}

func createDidDocument(id string, didMethod string, implementation string, implementer string, supportedContentTypes []string, datamodels map[string]DataModel) DidDoc {
	/* Requires: dids is a list of n strings (1 for each did)
	datamodels is a dict having 1 key for each element of dids
	*/
	didDocument := make(DidDoc)
	didDocument["id"] = id
	didDocument["didMethod"] = didMethod
	didDocument["implementation"] = implementation
	didDocument["implementater"] = implementer
	didDocument["supportedContentTypes"] = supportedContentTypes
	// Getting keys from datamodels
	var keys []string

	for key, value := range datamodels {
		keys = append(keys, key)
		didDocument[key] = value
	}
	didDocument["dids"] = keys

	return didDocument

}
func createDataModel(property Properties) DataModel {

	datamodel := DataModel{
		Property: property,
	}

	return datamodel
}

func createProperty(context []string, id string, methods []VerMethod, auth []string) Properties {

	property := Properties{
		Context: context,
		Id:      id,
		Method:  methods,
		Auth:    auth}

	return property
}

func createVerMethod(id string, typ string, controller string, publicKeyMB string) VerMethod {

	verMethod := VerMethod{
		Id:          id,
		Typ:         typ,
		Controller:  controller,
		PublicKeyMB: publicKeyMB}

	return verMethod
}

func printJson(doc any) {
	a1_json, err := json.Marshal(doc)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
	}
}
