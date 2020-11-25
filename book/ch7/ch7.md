
### 函数声明

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体

```
func name(parameter-list) (result-list) {
    body
}
```
形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由参数调用者提供，
返回值列表描述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值，返回
值列表的括号是可以省略的。如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会
返回任何值。
```
func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}

fmt.Println(hypot(3, 4))
```
x和y是形参名3和4是调用的传入实参，函数返回一个float64类型的值，返回值也可以像形式参数
一样被命名，在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其
初始化为零值。如果一个函数在声明时，包含返回值列表，该函数必须以return语句结尾。

函数的类型被称为函数的标识符，如果两个函数形式参数列表和返回值列表中的变量类型一一对应，
那么这两个函数有相同的标识符，形参和返回值的变量名不影响函数标识别

每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值），在函数调用时，Go语言没有
默认参数，也没有任何可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言
没有意义。

在函数体中，函数的形参作为局部变量，函数调用时以传入的值初始化，函数的形参和有名返回值
作为函数最外层的局部变量，被存储在相同的词法块中。

函数参数是通过传入值的拷贝，对参数进行修改不会影响传入值，但是，如果传入参数是引用类型
如指针、slice（切片）、map、function、channel等类型，传入值可能会被函数的间接修改。

### 递归

函数可以是递归的，这意味着函数可以直接或间接的调用自身，对许多问题而言，递归是一种强
有力的技术，例如处理递归的数据结构。

大部分编程语言使用的是固定大小的函数调用栈，常见的大小从64KB到2MB不等，固定大小栈会
限制递归的深度，当你用递归处理大量数据时，需要避免栈溢出，除此之外还会导致安全问题，
与之相反Go语言使用可变栈，栈的大小按需增加（初始时很小），这使得我们使用递归时不用
考虑溢出和安全问题。

### 多返回值

在Go中，一个函数可以返回多个值，许多标准库中的函数返回2个值，一个是期望得到的值，另一个是
函数出错时的错误信息。
```
func findLinks(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("getring %s: %s", url, resp.Status)
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    return visit(nil, doc), nil
}
```
我们必须确保resp.Body被关闭，释放网络资源，虽然Go的垃圾回收机制会回收不被使用的内存，
但是这不包括操作系统层面的资源，比如打开文件、网络连接，因此我们必须释放这些资源。
调用多返回值函数时，返回给调用者的是一组值，调用者必须显式的将这些值分配给变量
```
links, err := findLinks(url)
```
如果某个值不被使用，可以将其分配给blank identifier
```
links, _ := findLinks(url)
```

一个函数内部可以将另一个有多返回值的函数调用作为返回值
```
func findLinksLog(url string) ([]string, error) {
    log.Printf("findLinks %s", url)
    return findLinks(url)
}
```

如果一个函数的返回值有显式的变量名，那么该函数的return语句可以省略操作数，这称之为bare
return。
```
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
```

按照返回值列表的顺序，返回所有的返回值，每个一个return语句等价于
```
return words, images, err
```

当一个函数有多处return语句以及许多返回值时，bare return 可以减少代码的重复，但是
使得代码难被理解

### 错误

在Go中有一部分函数总是能成功运行，比如`strings.Contains`和`strconv.FormatBool`
函数，对各种可能的输入都做了良好的处理，使得运行时几乎不会失败，除非遇到灾难性的、不可
预料的情况，比如运行时的内存溢出，导致这种错误的原因很复杂，难以处理，从错误中恢复的
可能性也很低。

对大部分函数而言，永远无法确保能否成功运行，因为错误的原因超出了程序员的控制，举个例子，
任何进行I/O操作的函数都会面临出现错误的可能。当本该可信的操作出乎意料的失败后，我们
必须弄清楚导致失败的原因。

错误处理是软件包API和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为是几个预期
的结果之一。
在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常（exception）,
这使得Go有别于那些将函数运行失败看作是异常的语言。虽然Go有各种异常机制，但这些机制
仅被使用在处理那些未被预料到的错误。如果将某个应该在控制流程中处理的错误以异常的形式
抛出会混乱对错误的描述，这个错误会将堆栈跟踪信息返回给终端用户，这些信息复杂且无用，无法
帮助定位错误。

### 错误处理策略

最常见的方式是**传播错误**，函数中某个子程序的失败，会变成该函数的失败。
```
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nill {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```
fmt.Errorf函数使用fmt.Sprintf格式化错误信息并返回，我们使用该函数添加额外的前缀信息
到原始错误信息。

处理错误的第二种策略，如果错误的发生是偶然性的，或由于不可预知的问题导致的，一个明智的
选择是重新尝试失败的操作。
```
func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil // success
        }
        log.Printf("server not responding (%s);retrying...", err)
        time.Sleep(time.Second << uint(tries))
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
```

如果错误发生后，程序无法继续进行，我们可以采用第三种策略：输出错误并结束程序。需要注意
的是这种策略只应在main中执行，对于库函数而言，应该向上传播错误。
```
if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "site is down %v\n", err)
    os.Exit(1)
}
```
调用log.Fatalf可以更简洁的代码达到与上文相同的效果，log中的所有函数，都会在错误信息
之前输出时间信息
```
if err := WaitForServer(url); err != nil {
    log.Fatalf("site is down: %v\n", err)
}
```

第四种策略：有时我们只需要输出错误信息就足够了，不需要中断程序的运行
```
if err : Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled", err)
}
```
或者标准错误流输出错误信息
```
if err := Ping(); err != nil {
    fmt.Fpringt(os.Stderr, "ping failed: %v; nerworking disabled\n", err)
}
```
log包中的所有函数会为没有换行符的字符串增加换行符。


### 文件结尾错误（EOF）

如何从标准输入中读取字符，以及判断文件结束
```
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed: %v", err)
    }
}
```
















