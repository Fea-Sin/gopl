
## 工具

Go语言的工具箱集合了一系列功能的命令集，它可以看作是一个包管理器（类似于Linux中的apt和rpm工具），
用于包的查询、计算包的依赖关系、从远程版本控制系统下载它们等任务。它也是一个构建系统，计算文件的
依赖关系，然后调用编译器、汇编器和链接器构建程序，虽然它故意被设计成没有标准的make命令那么复杂，
它也是一个单元测试和基准测试的驱动程序。

你可以运行go或go help命令查看内置的帮助文档。为了达到零配置的设计目标，Go语言的工具箱很多地方
都依赖各种约束。根据给定的源文件的名称，Go语言的工具可以找到源文件对应的包，因为每个目录只包含了
单一的包，并且包的导入路径和工作区的目录结构是对应的，给定一个包的导入路径，Go语言的工具可以找到
与之对应的存储着实体文件的目录，它还可以根据导入路径找到存储代码的仓库的远程服务器URL。

### 工作区结构

对于大多数的Go语言用户，只需要配置一个名叫GOPATH的环境变量，用来指定当前工作目录即可。当需要
切换到不同工作区的时候，只需要更新GOPATH就可以了。
```
$ export GOPATH=$HOME/gobook
$ go get gopl.io
```

GOPATH对应的工作区目录有三个子目录，其中src子目录用于存储源代码，每个包被保存在与$GOPATH/src
的相对路径为包导入路径的子目录中，一个GOPATH工作区的src目录中可能有多个独立的版本控制系统。
其中pkg子目录用于保存编译后的包的目标文件，bin子目录用于保存编译后的可执行程序，例如helloworld
可执行程序。

第二个环境变量GOROOT用来指定Go的安装目录，还有它自带的标准库包的位置，GOROOT的目录结构和GOPATH
类似，因此存放fmt包的源代码对应目录应该为$GOROOT/src/fmt，用户一般不需要设置GOROOT，默认情况下
Go语言安装工具将其设置为安装的目录路径。

其中`go env`命令用于查看Go语言工具涉及的所有环境变量，包括未设置环境变量的默认值，GOOS环境变量
用于指定目标操作系统（如android、linux、darwin和windows），GOARCH环境变量用于指定处理器的类型，
例如amd64、386和arm等。

### 下载包

使用Go语言工具箱的Go命令，不仅可以根据包导入路径找到本地工作区的包，甚至可以从互联网上找到和更新包。
使用命令`go get`可以下载一个单一的包或者下载整个子目录里面的每个包。Go语言工具箱的go命令同时计算
并下载所依赖的每个包。

一旦`go get`命令下载了包，然后就是安装包或包对应的可执行程序。
```
$ go get  github.com/golang/lint/golint
```
``go get``命令支持当前流行的托管网站Github、Bitbucket和Launchpad，可以直接向它们的版本控制
系统请求代码。对于其它网站，你可能需要指定版本控制系统的具体路径协议。
`go get`命令获取的代码是真实的本地存储仓库，而不仅仅只是复制源文件，因此你依然可以使用版本管理
工具比较本地代码的变更或者切换到其它的版本。
需要注意的是导入路径含有的网站域名和本地Git仓库对应远程服务地址并不相同，这其实是Go语言工具的
一个特性，可以让包用一个自定义的导入路径，但是真实的代码却是由更通用的服务提供，例如go.googlesource.com
或github.com。

如果指定`-u`命令行参数，`go get`命令将确保所有的包和依赖的包的版本都是最新的，然后重新编译
和安装它们。如果不包含该标志参数的话，而且如果包已经在本地存在，那么代码将不会被自动更新。

`go get -u`命令只是简单地保证每个包是最新版本，如果是第一次下载包则是比较方便的，但是对于发布程序
则可能是不合适的，因为本地程序可能需要对依赖的包做精确的版本依赖管理。通常的解决方案是使用vendor
的目录用于存储依赖包的固定版本的源代码，对本地依赖的包的版本更新也是谨慎和持续可控的。

### 构建包

`go build`命令编译指定的包，如果包是一个库，则忽略输出结果，这可以用于检测包是可以正确编译的。
如果包的名字是main，`go build`将调用链接器在当前目录创建一个可执行程序，以路径的最后一段作为
可执行程序的名字。

由于每个目录只包含一个包，因此每个对应可执行程序或Unix术语中的命令的包，会要求放到一个独立的目录中。

每个包可以由它们的导入路径指定，就像前面看到的那样，或者用一个相对目录的路径名指定，相对路径
必须以`.`或`..`开头，如果没有指定参数，那么默认指定为当前目录对应的包。下面的命令由于构建同一个包
```
$ cd $GOPATH/src/gopl.io/ch1/helloworld
$ go build
```
或者
```
$ cd anywhere
$ go build gopl.io/ch1/helloworld
```
或者
```
$ cd $GOPATH
$ go build ./src/gopl.io/ch1/helloworld
```
但不能这样
```
$ cd $GOPATH
$ go build src/gopl.io/ch1/helloworld
Error: cannot find package "src/gopl.io/ch1/helloworld"
```

也可以指定包的源文件列表，这一般只用于构建一些小程序或做一些临时性的实验。如果是main包，将
会以第一个Go源文件的基础文件名作为最终的可执行程序的名字。
```
$ go build quoteargs.go
```
特别是对于这类一次性运行的程序，我们希望尽快构建并运行它可以使用`go run`命令。

默认情况下，`go build`命令构建指定的包和它依赖的包，然后丢弃除了最后的可行性文件之外所有中间
编译结果。依赖分析和编译过程虽然都很快的，但是随着项目增加到几十个包和成千上万行代码，依赖关系
分析和编译时间的消耗将变的可观。
`go install`命令和`go build`命令很相似，但是它会保存每个包的编译结果，而不是将它们都丢弃，
被编译的包被保存到$GOPATH/pkg目录下，目录路径和src目录路径对应，可执行程序被保存到$GOPATH/bin
目录。还有`go install`和`go build`命令都不会重新编译没有发生变化的包，这可以使后续构建更快捷。
为了方便编译依赖包，`go build -i`命令将安装每个目标所依赖的包。

因为编译对应不同的操作系统平台和CPU架构，`go install`命令会将编译结果安装到GOOS和GOARCH对应
的目录。在Mac系统，golang.org/x/net/html包将被安装到$GOPATH/pkg/darwin_amd64目录下的
golang.org/x/net/html.a文件。

更多细节，可以参考go/build包的构建约束部分
```
$ go doc go/build
```

### 包文档

Go语言的编码风格鼓励为每个包提供良好的文档，包中每个导出的成员和包声明前都应该包含目的和用法
说明注释。
Go语言中的文档注释一般是完整的句子，第一行通常是摘要说明，以被注释者的名字开头，注释中函数的
参数或其它的标识符并不需要额外的引号或者其它标记注明。
```
// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.

func Fprintf(w io.Writer, format string, a... interface{}) (int, error)
```
包注释可以出现在任何一个源文件中，如果包的注释内容比较长，一般会放到一个独立的源文件中，这个
专门用于保存包文档的源文件通常叫doc.go。

好的文档并不需要面面具到，文档本身应该是简洁但不可忽略的，并且文档也是需要像代码一样维护的。

`go doc`命令打印其后所指定的实体的声明与文档注释，该实体可能是一个包
```
$ go doc time
```
实体或者是某个具体的包成员
```
$ go doc time.Since
```
或者是一个方法
```
$ go doc time.Duration.Seconds
```
该命令并不需要输入完整的包导入路径或正确的大小写。

第二个工具，名字叫[godoc](https://godoc.org/)，它提供了包文档信息，包含了成千上万的开源包的
检索工具。

### 内部包

在Go语言中，包是最重要的封装机制，没有导出的标识符只在同一个包内部可以访问，而导出的标识符则是
**面向全宇宙是可见的**。

有时候一个中间状态可能也是有用的，标识符对于一小部分信任的包是可见的，但并不是对所有调用者都可见。
例如，当我们计划将一个大的包拆分为很多小的更容易维护的子包，但是我们并不想将内部的子包结构也完全
暴露出去，同时我们可能还希望在内部子包之间共享一些通用的处理包，或者我呢只是想实验一个新包的
还并不稳定的接口，暂时只暴露给一些受限的用户使用。

为了满足这些要求，Go语言的构建工具对包含internal名字的路径段的包导入路径做了特殊处理。
这种包加internal包，一个internal包只能被和internal目录有同一个父目录的包所导入。例如
```
net/http
net/http/internal/chunked
net/http/httputil
net/url
```
net/http/internal/chunked内部包只能被net/http/httputil和net/http包导入，但是不能被
net/url包导入，不过net/url包却可以导入net/http/httputil包。

### 查询包

`go list`命令可以查询可用包的信息，其最简单的形式，可以测试包是否在工作区并打印它的导入路径
```
$ go list github.com/go-sql-driver/mysql
github.com/go-sql-driver/mysql
```
`go list`命令的参数还可以用"..."表示匹配任意的包的导入路径，可以查看特定子目录下的所有包
```
$ go list gopl.io/ch3/...
```
或者是和某个主题相关的所有的包
```
$ go list ...xml...
```

`go list`命令还可以获取每个包完整的元信息，而不仅仅是导入路径，这写元信息可以以不同格式提供给用户，
其中`-json`命令行参数表示用JSON格式打印每个包的元信息

命令行参数`-f`则允许用户使用text/template的模版语言定义输出文本的格式。下面命令打印
compress子目录下所有包的导入包列表
```
$ go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...

compress/bzip2 -> bufio io sort
compress/flate -> bufio fmt io math math/bits sort strconv sync
compress/gzip -> bufio compress/flate encoding/binary errors fmt hash/crc32 io time
compress/lzw -> bufio errors fmt io
compress/zlib -> bufio compress/flate encoding/binary errors fmt hash hash/adler32 io
```

`go list`命令对于一次性交互查询或自动化构建或测试很有帮助，每个子命令的更多
信息，包括可设置的字段和意义，可以用`go help list`查看。












