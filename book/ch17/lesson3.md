

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















