# conc 库的例子

官网链接：[conc: better structured concurrency for go](https://github.com/sourcegraph/conc)

## conc 可用来

- 代替使用 go 启动 goroutine，不用担心 goroutine 泄漏。[例子](./wg.go)

- 作为 goroutine 池。从 slice 或者 chanel 读取数据进行处理。[例子](./pool.go)

- 用 iter 包原地操作 slice 中的元素。也可以生成新 slice。[例子](./iter.go)

- 串联并发的数据流处理。[例子](./stream.go)

以上的例子中，如果 goroutine 中有 panic，都会被接货并在 Wait() 方法
