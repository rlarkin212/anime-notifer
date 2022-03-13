#!/bin/bash

GOOS=linux go build -o bin/anime-notifer
zip -j anime-notifer.zip bin/anime-notifer config.yaml