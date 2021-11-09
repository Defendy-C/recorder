start /b etcd
start /b go run service/user/cmd/rpc/user.go -f service/user/cmd/rpc/etc/user.yaml
start /b go run service/video/cmd/rpc/video.go -f service/video/cmd/rpc/etc/video.yaml
start /b go run service/file_sys/cmd/rpc/filesys.go -f service/file_sys/cmd/rpc/etc/filesys.yaml
start /b go run service/gateway/cmd/api/gateway.go -f service/gateway/cmd/api/etc/gateway.yaml