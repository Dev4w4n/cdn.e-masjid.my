# Use an official Nginx slim image as base
FROM nginx:alpine

WORKDIR /usr/share/nginx/html/

# Copy custom nginx configuration files
COPY nginx.conf /etc/nginx/nginx.conf
COPY index.html /usr/share/nginx/html/index.html
COPY ./img-init/main.jpg /usr/share/nginx/html/images/demo/main.jpg
COPY ./img-init/main.jpg /usr/share/nginx/html/volume/demo/main.jpg

VOLUME /usr/share/nginx/html/

# Expose port 80
EXPOSE 80

# Set the default command to execute when the container starts
CMD ["nginx", "-g", "daemon off;"]
