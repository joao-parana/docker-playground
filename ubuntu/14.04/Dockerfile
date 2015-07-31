FROM ubuntu:14.04
MAINTAINER João Antonio Ferreira "joao.parana@gmail.com"

RUN apt-get update

# Upstart and DBus have issues inside docker. We work around that.
# RUN dpkg-divert --local --rename --add /sbin/initctl && ln -s /bin/true /sbin/initctl

# Basic build dependencies.
RUN apt-get install -y autoconf2.13 build-essential unzip yasm zip

# Libraries required for building.
RUN apt-get install -y libasound2-dev libcurl4-openssl-dev libdbus-1-dev libdbus-glib-1-dev libgtk2.0-dev libiw-dev libnotify-dev libxt-dev mesa-common-dev uuid-dev

# Gold linker is much faster than standard linker.
RUN apt-get install -y binutils-gold

# Developer tools.
RUN apt-get install -y bash-completion curl emacs git man-db python-dev python-pip vim
RUN pip install mercurial
RUN apt-get install -y git

# Create a user for development.
RUN useradd -m firefox
