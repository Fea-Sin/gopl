

## 竞争条件

在一个线性（就是说只有一个goroutine）的程序中，程序的执行顺序只是由程序的逻辑决定的。在有两个
或更多goroutine的程序中，每一个goroutine内的语句也是按照既定的顺序去执行的，但是一般情况下我们
没法知道分别位于两个goroutine的事件x和y的执行顺序，x是在y之前还是之后还是同时发生是没法判断的。
当我们没有办法自信地确认一个事件是在另一个事件的前面或者后面发生的话，就说明x和y这两个事件是
并发的。

一个函数在线性程序中可以正确地工作，如果在并发的情况下，这个函数依然可以正确地工作的话，那么我们
就说这个函数是并发安全的。并发安全的函数不需要额外的同步工作，我们可以把这个概念概括为一个特定类型的
一些方法和操作函数。对于某个类型来说，如果其所有可访问的方法和操作都是并发安全的话，那么该类型
便是并发安全的。

在一个程序中有非并发安全的类型的情况下，我们依然可以使这个程序并发安全，确实，并发安全的类型是例外，
只有文档中明确地说明了其是并发安全的情况下，你才可以并发地去访问它。我们会避免并发访问大多数的类型，
无论是将变量局限在单一的一个goroutine内，还是用互斥条件维持更高级别的不变性。

包级别的导出函数一般情况下都是并发安全的，由于package级的变量没法被限制在单一的goroutine，所以
修改这些变量必须使用互斥条件。

一个函数在并发调用时没法工作的原因太多了，比如死锁（deadlock）、活锁（livelock）和饿死
（resource starvation），我呢聚焦竞争条件上。

竞争条件指的是程序在多个goroutine交叉执行操作时，没有给出正确的结果，竞争条件是很恶劣的一种场景，
因为这种问题会一直潜伏在你的程序里，然后在非常少见的时候蹦出来，或许只是会在很大的负载时才会发生，
又或许是会在使用了某一个编译器、某一种平台或者某一种架构的时候才会出现，这些使得竞争条件带来的
问题非常难以复现而且难以分析诊断。

传统上经常用经济损失来为竞争条件做比喻，所以我们来看一个简单的银行账户程序
```
package bank

var balance int

func Deposit(amount int) {
    balance = balance + amount
}
func Balance() int {
    return balance
}

```
当我们并发地而不是顺序地调用这些函数的话，Balance就再也没办法保证结果正确了。

考虑以下两个goroutine
```
go func() {
    bank.Deposit(200)
    fmt.Println("=", bank.Balance())
}()

go bank.Deposit(100)
```

这个程序包含了一个特定的竞争条件，叫做数据竞争，无论任何时候，只要有两个goroutine并发访问同一
变量，且至少其中的一个是写操作的时候就会发生数据竞争。如果数据竞争的对象是一个比机器字（32位机器上
一个字是4个字节），事情就变得更麻烦了，比如interface、string、slice类型都是如此。
```
var x []int

go func() { x = make([]int, 10) }()

go func() { x = make([]int, 1000000) }()

x[999999] = 1
```
最后一个语句中的x的值可能是未定义的，其值是nil，或者也可能是一个长度为10的slice，也可能是一个长度是
1000000的slice。但是回忆以下slice的三个组层部分:指针（pointer）、长度（length）和容量
（capacity），如果指针是从第一个slice调用而来，而长度从第二make来，x就变成了一个混合体，一个长度
为1000000但实际上内部只有10个元素的slice，这种情况难以对值进行预测，而且debug也会变成噩梦。

并发程序的概念让我们知道并发并不是简单的语句交叉执行，数据竞争可能会有奇怪的结果。许多程序员还是偶尔
提出一些理由来允许数据竞争，比如互斥条件代价太高、这个逻辑只是用来做logging、我不介意丢失一些
消息等等。因为在他们的编译器或平台上很少遇到问题，可能会给了他们错误的信息。一个好的经验法则是根本
就没有所谓的良性数据竞争，所以我们一定要避免数据竞争。

#### 第一种方法是不要去写变量

```
var icons = map[string]image.Image {
    "spades.png": loadIcon("spades.png")
    "hearts.png": loadIcon("hearts.png")
    "diamonds.png": loadIcon("diamonds.png")
    "clubs.png": loadIcon("clubs.png")
}

// 并发安全
func Icon(name string) image.Image { return icons[name] }
```
上面的例子里icons变量在包初始化已经被赋值了，包的初始化是在程序main函数开始执行之前就完成了的，
只要初始化完成了，icons就再也不会被修改，数据结构如果从不被修改或是不变量则是并发安全的，无需
进行同步。

#### 第二种避免数据竞争的方法是，避免从多个goroutine访问变量
这也是前一章多数程序所采用的方法，例如聊天服务器中的broadcaster goroutine是唯一一个能访问
clients map的goroutine，这些变量都被限定在了一个单独的goroutine中。
由于其他的goroutine不能够直接访问变量，它们只能使用一个channel来发送请求给指定的goroutine
来查询跟新变量，这也是Go的口头禅：不要使用共享数据来通信；使用通信来共享数据。一个指定的变量通过
channel来请求的goroutine叫做这个变量的monitor（监控）goroutine，例如broadcaster goroutine
会监控clients map的全部访问。

```
package bank

var deposits = make(chan int)  // send amout to deposit
var balances = make(chan int)  // receive balance

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
    var balance int // balance is confined to teller goroutine
    for {
        select {
        case amount := <-deposits:
            balance += amount
        case balances <- balance:
        }
    }
}

func init() {
    go teller() // start the monitor goroutine
}

```

绑定依然是并发问题的一个解决方案，例如一条流水线的goroutine之间共享变量是很普遍的，在这两者间
会通过channel来传输信息，如果流水线的每一个阶段在将变量传送到下一个阶段后不再去访问它，那么对
这个变量的所有访问都是线性的。其效果是变量会被绑定到流水线的一个阶段，传送完之后被绑定到下一个，
一次类推，这种规则被成为串行绑定。

例子中，Cakes会被严格地顺序访问，先是baker goroutine，然后是icer goroutine
```
type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
    for {
        cake := new(Cake)
        cake.state = "cooked"
        caked <- cake
    }
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
    for cake := range cooked {
        cake.state = "iced"
        iced <- cake
    }
}
```

#### 第三种避免数据竞争的方法是允许很多goroutine去访问变量，但是在同一时刻
最多只有一个goroutine在访问，这种方式被称为互斥。


















