#!/bin/bash

echo "Starting server setup..."

# Kill any existing processes
echo "Cleaning up existing processes..."
./cleanup.sh

# Clean the database
rm -f ../../bugtracker-backend/bugs.db

# Start the backend
echo "Starting backend server..."
cd ../../bugtracker-backend
go run cmd/bugtracker/main.go & echo $! > ../tests-api/scripts/backend.pid
cd ..

echo "Server setup complete." 