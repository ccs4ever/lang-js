package main

import (
	"cuelang.org/go/cue"
	"github.com/gopherjs/gopherjs/js"
)

type Cue struct {
	runtime cue.Runtime
}

func New() *js.Object {
	return js.MakeWrapper(&Cue{})
}

func main() {
	js.Global.Set("cue", map[string]interface{}{
		"New": New,
	})
}

func (c *Cue) Compile(data string) string {
	instance, err := c.runtime.Compile("test", data)
	if err != nil {
		println(err)
		return ""
	}

	b, err1 := instance.Value().MarshalJSON()
	ret := string(b[:])
	println(ret)
	if err1 != nil {
		println(err1)
		return ""
	}

	return ret
}
