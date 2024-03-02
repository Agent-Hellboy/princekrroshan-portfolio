# Use the official Node.js image as the base image
FROM node:14-alpine

# Create the directory for the application in the container
RUN mkdir -p /app/my-blog-app

# Set the working directory in the container
WORKDIR /app/my-blog-app

# Copy package.json and package-lock.json to the working directory
COPY my-blog-app/package*.json ./

# Install dependencies
RUN npm install --production

# Copy the rest of the application code to the working directory
COPY my-blog-app .

# Expose port 3000 for the React development server
EXPOSE 3000

# Command to start the React development server
CMD ["npm", "start"]
