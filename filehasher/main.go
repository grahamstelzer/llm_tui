package main

import (
    "crypto/md5"
    "crypto/sha256"
	"hash"
	"flag"
	"fmt"
	"os"
	"io"
)

/*
	filehasher:
	- input: file path
		- ex "./filehasher --file notes.txt --algo sha256
	- read file contents
	- calculate hash (SHA256 or MD5)
	- print hash as string to terminal

	NOTE: to make this usable across terminal, must place
	 into directory in system PATH like ..local/bin or \system32

	NOTE: also, build via "go build -o filehasher main.go"
*/


func main() {

	// STEP 1: PARSE INPUT:
	//	- will use "flag" package for file path and has algo:

	// help check if cmd entered correclty
    if len(os.Args) == 1 {
        fmt.Println("Usage: ./filehasher --file=<input-file> --algo=<algoritm (md5/sha256) --output=<output-file>")
        flag.PrintDefaults()
        return
    }

	// flag.String() --> flag.String(name, defaultValue, description)
	filePath := flag.String("file", "", "path to file") // note: *filePath gets actual string
	algo := flag.String("algo", "sha256", "hash algorithm used")

	outputFile := flag.String("output", "", "file to save the hash")

	flag.Parse()
	// MUST be called after flag definitions
	// processes os.Args[1:]
	// fills the Flag Pointers with the value given from os
	//	otherwise shows defaultvalue

	fmt.Println("fp: ", *filePath)
	fmt.Println("algo: ", *algo)

	
	// STEP 2: open and read file
	//	use os package
	
	file, err := os.Open(*filePath)
	// immediate error checking as per go convention (i think?)
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close() // ensures file closed when main ends


	// STEP 3: HASHING FILE
	//	will use crypto/sha256 and ..md5

	// create hasher
	var hasher hash.Hash

	fmt.Println("note: this is not secure, do not use this hasher in actual production.")

	// switch
	switch *algo {
	case "sha256":
		hasher = sha256.New()
	case "md5":
		hasher = md5.New()
	default:
		fmt.Println("unsupported algo", *algo)
		return
	}

	// copy file contents into hasher:
	if _, err := io.Copy(hasher, file); err != nil {
		fmt.Println("error hashing file ", err)
		return
	}

	// calcs hash value
	hashSum := hasher.Sum(nil) 
	// NOTE: stores as series of bytes representing hash (byteslice)
	//	nil = empty slice

	fmt.Printf("hash (%s): %x\n", *algo, hashSum)
	// NOTE: %x converts byte slice to human readable string
	hashStr := fmt.Sprintf("%x", hashSum)


	// write to output file
    if *outputFile != "" {
        err := os.WriteFile(*outputFile, []byte(hashStr), 0644)
		// WriteFile expects byteslice []byte(...), 0644 is file perms
        if err != nil {
            fmt.Println("Error writing to output file:", err)
            return
        }
        fmt.Printf("Hash written to %s\n", *outputFile)
    }



}