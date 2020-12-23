

## sync.Once惰性初始化

如果初始化成本比较大的话，那么将初始化延迟到需要的时候再去做就是一个比较好的选择。如果
在程序启动的时候就去做这类初始化的话，会增加程序的启动时间，并且因为执行的时候可能也
并不需要这些变量，所以实际上有一些浪费。让我们看如下实例
```
var icons map[string]image.Image
```
这个版本的Icon用到了懒初始化（lazy initization）
```
func loadIcons() {
    icons = map[string]image.Image {
        "spades.png": loadIcon("spades.png"),
        "hearts.png": loadIcon("hearts.png"),
        "diamonds.png": loadIcon("diamonds.png"),
        "clubs.png": loadIcon("clubs.png"),
    }
}

func Icon(name string) image.Image {
    if icons == nil {
        loadIcons()
    }
    return icons[name]
}
```
如果变量只被一个单独的goroutine所访问的话，我们可以使用上面的这种模版，但这种模版在
Icon被并发调用时并不安全。Icon函数也是由多个步骤组成的，首先测试icons是否为空，然后
load这些icons，之后将icons更新为一个非空的值。直觉会告诉我们最差的情况是loadIcons
函数被多次反问会带来数据竞争。
不过这种直觉是错误的，我们希望你能从现在开始构建自己对并发的直觉，也就是说并发的直觉
总是不能被信任的！因为缺少显式的同步，编译器和CPU是可以随意地去更改访问内存的指令顺序，
以任意方式，只有保证每一个goroutine自己的执行顺序一致。其中一个种可能loadIcons的语句
重排是下面这样，它会在填写icons变量的值之前先用一个空map来初始化icons变量。
```
func loadIcons() {
    icons = make(map[string]image.Image)
    icons["spades.png"] = loadIcon("spades.png")
    icons["hearts.png"] = loadIcon("hearts.png")
    ...
}
```
因此，一个goroutine在检查icons是非空时，也不能假设这个变量的初始化流程已经走完了。
最简单且正确的保证所有goroutine能够观察到loadIcons效果的方式，是用一个mutex来
同步检查
```
var mu sync.Mutex
var icons map[string]image.Image

// 并发安全
func Icon(name string) image.Image {
    mu.Lock()
    defer mu.Unlock()
    if icons == nil {
        loadIcons()
    }
    return icons[name]
}
```
然而使用互斥访问icons的代价就是没有办法对该变量进行并发访问，即使变量已经被初始化
完毕，且再也不会进行变动，这里我们引入一个允许多读的锁
```
var mu sync.RWMutex
var icons map[string]image.Image

func Icon(name string) image.Image {
    mu.RLock()
    if icons != nil {
        icon := icons[name]
        mu.RUnlock()
        return icon
    }
    mu.RUnlock()

    // 没有初始化时
    mu.Lock()
    if icons == nil {
        loadIcons()
    }
    icon := icons[name]
    mu.Unlock()
    return icon
}
```
上面的模版使我们的程序能够更好的并发，但是有一点太复杂且容易出错，代码中有两个临界区。
sync包为我们提供了一个专门的方案来解决这种一次性初始化的问题：sync.Once。概念上讲，
一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成，互斥量
用来保护boolean变量和客户端数据结构。
```
var loadIconsOnce sync.Once
var icons map[string]image.Image

func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
```

在第一次调用时，boolean变量是false，Do会调用loadIcons并会将boolean变量
设置为true，并且是并发安全的，随后的调用什么都不会做，只是返回icon

### 竞争条件的检测

即使我们小心到不能再小心，但在并发程序中犯错误还是态容易了，幸运的是，Go的
runtime和工具链为我们装备了一个复杂但好用的动态分许工具，竞争检查器（the 
race detector）。
只要在go build，go run或者go test命令后面加上 -race 的flag，就会使编译器
创建一个你的应用的"修改"版或者一个附带了能够记录所有运行期对共享变量访问工具的
test，并且会记录下每一个读或者写共享变量的goroutine的身份信息，另外，修改
版的程序会记录下所有的同步事件，比如go语句，channel操作，以及对(*sync.Mutex).Lock，
(*sync.WaitGroup).Wait等等的调用。

修改版需要额外的记录，因此构建时加了竞争检测的程序跑起来会慢一些，且需要更大的
内存，即使是这样，这些代价对于很多成产环境的工作来说还是可以接受的，对于一些偶发
的竞争条件来说，让竞争检查器来干活可以节省无数日夜的debugging。
















