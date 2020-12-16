

### 并发的循环

像这种子问题都是完全彼此独立的问题被叫做易并行问题（embarrassingly parallel），并且最能够
享受到并发带来的好处，能够随着并行的规模线性地扩展。

可以预测循环次数的情况
```
func makeThumbnails(filenames []string) (thumbfiles []string, err error) {
    type item struct {
        thumbfile string
        err       error
    }
    
    ch := make(chan item, len(filenames))
    for _, f := range filenames {
        go func(f string) {
            var it item
            it.thumbfile, it.err = thumbnail.ImageFile(f)
        }(f)
    }
    
    for range filenames {
        it := <-ch
        if it.err != nil {
            return nil, it.err
        }
        thumbfiles = append(thumbfiles, it.thumbfile)
    }

    return thumbfiles, nil
}
```

无法预测循环次数的情况，为了知道最后一个goroutine什么时候结束（最后一个结束并不一定最后一个开始），
我们需要一个递增的计数器，在每个一个goroutine启动时加一，在goroutine退出时减一，这需要一种特殊
的计数器，这个计数器要在多个goroutine操作时做到安全并且提供在其减为零之前一直等待的一种方法，这种
计数器被称为sync.WaitGroup
```
func makeThumbnails(filenames <-chan string) int64 {
    size := make(chan int64)
    var wg sync.WatiGroup
    for f := range filenames {
        wg.Add(1)
        
        go func(f string) {
            defer wg.Done()
            thumb, err := thumbnail.ImageFile(f)
            if err := nil {
                log.Println(err)
                return
            }
            info, _ := os.Stat(thumb)
            sizes <- info.Size()
        }(f)
    }

    go func() {
        wg.Wait()
        close(sizes)
    }()

    var total int64
    for size := range sizes {
        total += size
    }
    return total
}
```












