#!/bin/bash
cd bugtracker-backend
go run cmd/bugtracker/main.go & echo $! > ../backend.pid 