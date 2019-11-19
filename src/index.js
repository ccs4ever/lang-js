const wasm = require('./main.go').default;

global.thing = wasm;
