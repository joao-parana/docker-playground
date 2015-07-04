# build e deploy

Criando a imagem

    docker build -t linux-solutions:dev .

Fazendo o build do programa GO no debian

Inicie o Contêiner para testar

    docker run --rm -i  linux-solutions:dev  

Inicie o Contêiner com shell bash para copiar o binário gerado

    docker run --rm -i -v `pwd`/bin-debian:/tmp/bin-debian  linux-solutions:dev  bash 
    # Copie o binário e saia
    cp /linux-solutions-debian-64 /tmp/bin-debian/
    exit

No seu host execute:

    cp bin-debian/linux-solutions-debian-64 ../prod/

Agora você pode gerar a imagem para produção já com o binário 
atualizado e depois criar um contêiner e executá-lo

    cd ../prod/
    docker build -t linux-solutions:prod .   
    docker run --rm -i  linux-solutions:prod

