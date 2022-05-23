#!/usr/bin/env bash

PROBLEM_NAME=$1
echo -e "package leetcode \n" >> leetcode/$1.go

echo -e "Now add the function stub, then run:\n gotests -all -w parallel leetcode/$1.go"

# gow test $1