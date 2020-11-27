
### 可变参数

参数数量可变的函数称为可变参数函数，典型的例子就是fmt.Printf和类似
函数，Printf首先接收一个必备的参数，之后接收任意个数的后续参数。
在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略号
"..."，这表示该函数会接收任意数量的该类型参数。
```
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}
```
在函数体中vals被看作是类型为[]string的切片，sum可以接收任意数量的int型参数
在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的
一个切片作为参数传给被调用的函数，如果原始参数已经是切片类型，只需要在最后一个
参数后面加上省略符。
```
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...))
```

### Deferred 函数

有时函数在多执行路径下调用诸如resp.Body.Close()，以释放网络资源，但随着函数
变得复杂，需要处理的错误也变多，维护清理逻辑变得越来越困难，而Go语言独有的defer
机制可以让事情变得简单。

你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需的语法，当执行
到该条语句时，defer后的函数参数的表达式得到计算，但直到包含该defer语句的函数
执行完毕，defer后的函数才会被执行，不论包含defer语句的函数是通过return正常
结束，还是由于panic导致的异常结束，可以在一个函数中执行多条defer语句，他们的
执行顺序与声明顺序相反。

defer语句经常被用于处理成对的操作，如打开、关闭，连接、断开连接，加锁、释放锁等，
通过defer机制，不论函数逻辑多复杂都能保证在任何执行路径下释放资源，释放资源的
defer应该直接跟在请求资源的语句后。

```
func title(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    ct := resp.Header.Get("Content-Type")
    if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
        return fmt.Errorf("%s has type %s, not text/html", url, ct)
    }
    doc, err := html.Parse(resp.Body)
    if err != nil {
        return fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    visitNode := func (n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
            fmt.Println(n.FirstChild.Data)
        }
    }
    forEachNode(doc, visitNide, nil)
    return nil
}
```

在调试复杂程序时，defer机制也常被用于记录何时进入和退出函数。
[实例](gb.go)

通过这种方式，我们可以只通过一条语句控制函数的入口和所有的出口，甚至可以记录函数
运行的时间。需要注意一点：不要忘记defer语句后的圆括号，否则本该在进入时执行的操作
会在退出是执行，本该在退出时执行的永远不被执行。


























