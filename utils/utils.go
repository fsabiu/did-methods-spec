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
	windowLen := 4
	windowNum := len(TxID) / windowLen
	startIdx := 0
	endIdx := windowLen
	currentChunk := ""
	for i := 0; i <= windowNum-1; i++ {

		currentChunk = TxID[startIdx:endIdx]
		fmt.Println(currentChunk)
		startIdx = startIdx + windowLen
		endIdx = endIdx + windowLen

	}
}

func indexGenerator(TxID string) int {

	h := md5.New()
	io.WriteString(h, TxID)
	var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))
	index := rand.Intn(32)
	return index
}
