GOCMD=/usr/local/bin/go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

DOCKER=docker
DOCKER_VM=or

PACKAGES=getsrv/core getsrv/server
SERVER=getsrv
BASEDIR=.
BINDIR=./bin
SRCDIR=./src
IMAGE=orichaud/getsrv
GOPATH=$(shell pwd):$(shell echo $$GOPATH)


all: clean build docker

build:
	mkdir -p $(BINDIR)/linux && GOOS=linux GOARCH=386 CGO_ENABLED=0 GOPATH=$(GOPATH) $(GOBUILD) -o $(BINDIR)/linux $(PACKAGES) 
	mkdir -p $(BINDIR)/darwin && GOOS=darwin GOARCH=386 CGO_ENABLED=0 GOPATH=$(GOPATH) $(GOBUILD) -o $(BINDIR)/darwin $(PACKAGES)

test:
	GOPATH=$(GOPATH) $(GOTEST) -v -timeout 30s $(PACKAGES)	

docker:
	eval $$(docker-machine env $(DOCKER_VM)) && $(DOCKER) build -t $(IMAGE) .
	eval $$(docker-machine env $(DOCKER_VM)) && $(DOCKER) images

clean:
	- rm -rf $(BINDIR)
	- eval $$(docker-machine env $(DOCKER_VM)) && $(DOCKER) rmi -f $(IMAGE)
	- $(GOCLEAN) $(PACKAGE)
