

### sort.Interface接口

排序操作和字符串格式化一样是很多程序经常使用的操作，尽管一个最短的快排程序只要15行
就可以搞定，但是一个健壮的实现需要更多的代码，并且我们不希望每次我们需要的时候都
重写或者拷贝这些代码。

幸运的是，sort包内置提供了根据一些排序函数来对任何序列排序的功能。在很多语言中，排序
算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。相比之下，Go语言的sort.Sort
函数不会对具体的序列和它的元素做任何假设，它使用了一个接口类型sort.Interface来
指定通用的排序算法和可能被排序序列类型之间的约定，这个接口的实现由序列的具体表示和
它希望排序的元素决定。

排序算法需要知道三个东西，序列的长度、表示两个元素比较的结果、一种交换两个元素的方式。
```
package sort

type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例
应用sort.Sort函数。
```
type StringSlice []string
func (p StringSlice) Len() int {
    return len(p)
}
func (p StringSlice) Less(i, j int) bool {
    return p[i] < p[j]
}
func (p StringSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
```

[实例1](gb.go)


### http.Handler 接口

我们粗略的了解怎么用net/http去实现网络客户端和服务器，我们会对那些基于http.Handler
接口的服务器API进一步学习。
```
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```
ListenAndServe函数需要一个例如"localhost:8000"的服务器地址，和一个所有请求都可以分派的
Handler接口实例。它会一直运行，直到这个服务器因一个错误而失败。

[实例2](gc.go)

mac下查看所有进程
```
ps -ax
```

查看指定端口的进程
```
sudo lsof -i :8000 
// or
lsof -i :8000
```

根据进程名称查看
```
ps -ef | grep nginx
```

根据PID杀进程
```
sudo kill -9 93736
// or 
kill -9 93736
```















