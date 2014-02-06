package main

import (
	"fmt"
)

func main() {
	m := New()
	m.Add("a", 123)
	m.Add("b", 456)
	m.Add("c", "789")

	fmt.Println(m.Get("a"))
	fmt.Println(m.Get("b"))
	fmt.Println(m.Get("c"))
	fmt.Println(m.Len())

	m.Remove("c")
	fmt.Println(m.Get("c"))
	fmt.Println(m.Len())

	m.Update("b", func(value interface{}, found bool) interface{} {
		if found {
			return value.(int) * 2
		}
		return ""
	})
	fmt.Println(m.Get("b"))

	fmt.Println(m.Close())
}

type SafeMap interface {
	Add(string, interface{})
	Remove(string)
	Len() int
	Get(string) (interface{}, bool)
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

type UpdateFunc func(interface{}, bool) interface{}

type safeMap chan commandData

type commandData struct {
	action  commandAction
	key     string
	value   interface{}
	result  chan<- interface{}
	data    chan<- map[string]interface{}
	updater UpdateFunc
}

type commandAction int

const (
	remove commandAction = iota
	end
	get
	add
	length
	update
)

func (sm safeMap) Add(key string, value interface{}) {
	sm <- commandData{action: add, key: key, value: value}
}

func (sm safeMap) Remove(key string) {
	sm <- commandData{action: remove, key: key}
}

type getResult struct {
	value interface{}
	found bool
}

func (sm safeMap) Get(key string) (interface{}, bool) {
	reply := make(chan interface{})
	sm <- commandData{action: get, key: key, result: reply}
	result := (<-reply).(getResult)
	return result.value, result.found
}

func (sm safeMap) Len() int {
	reply := make(chan interface{})
	sm <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (sm safeMap) Close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	sm <- commandData{action: end, data: reply}
	return <-reply
}

func (sm safeMap) Update(key string, updater UpdateFunc) {
	sm <- commandData{action: update, key: key, updater: updater}
}

func New() SafeMap {
	sm := make(safeMap)
	go sm.run()
	return sm
}

func (sm safeMap) run() {
	data := make(map[string]interface{})

	for command := range sm {
		switch command.action {
		case add:
			data[command.key] = command.value
		case remove:
			delete(data, command.key)
		case get:
			value, found := data[command.key]
			command.result <- getResult{value, found}
		case length:
			command.result <- len(data)
		case update:
			value, found := data[command.key]
			data[command.key] = command.updater(value, found)
		case end:
			close(sm)
			command.data <- data
		}
	}
}
