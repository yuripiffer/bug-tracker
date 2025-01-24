#!/bin/bash

echo "Starting server setup..."
cd "$(dirname "$0")"  # Change to script directory

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
cd ../..  # Go back to root for playwright
cd tests
npx playwright test --config playwright.config.ts

# Clean up
echo "Cleaning up..."
./scripts/cleanup.sh

echo "Tests completed." 