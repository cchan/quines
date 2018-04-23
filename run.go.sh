#!/bin/sh
go run quine.go > quine2.go && diff quine.go quine2.go
