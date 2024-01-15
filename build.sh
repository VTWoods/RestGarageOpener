#!/bin/bash

env GOARCH=arm GOOS=linux GOARM=6 go build -o bin/main main.go 
