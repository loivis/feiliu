# 飞流直下三千尺
![](https://raw.githubusercontent.com/loivis/feiliu/master/libai.jpg)

go cli tool to tail cloudwatch logs from selected log group

# Installation

```
go get https://github.com/loivis/feiliu
```

# Usage

```
Usage of feiliu:
  -g string
    	name or prefix of log group
```

# Examples

`feiliu -g group-name`

+ If `group-name` has an exact match, log events will be streamed from last minute

+ If `group-name` matches no log group, all available log groups will be listed

+ If `group-name` matched multiple groups, all groups with prefix `group-name` will be listed
