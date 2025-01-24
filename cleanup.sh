#!/bin/bash

echo "Starting cleanup..."

# Function to kill process and wait for it to end
kill_and_wait() {
  local pid_file=$1
  if [ -f "$pid_file" ]; then
    pid=$(cat "$pid_file")
    echo "Found PID file $pid_file with PID $pid"
    if kill "$pid" 2>/dev/null; then
      echo "Stopping process $pid"
      # Wait for the process to actually terminate
      while kill -0 "$pid" 2>/dev/null; do
        sleep 1
      done
    fi
  fi
}

echo "Checking for existing processes..."

# First try graceful shutdown
# Kill backend
kill_and_wait "backend.pid"

# Kill frontend
kill_and_wait "frontend.pid"

# Clean up PID files
rm -f backend.pid frontend.pid

echo "Checking for processes on ports..."

# Extra cleanup: kill any processes still using our ports
if lsof -ti:8080 > /dev/null; then
  echo "Found process on port 8080, killing..."
  # Kill any Go processes using this port
  ps aux | grep "main.go" | grep -v grep | awk '{print $2}' | xargs kill -9 2>/dev/null || true
  # Kill anything else using this port
  lsof -ti:8080 | xargs kill -9 2>/dev/null || true
fi

if lsof -ti:3000 > /dev/null; then
  echo "Found process on port 3000, killing..."
  lsof -ti:3000 | xargs kill -9 2>/dev/null || true
fi

# Double check ports are clear
sleep 2
if lsof -ti:8080 > /dev/null || lsof -ti:3000 > /dev/null; then
  echo "ERROR: Ports still in use after cleanup!"
  echo "Processes using port 8080:"
  lsof -i:8080
  echo "Processes using port 3000:"
  lsof -i:3000
  # Last resort: kill ALL Go processes
  echo "Attempting to kill all Go processes..."
  pkill -9 main 2>/dev/null || true
  exit 1
fi

echo "Cleanup complete" 