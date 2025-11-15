#!/bin/bash
cd cmd/auth-service && go run . &
cd cmd/user-service && go run . &
cd cmd/book-service && go run . &
cd cmd/discovery-service && go run . &
cd cmd/chat-service && go run . &
cd cmd/media-service && go run . &
cd cmd/api-gateway && go run . &
