# Printing 1 to 10 in random order
-----------------------
This project is used to create number from 1 to t 10 in random order

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
wget 
```
* Linux
```
wget
```

How to use
-----------------------
```
./randomorder -h
Usage of ./random:
  -d    Enable debug log
  -f string
        Output to file
```

File description
-----------------------
* [README.md](./README.md) : Readme file
* [random.go](./randomnumber.go) : Main file for this project
* [random_test.go](./random_test.go) : Unit test for main function

