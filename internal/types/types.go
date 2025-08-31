package types 

import (

	"encoding/json"
	"strconv"
	"strings"
	"reflect"
)

type KumiteInt int
type KumiteFloat float32
type KumiteString string

type KumiteSlice[T any] []T

type KumiteType interface {
	UnmarshalJSON(data []byte) error
  Convert(data string) error
}

 

func (kint *KumiteInt) UnmarshalJSON(data []byte) error {
	return processType(data, kint)
}

func (kint *KumiteSlice[T]) UnmarshalJSON(data []byte) error {
	return processType(data, kint)
}


func (kfloat *KumiteFloat) UnmarshalJSON(data []byte) error {
	return processType(data, kfloat)
}

func (kstring *KumiteString) UnmarshalJSON(data []byte) error {
	return processType(data, kstring)
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

func (kstring *KumiteString) Convert(data string) error {
  *kstring = KumiteString(data)
  return nil
}


func (kslice *KumiteSlice[T]) Convert(data string) error {
  var t T
	sliceType := reflect.TypeOf(t)

	elements := strings.SplitSeq(data, ",") 	
	switch sliceType {
	case reflect.TypeOf(KumiteInt(0)):
		for val := range elements {
			num, err := strconv.Atoi(val)
			if err != nil {
				return err
			}

			*kslice = append(*kslice, any(KumiteInt(num)).(T))
		}
	case reflect.TypeOf(KumiteString("")): 
		for val := range elements { 
			tmp := KumiteString(val)
			*kslice = append(*kslice, any(tmp).(T))
    }
	}
  


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


