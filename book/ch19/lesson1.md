
## 包的匿名导入

如果只是导入一个包而并不使用导入的包将会导致一个编译错误，但是有时候我们只是想
利用导入包而产生的副作用，它会计算包级变量的初始化表达式和执行导入包的init初始
化函数。这时候我们需要抑制"unused import"编译错误，我们可以使用下划线`_`来
重命名导入的包，像往常一样，下划线为空白标识符，并不能被访问。

```
import _ "image/png" // register PNG decoder
```
这个被称为包的匿名导入，它通常是用来实现一个编译时机制，然后通过在main主程序入口
选择性地导入附加的包。

标准库的image图像包包含了一个`Decode`函数，用于从`io.Reader`接口读取数据并
解码图像，它调用底层注册的图像解码器来完成任务，然后返回image.Image类型的图像，
使用`image.Decode`很容易编写一个图像格式的转换工具，读取一种格式的图像，然后
编码为另一种图像格式。

```
package png // image/png

func Decode(r io.Reader) (image.Image, error)
func DecodeConfig(r io.Reader) (image.Config, error)

func init() {
    const pngHeader = "\x89PNG\r\n\x1a\n"
    image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}
```
主程序只需要匿名导入特定图像驱动包就可以用image.Decode解码对应格式的图像了。

数据库包database/sql也是采用了类似的技术，让用户可以根据自己需要选择导入必要
的数据库驱动
```
import (
    "database/sql"
    _ "github.com/lib/pq" // enable support for Postgres
    _ "github.com/go-sql-driver/mysql" // enable support for MySQL
)

db, err = sql.Open("postgres", dbname)  // OK
db, err = sql.Open("mysql", dbname)     // OK
```

### 包和命名

我们将提供一些关于Go语言独特的包和成员命名的约定。当创建一个包，一般要用短小
的包名，但也不能太短导致难以理解，标准库中最常见的包有bufio、bytes、flag、
fmt、http、io、json、os、sort、sync、time包。要尽量避免包名使用可能
被经常用于局部变量的名字，这样可能导致用户重命名导入包，例如path包。

还有一些包只描述了单一的数据类型，例如html/template和math/rand等，只暴露
一个主要的数据结构和与它相关的方法，还有一个以New命名的函数用于创建实例。
```
package  rand // "math/rand"

type Rand struct{}

func New(source Source) *Rand
```

















