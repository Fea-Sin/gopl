
### 类型断言

类型断言是一个使用在接口值上的操作，语法上它看起来像x.(T)被称为断言类型，这里x表示一个接口的类型
和T表示一个类型，一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。

```
var w io.Writer
w = os.Stdout
f := w.(*os.File) // success f == os.Stdout
c := w.(*bytes.Buffer) // panic interface holds *os.File, not *bytes.Buffer
```

经常地，对一个接口值的动态类型我们是不确定的，并且我们更愿意去检验它是否是一些特定的类型。

```
var w io.Writer = os.Stdout

f, ok := w.(*os.File)      // success f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure f == nil
```

if语句的扩展格式让这个变的很简洁
```
if f, ok := w.(*os.File); ok {
    // ...
}
```

基于类型断言可以区别错误类型

通过类型断言询问行为

fmt.Fprintf函数怎么从其它所有值中区分满足error或者fmt.Stringer接口的值，在
fmt.Fprintf内部有一个将单个操作对象转换称一个字符串的步骤。
```
package fmt

func formatOneValue(x interface{}) string {
    if err, ok := x.(error); ok {
        return err.Error()
    }
    if str, ok := x.(Stringer); ok {
        return str.String()
    }
}
```
如果x满足这两个接口类型中的一个，具体满足的接口决定对值的格式化方式，如果不满足，默认
case会统一地使用反射来处理所有的其它类型。
