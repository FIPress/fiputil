package fiputil

import (
	"bytes"
	"strconv"
)

func Map(array interface{}, f func(a interface{}) interface{}) interface{} {
	arrInt, ok := array.([]int)
	if ok {
		for i, a := range arrInt {
			arrInt[i] = f(a).(int)
		}
		return arrInt
	}

	arrStr, ok := array.([]string)
	if ok {
		for i, a := range arrStr {
			arrStr[i] = f(a).(string)
		}
		return arrStr
	}

	//todo: other types
	return nil
}

func Filter() {
	//todo
}

func MkString(array interface{}, start string, sep string, end string) string {
	buf := new(bytes.Buffer)
	if start != "" {
		buf.WriteString(start)
	}

	arrInt, ok := array.([]int)
	if ok {
		for i, l := 0, len(arrInt); i < l; i++ {
			buf.WriteString(strconv.Itoa(arrInt[i]))
			if i != l-1 {
				buf.WriteString(sep)
			}
		}
	}

	arrStr, ok := array.([]string)
	if ok {
		for i, l := 0, len(arrStr); i < l; i++ {
			buf.WriteString(arrStr[i])
			if i != l-1 {
				buf.WriteString(sep)
			}
		}
	}

	arrInterface, ok := array.([]interface{})
	if ok {
		for i, l := 0, len(arrInterface); i < l; i++ {
			buf.WriteString(ToString(arrInterface[i]))
			if i != l-1 {
				buf.WriteString(sep)
			}
		}
		for _, a := range arrInterface {
			buf.WriteString(ToString(a))
		}
	}

	//todo: other types

	if end != "" {
		buf.WriteString(end)
	}

	return buf.String()
}
