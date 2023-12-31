protoc --go_out=plugins=grpc:./auth --go_opt=module=gitee.com/yuekukeji/protocol/auth auth/auth.proto
protoc --go_out=plugins=grpc:./conversation --go_opt=module=gitee.com/yuekukeji/protocol/conversation conversation/conversation.proto
protoc --go_out=plugins=grpc:./errinfo --go_opt=module=gitee.com/yuekukeji/protocol/errinfo errinfo/errinfo.proto
protoc --go_out=plugins=grpc:./friend --go_opt=module=gitee.com/yuekukeji/protocol/friend friend/friend.proto
protoc --go_out=plugins=grpc:./group --go_opt=module=gitee.com/yuekukeji/protocol/group group/group.proto
protoc --go_out=plugins=grpc:./msg --go_opt=module=gitee.com/yuekukeji/protocol/msg msg/msg.proto
protoc --go_out=plugins=grpc:./msggateway --go_opt=module=gitee.com/yuekukeji/protocol/msggateway msggateway/msggateway.proto
protoc --go_out=plugins=grpc:./push --go_opt=module=gitee.com/yuekukeji/protocol/push push/push.proto
protoc --go_out=plugins=grpc:./sdkws --go_opt=module=gitee.com/yuekukeji/protocol/sdkws sdkws/sdkws.proto
protoc --go_out=plugins=grpc:./third --go_opt=module=gitee.com/yuekukeji/protocol/third third/third.proto
protoc --go_out=plugins=grpc:./user --go_opt=module=gitee.com/yuekukeji/protocol/user user/user.proto
protoc --go_out=plugins=grpc:./wrapperspb --go_opt=module=gitee.com/yuekukeji/protocol/wrapperspb wrapperspb/wrapperspb.proto
protoc --go_out=plugins=grpc:./statistics --go_opt=module=gitee.com/yuekukeji/protocol/statistics statistics/statistics.proto
