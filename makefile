
build:
	go build -o bin/simple-file-server main.go 

# https://stackoverflow.com/questions/48973397/how-to-automatically-delete-intermediate-stage-docker-containers-from-my-system
build-docker-image:
	docker build --rm -t simple-file-server . && \
	docker rmi `docker images --filter label=builder=true -q`

run-docker-server:
	docker run -it --rm -p 8080:8080 -v ${PWD}:/app/root simple-file-server

.PHONY: build build-docker-image run-docker-server
