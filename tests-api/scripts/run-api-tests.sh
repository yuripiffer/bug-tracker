#!/bin/bash

echo "Starting API test setup..."
cd "$(dirname "$0")"  # Change to script directory

# Make scripts executable
chmod +x cleanup.sh
chmod +x start-server.sh

# Kill any existing processes
echo "Cleaning up existing processes..."
./cleanup.sh

# Start the backend server
echo "Starting backend server..."
./start-server.sh

# Wait for the backend to be ready
echo "Waiting for backend..."
npx wait-port http://localhost:8080/api/health -t 30000

echo "Running API tests..."
cd ..  # Go to tests-api directory
echo "Running Playwright with args: $@"
npx playwright test --config playwright.config.ts "$@"

# Clean up
echo "Cleaning up..."
./scripts/cleanup.sh

echo "API Tests completed." 