# Stage 1: Build Stage
FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies separately for better caching
COPY package.json package-lock.json ./
RUN npm install --frozen-lockfile

# Copy application files
COPY . .

# Build the Next.js application
RUN npm run build

# Stage 2: Production Stage (Minimal Image)
FROM node:20-alpine AS runner

WORKDIR /app

# Install only production dependencies for smaller image
COPY package.json package-lock.json ./
RUN npm install --omit=dev --frozen-lockfile

# Copy only the necessary built files from builder stage
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

# Expose the port Next.js runs on
EXPOSE 3000

# Run as a non-root user for security
USER node

# Start the Next.js app
CMD ["npm", "run", "start"]
