package tools

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/rand"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

const windowLen = 8 // used by idstring generator

// Used by Method 2
//const wordLength = 32 // Depends on the machine
const M = 5 // 5 bits to represent 32 chars

// Used for tests
const TxIDAlphabet = "abcdefghijklmnopqrstuvwxyz01233456789"

//this function returns the int64-converted checksum of a string s
func ComputeChecksum(s string) int64 {
	h := md5.New()                                  //init new new hash.Hash object computing the MD5 checksum
	io.WriteString(h, s)                            //write string into object
	checksum := binary.BigEndian.Uint64(h.Sum(nil)) //compute checksum and convert into uint64
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
	currentChunk := ""
	startIdx := 0
	endIdx := windowLen

	switch method {
	case 1:
		for endIdx <= len(TxID) {
			currentChunk = TxID[startIdx:endIdx]       // Get current 4 chars chunk of TxId
			idChar := EncodeStringBase32(currentChunk) // Use checksum as a seed for index generation, then convert to char
			idString = idString + idChar               // Append new base32 char to idString

			// Step to next chunk
			startIdx += windowLen
			endIdx += windowLen
		}
	case 2:
		for endIdx <= len(TxID) {
			currentChunk = TxID[startIdx:endIdx]
			asciiVect := String2Ascii(currentChunk)
			oddVect := OddArrayGenerator(currentChunk, 100)

			h := HashingVectors(asciiVect, oddVect) // Getting hash value for current chunk [0, 32)
			idString = idString + encodeStd[h:h+1]  // Appending char to idstring

			// Step to next chunk
			startIdx += windowLen
			endIdx += windowLen
		}
	}

	return idString
}

// Hash functions returning value in range [0, 2^M -1] = [0, 31]
func HashingVectors(x []int, y []int) int {
	// Dot product
	dp := DotProduct(x, y)
	// Hash -> [0,32)
	h := dp % int(math.Pow(2, M))
	return h
}

// OddArrayGenerator returns a pseudo-random integer array of len(chunk) generated using the string "chunk" as a seed.
// The (odd) integers of the generated array "a" are included in [1,randRange].
func OddArrayGenerator(chunk string, randRange int) []int {

	seed := ComputeChecksum(chunk) // Compute checksum of the block
	rand.Seed(seed)                // Setting seed

	permutation := rand.Perm(randRange) // Generating permutation of array [0, randRange)
	a := permutation[0:len(chunk)]      // Fetching first len(chunk) elements

	// Getting only odd numbers (if even then sum 1)
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

// n = number of Transaction IDs
func TestIdStringGenerators(n int) {

	var TxIDs []string // Slice of TxIDs

	// Generating n TxIDs
	for i := 0; i < n; i++ {
		TxIDs = append(TxIDs, generateString(TxIDAlphabet, 64))
	}

	// Mapping idstrings -> counter
	idstrings1 := make(map[string]int)
	idstrings2 := make(map[string]int)

	// Collision counters
	collisions1 := 0
	collisions2 := 0

	// Counting collisions
	idstring := ""
	for i := 0; i < n; i++ {
		// Method 1
		idstring = GenerateIdString(TxIDs[i], 1)
		// Checking
		if i%1000 == 0 {
			fmt.Println("TxID #", i, ": ", TxIDs[i])
			fmt.Println("Method 1 idstring #", i, ": ", idstring)
		}

		if idstrings1[idstring] == 0 { // no collision
			idstrings1[idstring] = 1
		} else {
			idstrings1[idstring] += 1 // Collision
			collisions1 += 1          // Updating counter
		}

		// Method 2
		idstring = GenerateIdString(TxIDs[i], 2)
		if i%1000 == 0 {
			fmt.Println("Method 2 idstring #", i, ": ", idstring)
		}

		if idstrings2[idstring] == 0 { // no collision
			idstrings2[idstring] = 1
		} else {
			idstrings2[idstring] += 1 // Collision
			collisions2 += 1          // Updating counter
		}
	}

	fmt.Println("Collisions 1: ", collisions1)
	fmt.Println("Collisions 2: ", collisions2)
}

func generateString(alphabet string, length int) string {
	txID := ""
	alphabetLen := len(alphabet)

	for i := 0; i < length; i++ {
		index := rand.Intn(alphabetLen)       // generating random char
		txID = txID + alphabet[index:index+1] // appending char
	}

	return txID
}
