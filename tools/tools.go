package tools

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math"
	"math/rand"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

const windowLen = 4 // used by idstring generator

// Used by Method 2
//const wordLength = 32 // Depends on the machine
const M = 5 // 5 bits to represent 32 chars

//this function returns the int64-converted checksum of a string s
func ComputeChecksum(s string) int64 {
	h := md5.New()       //init new new hash.Hash object computing the MD5 checksum
	io.WriteString(h, s) //write string into object
	//fmt.Println("h sum: ", h.Sum(nil))
	checksum := binary.BigEndian.Uint64(h.Sum(nil)) //compute checksum and convert into uint64
	//fmt.Println("Checksum: ", checksum)
	return int64(checksum)
}

//IndexGenerator takes as input a string generates an integer in the interval [0,31]
//The string is hashed into an integer and used as seed of the random generation
func IndexGenerator(s string) int {
	seed := ComputeChecksum(s) //compute checksum and assign to seed
	rand.Seed(seed)            //set seed
	index := rand.Intn(32)
	return index
}

//EncodeStringBase32 takes as input a string of arbitrary length and maps it into a base32 char
func EncodeStringBase32(s string) string {
	index := IndexGenerator(s)
	base32Encoded := encodeStd[index : index+1]
	return base32Encoded
}

func GenerateIdString(TxID string, method int) string {

	idString := ""

	switch method {
	case 1:
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
	case 2:
		s := 0
		e := windowLen

		for e <= len(TxID) {
			chunk := TxID[s:e]
			//fmt.Println("Chunk: ", chunk)

			asciiVect := String2Ascii(chunk)
			//fmt.Println("ASCII: ", asciiVect)

			oddVect := OddArrayGenerator(chunk, 100)
			//fmt.Println("Odd vect: ", oddVect)

			h := HashingVectors(asciiVect, oddVect)
			//fmt.Println("Hash: ", h)

			// Appending char to idstring
			idString = idString + encodeStd[h:h+1]

			// Step to next chunk
			s += windowLen
			e += windowLen
		}
	}

	return idString
}

// Hash functions returning value in range [0, 2^M -1] = [0, 31]
func HashingVectors(x []int, y []int) int {

	// Dot product
	dp := DotProduct(x, y)
	//fmt.Println("Dot product: ", dp)

	// Hasing vectors formula: https://en.wikipedia.org/wiki/Universal_hashing#Hashing_vectors
	//h := int(dp%int(math.Pow(2, 2*wordLength))) / int(math.Pow(2, (2*wordLength)-M))
	//fmt.Println("par: ", int(dp%int(math.Pow(2, 2*wordLength))))
	//fmt.Println("div: ", int(math.Pow(2, (2*wordLength)-M)))
	//fmt.Println("h1: ", h)

	//h = h % int(math.Pow(2, M))
	//fmt.Println("h2: ", h)

	h := dp % int(math.Pow(2, M))
	return h
}

// OddArrayGenerator returns a pseudo-random integer array of len(chunk) generated using the string "chunk" as a seed.
// The (odd) integers of the generated array "a" are included in [1,randRange].
func OddArrayGenerator(chunk string, randRange int) []int {

	seed := ComputeChecksum(chunk) //compute checksum of the block
	rand.Seed(seed)

	permutation := rand.Perm(randRange)
	a := permutation[0:len(chunk)]

	// Getting only odd numbers
	for i := 0; i < len(chunk); i++ {
		if a[i]%2 == 0 {
			a[i]++
		}
	}

	return a
}

// Returns the integer vector representing a string according to ASCII standard
func String2Ascii(chunk string) []int {
	var ascii []int

	for i := 0; i < len(chunk); i++ {
		ascii = append(ascii, int(chunk[i]))
	}

	return ascii
}

// Retuns the dot products of two integer vectors
func DotProduct(x []int, y []int) int {

	dp := -1
	if len(x) == len(y) {
		dp = 0
		for i := 0; i < len(x); i++ {
			dp += x[i] * y[i]
		}
	}

	return dp
}
