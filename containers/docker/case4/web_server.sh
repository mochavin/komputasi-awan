#!/bin/sh
docker container run \
	-dit \
	--name webserver1 \
	--volume $(pwd)/files:/data \
	--publish 8000:8000 \
	python:3.13.0a1-alpine3.17 \
	python3 -m http.server 8000 -d /data

