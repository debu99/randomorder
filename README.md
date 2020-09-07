# Printing 1 to 10 in random order
-----------------------
This project is used to create number from 1 to t 10 in random order


Test
-----------------------
```
go test -v
```

Build
-----------------------
* To run under MacOs
```
go build 
```
* To run under Linux
```
CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"
```

Download
-----------------------
* MacOs
```
wget https://github.com/debu99/randomorder/releases/download/0.0.1/randomorder_darwin.zip
unzip randomorder_darwin.zip
```
* Linux
```
wget https://github.com/debu99/randomorder/releases/download/0.0.1/randomorder_linux.zip
unzip randomorder_linux.zip
```

How to use
-----------------------
```
./randomorder -h
Usage of ./randomorder:
  -d    Enable debug log
  -f string
        Output to file
```

File description
-----------------------
* [README.md](./README.md) : Readme file
* [randomorder.go](./randomorder.go) : Main file for this project
* [randomorder_test.go](./randomorder_test.go) : Unit test for main function

