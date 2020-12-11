
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

### 类型分支

接口被以两种不同的方式使用。
第一种方式，以io.Reader、io.Writer、fmt.Stringer、sort.Interface、
http.Handler和error为典型，一个接口的方法表达了实现这个接口的具体类型的相似性，
但是隐藏了代码的细节和这些具体类型本身的操作，重点在方法生，而不是具体的类型上。

第二个方式，利用一个接口值可以持有各种具体类型的能力，将这个接口认为是这些类型的
联合，类型断言用来动态区别这些类型，使得对每一种情况都不一样。在这个方式中，重点
在于具体的类型满足这个接口，而不在于接口的方法，并且没有任何的信息隐藏，我们将
以这种方式使用的接口描述为可变识联合。

如果你熟悉面向对象编程，你可能会将这两种方式当作是子类型多态和非参数多态，但是你
不需要去记住这些术语。

```
func sqlQuote(x interface{}) string {
    switch x := x.(type) {
        case nil:
            return "NULL"
        case int, unit:
            return fmt.Sprintf("%d", x)
        case bool:
            if x {
                return "TRUE"
            }
            return "FALSE"
        case string:
            return sqlQuoteString(x)
        default:
            panic(fmt.Sprintf("unexpected type %T: %v", x, x))
    }
}
```

### 一些建议

当设计一个新的包时，新手Go程序员总是先创建一套接口，然后在定义一些满足它们的具体类型。
这种方式的结果就是有很多接口，它们中的每一个仅只有一个实现，这么做是不必要的的抽象，
它们也有一个运行时的损耗。你可以使用导出机制来限制一个类型的方法或一个结构体的字段
是否在包外可见。**接口只有当两个或两个以上的具体类型必须以相同的方式进行处理时才需要**。

当一个接口只被一个单一的具体类型实现时有一个例外，就是由于它的依赖，这个具体类型
不能和这个接口存在一个相同的包中，这种情况下，一个接口是解耦这两个包的一个好方式。

因为在Go语言中只有当两个或更多的类型实现一个接口时才使用接口，它们必定会从任意的
实现细节中抽象出来，结果就是有更少和更简单方法的更小的接口。当新的类型出现时，小的
接口更容易满足，对于接口设计的一个好的标准就是 ask only for what you need

我们完成了对方法和接口的学习过程，Go语言对面向对象风格的编程支持良好，但并不意味着
你只能使用这一风格，不是任何事务都需要被当做一个对象，独立的函数有它们自己的用处，
未封装的数据类型也是这样。观察一下，像input.Scan这样的方法被调用的并不多，与之
相反普遍调用的是函数fmt.Printf。








