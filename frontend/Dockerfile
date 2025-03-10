FROM node:18-alpine

WORKDIR /app

# Copy package files and install dependencies
COPY package.json package-lock.json* ./
RUN npm ci

# Copy all frontend files
COPY . .

# Build the Next.js application
RUN npm run build

# Start the application
CMD ["npm", "start"]
