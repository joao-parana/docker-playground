## FROM golang:onbuild
## Os comandos Docker abaixo são os contidos na imagem golang:onbuild
## que servem para criar um ambiente de desenvolvimento no momento do 
## build e copiar o código fonte pro lugar correto e em seguida invocar 
## o build à partir dos fontes.
## 
## FROM golang:1.4.2
## 
## RUN mkdir -p /go/src/app
## WORKDIR /go/src/app
## 
## # this will ideally be built by the ONBUILD below ;)
## CMD ["go-wrapper", "run"]
## 
## ONBUILD COPY . /go/src/app
## ONBUILD RUN go-wrapper download
## ONBUILD RUN go-wrapper install

FROM golang:latest 

MAINTAINER João Antonio Ferreira "joao.parana@gmail.com"

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
# RUN go build -o main . 
ENV CGO_ENABLED   0 
ENV GOOS          linux 
RUN rm -rf /linux-solutions-debian-64 && \
    go build -a -installsuffix cgo -o main . &&  \
    cp main /linux-solutions-debian-64

ENV blog_addr   http://64.111.112.150
CMD ["/app/main"]

# Baseado no post https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
