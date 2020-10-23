# Topics

## Day 1: Hello Docker, Nice to meet you

1. Problem of Software Testing

   - [Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html)
   - Pain Point
   - Shopping Cart: High-level Structure (Diagram)
   - Internal Architecture
     - Resource (Controller)
     - Domain
     - Repository
     - Gateway
   - Tests
     - Unit
     - Integration
     - Contract
     - **_UI/API_**
     - **_End-to-End_**
     - **_Acceptance_**
     - Exploratory
   - The Confusion About Testing Terminology
   - Putting Tests Into Your Deployment Pipeline
   - Avoid Test Duplication
   - Writing Clean Test Code
   - Conclusion\*\*\*

2. Several Specific Benefits of Containers

   - What is Container?
   - Language Flexibility

     service A jdk8, service B jdk7, service C go1.5

   - Isolation Without Overhead

     light weight

   - Developer Efficiency

     Isolating Dependencies(libs, configuration)

   - Reproducibility

     Containers make it easier to reproduce your application environment.

3. On local machine

   - Install Docker and Editor
     1. [get-docker](https://docs.docker.com/get-docker/)
     2. [install windows home](https://docs.docker.com/docker-for-windows/install-windows-home/)  
        [WSL2](https://docs.docker.com/docker-for-windows/wsl/) must be installed before you can install and use Docker.
     3. Install VSCode & Docker Plugin
   - Hello World Docker

     ```sh
     docker run hello-world
     ```

   - Docker run: behind the scenes

     - Docker looks for the image on this computer
     - Is it installed
     - Docker searches Docker Hub for the image.
     - Is it on Docker Hub
     - Docker downloads the image
     - The image layers are installed on the computer
     - Docker create a new container and starts the program.
     - The container is running.

   - Basic Docker commands
     - run, ls, rm, ps, start, stop, help
   - Docker Image
   - Docker Container
   - Create Your own Docker Image:
     > [lab 01. Create Docker Image from Scratch](lab01-hello-world/README.md)
     - Modify, Commit, and Tag
     - Sign Up [https://hub.docker.com](https://hub.docker.com)
     - Push Image to hub.docker.com
     - Pull and Run
   - Dockerfile
     > [lab 02. Create Docker Image from Dockerfile](lab02-create-docker-image-from-Dockerfile/README.md)
     - [Reference](https://docs.docker.com/engine/reference/builder/)
       - FROM, COPY, RUN, CMD, ENTRYPOINT
     - [Dockerfile best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
     - Modify, Build + Tag, Push
     - Pull and Run

4. On local machine (Running other person implementation)

   > [lab 03. Provide Http API Service via Container](lab03-working-with-api/README.md)

   - Pull docker images to run API testing
   - Bind port
   - Run
   - Mount disk volume

5. docker-compose(intro)

   > [lab 03. Start Http API Service via docker-compose](lab03-working-with-api/README.md)

   - One service(Store Service)
   - Command: Up/Down

6. Running Frontend and Backend Service without docker-compose
   > [lab 04. Communicate via Host Network](lab04-communicate-to-other-containers-via-host-network/README.md)
   - Hello, World API v.1
   - Time API(dependency)
7. Running Frontend and Backend Service with docker-compose
   > [lab 05. Communicate via Docker Network](lab05-communicate-via-docker-compose/README.md)
   - Hello, World API v.1
   - Time API(dependency)

## Day 2: Hi Docker, You're so beautiful

1. Study the web application architecture
2. Analysis and design how to use docker
3. Define the implementation steps of using docker
   > [lab 06. From Dockerfile, Docker Image to docker-compose](lab06-docker-compose-and-Dockerfile/README.md)
   - Build Steps: Go, Node.js
   - Hello(Go) API v.2, Time(Node.js) API, and Mysql
   - HTTP
   - PORT
   - ENV(Configuration)
   - Volume(Storage)
4. Lab on local machine (Running with Docker compose)
   > lab 07. shopping cart **_images_**
   - Run Docker compose
   - Docker compose commands
   - Docker compose file
5. Lab on local machine (Build your own Docker compose and Docker file)
   > lab 08. shopping cart **_Dockerfile_**
   - Create docker images from Docker file
   - Tag docker images
   - Push Docker images to Docker hub
   - Create Docker compose
   - Run Docker compose
6. On cloud (Run your own)
   > lab 09. shopping cart on cloud
   - Create servers on cloud
   - Install Docker
   - Copy docker-compose.yaml to Cloud Server
   - Run docker compose
   - Demonstrate to execute API Testing with Robotframework + RequestsLibrary and Postman(by instructor team)
   - Demonstrate to execute workflow thought User Interface with Robotframework + SeleniumLibrary(by instructor team)
7. Docker with CI
