---
title: Golang静态代码分析工具
date: 2017-04-15 12:37:36
tags: Golang
---

###[gosimple](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/) [github.com](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple)


| count | Description                               | Suggestion                             |
|-------|-------------------------------------------|----------------------------------------|
| 128   | [if err != nil { return err }; return nil](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336507/#40)|  `return err`                          |
| 110   | [if alerts == nil || len(alerts) == 0](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336503/#668)  | `if len(alerts) == 0`                  |                              
| 25    | [t.Sub(time.Now()](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336489/#91)                      |  `time.Until`                          |
| 41    | [code,_:= codeMap\[v.CouponCodeId\]](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336477/#262)      | ` code := codeMap[v.CouponCodeId] `    |
| 61    | [`if bankAvailable == false`](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336430/#26)               | `if !bankAvailable`                    |
| 27    | [strconv.FormatInt(int64(day), 10)](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336445/#2705)       | strconv.itoa(i)是FormatInt(i, 10)的简写 |
| 16    | [循环拼接slice](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336337/#143)                   | `append(s1, s2...)`                    |
|1|[select 里面只有一个case](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526335917/#231)|`Use a plain channel send or receive` |
|1| [strings.Index(path,pattern)>0](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526335904/#304)|`用strings.Contains(path,pattern)代替`|
|2|[for _ := range XXX](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336193/#215)|`简写成 for range XXX{}`|
|1|[字符串中包含两个\\](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526335894/#3)|`用raw string写法 `|
|18| [`if <expr> { return <bool> }; return <bool>`](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336420/#360)|`return <expr>` |
|14| [`s[a:len(s)]`](http://jenkins.yxapp.xyz/job/stticcheck0410/218/checkstyleResult/source.-5323513831526336159/#126) | s[a:]|
|60| [time.Now().Sub](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336485/#3959) |time.Since()|
|10|[`var x uint; x = 1`](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336166/#35) | `var x uint = 1`|
|108|[amountMap := make(map[int64]int64, 0)](http://jenkins.yxapp.xyz/job/staticcheck0410/218/checkstyleResult/source.-5323513831526336502/#672) | `amountMap := make(map[int64]int64)`   |




###[staticcheck](http://jenkins.yxapp.xyz/job/staticcheck0410/219/checkstyleResult/)  [github.com](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck)

| count | Description                               | Suggestion                             |
|-------|-------------------------------------------|----------------------------------------|
|117    |[this value of XXXX is never used](http://jenkins.yxapp.xyz/job/staticcheck0410/219/checkstyleResult/source.-5323513831526335882/#40)   ||
|13     |[if 里面为 空](http://jenkins.yxapp.xyz/job/staticcheck0410/219/checkstyleResult/source.-5323513831526335870/#740)                   ||
|4|    [Called testing.T.FailNow or SkipNow in a goroutine, which isn't allowed](http://jenkins.yxapp.xyz/job/staticcheck0410/219/checkstyleResult/source.-5323513831526335778/#2899)||
|1|[Using regexp.Match or related in a loop, should use regexp.Compile](http://jenkins.yxapp.xyz/job/staticcheck0410/219/checkstyleResult/source.-5323513831526335764/#499)||


###[unused](http://jenkins.yxapp.xyz/job/staticcheck0410/220/checkstyleResult/) [github.com](https://github.com/dominikh/go-tools/tree/master/cmd/unused)

| count | Description                               | Suggestion                             |
|-------|-------------------------------------------|----------------------------------------|
|56     |[func XXXX is unused](http://jenkins.yxapp.xyz/job/staticcheck0410/220/checkstyleResult/source.-5323513831526335744/#3612)                    |                                        |
|3      |[var XXX is unused](http://jenkins.yxapp.xyz/job/staticcheck0410/220/checkstyleResult/source.-5323513831526335742/#20)|                                                            |
|177|  [field XXX is unused](http://jenkins.yxapp.xyz/job/staticcheck0410/220/checkstyleResult/source.-5323513831526335698/#814)||


###[golint](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/) [github.com](https://github.com/golang/lint)

| count | Description                               | Suggestion                             |
|-------|-------------------------------------------|----------------------------------------|
|5727   | [变量里字母简写错误](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/source.-5323513831526335506/#34)   |缩写全大写|
|2205   | [变量不要有下划线](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/source.-5323513831526335486/#10)  |要用驼峰命名法|
| 5192  | [导出的变量要有注释](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/source.-5323513831526335499/#21)||
|931    | [给函数注释的格式不对](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/source.-5323513831526335489/#74)||
|420    | [if 语句里有return,不要写else](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/source.-5323513831526335220/#89)
|107    |[should drop = 0 from declaration of var trailAssetInterest; it is the zero value](http://jenkins.yxapp.xyz/job/staticcheck0410/222/checkstyleResult/source.-5323513831526334979/#4011)



###[go tool vet](http://jenkins.yxapp.xyz/job/staticcheck0410/226/checkstyleResult/) 全是high priority[链接](https://golang.org/cmd/vet/)  

| count | Description                               | Suggestion                             |
|-------|-------------------------------------------|----------------------------------------|
| 126   | [unkeyed field](http://jenkins.yxapp.xyz/job/staticcheck0410/226/checkstyleResult/source.-5323513831526322001/#15)|声明变量是 struct 要加key|
|43| [unreachable code](http://jenkins.yxapp.xyz/job/staticcheck0410/226/checkstyleResult/source.-5323513831526321994/#802)|删掉呗|
|8| [Sprintf之类的错误](http://jenkins.yxapp.xyz/job/staticcheck0410/226/checkstyleResult/source.-5323513831526321997/#330)||
|170|[有json tag 但是没导出](http://jenkins.yxapp.xyz/job/staticcheck0410/226/checkstyleResult/source.-5323513831526321798/#121)||
|64 |[possible formatting directive in Sprint call](http://jenkins.yxapp.xyz/job/staticcheck0410/226/checkstyleResult/source.-5323513831526321971/#56)||





