function Value(value) {
  this.__value__ = value;
}

Value.prototype.marshalJSON = async function marshallJSON() {
  [json, err] = this.__value__.MarshalJSON();
  if (err) { throw err; }
  return json;
}

Value.prototype.toString = async function toString() {
  [json, err] = this.__value__.ToString();
  if (err) { throw err; }
  return json;
}

module.exports = Value;
