# study-docker
Repositório criado com o intuito de armazenar e descrever os estudos feitos a cerca da ferramenta de conteinerização docker

➜ docker push gabriel1/meu-apache:1.0.0 (PUSH para o docker hub)
➜ docker image ls | grep gabriel1

--restart=always (sempre que houver um problema o container irá reiniciar)


# Dokcer registry

* docker image inspect debian
* docker history linuxtips/apache:1.0
* docker login
* docker login registry.suaempresa.com
* docker push linuxtips/apache:1.0
* docker pull linuxtips/apache:1.0
* docker image ls
* docker container run -d -p 5000:5000 --restart=always --name registry registry:2
* docker tag IMAGEMID localhost:5000/apache

# Docker Machine

Possui o proposíto de gerenciar/criar/operar máquinas virtuais que executaram o docker

# Docker swarm

Orquestrador de container nativos do docker