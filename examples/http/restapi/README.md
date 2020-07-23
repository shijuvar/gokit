#### Compile the Dockerfile
docker build . -t restapi-server

#### Run the REST API Server
docker run -p 3000:8080 restapi-server