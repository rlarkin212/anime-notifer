#!/bin/bash
GOARCH=amd64 GOOS=linux go build -o bin/bootstrap main.go
cp config.yaml bin/
# zip -j anime-notifer.zip bin/anime-notifer config.yaml