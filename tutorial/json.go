package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Msg struct {
	Name string
	Body string
	Time int64
}

func main()  {
	//json serialization
	msg := Msg{"Alice", "Hello", 1294706395881547000}
	marshal, e := json.Marshal(msg)
	if e != nil {
		panic(e)
	}
	s := string(marshal)
	fmt.Println(s)

	//json deserilization
	var m Msg
	err := json.Unmarshal(marshal, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	//generic decode json
	var i interface{}
	i = "string"
	i = 2011
	i = 2.77
	f := i.(float64)
	fmt.Println("area", math.Pi * f * f)

	var mm interface{}
	err1 := json.Unmarshal(marshal, &mm)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(mm)
}


