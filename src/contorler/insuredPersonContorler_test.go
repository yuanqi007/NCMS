package contorler

import (
	"baseType"
	"testing"
)

func TestLen(t *testing.T) {
	i := NewInsurePersonContorler()
	i.InsuredPerson["1"] = baseType.InsuredPerson{NcmsId: "1", Name: "yuanqi", Id: "210902198202131533"}
	i.InsuredPerson["2"] = baseType.InsuredPerson{NcmsId: "2", Name: "wangying", Id: "210902198202131533"}

	if i.Len() != 2 {
		t.Errorf("Len() is failed. Got %d,expected 2", i.Len())
	}
}

func TestAdd(t *testing.T) {
	i := NewInsurePersonContorler()
	person := baseType.InsuredPerson{NcmsId: "1", Name: "yuanqi", Id: "210902198202131533"}
	i.Add(person)
	err := i.Add(person)

	if err == nil {
		t.Errorf("aaa")
	}

	//	if i.Len() == 1 && i.InsuredPerson["1"].Name == "yuanqi" {

	//	} else {
	//		t.Errorf("Len() is failed. Got %d,expected 2", i.Len())
	//	}
}

//func TestModify(t *testing.T) {
//	i := NewInsurePersonContorler()

//}
