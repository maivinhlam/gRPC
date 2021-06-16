@echo off
if %1 == proto (
    echo Generate Protocol Buffer
    protoc --go_out=plugins=grpc:. .\proto\contact.proto
)

if %1 == run (
    if %2 == client (
        echo Run client..
        go run client/main.go
    )
    if %2 == server (
        echo Run server..
        go run server/main.go
    )
)
