Omit:: {
  orig: {}
  pathsToOmit: {[_]: true | {} }
  nestedOmits: { for k, v in data {[k]: Omit & { orig: v, pathsToOmit: pathsToOmit[k] }} }
  result: {
    for k, v in orig
      let checks: {
        valueIsStruct: (isValueStruct & { value: v }).result
        inPathToOmit: (isKeySet & {struct: pathsToOmit, key: k}).result
        pathToOmit: inPathToOmit && (isValueBool & {value: pathsToOmit[k]}).result
      }
      if !checks.pathToOmit
        { "\(k)": if checks.valueIsStruct then (nestedOmits & {"/(k): _ })[k].result else v }
  }
}

data: {a: 3, c: d: 4, c: e: 5}
dataOmitted: { orig: data, pathsToOmit: {c: d: true} }

test: {
  keySetFoo: (isKeySet & {struct: { ban: "something" }, key: "foo"}).result
  keySetBan: (isKeySet & {struct: { ban: "something" }, key: "ban"}).result
  keySetFooStruct: (isKeySet & {struct: { ban: { asdfafoo: "something"} }, key: "foo"}).result
  keySetBanStruct: (isKeySet & {struct: { ban: { asdfasfoo: "something"} }, key: "ban"}).result
  keySetBanStruct: (isKeySet & {struct: { ban: { asdfasfoo: "something"} }, key: "ban"}).result
  keySetBanStructAnything: (isKeySet & {struct: { ban: _ }, key: "ban"}).result
}

isKeySet:: {
  struct: {...} // set the key we're checking
  key: string
  result: (struct & {[_]: _ | *_|_})[key] != _|_
}

isValueStruct:: {
  value: _
  result: ((value & {}) | *_|_) != _|_
}

isValueBool:: {
  value: _
  result: ((value & bool) | *_|_) != _|_
}

test: {
  isValueStructStruct: (isValueStruct & { value: { foo: "bar" } }).result
  isValueStructNull: (isValueStruct & { value: null }).result
  isValueStructInt: (isValueStruct & { value: int }).result
  isValueStructInt4: (isValueStruct & { value: 4 }).result
  isValueStructString: (isValueStruct & { value: "foo" }).result
  isValueStructBool: (isValueStruct & { value: bool }).result

  isValueBoolBoolBool: (isValueBool & { value: true }).result
  isValueBoolBoolStruct: (isValueBool & { value: { foo: "bar" } }).result
  isValueBoolStruct: (isValueBool & { value: { foo: "bar" } }).result
  isValueBoolNull: (isValueBool & { value: null }).result
  isValueBoolInt: (isValueBool & { value: int }).result
  isValueBoolInt4: (isValueBool & { value: 4 }).result
  isValueBoolString: (isValueBool & { value: "foo" }).result
}

something: 5 & {}

//keysetstruct: { [key=_]: (struct & {[_]: _ | *_|_})[key] != _|_ }
