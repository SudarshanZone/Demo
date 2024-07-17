module github.com/krishnakashyap0704/microservices/webserver

go 1.22.5

replace github.com/krishnakashyap0704/microservices/squareoff => ../squareoff

replace github.com/krishnakashyap0704/microservices/tradelist => ../tradelist

replace github.com/krishnakashyap0704/microservices/userlogin => ../userlogin

require (
	github.com/krishnakashyap0704/microservices/squareoff v0.0.0
	github.com/krishnakashyap0704/microservices/tradelist v0.0.0
	github.com/krishnakashyap0704/microservices/userlogin v0.0.0
	google.golang.org/grpc v1.65.0
)

require (
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240709173604-40e1e62336c5 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
