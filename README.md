## 包装dao层中的error

### 问题
在数据库操作的时候，dao层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该如何实现

### 方案
由于dao层获取到的sql.ErrNoRows是一个sentinel error，为了方便调试，我们会希望知道dao层的某个方法具体执行了什么sql语句导致了ErrNoRows的产生，因此我们应该在dao层在ErrNoRows的基础上，wrap上具体执行的sql的信息，与额外的一些信息（比如在哪一层发生的错误）

### 实现
```
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── README.md
└── src
    ├── controller.go
    ├── dao.go
    └── db.go
```
* src/controller.go
  模拟路由调用dao
* src/dao.go
  模拟一个dao层, 使用errors.Wrap来包装db.go中抛出的错误
* src/db.go
  模拟一个数据库驱动对象，内含一个Query方法，随机抛出sql.ErrNoRows或者未知错误或者为nil的error

### 运行结果
* 当没有发生错误时：
```bash
# go run ./cmd/
User[8]'s name is [user]
```

* 当发生sql.ErrNoRows错误时：
```bash
# go run ./cmd/
2021/10/27 01:06:29 sql: no rows in result set
dao: SELECT id, name FROM user WHERE id = 4
go-wrap-error-demo/src.(*dao).FindUserByID
        /root/playground/go/go-wrap-error-demo/src/dao.go:27
go-wrap-error-demo/src.(*controller).FindUserNameByID
        /root/playground/go/go-wrap-error-demo/src/controller.go:18
main.main
        /root/playground/go/go-wrap-error-demo/cmd/main.go:16
runtime.main
        /usr/local/go/src/runtime/proc.go:225
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1371
controller: error occurs when find name of user who owns id [3]
go-wrap-error-demo/src.(*controller).FindUserNameByID
        /root/playground/go/go-wrap-error-demo/src/controller.go:20
main.main
        /root/playground/go/go-wrap-error-demo/cmd/main.go:16
runtime.main
        /usr/local/go/src/runtime/proc.go:225
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1371
```

* 当发生其它未知错误时：
```bash
# go run ./cmd/
2021/10/27 01:11:47 sql: unkonwn error
go-wrap-error-demo/src.init
        /root/playground/go/go-wrap-error-demo/src/db.go:11
runtime.doInit
        /usr/local/go/src/runtime/proc.go:6265
runtime.doInit
        /usr/local/go/src/runtime/proc.go:6242
runtime.main
        /usr/local/go/src/runtime/proc.go:208
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1371
```

可以看到，当发生sql.ErrNoRows时，我们可以从看到错误的堆栈信息，并能获取到是具体执行了什么sql语句导致该错误的发生的