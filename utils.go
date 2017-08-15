package main

import (
	"os"
	"log"
	"unicode/utf8"
)

// A Collection of Golang Wrapper Util Functions

func Exists(path string) bool{
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// got error, panic it
		panic(err)
	}
	return true
}


// determine a path is directory or file
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	mode := fileInfo.Mode()
	if mode.IsDir(){
		return true
	} else {
		return false
	}
}

func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	mode := fileInfo.Mode()
	if mode.IsRegular(){
		return true
	} else {
		return false
	}
}

func IsPlainTextFile(path string) bool{
	// this method judge a file is text file or binary file
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	b1 := make([]byte, 40)
	f.Read(b1)
	if utf8.Valid(b1){
		return true
	} else {
		// some txt file maybe not utf-8, we have to find them out
		// or even convert them into utf-8
		return false
	}
	return false
}