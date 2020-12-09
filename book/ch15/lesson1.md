
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