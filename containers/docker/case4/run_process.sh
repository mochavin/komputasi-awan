#!/bin/sh
docker rm -f jokefetcher1

docker container run \
	--name jokefetcher1 \
	-dit \
	-e DELAY=8 \
	-v $(pwd)/files:/data \
	-v $(pwd)/script:/script \
	--workdir /data \
	alpine:3.18 \
	/bin/sh /script/getjokes.sh


