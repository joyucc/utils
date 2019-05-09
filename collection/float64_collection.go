package collection

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

type Float64Collection struct{
	AbsCollection
	objs []float64
}

func NewFloat64Collection(objs []float64) *Float64Collection {
	objs2 := make([]float64, len(objs))
	reflect.Copy(reflect.ValueOf(objs2), reflect.ValueOf(objs))
	arr := &Float64Collection{
		objs:objs2,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.compare = func(i interface{}, i2 interface{}) int {
		int1 := i.(float64)
		int2 := i2.(float64)
		if int1 > int2 {
			return 1
		}
		if int1 < int2 {
			return -1
		}
		return 0
	}
	return arr
}

func (arr *Float64Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(float64); ok {
		length := len(arr.objs)

		// 如果是append操作，直接调用系统的append，不新创建collection
		if index == length {
			arr.objs = append(arr.objs, i)
			return arr
		}

		new := arr.objs[0: index]
		new = append(new, i)
		new = append(new, arr.objs[index:length]...)
		arr.objs = new
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}

func (arr *Float64Collection) Remove(i int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1: len]...)
	return arr
}

func (arr *Float64Collection) NewEmpty(err ...error) ICollection {
	intArr := NewFloat64Collection([]float64{})
	if len(err) != 0 {
		intArr.err = err[0]
	}
	return intArr
}


func (arr *Float64Collection) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *Float64Collection) Count() int {
	return len(arr.objs)
}

func (arr *Float64Collection) DD() {
	ret := fmt.Sprintf("Float64Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%f\n",k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}