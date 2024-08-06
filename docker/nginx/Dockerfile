# Use the official Nginx image as base
FROM nginx:1.27.0-alpine

RUN adduser -S www-data -G www-data -u 82

USER www-data

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]
