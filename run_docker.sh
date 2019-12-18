#!/bin/bash 

IMAGE=orichaud/getsrv

AWS_DIR=~/.aws

AWS_DEFAULT_REGION=$(grep "^region " $AWS_DIR/config | cut -d' ' -f3-)
AWS_ACCESS_KEY_ID=$(grep "^aws_access_key_id " $AWS_DIR/credentials | cut -d' ' -f3-)
AWS_SECRET_ACCESS_KEY=$(grep "^aws_secret_access_key " $AWS_DIR/credentials| cut -d' ' -f3-)

eval $(docker-machine env or)
docker run  -p 8080:8080 -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY -e AWS_DEFAULT_REGION=$AWS_DEFAULT_REGION $IMAGE 