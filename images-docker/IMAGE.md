# Image

"Imagem é um container parado".

Nela, contém todas as instruções necessárias para a criação de um container.

O container é baseado em uma imagem.

Não utilize imagens de terceiros em produção, faça uma própria. Pois há uma grande possibilidade de que as imagens tenham vulnerabilidade. (Tenha cuidado com as imagens)

Alpaine -> Possui gerênciador de pacotes e é super enxuto. Versão simples do linux.

# COMMAND

docker image build -t "name":"version" local-dockerfile
docker container run -it meu-apache:1.0.0 

--no-cache cria imagens sem cache