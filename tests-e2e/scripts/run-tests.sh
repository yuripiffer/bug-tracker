#!/bin/bash

echo "Starting server setup..."
echo "Running integration tests...."
cd "$(dirname "$0")"  # Change to script directory

# Make scripts executable
chmod +x cleanup.sh
chmod +x start-servers.sh

# Kill any existing processes
echo "Cleaning up existing processes..."
./cleanup.sh

# Start the servers
echo "Starting servers..."
./start-servers.sh

# Wait for the servers to be ready
echo "Waiting for servers to be ready..."
echo "Waiting for backend..."
npx wait-port http://localhost:8080/api/health -t 30000
echo "Waiting for frontend..."
npx wait-port http://localhost:3000 -t 30000

echo "Running tests..."
cd ..  # Go to tests directory
echo "Running Playwright with args: $@"
if [[ "$*" =~ "--headless" ]]; then
  export CI=1
fi
npx playwright test --config playwright.config.ts integration.spec.ts "$@"

# Clean up
echo "Cleaning up..."
./scripts/cleanup.sh

echo "Tests completed." 