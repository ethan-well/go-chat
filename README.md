# chat room
The chat room developed by Golang, It's a good project to learn Golang and Golang netWork developemt

# usage

## clone this project to your GOPATH

```
cd ${GOPATH}/src
git clone git@github.com:ItsWewin/go-chat.git
```

## build server and run it

```
go build -o server go-chat/server/main
./server
```

## build client and run it
```
go build -o client go-chat/client
./client
```

## Function of this demo
1. user register
2. user login
3. user send group message
4. show online user list
5. point-to-point communication
