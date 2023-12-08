# Mac 

在mac系统下,我们首先需要安装protoc,下面我们使用brew安装

```go
brew install protobuf
```

接下来我们从github中,拉取最新的protoc-gen-go代码,用于生成protoc-gen-go 插件

```go
go install github.com/golang/protobuf/protoc-gen-go@latest
```

接下来,切换到gopath路径中,找到该源码,在protoc-gen-go目录下执行编译语句,执行完毕会获得protoc-gen-go文件用于生成 go文件

```go
go build -o protoc-gen-go main.go
```

最后我们尝试一下生成proto文件和go文件

```protobuf
syntax = "proto3"; //指定版本信息
option go_package="./pb;login_pb"; //指定生成的路径为pb,生成的package名为login_pb
message LoginReq{
  string account =1;
  string password =2;
}

message LoginRes{
  string token =1;
  uint32 error =2;
}

```

执行生成go文件语句,其中 

go_out:	生成的go文件输出路径

plugin:	上面提到的protoc-gen-go的路径

Login.proto:	我们刚才编辑的proto文件

```sh
protoc --go_out=./.. --plugin=./../../kit/proto/mac/protoc-gen-go login.proto
```

最终得到生成的login.pb.go文件

# Win

