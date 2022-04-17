FROM -> de qual image você vai se basear

RUN -> Enquanto estiver no build (criação da imagem) execute "x" instrução

Cada instrução representa uma camada diferente, tenha cuidado com a formularização de sua instrução.

Exceto a última camada, as camadas anteriores são de leitura, não conseguem interagir entre si.

ENV -> cria variáveis de ambiente

LABEL -> label cria uma chave valor u=com uma descrição, version, description e etc.

VOLUME -> dentro do dockerfile, na execução o docker cria um volume com uma hash automática, a declarção auxilia na descrição do volume.

EXPOSE -> Docker é esperado que a porta "x" que está exporta seja executado pelo container.

ENTRYPOINT -> principal processo do container, como um "init"

CMD -> Com o entrypoint, ele não tem a necessidade de chamar qualquer comando. 
    -> CMD, só passa parâmetros para o principal processo no container

COPY -> pega um determinado arquivo e o adciona em um diretório

ADD -> ele como os arquivos de diretório como o COPY. No entanto, mais completo, adicionando mais 2 pontos, como: arquivos remotos e arquivos tar "explodido" ou seja, só o conteúdo

USER -> pedindo para executar com um usuário expecífico e não mais como o root  (Evite usar o root)

WORKDIR -> diretório default quando for inicializado