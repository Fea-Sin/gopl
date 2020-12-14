
### Goroutines

在Go语言中，每一个并发的执行单元叫作一个goroutine，设想这里的一个程序有两个函数，一个函数做计算，
另一个输出结果，假设两个函数没有相互之间的调用关系，一个线性的程序会先调用其中的一个函数，然后再
调用另一个，如果程序中包含多个goroutine，对两个函数的调用则可能发生在同一时刻。
如果你使用过操作系统或者其它语言提供的线程，那么你可以简单地把goroutine类比作一个线程，这样你就
可以写出一些正确的程序了。

当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它 main goroutine，新的goroutine
会用go语句来创建，在语法上go语句是一个普通的函数或方法调用前加上关键字go，go语句会使其语句中的
函数在一个新创建的goroutine中运行，而go语句本省会迅速地完成。
```
f() // call f() wait for it to return

go f() // create a new goroutine that calls f() don't wait
```

主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或直接终止
程序之外，没有其他的编程方法能够让一个goroutine来打断另一个的执行，但是之后可以看到
一种方式来实现这个目的，通过goroutine之间的通信来让一个goroutine请求其他的
goroutine，并让被请求的goroutine执行结束执行。

网络编程是并发大显身手的一个领域，由于服务器是最典型的需要同时处理很多连接的程序，这些连接一般来自
于彼此独立的客户端。go语言的net包提供了编写一个网络客户端或者服务器程序的基本组件，无论两者间通信
是使用TCP、UDP或者Unix domain sockets。

killall命令，是一个Unix命令行工具，可以用给定的进程名来杀掉所有名字匹配的进程
```
killall gb
```




















