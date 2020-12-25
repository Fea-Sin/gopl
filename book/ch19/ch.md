
## 包和工具

现在随便一个小程序的实现都可能包含超过10000个函数，然而作者一般只考虑其中很小的
一部分和做很少的设计，因为绝大部分代码都是由他人编写的，它们通过类似包或模块的
方式被重用。

Go语言有超过100个的标准包（可以使用`go list std | wc -l`命令查看标准包的具体
数目），标准库为大多数的程序提供了必要的基础构件，在Go的社区，有很多成熟的包被设计、
共享、重用和改进，目前互联网上已经发布了非常多的Go语言开源包，可以通过[godoc](https://godoc.org/)
检索。

### 包简介

任何包系统设计的目的都是为了简化大型程序的设计和维护工作，通过将一组相关的特性
放进一个独立的单元以便于理解和更新，在每个单元更新的同时保持和程序中其它单元的
相对独立性，这种模块化的特性允许每个包可以被其它的不同项目共享和重用，在项目
范围内、甚至全球范围统一的分发和复用。

每个包一般都定义了一个不同的名字空间用于它内部的每个标识符的访问，每个名字空间关联到
一个特定的包，让我们给类型、函数等选择简短明了的名字，这样就可以减少和其它部分
名字的冲突。

每个包还通过控制包内名字的可见性和是否导出来实现封装特性。通过限制包成员的可见性并
隐藏包API的具体实现，将允许包的维护者在不影响外部包用户的前提下调整包的内部实现。
通过限制包内变量的可见性，还可以强制用户通过某些特定函数来访问和更新内部变量，这样
可以保证内部变量的一致性和并发时的互斥约束。

当我们修改了一个源文件，我们必须重新编译源文件对应的包和所有依赖该包的其它包，即使
是从头构建，Go语言编译器的编译速度也明显快于其它编译语言。Go语言的闪电般的编译
速度主要得益于三个语言特性。

第一点，所有导入的包必须在每个文件的开头显示声明，这样的话编译器就没有必要读取和
分析整个源文件来判断包的依赖关系。

第二点，禁止包的环状依赖，因为没有循环依赖，包的依赖关系形成一个有向无环图，
每个包可以独立编译，而且很可能是被并发编译。

第三点，编译后包的目标文件不仅仅记录包本身的导出信息，目标文件同时还记录了包的
依赖关系，因此，在编译一个包的时候，编译器只需读取每个直接导入包的目标文件，
而不需要遍历所有依赖的文件（很多都是重复的间接依赖）。

### 导入路径

每个包是由一个全局唯一的字符串所标识的导入路径定位，出现在import语句中的导入
路径也是字符串。
```
import {
    "fmt"
    "math/rand"
    "encoding/json"

    "golang.org/x/net/html"
    "github.com/go-sql-driver/mysql"
}
```
Go语言的规范并没有指明包的导入路径字符串的具体含义，导入路径的具体含义是由
构建工具来解释的。

如果你计划分享或发布包，那么导入路径最好是全球唯一的，为了避免冲突，所有非标准
库包的路径建议以所在组织的互联网域名为前缀，而且这样也有利于包的检索。

### 包声明

在每个Go语言源文件的开头都必须有包声明语句，包声明语句的主要目的是确定当前包被
其它包导入时默认的标识符。

例如，math/rand包每个源文件的开头都包含`package rand`包声明语句，所以当你
导入这个包，你就可以用rand.Int，rand.Float64类似的方式访问包的成员。
`fmt.Println(rand.Int())`

通常来说，默认的包名就是包导入路径名的最后一段，因此即使两个包的的导入路径不同，
它们依然可能有一个相同的包名，例如math/rand和crypto/rand包的包名都是rand。
关于默认包名一般采用导入路径的最后一段的约定也有三种例外情况。

第一个例外，包对应一个可执行程序，也就是main包，名字main的包是给go build
构建命令一个信息，这个包编译完之后必须调用连接器生成一个可执行程序。

第二个例外，包所在的目录中可能有一些文件名是以`_test.go`为后缀的Go源文件（前面
必须有其它的字符，因为`_`或`.`开头的源文件会被构建工具忽略）。

第三个例外，一些依赖版本号的管理工具会在导入路径后追加版本信息，例如`gopkg.in/yaml.v2`，
这种情况下包的名字并不包含版本号后缀，而是yaml。

### 导入声明

可以在一个Go语言源文件包声明之后，其它非导入声明语句之前，包含零到多个导入包
声明语句。导入的包之间可以通过空行开分组，通常将来自不同组织的包独自分组，包的
导入顺序无关紧要。

如果我们想同时导入两个有着名字相同的包，例如math/rand和crypto/rand包，那么
导入声明必须至少为一个同名包指定一个新的包名以避免冲突，这叫做包导入的重命名。
```
import (
    "crypto/rand"
    mrand "math/rand"
)
```
导入包的重命名只影响当前的源文件，其它的源文件如果导入了相同的包，可以用导入原本
默认的名字或重命名为另一个完全不同的名字。

导入包重命名是一个有用的特性，它不仅仅只是为了解决名字冲突，如果导入的一个包名
很笨重，特别是一些自动生成的代码中，这时候用一个简短名称会更方便。选择另一个
包名可以帮助避免和本地普通本地变量名产生冲突，例如，如果文件中已经有了一个名
为path的变量，那么我们可以将path标准包重命名为pathpkg。

每个导入声明语句都明确指定了当前包和被导入包之间的依赖关系，如果遇到包循环导入
的情况，Go语言的构建工具将报告错误。































