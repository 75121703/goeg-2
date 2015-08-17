#!/bin/bash

root=$GOPATH/src/github.com/jvillasante/goeg
project=github.com/jvillasante/goeg

echo Installing Packages...
go install $project/lib/numbers
go install $project/lib/stringutils
go install $project/lib/slices
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

# Chapter04
go build -o $root/bin/guess_separator $project/ch04/guess_separator
go build -o $root/bin/wordfrequency   $project/ch04/wordfrequency
go build -o $root/bin/ch04_ans        $project/ch04.exercises/answers

# Chapter05
go build -o $root/bin/counter               $project/ch05/counter
go build -o $root/bin/palindrome            $project/ch05/palindrome
go build -o $root/bin/generics              $project/ch05/generics
go build -o $root/bin/memoization           $project/ch05/memoization
go build -o $root/bin/indent_sort           $project/ch05/indent_sort
go build -o $root/bin/archive_file_list     $project/ch05/archive_file_list
go build -o $root/bin/archive_file_list_ans $project/ch05.exercises/archive_file_list_ans
go build -o $root/bin/palindrome_ans        $project/ch05.exercises/palindrome_ans
go build -o $root/bin/common_prefix         $project/ch05.exercises/common_prefix
echo Done.
