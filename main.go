// ncms project main.go
package main

import (
	//	"baseType"
	"contorler"
	"fmt"
)

func main() {
	i := contorler.NewInsurePersonContorler()

	//	i.InsuredPerson["1"] = baseType.InsuredPerson{NcmsId: "1", Name: "yuanqi", Id: "210902198202131533"}
	//	i.InsuredPerson["2"] = baseType.InsuredPerson{NcmsId: "2", Name: "wangying", Id: "210902198202131533"}

	//	err := i.WriteInfoToFile()

	//	if err != nil {
	//		fmt.Println(err.Error())
	//	} else {
	//		fmt.Println("ok")
	//	}

	err := i.LoadInfoFromFile()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(i.InsuredPerson)
	}

}
