package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

const (
	inFilename       = "input.txt"
	expectedFilename = "expected.txt"
	outFilename      = "ouput.txt"
)

func Test_americanise_americanise(t *testing.T) {
	var inFile, outFile *os.File
	var err error

	if inFile, err = os.Open(inFilename); err != nil {
		t.Fatal(err)
	}
	defer inFile.Close()

	if outFile, err = os.Create(outFilename); err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	defer os.Remove(outFilename)

	if err = americanise(inFile, outFile); err != nil {
		t.Fatal(err)
	}

	var inBytes, expectedBytes []byte
	if inBytes, err = ioutil.ReadFile(outFilename); err != nil {
		t.Fatal(err)
	}
	if expectedBytes, err = ioutil.ReadFile(expectedFilename); err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(inBytes, expectedBytes) != 0 {
		t.Error("output.txt is not as expected!")
	}
}
