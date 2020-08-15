# grpc-hello-world
A simple hello-world grpc application. 

### Running the application with docker:
- To build the docker image from the Dockerfile, run `docker build -t hello-world-image .` from the root directory of the application.
- To run a container from this image, run `docker run -d -p 8080:8080 hello-world-image`. 

Note:
- Argument `-d` runs the application is the background.
- Argument `-p` is used to map the host operating system port 8080 to the container port 8080, where the gRPC service is listening at.
- You can verify the running container using `docker ps`.
    
### Running the application with docker compose:
 - Run `docker-compose up -d --build` which builds the images, creates and starts the containers with a single command. 
 
 
 *Happy learning!*
