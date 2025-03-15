# Competitions Manager API

API de gerenciamento de campeonatos esportivos - Iotnest.ai.

## Rodando o projeto

### Instalando as dependencias

Após baixar o projeto na sua máquina, rode o seguinte comando para instalar as dependencias do mesmo:

```
go mod tidy
```

### Docker Compose

Para criar a instância do ```Postgres``` e do ```RabbitMQ``` com Docker Compose, deve ser utilizado o seguinte comando:

```
docker-compose up -d
```

### Variaveis de ambiente

No projeto, há um arquivo chamado ```cmd/.env-example``` em que as informações devem ser copiadas para um arquivo chamado ```cmd/.env```.

Caso necessário, poderá alterar as informações do .env para apontar para sua aplicação, banco de dados ou ferramenta.

### Inicializando o projeto sem docker

Para rodar o projeto, utilize o comando:

```
go run cmd/main.go
```

### Inicializando o projeto com docker

Para construir o container docker, utilize o comando:

```
docker build -t competitions-manager-api .
```

Após isso, para rodar o container, utilize o comando:

```
docker run -p 8080:8080 competitions-manager-api
```

## Testes

### Testes unitários

Para rodar os testes unitários, você pode rodar o comando:

```
go test -cover ./...
```

E para gerar um relatório sobre a cobertura de testes, você pode gerar os seguintes comandos:

```
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

E após isso, rode o seguinte comando:

```
go tool cover -html=coverage.out -o coverage.html
```

## Endpoints

Para realizar os testes nos endpoints é importante registrar um novo usuário e após isso realizar o login com esse usuário, em que no retorno dessa requisição você deve receber um token para utilizar nos demais endpoints.

### Postman

Caso utilize o postman, você pode baixar a collection chamada X que está no projeto e importa-lá na aplicação do postman para rodar os endpoints.

### cURLs

Caso utilize cURLs para realizar requisições, você pode verificar abaixo um exemplo dos endpoints dessa API.

Listar campeonatos

```
curl --location 'http://localhost:8080/campeonatos' \
--header 'Authorization: Bearer JWT_TOKEN'
```

Listar partidas por campeonato (com filtro por equipe e rodada)

```
curl --location 'http://localhost:8080/campeonatos/2000/partidas?equipe=England&rodada=2' \
--header 'Authorization: Bearer JWT_TOKEN'
```

Cadastrar usuário

```
curl --location 'http://localhost:8080/auth/registar' \
--header 'Content-Type: application/json' \
--data '{
  "usuario": "gabriel",
	"senha": "12345",
	"confirmarSenha": "12345"
}'
```

Logar usuário

```
curl --location 'http://localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data '{
  "usuario": "gabriel",
	"senha": "12345"
}'
```

Cadastrar torcedor

```
curl --location 'http://localhost:8080/torcedores' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer JWT_TOKEN' \
--data '{
  "nome": "gabriel",
	"email": "gabriel@gmail.com",
	"time": "Coritiba"
}'
```

Broadcast

```
curl --location 'http://localhost:8080/broadcast' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer JWT_TOKEN' \
--data '{
  "tipo": "fim",
	"time": "Coritiba",
	"placar": "2-2",
	"mensagem": "O jogo terminou com placar 2-1""
}'
```