FROM golang:latest 

RUN go get github.com/docker/engine-api/client
RUN go get github.com/docker/engine-api/types
RUN go get github.com/docker/engine-api/types/container

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go build -o main .

EXPOSE 3000

ENTRYPOINT ["/app/main"]