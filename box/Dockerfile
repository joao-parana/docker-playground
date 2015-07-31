FROM busybox

ENV DO_ZERO TRUE
RUN mkdir /compartilhado
VOLUME ["/compartilhado"]

CMD ["ls", "-t", "/"]

# docker build -t parana/partindo-do-busybox .
