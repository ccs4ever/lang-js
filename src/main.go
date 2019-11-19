//+ build js,wasm

package main

import (
	"errors"
	"strconv"
	"syscall/js"

	"cuelang.org/go/cue"
	"github.com/aaronpowell/webpack-golang-wasm-async-loader/gobridge"
)

var global = js.Global()

func add(this js.Value, args []js.Value) (interface{}, error) {
	ret := 0

	for _, item := range args {
		val, _ := strconv.Atoi(item.String())
		ret += val
	}

	return ret, nil
}

func err(this js.Value, args []js.Value) (interface{}, error) {
	return nil, errors.New("This is an error")
}

func cueEval(this js.Value, args []js.Value) (interface{}, error) {
	var r cue.Runtime

	println(args[0].String())

	instance, err := r.Compile("test", args[0].String())
	if err != nil {
		return nil, err
	}
	b, err1 := instance.Value().MarshalJSON()
	ret := string(b[:])
	println(ret)
	if err1 != nil {
		return nil, err1
	}

	return ret, nil
}

func main() {
	c := make(chan struct{}, 0)
	println("Web Assembly is ready")
	gobridge.RegisterCallback("add", add)
	gobridge.RegisterCallback("cue", cueEval)
	gobridge.RegisterCallback("raiseError", err)
	gobridge.RegisterValue("someValue", "Hello World")

	<-c
}
