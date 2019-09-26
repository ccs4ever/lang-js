package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"runtime"
)

func main() {
	d3root := Require("d3")
	println(d3root.Call("max", []int{10, 20}).Float())
	js.Global.Set("pet", map[string]interface{}{
		"New": New,
	})
}

type Pet struct {
	name string
}

func New(name string) *js.Object {
	fmt.Println("calling name")
	return js.MakeWrapper(&Pet{name})
}

func (p *Pet) Name() string {
	return p.name
}

func (p *Pet) SetName(name string) {
	p.name = name
}

func Require(module string) *js.Object {
	if runtime.GOARCH != "js" {
		return nil
	}

	switch {
	case js.Global.Get("gopm_modules") != js.Undefined:
		return js.Global.Get("gopm_modules").Get(module)
	case js.Global.Get("require") != js.Undefined:
		return js.Global.Call("require", module)
	default:
		return nil
	}
}
