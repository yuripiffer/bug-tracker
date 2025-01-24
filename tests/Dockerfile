FROM mcr.microsoft.com/playwright:v1.50.0-jammy

WORKDIR /app

# Install debugging tools
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    iputils-ping \
    net-tools \
    && rm -rf /var/lib/apt/lists/*

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Install Playwright browsers
RUN npx playwright install --with-deps chromium

# Copy test files
COPY . .

# Force headless mode in container
RUN sed -i 's/headless: process.env.CI ? true : false/headless: true/' playwright.config.ts

# Create a startup script
RUN echo '#!/bin/bash\nset -x\necho "Starting test sequence..."\nping -c 1 frontend\nping -c 1 backend\nwget -q --spider http://frontend:3000 || echo "Frontend not reachable"\nwget -q --spider http://backend:8080/api/health || echo "Backend not reachable"\nnpx playwright test' > startup.sh \
    && chmod +x startup.sh

# Wait for services and run tests
CMD ["./startup.sh"] 