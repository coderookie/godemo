# windows下golang使用protobuf的示例
1、下载proto.exe, xxx.proto文件的解释器. 地址: https://github.com/protocolbuffers/protobuf/releases/tag/v21.1
2、下载protoc-go-gen, 将xxx.proto文件转换成go代码的插件. go install github.com/golang/protobuf/protoc-gen-go, 下载到了C:\Users\zhanglei16312\go\bin中(下载到的实际上是GOPATH下)
3、编写proto文件
4、通过protoc和proto的go插件生成xxx.pb.go. cd /d/program/godemo/conf/proto/; /d/Program\ Files/protoc/bin/protoc.exe --plugin=protoc-gen-go=/c/Users/zhanglei16312/go/bin/protoc-gen-go.exe --go_out=./ user.proto
5、mv godemo/models/protobuf_package/user ../../models/protobuf_package/
6、编写go代码. see models/protobuf/user.go