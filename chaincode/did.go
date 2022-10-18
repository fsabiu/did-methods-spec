package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
	//type DataModel struct {
	//DidDocumentDataModel string `json:"didDocumentDataModel"`}
	type Authentication struct {
	Authentication string[] `json:"authentication"`}
	
	
	
	
	//type context map[string] []string
*/
	s := "did:orcl:123"
	type VerMethod struct {
		Id string `json:"id"`
		Typ string `json:"type"`
		Controller string `json:"controller"`
		PublicKeyMB string `json:"publicKeyMultibase"`}

	type Properties struct {
		Context []string `json:"@context"`
		Id string `json:"id"`
		Method VerMethod `json:"verificationMethod"`
		Auth []string `json:"authentication"`
	}
	/*
	type Property struct {
	Context []string `json:"@context"`}
*/
	//type a map[string]DataMode
/*
	type Asset struct {
 	DidMethod string    `json:"didMethod"`
 	Implementation string `json:"implementation"`
  	Implementer string `json:"implementer"`
  	SupportedContentType [2]string `json:"supportedContentType "`
  	Mapexample map[string]DataModel `json:"did_blabla "`
	}

	a := make(map[string]DataModel)
	a["did:orcl:QC5S8KGCFN37Z5VP"] = DataModel{DidDocumentDataModel: "Test"}

	prova := Asset{DidMethod : "did:orcl",
					Implementation: "DID Oracle Test",
					Implementer : "Oracle",
					SupportedContentType : [2]string{"application/did+json","application/did+ld+json"},
    				Mapexample:a}
	//fmt.Println(prova)
	/*
	ja, err := json.Marshal(prova)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(ja))
	}
*/

/*
	s := "did:orcl:123"
	a1 := make(map[string]interface{})
	a1["didMethod"] = "didMethod"
	a1["implementer"] = "implementer"
	a1["dids"]= []string{s}
	a1[s] = DataModel{DidDocumentDataModel: "Properties"}
	a1["supportedContentTypes"] = []string {"application/did+json",
    "application/did+ld+json"}
    a1[s].Property
	a1_json, err := json.Marshal(a1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
	}
	//Authentication []string
	// Verificationmethod []
*/

	prova := Properties{Context : []string {"https://www.w3.org/ns/did/v1",
          				"https://w3id.org/security/suites/ed25519-2020/v1"},
    				Id: s,
    				Method : VerMethod {
            Id : "key-1",
            Typ : "Ed25519VerificationKey2020",
            Controller: s,
            PublicKeyMB: "cTx0CiPUqsTLr2xy53VAQUVfOn7dvFqHeeC1k"},
            Auth : []string {s} }

	a1_json, err := json.Marshal(prova)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
}

}
