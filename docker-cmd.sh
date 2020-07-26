#!/bin/sh

go mod download && yarn install && go run src/backend/server.go
