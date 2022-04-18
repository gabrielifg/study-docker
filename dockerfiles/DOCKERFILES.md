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

\ -> utilizado para identar comandos que são executados na mesma instrução/linha para deixar mais organizados

.dockerignore -> utilize para minimizar arquivos desnecessários nas imagens

Use RUN com moderação para minimizar a quantidade de camada, evitando danos

NON-ROOT -> evita que usuário não autorizados execute imagens em modo root

CMD e ENTRYPOINTS (modos shell e exec)
    EXEC -> ENTRYPOINT ["GIROPOPS"]
    SHELL -> ENTRYPOINT "GIROPOPS"
    quando executar um comando com entrypoint e cmd no mesmo dockerfile (obrigado a usar no modo exec) o cmd somente adicionará parâmetros


--no-cache -> use ou não, depende do caso

START -> NÃO É ELEGANTE, se possível não use
    ENTRYPOINT ["python", "start.py"]

Variáveis -> exemplo

    ENV app_dir /opt/app
    WORKDIR ${app_dir}
    ADD .$app_dir

ARG -> ou seja, argumento, te possibilita mudar no momento do build o valor da variável "--build-arg"

HEALTCHECK -> te possibilita definir no dockerfile instruções para verificar a saúde do container 
    OBS: melhor colocar no compose

 docker container rm -f $(docker ps -q)    -> remove os containers através do ID


Criando imagens a partir de um container com mudanças

    ➜ docker commit -m "Ubuntu vim + curl" e0d7db611009
    ➜ docker image tag 7cc4ed5f5092 ubuntu_vim_curl:1.0
    ➜ docker container run -it ubuntu_vim_curl:1.0

--- Definição curso descomplicando docker ---

ADD => Copia novos arquivos, diretórios, arquivos TAR ou arquivos remotos e os adicionam ao filesystem do container;

CMD => Executa um comando, diferente do RUN que executa o comando no momento em que está "buildando" a imagem, o CMD executa no início da execução do container;

LABEL => Adiciona metadados a imagem como versão, descrição e fabricante;

COPY => Copia novos arquivos e diretórios e os adicionam ao filesystem do container;

ENTRYPOINT => Permite você configurar um container para rodar um executável, e quando esse executável for finalizado, o container também será;

ENV => Informa variáveis de ambiente ao container;

EXPOSE => Informa qual porta o container estará ouvindo;

FROM => Indica qual imagem será utilizada como base, ela precisa ser a primeira linha do Dockerfile;

MAINTAINER => Autor da imagem; 

RUN => Executa qualquer comando em uma nova camada no topo da imagem e "commita" as alterações. Essas alterações você poderá utilizar nas próximas instruções de seu Dockerfile;

USER => Determina qual o usuário será utilizado na imagem. Por default é o root;

VOLUME => Permite a criação de um ponto de montagem no container;

WORKDIR => Responsável por mudar do diretório / (raiz) para o especificado nele;