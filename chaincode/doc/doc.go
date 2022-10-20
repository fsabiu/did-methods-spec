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
func CreateDidDocument(id string, didMethod string, implementation string, implementer string, supportedContentTypes []string, datamodels map[string]DataModel) DidDoc {
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

func PrintJson(doc any) {
	a1_json, err := json.Marshal(doc)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
	}
}
