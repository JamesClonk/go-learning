package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
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
	// case ".gob":
	// 	unmarshaler = GobMarshaler{}
	// case ".inv":
	// 	unmarshaler = InvMarshaler{}
	case ".jsn", ".json":
		unmarshaler = JSONMarshaler{}
		// case ".txt":
		// 	unmarshaler = TxtMarshaler{}
		// case ".xml":
		// 	unmarshaler = XMLMarshaler{}
	}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalInvoices(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}

type JSONMarshaler struct{}

type JSONInvoice struct {
	Id         int
	CustomerId int
	Raised     string // time.Time in Invoice struct
	Due        string // time.Time in Invoice struct
	Paid       bool
	Note       string
	Items      []*Item
}

// for custom handling of time.Time
func (invoice Invoice) MarshalJSON() ([]byte, error) {
	jsonInvoice := JSONInvoice{
		invoice.Id,
		invoice.CustomerId,
		invoice.Raised.Format(dateFormat),
		invoice.Due.Format(dateFormat),
		invoice.Paid,
		invoice.Note,
		invoice.Items,
	}
	return json.Marshal(jsonInvoice)
}

// for custom handling of time.Time
func (invoice *Invoice) UnmarshalJSON(data []byte) (err error) {
	var jsonInvoice JSONInvoice
	if err = json.Unmarshal(data, &jsonInvoice); err != nil {
		return err
	}

	var raised, due time.Time
	if raised, err = time.Parse(dateFormat, jsonInvoice.Raised); err != nil {
		return err
	}
	if due, err = time.Parse(dateFormat, jsonInvoice.Due); err != nil {
		return err
	}

	*invoice = Invoice{
		jsonInvoice.Id,
		jsonInvoice.CustomerId,
		raised,
		due,
		jsonInvoice.Paid,
		jsonInvoice.Note,
		jsonInvoice.Items,
	}
	return nil
}

func (JSONMarshaler) MarshalInvoices(writer io.Writer, invoices []*Invoice) error {
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(fileType); err != nil {
		return err
	}
	if err := encoder.Encode(fileVersion); err != nil {
		return err
	}
	return encoder.Encode(invoices)
}

func (JSONMarshaler) UnmarshalInvoices(reader io.Reader) ([]*Invoice, error) {
	decoder := json.NewDecoder(reader)

	var kind string
	if err := decoder.Decode(&kind); err != nil {
		return nil, err
	}
	if kind != fileType {
		return nil, errors.New(fmt.Sprintf("cannot read non-invoices json file: [%v]", kind))
	}

	var version int
	if err := decoder.Decode(&version); err != nil {
		return nil, err
	}
	if version > fileVersion {
		return nil, fmt.Errorf("version %d is not supported", version)
	}

	var invoices []*Invoice
	err := decoder.Decode(&invoices)
	return invoices, err
}
