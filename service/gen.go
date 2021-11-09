package service

// gateway api
//go:generate goctl api go -api gateway/cmd/api/gateway.api -dir gateway/cmd/api/

// user rpc
//go:generate goctl rpc proto -src user/cmd/rpc/user.proto -dir user/cmd/rpc/

//video rpc
//go:generate goctl rpc proto -src video/cmd/rpc/video.proto -dir video/cmd/rpc/

// video model
//go:generate goctl model mysql ddl -c -src video/model/video.sql -dir video/model/

// file_sys rpc
//go:generate goctl rpc proto -src file_sys/cmd/rpc/filesys.proto -dir file_sys/cmd/rpc/

// file_sys model
//go:generate goctl model mysql ddl -c -src file_sys/model/filesys.sql -dir file_sys/model/
