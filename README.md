# nossobr
O nossobr é um microserviço que tem a funcionalidade de retornar dados referente a uma cidade ou estado do 🇧🇷

A busca será feita diretamento no banco de dados (postgresql) do microserviço, caso a cidade ainda não esteja cadastrado, será feito uma requisição no endpoint do BrasilAPI usando a lib [brasilapi-go](https://github.com/isaqueveras/brasilapi-go) e aproveitando as informações para cadastrar os dados no microserviço.

Esse é apenas o primeiro microserviço que devo construir para criar uma aplicação que disponibilizará os dados público do Brasil atrelados a cidades e estados.

Essa aplicação será feito usando arquitetura limpa , de acordo com a imagem abaixo usando gRPC para outras aplicações se comunicar com o nossobr.
![arquitetura limpa](./clean-arch.png "Arquitetura Limpa")
