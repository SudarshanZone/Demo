# Microservices
@ICICI Securities || IRRA Project || Application Server

### gRpc and ProtoBuffer Installation
ProtoBuffer: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

gRPC: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Install proto from github https://github.com/protocolbuffers/protobuf/releases and set the path variable

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

go run greeter_server/main.go

### Important Commands

Github:
1. git init
2. git remote add origin https://github.com/krishnakashyap0704/microservices.git 
3. git checkout -b "New_Branch"
4. git checkout Branch_Name (if another branch present)
5. git pull origin Branch_Name
6. git add . 
7. git commit -m "Message"
8. git push -u origin Branch_Name

Extra: 
1. git status 
2. git log
3. git remote -v
4. git remote remove origin
5. git pull --rebase origin Branch_Name

---

### Initial Project Command

1. go mod init github.com/krishnakashyap0704/microservices
2. go get google.golang.org/grpc (used to import)
3. go mod tidy (used to change indirect to direct)
4. go run to/path/main.go

### Components

Squareoff:
1. protoc --go_out=. --go-grpc_out=. internal\proto\squareoff.proto
2. go run .\cmd\main.go

Tradelist:
1. protoc --go_out=. --go-grpc_out=. internal\proto\tradelist.proto
2. go run .\cmd\main.go

---

### Database Connection

1. go run to/path/db.go





MICROSERVICES
    |--Squareoff
       |--cmd
       |   |--main.go
       |--internal
       |   |--proto
       |   |   |--squareoff.proto
       |   |--db
       |    |   |--db.go
       |    |--generated
       |       |--squareoff_grpc.pb.go
       |       |--squareoff.pb.go  
       |--go.mod
     
MICROSERVICES
    |--Tradelist
       |--cmd
       |   |--main.go
       |--internal
       |   |--proto
       |   |   |--tradelist.proto
       |   |--db
       |    |   |--db.go
       |    |--generated
       |       |--tradelist_grpc.pb.go
       |       |--tradelist.pb.go  
       |--go.mod

MICROSERVICES
    userlogin
       |--cmd
       |   |--main.go
       |--internal
       |   |--proto
       |   |   |--userlogin.proto
       |   |--db
       |    |   |--db.go
       |    |--generated
       |       |--userlogin_grpc.pb.go
       |       |--userlogin.pb.go  
       |--go.mod


MICROSERVICES
    webserver
       |--cmd
       |   |--main.go
       |--internal
       |   |--grpcclient
       |   |   |--grpcclient.go
       |   |--handlers
       |      |--handlers.go
       |--go.mod

This is My Structer and I wanted to Import all the microservices like Squareoff, Tradelist, Userlogin, Webserver in one project. I wanted to import pbSquareOff "github.com/krishnakashyap0704/microservices/squareoff/internal/generated"
	pbTradeList "github.com/krishnakashyap0704/microservices/tradelist/internal/generated"
	pbUserLogin "github.com/krishnakashyap0704/microservices/userlogin/internal/generated"
in grpcclient.go and handlers.go. But I am getting error like "undefined: pbSquareOff" and "undefined: pbTradeList" and "undefined: pbUserLogin". I am not able to import these packages. This Issue Will Be Fixed Using go.mod file of web server I have to add all These Dependacies using Replace Command