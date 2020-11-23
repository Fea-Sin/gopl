
### Map

哈希表是一种巧妙且实用的数据结构。它是一个无序的`key/value`对的集合，其中所有的`key`都是不同的，
然后通过给定的`key`可以在常数时间复杂度内检索、更新和删除对应的`value`

在Go语言中，一个map就是一个哈希表的引用，map类型可以写为map[K]V形式。map中所有的key都是相同的类型，
所有的value也有相同的类型，但是key和value之间可以是不同的数据类型。其中K对应的key必须是支持==比较
运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。对于V对应value数据类型则没有任何
的限制。

内置的`make`函数可以创建一个map

```
// mapping from strings to ints
ages := make(map[string]int)
```

我们也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value

```
ages := map[string]int {
    "alice": 31,
    "charlie": 34,
}
```
这相当于

```
ages := make(map[string]int)

ages["alice"] = 31
ages["charlie"] = 34
```

另一种创建空map的表达式是 `map[string]int {}`

Map中的元素y对通过ke应的下标语法访问[实例1](./ga.go)

实用内置的delete函数可以删除元素

```
// remove element ages["alice"]
delete(ages, "alice")
```
所有这些操作是安全的，即使这些元素不在map中也没有关系，如果一个查找失败将返回value类型的零
值。

而且``x += y` 和 `x++`等简短赋值语法也可以用在map上
```
ages["bob"] += 1

 or 

ages["bob"]++
```

map中的元素并不是一个变量，因此我们不能对map元素进行取址操作

```
_ = &ages["bob"]
```
禁止对map元素取址的原因是map可能随着元素的增长而重新分配内存空间，从而导致之前的
地址无效。

要想遍历map中的全部的key/value的话，可以实用range风格的for循环实现[实例2](./ga.go)

Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中，遍历的顺序
是随机的，每一次遍历的顺序都是不同的。如果要按顺序遍历key/value对，我们必须显式地对key进行排序
可以使用sort包地String函数对字符串slice进行排序

[实例3](./gb.go)

map类型的零值是nil，也就是没有引用任何哈希表。

map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，他们的行为和
一个空的map类似，但是向一个nil值的map存入元素将导致一个panic异常
```
ages["carol"] = 21 // panic：assignment to entry in nil map
```
在向map存数据前必须先创建map
```
age, ok := ages["bob"]
if !ok {
    /* "bob" is not a key in this map; age == 0 */
}
```
经常会将这两个结合起来使用
```
if age, ok := ages["bob"]; !ok { /* ... */ }
```
map下标语法将产生两个值，第二个是一个布尔值，用于报告元素是否是真的存在，布尔变量一般命名为ok,
特别适合马上用于if条件判断部分。

和slice一样，map之间也不能进行相等比较，唯一例外是和nil进行比较。要判断两个map是否包含相同的
key/value，我们必须通过一个循环实现[实例4](gc.go)


Go语言中并没有提供一个set类型，但是map中的key也是不相同的，可以用map实现类似的set的功能。
```
seen := make(map[string]int)
```

map 的value类型也可以是一个聚合类型类型，比如是一个map或slice，











