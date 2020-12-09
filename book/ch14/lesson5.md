
### error接口

从开始我们就已经创建和使用过神秘的预定义error类型，而且没有解释它究竟是什么，实际上它就是
interface类型，这个类型有一个返回错误信息的单一方法。
```
type error interface {
    Error() string
}
```
创建一个error最简单的方法就是调用errors.New函数，它会根据传入的错误信息返回一个新的
error，整个errors包仅寥寥数行。
```
package errors

type errorString struct {
    text string
}

func New(text string) error {
    return &errorString{text}
}

func (e *errorString) Error() string {
    return e.text
}
````

调用errors.New函数是非常稀少的，因为有一个方便的封装函数fmt.Errorf，它还会
处理字符串格式化
```
package fmt

import "errors"

func Errorf(format string, args... interface{}) error {
    return errors.New( Sprintf(format, args...) )
}
```

syscall包提供了Go语言底层系统调用API，在多个平台，它定义了一个实现error接口
的数字类型Errno，并且在Unix平台上，Errno的Error方法会从一个字符串表中查找
错误消息。
```
package syscall

type Errno unitptr // operating system error code

var errors = [...]string {
    1: "operation not permitted",
    2: "no such file or directory"
    3: "no such process"
}

func (e Errno) Error() string {
    if 0 <= int(e) && int(e) < len(errors) {
        return errors[e]
    }
    
    return fmt.Sprintf("errno %d", e)
}
```
















