package doc

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
	Property *Properties `json:"didDocumentDataModel"`
}

type DidDoc map[string]interface{}

// Model
func (didDocument DidDoc) AddDataModel(datamodels map[string]DataModel) {
	//var keys []string

	keys := didDocument["dids"].([]string)

	for key, value := range datamodels {
		didDocument[key] = value
		keys = append(keys, key)
	}
	didDocument["dids"] = keys
}

func (didDocument DidDoc) AddAuthMethod(methods map[string][]VerMethod) {

	for key, value := range methods {
		p := append(didDocument[key].(DataModel).Property.Method, value...)
		authDids := append(didDocument[key].(DataModel).Property.Auth, key)
		didDocument[key].(DataModel).Property.Method = p
		didDocument[key].(DataModel).Property.Auth = authDids

	}
}

func CreateDidDocument(did string, didMethod string, implementation string, implementer string, supportedContentTypes []string, publicKey string) DidDoc {

	didDocument := make(DidDoc)
	didDocument["id"] = did
	didDocument["didMethod"] = didMethod
	didDocument["implementation"] = implementation
	didDocument["implementer"] = implementer
	didDocument["supportedContentTypes"] = supportedContentTypes
	// Getting keys from datamodels
	//var keys []string

	verMeth := CreateVerMethod(did+"#key-1", "Ed25519VerificationKey2020", did, publicKey)

	contexts := []string{"https://www.w3.org/ns/did/v1", "https://w3id.org/security/suites/ed25519-2020/v1"}

	prop := CreateProperty(contexts, did+"#key-1", []VerMethod{verMeth}, []string{did})

	dm := CreateDataModel(prop)

	didDocument[did] = dm

	didDocument["dids"] = []string{did}

	return didDocument

}
func CreateDataModel(property Properties) DataModel {

	datamodel := DataModel{
		Property: &property,
	}

	return datamodel
}

func CreateProperty(context []string, id string, methods []VerMethod, auth []string) Properties {

	property := Properties{
		Context: context,
		Id:      id,
		Method:  methods,
		Auth:    auth}

	return property
}

func CreateVerMethod(id string, typ string, controller string, publicKeyMB string) VerMethod {

	verMethod := VerMethod{
		Id:          id,
		Typ:         typ,
		Controller:  controller,
		PublicKeyMB: publicKeyMB}

	return verMethod
}

// Controller
/* func Create(publicKey string) {

	// call to get idstring
	idstring := "QC5S8KGCFN37Z5VP"

} */

func PrintJson(doc any) {
	a1_json, err := json.Marshal(doc)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
	}
}
