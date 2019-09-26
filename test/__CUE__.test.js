const test = require('tape');

require('../build/cue');

test('__CUE__', (t) => {
  t.test('MarshallJSON', (assert) => {
    assert.equal(
      new __CUE__().Compile('test', 'foo: "bar"')[0].Value().MarshalJSON()[0],
      '{"foo":"bar"}'
    );
    assert.end();
  });
  t.test('MarshallJSON', (assert) => {
    assert.equal(
      new __CUE__().Compile('test', 'foo: "bar"')[0].Value().ToString()[0],
      `{\n\tfoo: "bar"\n}`
    );
    assert.end();
  });
});
