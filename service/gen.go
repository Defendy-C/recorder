package service

// gateway api
//go:generate goctl api go -api gateway/cmd/api/gateway.api -dir /gateway/cmd/api/

// user rpc
//go:generate goctl rpc proto -src user/cmd/rpc/user.proto -dir /user/cmd/rpc/
