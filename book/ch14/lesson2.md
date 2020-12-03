
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

一个类型如果拥有一个接口所需要的所有方法，那么这个类型就实现了这个接口。

接口指定的规则非常简单，表达一个类型属于某个接口只要这个类型实现这接口。

ReadWriter和ReadWriteCloser包含所有Writer的方法，所以任何实现了
ReadWriter和ReadWriteCloser的类型也必定实现了Writer接口























