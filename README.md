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

#### 编译 *.proto 方式1 <未成功>
-  ./compiler/protoc --grpc_out=. --plugin=protoc-gen-grpc=./compiler/protoc_gen_go ./proto/myproto.proto
- 生成 myproto.pb.go  
- --grpc_out=. 中的. 指 myproto.proto 所在的目录 为当前目录
- 至此 proto 编译生成go文件 成功
- 此种方式没有生成带有服务器代码的go文件
- 未找到为什么没有生成的原因
#### 编译 方式2 
- 1.  glide get github.com/golang/protobuf/protoc-gen-go
- 2.  go install github.com/golang/protobuf/protoc-gen-go
- 3.  ./compiler/protoc  --go_out=plugins=grpc:.   ./proto/myproto.proto
- 用这种方式就能正常生成待遇服务器代码的go文件

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
- 要想编译生成的有服务 支持grpc调用 需要类似这么写
```
message AddressBook{
    repeated Person person = 1;
}

message AddressBook{
    repeated Person people = 1;
}

message AddPersonRequest{
    Person person = 1;
}

message AddPersonResponse{
    int32 id = 1;
}

service AddressBookStore{
    rpc AddPerson(AddPersonRequest) returns(AddPersonResponse);
}
```


