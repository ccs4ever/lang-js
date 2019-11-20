//data: {a: 3, c d: 4, c: e: 5}
//
//data_: { for k, v in data {[k]: v }}
//
//Omit:: {
//  orig: {}
//  pathsToOmit: {[_]: true | {} }
//  nestedOmits: { for k, v in data {[k]: Omit }}
//  result: { 
//    for k, v in orig
//      let shouldReove = (
//        v & bottomOmitValues[k]
//        nestedOmit[k] &
//        { orig: v, bottomOmitValues: (bottomOmitValues[k] | *_) }
//      ).result | *v
//      if (orig & bottomOmitValues) != _|_ && orig & {}
//        { "\(k)": v } }
//}
//
//dataOmitted: { orig: data, bottomOmitValues:

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


//keysetstruct: { [key=_]: (struct & {[_]: _ | *_|_})[key] != _|_ }
