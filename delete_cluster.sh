#!/bin/bash
source ecs.env

echo "[-] Delete key pair"
aws ec2 delete-key-pair --key-name $KEY_PAIR_NAME
echo "[+] DONE"

echo "[-] Delete ECR repository"
aws ecr delete-repository --repository-name $REPO_NAME

echo "[-] Create a cluster configuration"
ecs-cli compose service rm --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME
ecs-cli down --force --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME
echo "[+] DONE"
