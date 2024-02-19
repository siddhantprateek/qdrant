#!/bin/bash

check_docker() {
  if command -v docker &> /dev/null; then
    echo "Docker is already installed."
  else
    echo "Docker is not installed. Please install Docker before running this script."
    echo "Refer to the Docker documentation for instructions: https://docs.docker.com/get-docker/"
    exit 1
  fi
}
check_docker

docker network create jenkins

# Run docker:dind Docker image
docker run --name jenkins-docker --rm --detach \
  --privileged --network jenkins --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume $(pwd)/jenkins/jenkins-docker-certs:/certs/client \
  --volume $(pwd)/jenkins/jenkins-data:/var/jenkins_home \
  --publish 2376:2376 \
  docker:dind --storage-driver overlay2


# Create a Dockerfile for customizing Jenkins image
cat <<EOF > Dockerfile.jenkins
FROM jenkins/jenkins:2.426.2-jdk17
USER root
RUN apt-get update && apt-get install -y lsb-release
RUN curl -fsSLo /usr/share/keyrings/docker-archive-keyring.asc \
  https://download.docker.com/linux/debian/gpg
RUN echo "deb [arch=\$(dpkg --print-architecture) \
  signed-by=/usr/share/keyrings/docker-archive-keyring.asc] \
  https://download.docker.com/linux/debian \
  \$(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list
RUN apt-get update && apt-get install -y docker-ce-cli
USER jenkins
RUN jenkins-plugin-cli --plugins "blueocean docker-workflow"
EOF

# Build a new docker image
docker build -t myjenkins-blueocean:2.426.2-1 -f Dockerfile.jenkins .


# Remove the Dockerfile
rm -f Dockerfile.jenkins

# Run the customized Jenkins image as a container
docker run --name jenkins-blueocean --restart=on-failure --detach \
  --network jenkins --env DOCKER_HOST=tcp://docker:2376 \
  --env DOCKER_CERT_PATH=/certs/client --env DOCKER_TLS_VERIFY=1 \
  --publish 8082:8080 --publish 50000:50000 \
  --volume $(pwd)/jenkins/jenkins-data:/var/jenkins_home \
  --volume $(pwd)/jenkins/jenkins-docker-certs:/certs/client:ro \
  myjenkins-blueocean:2.426.2-1

# Wait for Jenkins to start
sleep 10

echo "Jenkins is Ready"
