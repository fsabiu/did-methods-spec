package main

import (
	"fmt"

	"github.com/fsabiu/did-methods-spec/tools"
)

const windowLen = 4

func main() {

	TxID := "bc5221c648533646877505288fc50b6c6100394213694bf111f7a3183074a329"
	//TxID_similar := "bc52bc5348533646877505288fc50b6c6100394213694bf111f7a3183074a329"

	idString := tools.GenerateIdString(TxID, 1)
	fmt.Println("Idstring method 1: ", idString)

	idString = tools.GenerateIdString(TxID, 2)
	fmt.Println("Idstring method 2: ", idString)
}
