# bazel-intelij-debug-repro

Bazel plugin version: `2022.02.23.0.0-api-version-212`

This repository was created for the purpose of reproducing some behaviour to share with
https://github.com/bazelbuild/intellij. There is different behaviour when running versus
debugging a target.

## from commandline

```
bazel run //:demo
...
2022/06/14 15:16:47 MY_VARIABLE: hello world!
```

## from Intelij (clicking run on target)

```
...
2022/06/14 15:17:56 MY_VARIABLE: hello world!
```

## from Intelij (clicking debug on target)

```
...
2022/06/14 15:18:32 MY_VARIABLE:
```
