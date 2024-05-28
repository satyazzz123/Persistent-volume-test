# Use the official Node.js base image with version 21.1.0
FROM node:21.1.0

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code to the working directory
COPY . .


# Expose the application port


# Command to run the application
CMD ["node", "server.js"]
