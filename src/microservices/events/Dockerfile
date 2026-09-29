FROM node:22-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy the source code
COPY . .

# Use a smaller image for the final application
FROM node:22-alpine

WORKDIR /app

# Copy application from builder
COPY --from=builder /app ./

# Expose the application port
EXPOSE 8082

# Command to run the application
CMD ["npm", "start"]