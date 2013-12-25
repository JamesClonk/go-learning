package stack

import (
	"fmt"
	"testing"
)

func Test_stack_Push(t *testing.T) {
	var stack Stack

	stack.Push(007)
	result, err := stack.Top()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == 007, fmt.Sprintf("Expected stack.Top() to be [%v], but was [%v]", 007, result))

	stack.Push("Mr. Blond")
	result, err = stack.Top()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == "Mr. Blond", fmt.Sprintf("Expected stack.Top() to be [%v], but was [%v]", "Mr. Blond", result))
}

func Test_stack_Pushf(t *testing.T) {
	var stack Stack

	stack.Pushf(007)

	result, err := stack.Top()
	assertTrue(t, result == nil && err != nil, fmt.Sprintf("Expected stack.Top() to be return an error [%v], but got result [%v]", err, result))

	stack = stack.Pushf(007)
	result, err = stack.Top()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == 007, fmt.Sprintf("Expected stack.Top() to be [%v], but was [%v]", 007, result))

	stack = stack.Pushf("Mr. Blond")
	result, err = stack.Top()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == "Mr. Blond", fmt.Sprintf("Expected stack.Top() to be [%v], but was [%v]", "Mr. Blond", result))

	stack.Pushf(1234567890)
	assertTrue(t, stack.Len() == 2, fmt.Sprintf("Expected stack.Len() to be [%v], but was [%v]", 2, stack.Len()))
}

func Test_stack_Top(t *testing.T) {
	var stack Stack

	result, err := stack.Top()
	assertTrue(t, result == nil && err != nil, fmt.Sprintf("Expected stack.Top() to be return an error [%v], but got result [%v]", err, result))

	stack.Push(007)
	stack.Push("Mr. Blond")
	result, err = stack.Top()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == "Mr. Blond", fmt.Sprintf("Expected stack.Top() to be [%v], but was [%v]", "Mr. Blond", result))
}

func Test_stack_Pop(t *testing.T) {
	var stack Stack

	result, err := stack.Pop()
	assertTrue(t, result == nil && err != nil, fmt.Sprintf("Expected stack.Pop() to be return an error [%v], but got result [%v]", err, result))

	stack.Push(007)
	stack.Push("Mr. Blond")
	result, err = stack.Pop()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == "Mr. Blond", fmt.Sprintf("Expected stack.Pop() to be [%v], but was [%v]", "Mr. Blond", result))
	assertTrue(t, stack.Len() == 1, fmt.Sprintf("Expected stack.Len() to be [%v], but was [%v]", 1, stack.Len()))

	result, err = stack.Pop()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == 007, fmt.Sprintf("Expected stack.Pop() to be [%v], but was [%v]", 007, result))
	assertTrue(t, stack.IsEmpty(), "Expected stack.IsEmpty() to be true")
}

func Test_stack_Popf(t *testing.T) {
	var stack Stack

	_, result, err := stack.Popf()
	assertTrue(t, result == nil && err != nil, fmt.Sprintf("Expected stack.Popf() to be return an error [%v], but got result [%v]", err, result))

	stack.Push(007)
	stack.Push("Mr. Blond")
	_, result, err = stack.Popf()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == "Mr. Blond", fmt.Sprintf("Expected stack.Popf() to be [%v], but was [%v]", "Mr. Blond", result))
	assertTrue(t, stack.Len() == 2, fmt.Sprintf("Expected stack.Len() to be [%v], but was [%v]", 2, stack.Len()))

	result, err = stack.Pop()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == "Mr. Blond", fmt.Sprintf("Expected stack.Pop() to be [%v], but was [%v]", "Mr. Blond", result))
	assertTrue(t, stack.Len() == 1, fmt.Sprintf("Expected stack.Len() to be [%v], but was [%v]", 1, stack.Len()))

	stack, result, err = stack.Popf()
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, result == 007, fmt.Sprintf("Expected stack.Popf() to be [%v], but was [%v]", 007, result))
	assertTrue(t, stack.IsEmpty(), "Expected stack.IsEmpty() to be true")
}

func Test_stack_Len(t *testing.T) {
	var stack Stack

	assertTrue(t, stack.Len() == 0, fmt.Sprintf("Expected stack.Len() to be [%v], but was [%v]", 0, stack.Len()))

	stack.Push(007)
	stack.Push("Mr. Blond")
	assertTrue(t, stack.Len() == 2, fmt.Sprintf("Expected stack.Len() to be [%v], but was [%v]", 2, stack.Len()))
}

func Test_stack_Cap(t *testing.T) {
	var stack Stack

	assertTrue(t, stack.Cap() == 0, fmt.Sprintf("Expected stack.Cap() to be [%v], but was [%v]", 0, stack.Cap()))

	stack.Push(007)
	stack.Push("Mr. Blond")
	assertTrue(t, stack.Cap() == 2, fmt.Sprintf("Expected stack.Cap() to be [%v], but was [%v]", 2, stack.Cap()))

	if _, err := stack.Pop(); err != nil {
		t.Fatal(err)
	}
	assertTrue(t, stack.Cap() == 2, fmt.Sprintf("Expected stack.Cap() to be [%v], but was [%v]", 2, stack.Cap()))
}

func Test_stack_IsEmpty(t *testing.T) {
	var stack Stack

	assertTrue(t, stack.IsEmpty(), "Expected stack.IsEmpty() to be true")

	stack.Push(007)
	stack.Push("Mr. Blond")
	assertFalse(t, stack.IsEmpty(), "Expected stack.IsEmpty() to be false")
}

func assertTrue(t *testing.T, assertion bool, msg string) {
	if !assertion {
		t.Errorf("Test:  %s", msg)
	}
}
func assertFalse(t *testing.T, assertion bool, msg string) {
	assertTrue(t, !assertion, msg)
}
