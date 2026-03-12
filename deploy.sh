#!/bin/bash
set -e

# ==== 1. デプロイパッケージの作成 ====
echo "======================================"
echo "1. Creating deployment package..."
echo "======================================"
cd source
./build.sh
cd ..

# ==== 2. Terraformによるインフラストラクチャのデプロイ ====
echo "======================================"
echo "2. Deploying infrastructure with Terraform..."
echo "======================================"
cd terraform

echo "Running: terraform init"
terraform init

echo "Running: terraform plan"
terraform plan -out=/tmp/$$.tfplan

echo "Running: terraform apply"
terraform apply /tmp/$$.tfplan

echo "======================================"
echo "Deployment completed!"
echo "======================================"
