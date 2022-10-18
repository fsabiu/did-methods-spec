package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type DataModel struct {
		DidDocumentDataModel string
	}
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
	s := "did:orcl:123"
	a1 := make(map[string]interface{})
	a1["didMethod"] = "didMethod"
	a1["implementer"] = "implementer"
	a1[s] = DataModel{DidDocumentDataModel: "Test"}
	a1["supportedContentTypes"] = []string{"application/did+json",
		"application/did+ld+json"}

	a1_json, err := json.Marshal(a1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(a1_json))
	}

}
