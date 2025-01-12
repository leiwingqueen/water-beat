# 羽毛球-打水小程序

## 项目介绍

这是一个羽毛球打水小程序，主要功能有：水费计算

## 使用方式

### 1. 安装

```shell
go install github.com/leiwingqueen/water-beat@latest
```

### 2. 输入对局结果

参考 [test.csv](./test.csv)

格式：队伍1-参赛人员1,参赛人员2,队伍2-参赛人员1,参赛人员2,比分,水费

### 3. 运行程序

```shell
water-beat ./test.csv
```