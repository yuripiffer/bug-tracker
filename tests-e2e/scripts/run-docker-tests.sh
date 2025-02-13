#!/bin/bash

echo "Starting Docker test setup..."

# Ensure we're in the project root directory
cd "$(dirname "$0")/../.."  # Go up two levels: from scripts to tests to root

# Clean up any existing containers
docker-compose down

# Build and run the tests
docker-compose up --build --abort-on-container-exit tests

# Capture the exit code
test_exit_code=$?

# Clean up
docker-compose down

# Exit with the test result
exit $test_exit_code 