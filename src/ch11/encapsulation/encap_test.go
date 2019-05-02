package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

//func (e *Employee) String() string {
//	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
//}

func (e Employee) String() string {
		fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))

	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}

func TestStructOperations(t *testing.T)  {
	e := Employee{"0","Bob",20}
	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
	t.Log(e.String())
}