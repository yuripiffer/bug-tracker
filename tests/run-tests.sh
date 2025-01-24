#!/bin/bash
echo "Starting server setup..."
cd "$(dirname "$0")"  # Change to script directory
./cleanup.sh
./start-servers.sh
echo "Running tests..."
cd ..  # Go back to root for playwright
npx playwright test --config tests/playwright.config.ts
echo "Cleaning up..."
cd "$(dirname "$0")"  # Change back to script directory
./cleanup.sh
echo "Tests completed." 