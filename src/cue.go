package main

import (
	"log"
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/format"
	_ "cuelang.org/go/cue/errors"
	_ "cuelang.org/go/encoding"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("__CUE__", js.MakeWrapper(&Cue{}))
}

type Cue struct{}

func (_ *Cue) New() *js.Object {
	return js.MakeWrapper(&CueRuntime{})
}

func (_ *Cue) Merge(inst ...*CueInstance) *js.Object {
	var instances []*(cue.Instance)
	for i := range inst {
		instances = append(instances, inst[i].instance)
	}
	instance := cue.Merge(instances[:]...)
	return js.MakeWrapper(&CueInstance{instance})
}

func (r *CueRuntime) ValidateJSON(source string, v *CueValue) error {
	inst, errC := r.runtime.Compile("ValidateJSON", source);
	if errC != nil { return errC; }
	uv := v.value.Unify(inst.Value());
	if uv.Err() != nil { return uv.Err(); }
	err := uv.Validate(cue.Concrete(true));
	log.Println("go Validate error: ");
	log.Println(err);
	return err;
}

type CueRuntime struct {
	runtime cue.Runtime
}

func (r *CueRuntime) Compile(filename string, source string) (res *js.Object, err error) {
	instance, err := r.runtime.Compile(filename, source)
	if err != nil {
		return
	}
	res = js.MakeWrapper(&CueInstance{instance})
	return
}

type CueInstance struct {
	instance *cue.Instance
}

func (i *CueInstance) Value() *js.Object {
	value := i.instance.Value()
	return js.MakeWrapper(&CueValue{value})
}

type CueValue struct {
	value cue.Value
}

func (v *CueValue) MarshalJSON() (res string, err error) {
	b, err := v.value.MarshalJSON()
	if err != nil {
		return
	}
	res = string(b[:])
	return
}

func (v *CueValue) ToString(option ...string) (res string, err error) {
	var opts []cue.Option
	for _, opt := range option {
		switch opt {
		case "All":
			opts = append(opts, cue.All())
		case "Attributes":
			opts = append(opts, cue.Attributes(true))
		case "Definitions":
			opts = append(opts, cue.Definitions(true))
		case "DisallowCycles":
			opts = append(opts, cue.DisallowCycles(true))
		case "Hidden":
			opts = append(opts, cue.Hidden(true))
		case "Optional":
			opts = append(opts, cue.Optional(true))
		case "Raw":
			opts = append(opts, cue.Raw())
		}
	}

	b, err := format.Node(v.value.Syntax(opts[:]...))
	if err != nil {
		return
	}
	res = string(b[:])
	return
}
