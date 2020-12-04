
### 接口类型

接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型
的实例。

io.Writer类型是用的最广泛的接口之一，因为它提供了所有类型的写入bytes的抽象，
包括文件类型、内存缓冲区、网络连接、HTTP客户端、压缩工具、哈希等等，io包中定义
了很多其它有用的接口类型，Reader可以代表任意可以读取bytes的类型，Closer可以
是任意可以关闭的值，例如一个文件或是网络链接。

```
package io

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}
```

有些新的接口类型通过组合已有的接口来定义
```
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

上面用到的语法和结构体内嵌相似。我们可以用这种方式以一个简写命名一个接口，而
不用声明它所有的方法，这种方式称为接口内嵌，方法顺序的变化没有影响，唯一
重要的就是这个集合里面的方法。

### 实现接口的条件

一个类型如果**拥有一个接口所需要的所有方法**，那么这个类型就**实现了这个接口**。

**接口指定**的规则非常简单，表达一个类型属于某个接口只要这个类型实现这接口。

ReadWriter和ReadWriteCloser包含所有Writer的方法，所以任何实现了
ReadWriter和ReadWriteCloser的类型也必定实现了Writer接口

T类型的值不拥有所有`*T`指针的防范，举例说明
IntSet类型的String方法的接收者是一个指针类型，所以我们不能在一个不能寻址
的IntSet值上调用这个方法。
```
type IntSet struct {}
func (*IntSet) String() string

var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
```
然而，由于只有`*IntSet`类型有String方法，所以也只有`*IntSet`类型实现了
fmt.Stringer接口
```
var s IntSet
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s  // compile error: IntSet lacks String method
```

一个有更多方法的接口类型，比如io.ReadWriter和少一些方法的接口类型io.Reader进行
对比，更多方法的接口类型会告诉更多关于它的值持有的信息，并且实现它的类型要求更加
严格。
那么关于interface{}类型，它没有任何方法，空类型对实现它的类型没有要求，所以
我们可以将任意一个值赋给空接口类型。
```
var any interface{}
any = true
any = 12.34
any = "hello"
any = map[string]int{"one": 1}
any = new(bytes.Buffer)
```

每一个具体类型的组基于他们相同的行为可以表示成一个接口类型。不像基于类的语言，他们一个
类实现的接口集合需要进行显式定义。Go语言中我们可以在需要的时候定义一个新的抽象或者
特定特点的组，而不需要修改具体类型的定义，当具体的类型来自不同的作者时这种方式会特别
有用。






















