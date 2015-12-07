/*
	参保人员信息控制器 2015.12.01 created by yq
*/
package contorler

import (
	"baseType"
	"bufio"
	"commonTools"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

type InsuredPersonContorler struct {
	InsuredPerson map[string]baseType.InsuredPerson
}

func NewInsurePersonContorler() *InsuredPersonContorler {
	return &InsuredPersonContorler{make(map[string]baseType.InsuredPerson)}
}

/*
	返回参合人员信息人数
*/
func (i *InsuredPersonContorler) GetNum() int {
	return len(i.InsuredPerson)
}

/*
	按新农合个人编码查找人员信息
*/
func (i *InsuredPersonContorler) Find(ncmsId string) (value baseType.InsuredPerson, err error) {
	if _, ok := i.InsuredPerson[ncmsId]; ok {
		value = i.InsuredPerson[ncmsId]
	} else {
		err = errors.New(baseType.ErrInsuredPersonNotExist)
	}
	return
}

/*
	增加人员信息
*/
func (i *InsuredPersonContorler) Add(person baseType.InsuredPerson) error {

	if _, ok := i.InsuredPerson[person.NcmsId]; ok {
		return errors.New(baseType.ErrInsuredPersonExist)
	} else {
		i.InsuredPerson[person.NcmsId] = person
	}
	return nil
}

/*
	修改人员信息
*/
func (i *InsuredPersonContorler) Modfiy(person baseType.InsuredPerson) error {

	if _, ok := i.InsuredPerson[person.NcmsId]; ok {
		i.InsuredPerson[person.NcmsId] = person
	} else {
		return errors.New(baseType.ErrInsuredPersonNotExist)
	}
	return nil
}

/*
	将人员信息写入文件中
*/
func (i *InsuredPersonContorler) WriteInfoToFile() error {

	var (
		path, fileName string
		err            error
		file           *os.File
		b              []byte
	)

	path, err = commonTools.GetCurrentPath()
	fileName = strings.Replace(path, "\\", "/", -1) + "/data/InsuredPerson_" + time.Now().Format("20060102") + ".dat"
	file, err = os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, value := range i.InsuredPerson {

		b, err = json.Marshal(value)

		if err != nil {
			return err
		}

		b = append(b, '\n')

		file.Write(b)
	}

	return nil

}

/*
	从文件中加载人员信息
*/
func (i *InsuredPersonContorler) LoadInfoFromFile() error {
	var (
		path, fileName string
		err            error
		file           *os.File
		bData          []byte
		person         baseType.InsuredPerson
	)

	path, err = commonTools.GetCurrentPath()

	if err != nil {
		return err
	}

	fileName = strings.Replace(path, "\\", "/", -1) + "/data/InsuredPerson_" + time.Now().Format("20060102") + ".dat"

	file, err = os.Open(fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	br := bufio.NewReader(file)

	bData, err = br.ReadBytes('\n')

	for ; err == nil; bData, err = br.ReadBytes('\n') {

		err = json.Unmarshal(bData, &person)
		if err != nil {
			return err
		} else {
			err = nil
		}
		i.InsuredPerson[person.NcmsId] = person
	}

	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
