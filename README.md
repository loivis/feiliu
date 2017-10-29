# 飞流直下三千尺
![](https://raw.githubusercontent.com/loivis/feiliu/master/libai.jpg)

go cli tool to tail cloudwatch logs from selected log group

# Installation

```
go get https://github.com/loivis/feiliu
```

# Help

```
cli tool to stream aws cloudwatch logs

Flags:
      --help                    Show context-sensitive help (also try --help-long and --help-man).
      --version                 Show application version.
  -g, --group-prefix=kinase     prefix of log groups
  -a, --api-gateway=k2i0n1a5se  id of api gateway
  -s, --stage=prod              name of api gateway stage
  -l, --lambda=function-name    name or prefix of lambda function
  -m, --match=something         substring to match name of log groups
  -t, --start=10m               when to start streaming, default 1 minute
```
