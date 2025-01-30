#!/bin/bash

echo "Starting performance test setup..."
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

echo "Running k6 performance tests..."
cd ..  # Go to tests-perf directory
k6 run --out json=test-results/raw-perf-test-results.json script.js

# Store the exit code
TEST_EXIT_CODE=$?

# Clean up
echo "Cleaning up..."
./scripts/cleanup.sh

# Exit with the stored test result
exit $TEST_EXIT_CODE 