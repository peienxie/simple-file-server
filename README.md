# Simple File Server

a simple file server written in golang

## Build application 

run below command to build the executable.

```bash
go build -o bin/simple-file-server main.go 
```

## Usage

```bash
-p, --port int      listening port (default 8080)

-r, --root string   root directory of server. (default "./root")
```

## Build with docker

### Create docker image

```
docker build --rm -t simple-file-server . && \
docker rmi `docker images --filter label=builder=true -q`
```

### Run docker image

run below command on root of the folder you want to host.

```
docker run -it --rm -p ${YOUR_PORT_NUM}:8080 -v ${PWD}:/app/root simple-file-server
```

