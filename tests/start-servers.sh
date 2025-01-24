#!/bin/bash
# Kill any existing processes
./cleanup.sh

# Clean the database
rm -f ../bugtracker-backend/bugs.db

# Start the backend
cd ../bugtracker-backend
go run cmd/bugtracker/main.go & echo $! > ../tests/backend.pid
cd ..

# Start the frontend
cd bugtracker-frontend
npm install
npm run dev & echo $! > ../tests/frontend.pid
cd .. 