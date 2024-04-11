#!/bin/bash
docker build . -t jcaneira/tls-app:latest
docker push jcaneira/tls-app:latest

