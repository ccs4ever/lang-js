package main

import (
	"log"
	//"cuelang.org/go/cue"
	"cuelang.org/go/internal/legacy/cue"
	"cuelang.org/go/cue/format"
	_ "cuelang.org/go/cue/errors"
	_ "cuelang.org/go/encoding"
	"cuelang.org/go/encoding/json"
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
	errV := json.Validate([]byte(source), v.value);
	log.Println("go json validate:");
	log.Println(errV);
	return errV;
}

type CueRuntime struct {
	runtime cue.Runtime
}

func (r *CueRuntime) Compile(filename string, source string) (res *js.Object, err error) {
	instance, err := r.runtime.Compile(filename, source)
	log.Println("go Compile Err:");
	log.Println(err);
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

func (i *CueInstance) LookupField(field string) (res *js.Object, err error) {
	s, errS := i.instance.Value().Struct();
	if errS != nil { log.Println("Struct err:"); log.Println(errS); 
		return nil, errS; }
	it := s.Fields(cue.All());
	log.Println("Lookup fields:");
	for it.Next() {
		log.Println(it.Label());
		log.Println(it.Value());
	}
	info, err := s.FieldByName(field, true);
	log.Println("Lookup err: "); 
	log.Println(err); 
	res = js.MakeWrapper(&CueValue{info.Value});
	return;
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
