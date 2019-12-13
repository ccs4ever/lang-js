const test = require('tape');

require('../build/__CUE__.js');

test('__CUE__', (t) => {
  t.test('New', (t) => {
    const cue = new (__CUE__.New)();
    t.ok(cue);

    t.test(`Cue#Compile '{}'`, (t) => {
      const [instance, err] = cue.Compile('test', '{}');
      t.ok(instance);
      t.equal(err, null);

      t.test('Instance#Value', (t) => {
        value = instance.Value();
        t.ok(value);

        t.test('Value#MarshallJSON', (t) => {
          const [json, err] = value.MarshalJSON();
          t.equal(json, '{}');
          t.equal(err, null);
          t.end();
        });

        t.test('Value#ToString', (t) => {
          const [string, err] = value.ToString();
          t.equal(string, "{\n}");
          t.equal(err, null);
          t.end();
        });
        t.end();
      });
      t.end();
    });
    t.test(`Cue#Compile 'foo: "bar"'`, (t) => {
      const [instance, err] = cue.Compile('test', 'foo: "bar"');
      t.ok(instance);
      t.equal(err, null);

      t.test('Instance#Value', (t) => {
        value = instance.Value();
        t.ok(value);

        t.test('Value#MarshallJSON', (t) => {
          const [json, err] = value.MarshalJSON();
          t.equal(json, '{"foo":"bar"}');
          t.equal(err, null);
          t.end();
        });

        t.test('Value#ToString', (t) => {
          const [string, err] = value.ToString();
          t.equal(string, `{\n\tfoo: "bar"\n}`);
          t.equal(err, null);
          t.end();
        });
        t.end();
      });
      t.end();
    });
    t.end();
  });
  t.end();
});
