

## 文本和HTML模版

只是简单的格式化，使用Printf是完全足够的，但是有时候需要复杂的打印表格，这时候一般需要将
格式化代码分离出来以便更安全地修改，这写功能是由text/template和html/template等模版
包提供的，他们提供了一个将变量填充到一个文本和HTML格式的模版机制。

一个模版是一个字符串或一个文件，里面包含一个或多个由双花括号包含的`{{action}}`对象。
大部分的字符串只是按字面量打印，但是对于actions部分将触发其他行为，每个actions都包
含一个用模版语言书写的表达式，一个action虽然简短但是可以输出复杂的打印值，模版语言包含
选择结构体的成员、调用函数或方法、表达式控制流if-else语句和range循环语句等诸多特性。

```
const templ = `
{{.TotalCount}} issues:
{{range .Items}}-------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}
`
```

在action中`|`操作符表示将前一个表达式的结果作为后一个函数的输入，类似UNIX中管道的概念，
在Title这一行的action中第二操作是一个printf函数，是一个基于fmt.Sprintf实现的内置
函数，所有模版都可以直接使用，对于Age部分，第二个操作是daysAgo的函数，通过time.Since
函数将CreatedAt成员转换为过去的时间长度

```
func daysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}
```

生成模版的输出需要两个步骤，第一步是要分析模版并转为内部表示，然后基于指定的输入执行模版。
分析模版部分一般只需要执行一次。
```
report, err := template.New("report").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ)
if err != nil {
    log.Fatal(err)
}
```










