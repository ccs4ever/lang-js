const Value = require('./Value.js');

function Instance(instance) {
  this.__instance__ = instance;
}

Instance.prototype.value = async function value() {
  return new Value(this.__instance__.Value());
}

module.exports = Instance
