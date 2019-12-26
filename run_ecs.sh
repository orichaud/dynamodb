#!/bin/bash
source ecs.env

eval $(aws ecr get-login --region us-east-1 --no-include-email)
docker tag $IMAGE $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$REPO_NAME
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$REPO_NAME

#ecs-cli compose up --project-name  --create-log-groups --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME -region $AWS_DEFAULT_REGION --file docker-compose.yml --ecs-params ecs-params.yml
#ecs-cli ps --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME -region $AWS_DEFAULT_REGION
#ecs-cli compose service up --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME -region $AWS_DEFAULT_REGION 

#aws iam --region $AWS_DEFAULT_REGION create-role --role-name $TASK_ROLE_NAME --assume-role-policy-document file://task-execution-assume-role.json
#aws iam --region $AWS_DEFAULT_REGION attach-role-policy --role-name $TASK_ROLE_NAME --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy


sed "s/\$\$IMAGE\$\$/$AWS_ACCOUNT_ID/g" getsrv-compose.yml.template
ecs-cli compose --verbose --file getsrv-compose.yml --ecs-params getsrv-ecs-params.yml --project-name $PROJECT_NAME up --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME -region $AWS_DEFAULT_REGION 