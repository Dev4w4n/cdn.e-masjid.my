# Use an official Nginx slim image as base
FROM nginx:alpine

# Copy custom nginx configuration files
COPY nginx.conf /etc/nginx/nginx.conf
COPY index.html /usr/share/nginx/html/index.html

# Expose port 80
EXPOSE 80

# Set the default command to execute when the container starts
CMD ["nginx", "-g", "daemon off;"]
