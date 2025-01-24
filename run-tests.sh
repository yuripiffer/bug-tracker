#!/bin/bash

echo "Starting server setup..."

# Make scripts executable
chmod +x cleanup.sh
chmod +x start-servers.sh

# Kill any existing processes
echo "Cleaning up existing processes..."
./cleanup.sh

# Start the servers
echo "Starting servers..."
./start-servers.sh &

# Wait for the servers to be ready
echo "Waiting for servers to be ready..."
npx wait-port 3000

echo "Running tests..."
npx playwright test

# Clean up
echo "Cleaning up..."
./cleanup.sh

echo "Tests completed." 