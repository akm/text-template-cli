# Usages

## Without input object

```
$ text-template-cli example1.txt.tmpl
USER: takeshi
FOO: <no value>
BAR: <no value>
BAZ:
```

An environment variable `USER` is used to render.

## Use redirection to output to a file

```
$ text-template-cli example1.txt.tmpl dummy.env >output.txt
```

```
$ cat output.txt
USER: takeshi
FOO: foo_from_env
BAR: bar_from_env
BAZ:
```

## .env file

```
$ text-template-cli example1.txt.tmpl dummy.env
USER: takeshi
FOO: foo_from_env
BAR: bar_from_env
BAZ:
```

## JSON file

```
$ text-template-cli example1.txt.tmpl dummy.json
USER: takeshi
FOO: foo_from_json
BAR: <no value>
BAZ:
  - baz1
  - baz2
```

## YAML file

```
$ text-template-cli example1.txt.tmpl dummy.yaml
USER: takeshi
FOO: foo_from_yaml
BAR: <no value>
BAZ:
QUX.FOO: [baz1 baz2]
```

## Multiple input files

```
$ text-template-cli example1.txt.tmpl dummy.env dummy.json
Overwriting FOO: foo_from_env -> foo_from_json by dummy.json
USER: takeshi
FOO: foo_from_json
BAR: bar_from_env
BAZ:
  - baz1
  - baz2
```

```
$ text-template-cli example1.txt.tmpl dummy.env dummy.json dummy.yaml
Overwriting FOO: foo_from_env -> foo_from_json by dummy.json
Overwriting FOO: foo_from_json -> foo_from_yaml by dummy.yaml
USER: takeshi
FOO: foo_from_yaml
BAR: bar_from_env
BAZ:
  - baz1
  - baz2
QUX.FOO: [baz1 baz2]
```

## Suppress warning messages

Use redirection for stderr to /dev/null or somewhere.

```
$ text-template-cli example1.txt.tmpl dummy.env dummy.json dummy.yaml 2>/dev/null
USER: takeshi
FOO: foo_from_yaml
BAR: bar_from_env
BAZ:
  - baz1
  - baz2
QUX.FOO: [baz1 baz2]
```

## Sprig functions

Use functions `default`, `contains` and `ternary` from [sprig](https://masterminds.github.io/sprig/).

```
$ text-template-cli sprig.tmpl dummy.env
FOO: foo_from_env
FOO is from ENV
```

```
$ text-template-cli sprig.tmpl dummy.json
FOO: foo_from_json
FOO is not from ENV
```
