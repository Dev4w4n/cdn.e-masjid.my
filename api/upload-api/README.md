1. Create the image
```
docker build --build-arg GO_ENV=dev -t image-server-image .
```

2. Run the container
```
docker run -d \
  -p 8080:8080 \
  --name image-server \
  image-server-image
```