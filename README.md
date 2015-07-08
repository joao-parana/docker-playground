parana/docker-playground
====================

This **Dockerfile** is a [trusted build](https://registry.hub.docker.com/u/parana/docker-playground/) of [Docker Registry](https://registry.hub.docker.com/search?q=parana/).

### Based on debian Image

Play with Debian running as a command

Get the Image

    docker pull parana/docker-playground


Running some commands:

    docker run parana/docker-playground ls -lA /
    docker run parana/docker-playground help
    docker run parana/docker-playground find /etc -type f | grep bash
  

More complex use case:

    docker run -i --name debian-plus-nginx parana/docker-playground bash
    echo "deb http://nginx.org/packages/mainline/debian/ jessie nginx" >> /etc/apt/sources.list && cat /etc/apt/sources.list 
    apt-key adv --keyserver hkp://pgp.mit.edu:80 --recv-keys 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62 && echo "•••• "
    apt-get update  && echo "•••• "
    apt-get install -y ca-certificates nginx=1.9.2-1~jessie   && echo "•••• "
    ls -lA /var/lib/apt/lists  
    rm -rf /var/lib/apt/lists/*
    ls -lA /var/lib/apt
    # forward request and error logs to docker log collector
    mkdir /var/log/nginx
    ln -sf /dev/stdout /var/log/nginx/access.log
    ln -sf /dev/stderr /var/log/nginx/error.log
    ls -la /var/log/nginx
    exit

Now the debian-plus-nginx conteiner have a NGINX WebServer on Debian Machine

    docker ps -a 
    docker start debian-plus-nginx  
    docker ps
    docker exec debian-plus-nginx find / -name nginx -print 
    docker exec debian-plus-nginx ls -lA /usr/sbin/nginx
    docker exec debian-plus-nginx which nginx
    docker exec debian-plus-nginx /usr/sbin/nginx 
    docker exec debian-plus-nginx ip addr | grep eth0
    docker stop debian-plus-nginx

Starting with eth0 on Bridge mode

We can use Docker Compose to start a container with eth0 on Bridge mode 

For this you need save this contêiner as an new image and create a docker-compose.yml file.

Others Commands 

    docker version 
    docker info
    docker inspect
    docker save
    docker load
    
    
    




