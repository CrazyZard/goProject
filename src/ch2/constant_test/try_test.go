package constant_test

import (
	"testing"
)

const (
	Monday =  1 + iota
	Tuesday
	Wednesday
)


func TestConstantTry(t *testing.T)  {
	t.Log(Monday,Tuesday)
}

const(
	Readable = 1<<iota
	Writable
	Executable
)

func TestConstantTry1(t *testing.T)  {
	a:=7
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable ==Executable)
}
