package utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

type KumiteInt int
type KumiteFloat float32

type KumiteType interface {
	UnmarshalJSON(data []byte) error
  Convert(data string) error
}


func (kint *KumiteInt) UnmarshalJSON(data []byte) error {
	return processType(data, kint)
}

func (kfloat *KumiteFloat) UnmarshalJSON(data []byte) error {
	return processType(data, kfloat)
}


func (kint *KumiteInt) Convert(data string) error {
	num, err := strconv.Atoi(data)
	if err != nil {
	  return err
	}

  *kint = KumiteInt(num)
  return nil
}



func (kfloat *KumiteFloat) Convert(data string) error {
	data = strings.ReplaceAll(data, ",", ".")
	num, err := strconv.ParseFloat(data, 32)
	if err != nil {
	  return err
	}

  *kfloat = KumiteFloat(num)
  return nil
}


func processType(data []byte, val KumiteType) error {
	var tmp string

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

  return val.Convert(tmp)
}


