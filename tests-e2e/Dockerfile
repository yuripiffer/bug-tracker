FROM mcr.microsoft.com/playwright:v1.50.0-jammy

WORKDIR /tests

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy test files
COPY . .

# Run tests
CMD ["npx", "playwright", "test"] 