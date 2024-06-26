# Titan - A User Microservice

Este é um projeto simples em Go utilizado para praticar a criação de uma API. A API permite cadastrar, listar, obter usuários por ID, atualizar um usuário e realizar um patch em um usuário.

## Funcionalidades

- Cadastrar usuário
- Listar usuários
- Obter usuário por ID
- Atualizar usuário
- Realizar patch em um usuário

## Requisitos

- [Docker](https://www.docker.com/get-started) instalado na máquina
- [Make](https://www.gnu.org/software/make/) instalado na máquina

## Executando o Projeto

1. Certifique-se de que o Docker está em execução na sua máquina.

   ```brew
   docker compose up
   ```

3. Clone o repositório:

   ```sh
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```
   
4. Execute o comando make para rodar o projeto:

   ```sh
   make run
   ```
   
   ## Executando o Projeto

   ##Configuração

   Você pode alterar as portas utilizadas pela aplicação editando o arquivo .env ou diretamente no arquivo docker-compose.yml.

   ## Configuração

Você pode alterar as portas utilizadas pela aplicação editando o arquivo `.env` ou diretamente no arquivo `docker-compose.yml`.

### Exemplo de `.env`

  ```env
  PORT=8080
  DSN="host=postgres port=5432 user=postgres dbname=postgres password=password sslmode=disable"
  ```

### Exemplo de `docker-compose.yml`

   ```yaml
   version: '3'
   services:
     app:
       build: .
       ports:
         - "${PORT}:8080"
     db:
       image: postgres:latest
       ports:
         - "${DB_PORT}:5432"
   ```

## Documentação da API

A API possui documentação Swagger acessível em [http://localhost:{PORT}/api/swagger/index.html](http://localhost:{PORT}/api/swagger/index.html), onde `{PORT}` é a porta configurada no `.env` ou no `docker-compose.yml`.

## Endpoints da API

- **POST /users**: Cadastrar um novo usuário
- **GET /users**: Listar todos os usuários
- **GET /users/{id}**: Obter usuário por ID
- **PUT /users/{id}**: Atualizar um usuário existente
- **PATCH /users/{id}**: Realizar um patch em um usuário existente

## Contribuição

Sinta-se à vontade para abrir issues e pull requests.

# Titan - A User Microservice
This is a simple Go project used to practice creating an API. The API allows you to register, list, retrieve users by ID, update a user, and perform a patch on a user.

## Features
- Register user
- List users
- Get user by ID
- Update user
- Perform patch on a user
  
## Requirements

- [Docker](https://www.docker.com/get-started) Docker installed on the machine
- [Make](https://www.gnu.org/software/make/) Make installed on the machine
  
## Running the Project

1. Ensure Docker is running on your machine.

   ```brew
   docker compose up
   ```
2. Clone the repository:
   
   ```sh
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   ```
   
3. Run the make command to run the project:

   ```sh
   make run
   ```

## Configuration

You can change the ports used by the application by editing the `.env` file or directly in the `docker-compose.yml` file.

Example .env

   ```env
   PORT=8080
   DSN="host=postgres port=5432 user=postgres dbname=postgres password=password sslmode=disable"
   ```

## Example docker-compose.yml
   
   ```yaml
   Copy code
   version: '3'
   services:
     app:
       build: .
       ports:
         - "${PORT}:8080"
     db:
       image: postgres:latest
       ports:
         - "${DB_PORT}:5432"
   ```

## API Documentation
The API has Swagger documentation accessible at http://localhost:{PORT}/swagger/index.html, where `{PORT}` is the port configured in the `.env` or `docker-compose.yml` file.

## API Endpoints

- **POST /users:** Register a new user
- **GET /users:** List all users
- **GET /users/{id}:** Get user by ID
- **PUT /users/{id}:** Update an existing user
- **PATCH /users/{id}:** Perform a patch on an existing user

## Contribution
Feel free to open issues and pull requests.