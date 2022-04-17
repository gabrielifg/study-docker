FROM -> de qual image você vai se basear

RUN -> Enquanto estiver no build (criação da imagem) execute "x" instrução

Cada instrução representa uma camada diferente, tenha cuidado com a formularização de sua instrução.

Exceto a última camada, as camadas anteriores são de leitura, não conseguem interagir entre si.

ENV -> cria variáveis de ambiente

LABEL -> label cria uma chave valor u=com uma descrição, version, description e etc.

VOLUME -> dentro do dockerfile, na execução o docker cria um volume com uma hash automática, a declarção auxilia na descrição do volume.

EXPOSE -> Docker é esperado que a porta "x" que está exporta seja executado pelo container.