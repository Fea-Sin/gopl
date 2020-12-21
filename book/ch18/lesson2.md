

## sync.Mutex 互斥锁

我们使用了一个buffered channel作为一个计数信号量，来保证最多只有20个goroutine会同时执行HTTP
请求。同理，我们可以用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个
共享变量。一个只能为1和0的信号量叫做二元信号量（binary semaphore）

```
var (
    sema = make(chan struct{}, 1)
    balance int
)
func Deposit(amount int) {
    sema <- struct{}{}
    balance = balance + amount
    <-sema
}

func Balance() int {
    sema <- struct{}{}
    b := balance
    <-sema
    return b
}
```
这种互斥很实用，而且被sync包里的Mutex类型直接支持，它的Lock方法能够获取到token(这里叫锁)，并且
Unlock方法会释放这个token。

```
import "sync"

var (
    mu sync.Mutex
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}
func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}
```
每次一个goroutine访问bank余额变量时，它都会调用mutex的Lock方法来获取一个互斥锁，如果其它的
goroutine已经获得了这个锁的话，这操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态，
mutex会保护共享变量。按照惯例，被mutex锁保护的变量是在mutex变量声明之后立刻声明的。

在Lock和Unlock之间的代码段中的内容goroutine可以随便读取或修改，这个代码段叫做临界区，锁的
持有者在其他goroutine获取该锁之前需要调用Unlock，goroutine在结束后释放锁是必要的，无论以哪条
路径通过函数都需要释放，即使是在错误路径中，也要记得释放。

上面的bank程序例证了一种通用的并发模式，一系列的导出函数封装一个或多个变量，那么访问这些变量唯一的
方式就是通过这些函数来做（或者方法，对于一个对象的变量来说）。每一个函数在一开始就获取互斥锁并在最后
释放锁，从而保证共享变量不会被并发访问。这种函数、互斥锁和变量的编排，是用一个代理人保证变量被
顺序访问。

在更复杂的临界区的应用中，很难去靠人判断对Lock和Unlock的调用是在所有路径中都能够严格配对了，
Go语言里的defer简直就是这种情况下的救星，我们用defer来调用Unlock。一个deferred Unlock即使
在临界区发生panic时依然会执行。
```
func Balance() int {
    mu.Lock()
    defer mu.Unlock
    return balance
}
```

考虑一下取款函数（Withdraw）函数，成功的时候，它会正确地减掉余额并返回true，但如果银行记录资金对
交易来说不足，那么取款就会恢复余额，并返回false。
```
func Withdraw(amount int) bool {
    Deposit(-amount)
    if Balance() < 0 {
        Deposit(amount)
        return false
    }
    return true
}
```
以上程序有一点副作用，当过多的取款操作同时执行时，balance可能会瞬间被减到0以下，这可能会引起
一个并发的取款被不合逻辑地拒绝。这里的问题是取款不是一个原子操作，它包含了三步骤，每一步都需要
获取并释放互斥锁，但任何一次锁都不会锁上整个取款流程。

理想情况下，取款应该只在整个操作中获得一次互斥锁，但下面这样是错误的
```
func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    Deposit(-amount)
    if Balance() < 0 {
        Deposit(amount)
        return false
    }
    return true
}
```
上面这个例子，Deposit会调用mu.Lock()第二次去获取互斥锁，但mutex已经锁上了，而无法被重入，go
里没有重入锁的概念，也就是说没法对一个已经锁上的mutex来再次上锁，这会导致程序死锁。

一个通用的解决方案是将一个函数分离为多个函数，比如我们把Deposit分离成两个，让取款（Withdraw）
操作在不影响其他API的情况下变成并发安全。
```
func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    if balance < 0 {
        deposit(amount)
        return false
    }
    return true
}

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    deposit(amount)
}
func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}

func deposit(amount int) {
    balance += amount
}
```














