package main

import (
	"fmt"
	"strings"
)

type Count int
type StringMap map[string]string
type FloatChan chan float64

type Part struct {
	Id   int
	Name string
}

func main() {
	var i Count
	i = 7
	fmt.Printf("%v [%T] \n", i, i)

	sm := make(StringMap)
	sm["1"] = "1"
	sm["2"] = "2"
	sm["3"] = "3"
	fmt.Printf("%v [%T] \n", sm, sm)

	fc := make(FloatChan, 1)
	fc <- 1.234567890
	fmt.Printf("%v [%T] -> %v \n", fc, fc, <-fc)

	i = 1
	i.Inc()
	fmt.Printf("%v [%T] \n", i, i)
	i.Dec()
	i.Dec()
	fmt.Printf("%v [%T] \n", i, i)
	fmt.Printf("%v \n", i.IsZero())
	fmt.Printf("%v [%T] \n", int(i), int(i))

	// ---------------------------------------------
	part := Part{7, "screwdriver"}
	part.UpperCase()
	part.Id += 2
	fmt.Println(part, part.HasPrefix("s"))

	// --------------------------------------------- embedded
	special := SpecialItem{Item{"Green", 3, 5}, 200}
	fmt.Println(special.id, special.Item.id, special.price, special.quantity, special.catalogId, special.Cost()) // methods from embedded field available
	luxury := LuxuryItem{Item{"Green", 3, 5}, 4}
	fmt.Println(luxury.id, luxury.Item.id, luxury.price, luxury.quantity, luxury.markup, luxury.Cost()) // methods from embedded field available
}

func (c *Count) Inc()        { *c++ }
func (c *Count) Dec()        { *c-- }
func (c Count) IsZero() bool { return c == 0 }

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("<%d %q>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

// -------- embedding ---------
type Item struct {
	id       string
	price    float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item      // anonymous field (embedding)
	catalogId int
}
type LuxuryItem struct {
	Item   // anonymous field (embedding)
	markup float64
}

func (item *LuxuryItem) Cost() float64 { // overwrite Item's Cost method
	return item.Item.Cost() * item.markup
}
