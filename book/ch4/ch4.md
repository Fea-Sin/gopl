

### 结构体

结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体，每个值称为结构体的
成员。用结构体的经典案例是处理公司的员工信息，每个员工信息包含一个唯一的员工编号、员工
的名字、家庭住址、出生日期、工作岗位、薪资、上级领导等等。所有的这些信息都要绑定到一个
实体中，可以作为一个整体单元被复制，作为函数的参数或返回值，或者是被存储到数组中等等。

下面两个语句声明了一个叫Employee的命名的结构体类型，并且声明了一个Employee类型的变量
dilbert

```
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var dilbert Employee
```
通常一行对应一个结构体成员，成员的名字在前，类型在后，如果相邻成员类型相同的话可以被合并
都一行：
```
type Employee struct {
    ID            int
    Name, Address string
    DoB           time.Time
    Position      string
    Salary        int
    ManagerID     int
}
```
dilbert结构体变量的成员可以通过点操作符访问，比如dilbert.Name 和 dilbert.DoB，因为
dilbert是一个变量，它所有的成员也同样是变量，我们可以直接对每个成员赋值
```
dilbert.Salary -= 5000
```
或者是对成员取地址，然后通过指针访问
```
position := &dilbert.Position
*position = "Senior" + *position
```

点操作符也可以和指向结构体的指针一起工作
```
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
```
相当于下面的语句
```
(*employeeOfTheMonth).Position += " (proactive team player)"
```

结构体成员的输入顺序也有重要的意义，我们也可以将Position成员合并（因为是字符串类型），
或者是交换Name和Address出现的先后顺序，那样的话就是定义了不同的结构体类型。通常我们
是将相关的成员写到一起。

如果结构体成员名字是以大些字母开头的，那么该成员就是导出的，这是Go语言导出规则决定的，
一个结构体可能同时包含导出和未导出的成员。

结构体类型往往是冗长的，因为它的每个成员可能都会占一行，虽然我们每次都可以重写整个结构
体成员，但是重复会令人厌烦，因此，完整的结构体写法通常只在类型声明语句的地方出现。

一个命名为S的结构体类型将不能在包含S类型的成员，因为一个聚合的值不能包含它自身，该
限制同样适用于数组，但S类型的结构体可以包含`*s`指针类型的成员，这可以让我呢创建递归的
数据结构，比如链表和树结构等

结构体类型的零值是每个成员都是零值，通常会将零值作为最合理的默认值，例如bytes.Buffer
类型，结构体初始值就是一个可用的空缓存。

#### 结构体字面量

结构体值也可以用结构体字面量表示，结构体字面值可以指定每个成员的值。
```
type Point struct {
    X, Y int
}

p := Point{1, 2}
```
这里有两种形式的结构体字面量语法，上面的是第一种写法，要求以结构体成员定义的顺序为每个
结构体成员指定一个字面量。它要求写代码和读代码的人要记住结构体的每个成员的类型和顺序，
不过结构体成员有细微的调整就可能导致上述代码不能编译。因此，上述的语法一般只在定义结构体
的包内部使用，或者是在较小的结构体中使用，这些结构体的成员排列比较规则，比如image.Point{x, y}
或color.RGBA{red, green, blue, alpha}

其实更常用的是第二种写法，以成员名字和响应的值来初始化
```
anim := gif.GIF{LoopCount: nframes}
```
这种形式的结构体字面量值写法中，如果成员被忽略的话将默认用零值，因为提供了成员的名字，
所以成员出现的顺序并不重要。

结构体可以作为函数的参数和返回值，
```
func Scale(p Point, factor int) Point {
    return Point{p.X * factor, p.Y * factor}
}
```
如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回
```
func Bonus(e *Employee, percent int) int {
    return e.Salary * percent / 100
}
```

如果要在函数内部修改结构体成员的话，必须用指针传入，因为在Go语言中，所有的函数参数都是
值拷贝传入的，函数参数将不再是函数调用时的原始变量。

#### 结构体比较
入股结构体的全部成员都是可比较的，那么结构体也是可比较的，那样的话两个结构体可以使用==
或!=运算符进行比较，相等比较运算符==将比较两个结构体的每个成员。可比较的结构体类型和
其他可比较的类型一样，可以用于map的key类型。
```
type address struct {
    hostname sting
    port     int
}

hits := make(map[address]int)
hists[address{"golang.org", 443}]++
```

#### 结构体嵌入和匿名成员

Go语言有一个特性让我们只声明一个成员对应的数据类型而不指明成员的名字，这类成员就叫匿名
成员，匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。

[实例](./gb.go)

因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员


















