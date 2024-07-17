
//
go mod init github.com/SudarshanZone/MICROSERVICES/webserver

go mod init github.com/krishnakashyap0704/microservices/webserver





-----------------------------
go mod edit "-replace=github.com/krishnakashyap0704/microservices/squareoff=../squareoff"


go mod edit "-replace=github.com/krishnakashyap0704/microservices/tradelist=../tradelist"
go mod edit "-replace=github.com/krishnakashyap0704/microservices/userlogin=../userlogin"

------------------------------------------------------------------------


cd path/to/microservices/webserver

# Ensure webserver module is initialized
go mod init github.com/krishnakashyap0704/microservices/webserver



# Replace directives for dependencies
go mod edit -replace=github.com/krishnakashyap0704/microservices/squareoff=../squareoff
go mod edit -replace=github.com/krishnakashyap0704/microservices/tradelist=../tradelist
go mod edit -replace=github.com/krishnakashyap0704/microservices/userlogin=../userlogin

# Add dependencies with placeholder versions
go get github.com/krishnakashyap0704/microservices/squareoff@v0.0.0
go get github.com/krishnakashyap0704/microservices/tradelist@v0.0.0
go get github.com/krishnakashyap0704/microservices/userlogin@v0.0.0

# Clean up module dependencies
go mod tidy

# Run your application
go run main.go


