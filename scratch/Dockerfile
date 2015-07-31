FROM tianon/true

ENV DO_ZERO TRUE

CMD ["/true"]

# O Contêiner conterá apenas os arquivos: 
# resolv.conf, hostconfig.json, hosts, hostname, 
# config.json, resolv.conf.hash usados pela 
# Infraestrutura do Docker e o arquivo binário 
# executável /true ficará exposto aos usuários
#
# Criando a imagem e executando
# docker build -t parana/partindo-do-zero .
# docker run --name menor-possivel -h minimo.local parana/partindo-do-zero
