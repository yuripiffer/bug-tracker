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
go run cmd/bugtracker/main.go & echo $! > ../tests/scripts/backend.pid
cd ..

# Start the frontend
echo "Starting frontend server..."
cd bugtracker-frontend
cd ../tests && npm install && cd ../bugtracker-frontend
npm run dev & echo $! > ../tests/scripts/frontend.pid
cd ..

# Wait for the frontend to be ready
echo "Waiting for frontend server to be ready..."
npx wait-port http://localhost:3000 -t 30000

echo "Server setup complete." 