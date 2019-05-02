package condition

import "testing"

func TestIfMultiCase(t *testing.T)  {
	for i:=0;i<5;i++{
		switch i {
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}

	}

}

func TestSwitchCaseCondition(t *testing.T)  {

	for i:=0;i<5;i++{
		switch i {
		case i%2 ==0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("")
		default:

		}
	}
}