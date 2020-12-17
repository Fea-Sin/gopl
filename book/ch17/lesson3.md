

### 基于select的多路复用

```
select {
    case <-ch1:
        // ...
    case x := <-ch2:
        // use x ...
    case ch3<- y
        // ...
    default:
        // ...
}
```
上面是select语句的一般形式，和switch语句稍微有点相似，也会有几个case和最后的default选择分支，
每一个case代表一个通信操作（在某个channel上进行发送或者接收），并且会包含一些语句组成的一个语句
块。一个接收表达式可能只包含接收表达式自身，就像上面的第一个case，或者包含在一个简短的变量声明中，
像第二个case里一样，第二种形式让你能够引用接收到的值。

select会等待case中有能够执行的case时才去执行，当条件满足时，select才会去通信并执行case之后的
语句，这时候其它通信是不会执行的，一个没有任何case的select语句写作select{}，会永远等待下去。

如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的
机会。

[实例](gc.go)

有时候我们希望能够从channel中发送或者接收值，并避免因为发送或这接收导致阻塞，尤其是当channel没有
准备好写或者读时，select语句就可以实现这样的功能，select会有一个default来设置当其它的操作都不能
够马上被处理时程序需要执行那些逻辑。
```
select {
    case <-abort:
        fmt.Println("Lunch aborted!")
        return
    default:
        // do nothing
}

```

select语句会在abort channel中有值时，从其中接收值，无值时什么都不做，这是
一个非阻塞的接收操作，反复地这样的操作叫做"轮询channel"。

channel的零值是nil，nil的channel有时候也是有一些用处的，因为对一个nil的
channel发送和接收操作会永远阻塞，在select语句中nil的case永远都不会被select
到。

### 并发的目录遍历

[实例1](gd.go)

[实例2](ge.go)


### 并发的退出

有时候我们需要通知goroutine停止它正在干的事情，比如一个正在执行计算的web服务，然而它的客户端
已经断开和服务端的连接。

Go语言并没有提供在一个goroutine终止另一个goroutine的方法，由于这样会导致goroutine之间的共享
变量落在未定义的状态上。

我们想要退出两个或者任意多个goroutine怎么办呢？

一种可能的手段是向abort的channel里发送和goroutine数目一样多的事件来退出，如果这些goroutine
中已经有一些自己退出了，那么会导致我们的channel里的事件数比goroutine还多，这样导致我们的发送
直接被阻塞。另一方面，如果这些goroutine又生成了其它的goroutine，我们的channel里的数目又
太少了，所以有些goroutine可能会无法接收到退出的消息。一般情况下我们是很难知道在某一个时刻具体
有多少个goroutine在运行着的，另外，当一个goroutine从abort channel中接收到一个值的时候，他
会消费掉这个值，这样其它的goroutine就没法看到这条信息，为了能够达到我们退出goroutine的目的，
我们需要更靠谱的策略，来通过一个channel把消息广播出去，这样goroutine们能够看到这条事件消息，
并且在事件完成之后，可以知道这件事情已经发生过了。

我们不向channel发送值，而是用关闭一个channel来进行广播。

我们创建一个退出channel，不需要向这个channel发送任何值，但其所在的闭包内要写明程序需要退出，
我们同时还定义了一个工具函数cancelled，这个函数在被调用的时候会轮询退出状态。
```
var done = make(chan struct{})

func cancelled() bool {
    select {
        case <-done:
            return true
        default:
            return false
    }
}
```

下面我们创建一个从标准输入流中读取内容的goroutine，这是一个典型的连接到终端的程序，每当有
输入被读到（比如用户按了回车键），这个goroutine就会把取消消息通过关闭done的channel广播出去。

```
go func() {
    os.Stdin.Read(make([]byte, 1))
    close(done)
}()
```

[实例3](gf.go)















