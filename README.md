Golang Flight Crawler
--

## 一、关键字
* Golang
* Crawler
* Ryanair
* Go-chart

## 二、运行
直接运行main.go即可

```
Today: Wed, 21 Mar 2018 11:14:09 CET
Origin: HAM(Hamburg Airport)
Destination: BCN(Barcelona El Prat)

Time: 2018-05-14 08:00:00 +0000 UTC - 2018-05-14 10:30:00 +0000 UTC
Price: 43.85 EUR
Duration: 2h30m0s
FlightNumber: FR 129

Time: 2018-05-15 07:45:00 +0000 UTC - 2018-05-15 10:15:00 +0000 UTC
Price: 43.85 EUR
Duration: 2h30m0s
FlightNumber: FR 129

Time: 2018-05-16 08:00:00 +0000 UTC - 2018-05-16 10:30:00 +0000 UTC
Price: 43.85 EUR
Duration: 2h30m0s
FlightNumber: FR 129

Time: 2018-05-17 09:15:00 +0000 UTC - 2018-05-17 11:45:00 +0000 UTC
Price: 89.75 EUR
Duration: 2h30m0s
FlightNumber: FR 129

Time: 2018-05-18 08:00:00 +0000 UTC - 2018-05-18 10:30:00 +0000 UTC
Price: 179.51 EUR
Duration: 2h30m0s
FlightNumber: FR 129

Time: 2018-05-19 08:00:00 +0000 UTC - 2018-05-19 10:30:00 +0000 UTC
Price: 179.51 EUR
Duration: 2h30m0s
FlightNumber: FR 129

```

## 三、TODO
### 当前：
* 从HAM到BCN的某4周票价情况
* 输出打印在了命令行，生成了一个简单的折线图
* 数据是运行程序当天

### 期望：
* 时间更长
* 折线图上显示最低票价和具体的日期
* 收集每天的运行程序的数据，进行数据分析