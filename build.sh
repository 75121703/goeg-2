#!/bin/bash

root=$GOPATH/src/github.com/jvillasante/goeg
project=github.com/jvillasante/goeg

echo Installing Packages...
echo Done.

echo Building Programs...
rm -f $root/bin/*

# Chapter01
go build -o $root/bin/hello           $project/ch01/hello
go build -o $root/bin/bigdigits       $project/ch01/bigdigits
go build -o $root/bin/stacker         $project/ch01/stacker
go build -o $root/bin/americanise     $project/ch01/americanise
go build -o $root/bin/polar2cartesian $project/ch01/polar2cartesian
go build -o $root/bin/bigdigits_ans   $project/ch01.exercises/bigdigits
echo Done.
