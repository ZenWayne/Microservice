#!/bin/bash

docker build -t eth/nft:latest -f containers/Dockerfile .

# docker tag eth/nft:latest docker.io/i6o6i/emsvc
# kubectl rollout restart deployment/mysql