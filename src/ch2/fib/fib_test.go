package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T)  {
	//var a int = 1
	//var b int = 1
	a:=1
	b:=1
	t.Log(a)
	fmt.Print(a,"")
	for  i:=0 ;i<5;i++  {
		t.Log(" ",b)
		tmp:=a
		a=b
		b=tmp+a
	}
}

func TestExchange(t *testing.T)  {
	a:=1
	b:=2
	//tmp:=a
	//a=b
	//b=tmp
	//b=a
	a,b = b,a
	t.Log(a,b)
}