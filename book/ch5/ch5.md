
## JSON

JavaScript对象表示发（JSON）是一种用于发送和接收结构化信息的标准协议，类似的协议JSON
并不是唯一的标准协议，XML、ASN.1和Google的Protocol Buffers都是类似的协议，并且有
各自的特色，但是由于简洁性、可读性和流行程度等原因，JSON是应用最广泛的一个。

Go语言对这些标准格式的编码和解码都有良好的支持。

JSON是对JavaScript中各种类型的值，字符串、数字、布尔值和对象的Unicode文本编码。它可以
用有效可读的方式表示Go语言中基础数据类型、数组、slice、结构体和map等聚合数据类型。

一个JSON数组可以用于编码Go语言的数组和slice，一个JSON对象类型可以用于编码Go语言的map
类型和结构体

结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分割的key:"value"键值对
序列，成员Tag中json对应值的第一部分用于指定JSON对象的名字，比如将Go语言中的TotalCount
成员对应到JSON中的total_count，Color成员的Tag还带了一个额外的omitempty选项，表示
当Go语言结构体成员为空或零值时不生成该JSON对象（这里false是零值）

编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，
通过json.Unmarshal函数完成，将JSON格式的电影数据解码为一个结构体slice，结构体中只有
Title成员，通过定义合适的Go语言数据结构，我呢可以选择性地解码JSON中感兴趣的成员。




















