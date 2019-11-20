//data: {a: 3, c d: 4, c: e: 5}
//
//data_: { for k, v in data {[k]: v }}
//
//Omit:: {
//  orig: {}
//  bottomOmitValues: {[key=_]: *_|_ | *"remove"}
//  nestedOmits: { for k, v in data {[k]: Omit }}
//  result: { 
//    for k, v in orig
//      let nested = (
//        v & bottomOmitValues[k]
//        nestedOmit[k] &
//        { orig: v, bottomOmitValues: (bottomOmitValues[k] | *_) }
//      ).result | *v
//      if (orig & bottomOmitValues) != _|_ && orig & {}
//        { "\(k)": v } }
//}
//
//omit: data_ & { ["c"]: d: _|_  } & data

test: {
  keySetFoo: (isKeySet & {struct: { ban: "something" }, key: "foo"}).result
  keySetBan: (isKeySet & {struct: { ban: "something" }, key: "ban"}).result
  keySetFooStruct: (isKeySet & {struct: { ban: { asdfafoo: "something"} }, key: "foo"}).result
  keySetBanStruct: (isKeySet & {struct: { ban: { asdfasfoo: "something"} }, key: "ban"}).result
}

isKeySet: {
  struct: {} // set the key we're checking
  key: string
  result: (struct & {[_]: _ | *"someuniquestringnotSet"})[key] != _|_
}


//keysetstruct: { [key=_]: (struct & {[_]: _ | *"someuniquestringnotSet"})[key] != _|_ }
