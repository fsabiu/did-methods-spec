package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
)

func main() {

	TxID := "bc5221c648533646877505288fc50b6c6100394213694bf111f7a3183074a329"

	idString := GenerateIdString(TxID, 4)
	fmt.Println("Idstring : ", idString)

}

func indexGenerator(s string) int {

	h := md5.New()
	io.WriteString(h, s)
	var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))
	index := rand.Intn(32)
	return index
}

func EncodeStringBase32(s string) string {
	index := indexGenerator(s)
	const encodeStd = "abcdefghijklmnopqrstuvwxyz234567"
	base32Encoded := encodeStd[index : index+1]
	return base32Encoded
}

func GenerateIdString(TxID string, windowLen int) string {

	idString := ""
	windowNum := len(TxID) / windowLen
	startIdx := 0
	endIdx := windowLen
	currentChunk := ""

	for i := 0; i <= windowNum-1; i++ {

		currentChunk = TxID[startIdx:endIdx]
		//fmt.Println(currentChunk)
		startIdx = startIdx + windowLen
		endIdx = endIdx + windowLen
		base32Encoded := EncodeStringBase32(currentChunk)
		idString = idString + base32Encoded

	}
	return idString
}
