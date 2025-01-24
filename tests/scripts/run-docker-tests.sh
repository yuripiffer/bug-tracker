#!/bin/bash

echo "Starting Docker test setup..."

# Ensure we're in the project root directory
cd "$(dirname "$0")/../.."  # Go up two levels: from scripts to tests to root

# Clean up any existing containers
docker-compose down

# Start the application containers
docker-compose up -d frontend backend

# Wait for services to be healthy
echo "Waiting for services to be healthy..."
attempt_counter=0
max_attempts=30

echo "Checking backend health at http://localhost:8080/api/health"
until $(curl --output /dev/null --silent --fail http://localhost:8080/api/health); do
    if [ ${attempt_counter} -eq ${max_attempts} ];then
        echo "Max attempts reached. Backend not healthy."
        echo "Curl response for health check:"
        curl -v http://localhost:8080/api/health
        docker-compose logs
        docker-compose down
        exit 1
    fi

    printf '.'
    attempt_counter=$(($attempt_counter+1))
    sleep 1
done

echo "Backend is up"

until $(curl --output /dev/null --silent --head --fail http://localhost:3000); do
    if [ ${attempt_counter} -eq ${max_attempts} ];then
        echo "Max attempts reached. Frontend not healthy."
        docker-compose logs
        docker-compose down
        exit 1
    fi

    printf '.'
    attempt_counter=$(($attempt_counter+1))
    sleep 1
done

echo "Frontend is up"

# Run the Playwright tests
cd tests  # Go to tests directory from project root
npm install
CI=1 PLAYWRIGHT_TEST_BASE_URL=http://localhost:3000 npx playwright test --config=playwright.config.ts integration.spec.ts

# Store the test result
test_exit_code=$?

# Clean up
cd ..     # Go up one level to project root
docker-compose down

# Exit with the test result
exit $test_exit_code 