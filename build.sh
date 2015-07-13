#!/bin/bash

root=$GOPATH/src/books/programming-in-go
project=books/programming-in-go

echo Installing Packages...
echo Done.

echo Building Programs...
rm -f $root/bin/*

# Chapter01
go build -o $root/bin/hello     $project/ch01/hello
go build -o $root/bin/bigdigits $project/ch01/bigdigits
go build -o $root/bin/stacker   $project/ch01/stacker
echo Done.
