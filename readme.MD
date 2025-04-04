# Braza Tokens Dashboard

Braza Tokens Dashboard é um serviço para fornecer um painel para tokens Braza.

## Dependências

Este projeto utiliza as seguintes dependências principais:

- [github.com/go-playground/validator/v10](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [github.com/gofiber/fiber/v2](https://pkg.go.dev/github.com/gofiber/fiber/v2)
- [github.com/gofiber/swagger](https://pkg.go.dev/github.com/gofiber/swagger)
- [github.com/stretchr/testify](https://pkg.go.dev/github.com/stretchr/testify)
- [github.com/swaggo/swag](https://pkg.go.dev/github.com/swaggo/swag)
- [go.mongodb.org/mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap)

## Estrutura do Projeto

A estrutura do projeto é a seguinte:


## Como Executar

Para executar o projeto, siga os passos abaixo:

1. Instale as dependências:
    ```sh
    go mod tidy
    ```

2. Gere a documentação Swagger:
    ```sh
    swag init -g api/api.go -o api/docs --parseDependency true --parseInternal true --parseDepth 1
    ```

3. Execute o projeto:
    ```sh
    go run main.go
    ```

## Docker

Para construir e executar o projeto usando Docker, utilize os seguintes comandos:

1. Construa a imagem Docker:
    ```sh
    docker build -t braza-tokens-dashboard -f Dockerfile .
    ```

2. Execute o container Docker:
    ```sh
    docker run -p 8000:8000 braza-tokens-dashboard
    ```

## Contato

Para mais informações, entre em contato com qualquer membro da equipe Backend do getBRAZA:
- Email: [getbraza@braza.com.br](mailto:getbraza@braza.com.br)
- URL: [https://www.brazaon.com.br](https://www.brazaon.com.br)