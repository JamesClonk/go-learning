package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// JSON --------------------------------------------
	file, err := os.Open("invoice.json")
	if err != nil {
		log.Fatal(err)
	}
	invoices, err := readInvoices(file, ".json")
	if err != nil {
		log.Fatal(err)
	}
	for _, invoice := range invoices {
		fmt.Printf("%v \n", invoice)
	}

	// XML ---------------------------------------------
	file, err = os.Open("test.xml")
	if err != nil {
		log.Fatal(err)
	}
	invoices, err = readInvoices(file, ".xml")
	if err != nil {
		log.Fatal(err)
	}
	for _, invoice := range invoices {
		fmt.Printf("%v \n", invoice)
	}

	// GOB ---------------------------------------------
	file, err = os.Open("invoice.gob")
	if err != nil {
		log.Fatal(err)
	}
	invoices, err = readInvoices(file, ".gob")
	if err != nil {
		log.Fatal(err)
	}
	for _, invoice := range invoices {
		fmt.Printf("%v \n", invoice)
	}
}

const (
	fileType    = "INVOICES"
	magicNumber = 0x125D
	fileVersion = 100
	dateFormat  = "2006-01-02"
)

type Invoice struct {
	Id         int
	CustomerId int
	Raised     time.Time
	Due        time.Time
	Paid       bool
	Note       string
	Items      []*Item
}
type Item struct {
	Id       string
	Price    float64
	Quantity int
	Note     string
}

type InvoicesMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}
type InvoicesUnmarshaler interface {
	UnmarshalInvoices(reader io.Reader) ([]*Invoice, error)
}

func readInvoices(reader io.Reader, suffix string) ([]*Invoice, error) {
	var unmarshaler InvoicesUnmarshaler
	switch suffix {
	case ".gob":
		unmarshaler = GobMarshaler{}
	// case ".inv":
	// 	unmarshaler = InvMarshaler{}
	case ".jsn", ".json":
		unmarshaler = JSONMarshaler{}
	// case ".txt":
	// 	unmarshaler = TxtMarshaler{}
	case ".xml":
		unmarshaler = XMLMarshaler{}
	}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalInvoices(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}
