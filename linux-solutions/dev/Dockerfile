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
RUN go build -o main . 
CMD ["/app/main"]

# Baseado no post https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

# FROM scratch
# 
