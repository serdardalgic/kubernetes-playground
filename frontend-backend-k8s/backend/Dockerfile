FROM golang:1.10.1 as builder
ARG BUILDERDIR="/go/src/github.com/serdardalgic/kubernetes-playground/"
WORKDIR $BUILDERDIR
ARG mux="github.com/gorilla/mux"
ARG redigo="github.com/garyburd/redigo/redis"
RUN go get -d -v $mux
RUN go get -d -v $redigo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
ARG BUILDERDIR="/go/src/github.com/serdardalgic/kubernetes-playground/"
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder $BUILDERDIR/app .
CMD ["./app"]  
