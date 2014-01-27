package font

import (
	"fmt"
	"log"
)

type Font struct {
	family string
	size   int
}

func New(family string, size int) *Font {
	if !isValidFamily(family) {
		family = "serif"
	}
	if !isValidSize(size) {
		size = 11
	}
	return &Font{family, size}
}

func (f *Font) Family() string {
	return f.family
}

func (f *Font) SetFamily(family string) {
	if isValidFamily(family) {
		f.family = family
	}
}

func (f *Font) Size() int {
	return f.size
}

func (f *Font) SetSize(size int) {
	if isValidSize(size) {
		f.size = size
	}
}

func (f *Font) String() string {
	return fmt.Sprintf(`{font-family: "%s"; font-size: %dpt;}`, f.family, f.size)
}

func isValidFamily(family string) bool {
	if family == "" {
		log.Println("Invalid font-family given")
		return false
	}
	return true
}

func isValidSize(size int) bool {
	if size < 5 || size > 144 {
		log.Printf("Invalid font-size given [%d]\n", size)
		return false
	}
	return true
}
