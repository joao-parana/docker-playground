FROM debian
# 
MAINTAINER João Antonio Ferreira "joao.parana@gmail.com"

RUN apt-get update && apt-get install -y \
    aufs-tools \
    automake \
    btrfs-tools \
    build-essential \
    curl \
    dpkg-sig \
    git \
    iptables \
    libapparmor-dev \
    libcap-dev \
    libsqlite3-dev \
    lxc  \
    ssh-client \
    libc++-dev \
    libarchive-tar-perl \
    mercurial \
    parallel \
    reprepro \
    ruby2.1  \
    ruby2.1-dev \
    s3cmd=1.5.0~rc1-2 

# setting the image’s main command, allowing that image to be run as though it was that command
ENTRYPOINT ["s3cmd"]

# use --help as the default flag for s3cmd command
CMD ["--help"]
