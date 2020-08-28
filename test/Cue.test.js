const test = require('tape');

const { Cue } = require('../');

test('Cue', (t) => {
  t.test('MarshallJSON', async (assert) => {
    const cue = new Cue()
    const instance = await cue.compile('test', 'foo: "bar"');
    const value = await instance.value();
    const json = await value.marshalJSON();
    assert.equal(
      json,
      '{"foo":"bar"}'
    );
    assert.end();
  });
  t.test('ValidateJSON valid', async (assert) => {
    const cue = new Cue()
    const i = await cue.compile('test', 'foo: string');
    const v = await i.value();
    t.doesNotThrow(async ()=>{ await cue.validateJSON(
        '{"foo": "bar"}', v);}, null, "Validating valid JSON");
    assert.end();
  });
  t.test('ValidateJSON unification failure', async (assert) => {
    const cue = new Cue()
    const i = await cue.compile('test', 'foo: number');
    const v = await i.value();
    t.throws(async ()=>{ await cue.validateJSON(
        '{"foo": "bar"}', v);}, /foo/, "foo should fail to unify (string/number)");
    assert.end();
  });
  t.test('ToString', async (assert) => {
    const cue = new Cue()
    const instance = await cue.compile('test', 'foo: "bar"');
    const value = await instance.value();
    const json = await value.toString();
    assert.equal(
      json,
      `{\n\tfoo: "bar"\n}`
    );
    assert.end();
  });
  t.test('Merge', async (assert) => {
    const cue = new Cue()
    const i0 = await cue.compile('test', 'foo: "bar"');
    const i1 = await cue.compile('test', 'bar: "baz"');
    const json = await cue.merge(i0, i1).then(i => i.value()).then(v => v.marshalJSON());
    assert.equal(
      json,
      `{"foo":"bar","bar":"baz"}`
    );
    assert.end();
  });
  t.end();
});
