#!/bin/bash

set -x  # Print each command that is run
echo "Starting test sequence..."

echo "Checking network..."
echo "DNS resolution test:"
cat /etc/hosts

echo "Network connectivity test:"
ping -c 1 frontend
ping -c 1 backend

echo "Testing backend endpoints..."
echo "Trying to reach backend health endpoint..."
curl -v http://backend:8080/api/health || (echo "Backend health check failed (exit code: $?)" && exit 1)

echo "Testing frontend..."
curl -v http://frontend:3000 || (echo "Frontend check failed (exit code: $?)" && exit 1)

echo "Current directory and files:"
pwd
ls -la

echo "Both services ready, running tests..."
echo "Playwright config:"
cat playwright.config.ts
npx playwright test