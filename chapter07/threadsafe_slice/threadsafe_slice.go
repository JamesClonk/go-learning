package main

func main() {

}

type SafeSlice interface {
	Append(interface{})     // Append the given item to the slice
	At(int) interface{}     // Return the item at the given index position
	Close() []interface{}   // Close the channel and return the slice
	Delete(int)             // Delete the item at the given index position
	Len() int               // Return the number of items in the slice
	Update(int, UpdateFunc) // Update the item at the given index position
}

type UpdateFunc func(interface{}) interface{}

type safeSlice chan commandData

type commandData struct {
	action  commandAction
	index   int
	value   interface{}
	result  chan<- interface{}
	data    chan<- []interface{}
	updater UpdateFunc
}

type commandAction int

const (
	add commandAction = iota
	at
	end
	del
	length
	update
)

func (ss safeSlice) Append(value interface{}) {
	ss <- commandData{action: add, value: value}
}

func (ss safeSlice) At(index int) interface{} {
	reply := make(chan interface{})
	ss <- commandData{action: at, index: index, result: reply}
	return <-reply
}

func (ss safeSlice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- commandData{action: end, data: reply}
	return <-reply
}

func (ss safeSlice) Delete(index int) {
	ss <- commandData{action: del, index: index}
}

func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (ss safeSlice) Update(index int, updater UpdateFunc) {
	ss <- commandData{action: update, index: index, updater: updater}
}

func New() SafeSlice {
	ss := make(safeSlice)
	go ss.run()
	return ss
}

func (ss safeSlice) run() {
	data := []interface{}{}

	for command := range ss {
		switch command.action {
		case add:
			data = append(data, command.value)
		case at:
			if command.index >= 0 && command.index < len(data) {
				command.result <- data[command.index]
			} else {
				command.result <- nil
			}
		case end:
			close(ss)
			command.data <- data
		case del:
			if command.index >= 0 && command.index < len(data) {
				data = append(data[:command.index], data[command.index+1:]...)
				// copy(data[command.index:], data[command.index+1:])
				// data[len(data)-1] = nil
				// data = data[:len(data)-1]
			}
		case length:
			command.result <- len(data)
		case update:
			if command.index >= 0 && command.index < len(data) {
				data[command.index] = command.updater(data[command.index])
			}
		}
	}
}
