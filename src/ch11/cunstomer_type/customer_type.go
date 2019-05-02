package cunstomer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner func(op int) int) func(op int) int{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend :",time.Since(start))
		return ret
	}
}

func slowFun(op int) int  {
	time.Sleep(time.Second *1)
	return op;
}
v
func TestFn(t *testing.T)  {
	a,_ := returnMultiValues()
	t.Log(a)
	tsSF :=timeSpent(slowFun)
	t.Log(tsSF(10))
}
