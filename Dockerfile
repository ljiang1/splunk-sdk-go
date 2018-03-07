
#Starting from the latest Golang image
FROM golang:1.9
#RUN bash -c "echo something"

# INSTALL any further tools you need here so they are cached in the docker build

# Set the WORKDIR to the project path in your GOPATH, e.g. /go/src/github.com/go-martini/martini/
WORKDIR /go/src/first_project

# Copy the content of your repository into the image
COPY . ./

# Install dependencies through go get, unless you vendored them in your repository before
# Vendoring can be done with an external tool like godep or glide
# Go versions after 1.5.1 include support for a vendor directory
RUN go get


RUN bash -c "echo something"
RUN apt-get update && apt-get install -y curl
