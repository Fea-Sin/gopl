
### Channels

如果说goroutine是Go语言的并发体的话，那么channels则是它们之间的通信机制。一个channel是一个通信
机制，它可以让一个goroutine通过它给另一个goroutine发送值信息，每个channel都有一个特殊的类型，
也就是channel可发送数据的类型，一个可以发送int类型数据的channel一般写为chan int。使用内置的
make函数，可以创建一个channel
```
ch := make(chan int)
```
和map类似，channel也对应一个make创建的底层数据结构的引用，当我们复制一个channel或用于函数参数
传递时，我们只是拷贝一个channel引用，因此调用者和被调用者将引用同一个channel对象，和其他引用类型
一样，channel的零值也是nil。

两个相同类型的channel可以使用==运算符比较，如果两个channel引用的是相同的对象，那么比较的结果为
真，一个channel也可以和nil进行比较。

一个channel有发送和接收两个主要操作，都是通信行为，一个发送语句将一个值从一个goroutine通过
channel发送到另一个执行接收操作的goroutine，发送和接收两个操作都使用`<-`运算符。在发送语句
中`<-`运算符分割channel和要发送的值，在接收语句中，`<-`运算符写在channel对象前。一个不使用
接收结果的操作也是合法的。
```
ch <- x // send statement
x = <- ch // receive statement

<- ch // receive statement; result is discarded
```
channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。
对一个已经被close过的channel进行接收操作依然可以接收到之前已经成功发送的数据，如果channel中已经
没有数据的话将产生一个零值的数据。使用内置的close函数就可以关闭一个channel
```
close(ch)
```
以最简单方式调用make函数创建的是一个无缓存的channel，但是我们可以指定第二个整型参数，对应channel
的容量，如果channel的容量大于零，那么该channel就是带缓存的channel。
```
ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

### 不带缓存的 Channels

一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels
上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句，反之，
如果接收操作发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行
发送操作。

基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作，无缓存Channels有时候也被
称为同步Channels。当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine
之前（happens before 这是Go语言并发内存模型的一个关键术语）。

在讨论并发编程时，当我们说x事件在y事件之前发生（happens before），我们并不是说x事件在时间上比
y时间更早，我们要表达的意思是要保证在此之前的事件都已经完成，例如在此之前的更新某些变量的操作已经
完成，你可以放心依赖这些已经完成的事件了。

当我们说x事件既不是在y事件之前发生也不是在y事件之后发生，我们就说x事件和y事件是并发的，这并不
意味着x事件和y事件就一定是同时发生的，我们只是不能确定这两个事件发生的先后顺序。当两个goroutine
并发访问了相同的变量时，我们有必要保证某些事件的执行顺序，以避免出现某些并发问题。

基于channels发送消息有两个重要方面，首先每个消息都有一个值，但是有时候的事实和发生的时刻同样重要，
当我们希望强调通讯发生的时刻时，我们将它称为消息事件，有些消息事件并不携带额外的信息，它仅仅是用作
两个goroutine之间的同步，这时候我们可以用`struct{}`空结构体作为channels元素的类型，虽然也
可以使用bool和int类型实现同样的功能，`done <- 1`语句也比`struct{}{}`更短。

### 串联的Channels（Pipeline）

Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入，这种
串联的Channels就是所谓的管道（pipeline）。

[实例](ga.go)

其实你并不需要关闭每一个channel，只有当需要告诉接收者goroutine，所有的数据已经全部发送时才需要
关闭channel，不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。不要
将关闭一个打开文件的操作和关闭一个channel操作混淆，对于每个打开的文件，都需要在不使用的时候调用
对应的Close方法关闭文件。

### 单方向的Channel

Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel，类型`chan<- int`
表示一个只发送int的channel，只能发送不能接收，相反，类型`<-chan int`表示一个只接收int的channel，
只能接收不能发送。（箭头<-和关键字chan的相对位置表明了channel的方向）这种限制将在编译期检测。

关闭操作只用于不再向channel发送新的数据时使用，所以只有在发送者所在的goroutine才会调用close函数，
因此对一个只接收的channel调用close将是一个编译错误。

[实例](gb.go)
调用counter(naturals)时，naturals的类型将隐式地从`chan int`转换成`chan<- int`，调用
printer(squares)也会导致相似的隐式转换，这一次是转换为`<-chan int`类型（只接收的channel），
任何双向channel向单项channel变量的赋值操作都将导致该隐式转换，这里并没有反向转换的语法，也就
是不能将一个类似`chan<- int`类型的单项型的channel转换为`chan int`类型的双向的channel。

### 带缓存的Channels

向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素，如果
内部缓存队列是满的，那么发送操作将阻塞直到另一个goroutine执行接收操作而释放了新的队列空间，相反，
如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

内置的cap函数可以获取channel内部缓存的容量，内置len函数将返回channel内部缓存队列中有效元素的
个数。

Go语言新手有时候会将一个带缓存的channel当作同一个goroutine中的队列使用，虽然语法看似简单，但
实际上这是一个错误，Channel和goroutine的调度器机制是紧密相连的，如果没有其他goroutine从channel
接收，发送这或许是整个程序将会面临永远阻塞的风险。如果你只是需要一个简单的队列，使用slice就可以了。

多个goroutines并发地向同一个channel发送数据，或从同一个channel接收数据都是常见的用法。如果
多个goroutines向无缓存的channel发送数据，那么慢的goroutine将会因为没有人接收而被永远卡住，
这种情况称为goroutines泄露，这将是一个BUG和垃圾变量不同，泄露的goroutines并不会被自动回收，因此
确保每个不再需要的goroutine能正常退出是重要的。

关于无缓存的或带缓存channels之间的选择，或者是带缓存channels的容量大小的选择，都可能影响程序的
正确性。无缓存channel更强的保证了每个发送操作与相应的同步接收操作，但是对于缓存channel，这些操作
是解耦的。

















