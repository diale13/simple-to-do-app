# Use the official Node.js image as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the frontend application code into the container
COPY . .

# Install the necessary Node.js dependencies
RUN npm install

# Build the Svelte application
RUN npm run build

EXPOSE 5173

# Command to run the application
CMD ["npm", "start"]
