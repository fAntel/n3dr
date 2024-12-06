#!/bin/bash -e
export GITHUB_URL=https://github.com
export GS_URI=go-swagger/go-swagger/releases/download
export GS_VERSION=v0.30.4
export GS_URL=${GITHUB_URL}/${GS_URI}/${GS_VERSION}/swagger_darwin_arm64
export GS_DIR=internal/app/n3dr/goswagger
curl -L \
  ${GS_URL} \
  -o swagger
chmod +x swagger
mkdir -p "${GS_DIR}"
./swagger generate client \
  --name=nexus3 \
  --spec configs/swagger/nexus3.json \
  --target="${GS_DIR}" \
  --skip-validation && \
go mod tidy && \
rm swagger*