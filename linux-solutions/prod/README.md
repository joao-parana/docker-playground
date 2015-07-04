Atenção: o binário usado aqui para Debian 64 bits deve 
ser compilado e linkeditado no contêiner do ambiente dev
que está preparado para isso. Use a opção -v para ele 
para gerar o binário no diretório do host.

Para criar com o binário já disponível use:

    docker build -t linux-solutions:prod . 
    
Para executar o conteiner use:

    docker run --rm -i  linux-solutions:prod

Este contêiner serve apenas a proposito didático. Para fazer algo mais funcional seria necessário no mínimo ler o diretório corrente e fornecer os arquivos HTML, CSS e JS como conteúdo estático   

