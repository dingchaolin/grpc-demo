# grpc-demo

## 环境搭建
####  grpc编译器
- https://github.com/google/protobuf/releases
- https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-osx-x86_64.zip
- 下载的如目录protoc-3-2  
- 其中protoc-3-2/bin/protoc 就是编译器
- 把protoc挪到src/compiler下, 方便编译
####  protoc-gen-go
- glide get github.com/golang/protobuf/protoc-gen-go
- 编译protoc-gen-go
- 像配置普通项目一样编译生成一个二进制的可执行文件 
- 文件名随意 比如 protoc_gen_go
- 把protoc_gen_go挪到src/compiler下, 方便编译

#### 编译 *.proto
-  ./compiler/protoc --grpc_out=. --plugin=protoc-gen-grpc=./compiler/protoc_gen_go ./proto/myproto.proto
- 生成 myproto.pb.go  
- --grpc_out=. 中的. 指 myproto.proto 所在的目录 为当前目录
- 至此 proto 编译生成go文件 成功

#### protobuf 语法简介
```
message Person {
    int32 id = 1;
    string name = 2;
    string email = 3;
    repeated PhoneNumber phones = 4;
}

1, 2, 3, 4 是 编号 不能重复
repeated 表示列表/数组

syntax = "proto3"  默认是 proto2

```

