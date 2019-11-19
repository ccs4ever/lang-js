package main

import (
	"fmt"
	"math/big"

	"cuelang.org/go/cue"
	"github.com/aaronpowell/webpack-golang-wasm-async-loader/gobridge"
)

func add(i []js.Value) (interface{}, error) {
	const config = `
TimeSeries: {
  "2019-09-01T07:00:00Z": 36
}
TimeSeries: {
  "2019-09-01T07:10:59Z": 200
}
`

	const config1 = `
TimeSeries: {
  "2019-09-01T08:00:00Z": 36
}
TimeSeries: {
  "2019-09-01T08:10:59Z": 200
}
`
	var r cue.Runtime

	instance, err := r.Compile("test", config)
	if err != nil {
		fmt.Println(err)
	}

	instance1, err1 := r.Compile("test1", config1)
	if err1 != nil {
		fmt.Println(err1)
	}

	instance2 := cue.Merge(instance, instance1)

	var bigInt big.Int
	instance2.Lookup("TimeSeries", "2019-09-01T07:10:59Z").Int(&bigInt)
	fmt.Println(bigInt.String())
}

func main() {
	c := make(chan struct{}, 0)

	gobridge.RegisterCallback("add", add)
	gobridge.RegisterValue("someValue", "Hello World")

	<-c
}
