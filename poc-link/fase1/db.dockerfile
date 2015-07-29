FROM busybox:ubuntu-14.04
# 
MAINTAINER João Antonio Ferreira "joao.parana@gmail.com"

ENV REFRESHED_AT 2015-07-15

RUN ls -lAt /

ADD db.entrypoint.sh /db.entrypoint.sh
RUN chmod a+rx /db.entrypoint.sh
ENTRYPOINT ["/db.entrypoint.sh"]
CMD ["99999999"]