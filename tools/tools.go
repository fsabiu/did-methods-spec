package tools

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

//IndexGenerator takes as input a string generates an integer in the interval [0,31]
//The string is hashed into an integer and used as seed of the random generation
func IndexGenerator(s string) int {

	h := md5.New()                              //init new new hash.Hash object computing the MD5 checksum
	io.WriteString(h, s)                        //write string into object
	seed := binary.BigEndian.Uint64(h.Sum(nil)) //compute checksum and
	rand.Seed(int64(seed))                      //set seed
	index := rand.Intn(32)
	return index
}

//EncodeStringBase32 takes as input a string of arbitrary length and maps it into a base32 char
func EncodeStringBase32(s string) string {
	index := IndexGenerator(s)
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

		currentChunk = TxID[startIdx:endIdx] //get current 4 chars chunk of TxId
		startIdx = startIdx + windowLen
		endIdx = endIdx + windowLen
		idChar := EncodeStringBase32(currentChunk)
		idString = idString + idChar //append new base32 char to idString
	}

	return idString
}
