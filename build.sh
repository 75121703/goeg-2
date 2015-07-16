#!/bin/bash

root=$GOPATH/src/github.com/jvillasante/goeg
project=github.com/jvillasante/goeg

echo Installing Packages...
go install $project/lib/numbers
go install $project/lib/stringutils
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

# Chapter02
go build -o $root/bin/bitflag        $project/ch02/bitflag
go build -o $root/bin/pi_by_digits   $project/ch02/pi_by_digits
go build -o $root/bin/statistics     $project/ch02/statistics
go build -o $root/bin/statistics_ans $project/ch02.exercises/statistics_ans

# Chapter03
go build -o $root/bin/fmtexamples    $project/ch03/fmtexamples
go build -o $root/bin/stringexamples $project/ch03/stringexamples
go build -o $root/bin/m3u2pls        $project/ch03/m3u2pls
go build -o $root/bin/playlist       $project/ch03.exercises/playlist
go build -o $root/bin/soundex        $project/ch03.exercises/soundex
echo Done.
