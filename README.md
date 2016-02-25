# Study-Go-Lang
This repo is collaborated with theoneisneo to study Go language.

Getting started with Go
https://github.com/golang/go/wiki#getting-started-with-go

Install & Hello World:

https://golang.org/doc/install

hello world demo(Windows:)

    1.Setup GOPATH as workspace for an environment variable in Advanced Sysytem Setting
    ex:%GOPATH%=D:/wiser/go
    2.Create folder src/github.com/wiser/hello under %GOPATH% , then put hello.go demo code in it.
    3.CMD mode:
    C:\Users\wiser>go install github.com/wiserkuo/hello
    C:\Users\wiser>%GOPATH%\bin\hello
    hello, world

Your first library demo

https://golang.org/doc/code.html#Library

    1.Add stringutil/reverse.go in src/github.com/wiserkuo/.
    2.in hello.go , add import("fmt" "github.com/wiserkuo/stringutil") ,stringutil.Reverse("!oG.olleH")
    3.CMD mode:
        D:\wiser\Go>go build github.com\wiserkuo\stringutil
        D:\wiser\Go>go install github.com\wiserkuo\hello
        D:\wiser\Go\bin>hello
        Hello.Go!
    

The Go Programming Language Specification:
https://golang.org/ref/spec

Go Standard Library
https://golang.org/pkg/
