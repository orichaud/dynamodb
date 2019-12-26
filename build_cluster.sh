#!/bin/bash
source ecs.env

echo "[-] Create key pair"
aws ec2 create-key-pair --key-name $KEY_PAIR_NAME --query 'KeyMaterial' --output text > $KEY_PAIR_NAME.pem
aws ec2 describe-key-pairs --key-name $KEY_PAIR_NAME
echo "[+] DONE"

echo "[-] Create ECR repository"
aws ecr create-repository --repository-name $REPO_NAME
aws ecr describe-repositories --repository-names $REPO_NAME
echo "[+] DONE"

echo "[-] Create a cluster configuration"
ecs-cli configure --cluster $CLUSTER_NAME --default-launch-type EC2 --config-name $CONFIG_NAME --region $AWS_DEFAULT_REGION
echo "[+] DONE"

echo "[-] Create profile"
ecs-cli configure profile --access-key $AWS_ACCESS_KEY_ID --secret-key $AWS_SECRET_ACCESS_KEY --profile-name $PROFILE_NAME
echo "[+] DONE"

echo "[-] Create cluster"
ecs-cli up --force --keypair $KEY_PAIR_NAME --capability-iam --size 2 --instance-type t2.medium --cluster-config $CONFIG_NAME --ecs-profile $PROFILE_NAME
aws ecs describe-clusters --clusters $CLUSTER_NAME
echo "[+] DONE"
