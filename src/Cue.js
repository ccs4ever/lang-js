require('../build/__CUE__.js');
const Instance = require('./Instance.js');

function Cue() {
  this.__cue__ = new __CUE__.New()
}

Cue.merge = async function merge(...instances) {
  return __CUE__.Merge(...instances)
}

Cue.prototype.compile = async function compile(filename, source) {
  [instance, err] = this.__cue__.Compile(filename, source);
  if (err) { throw err; }
  return new Instance(instance);
}

module.exports = Cue;
