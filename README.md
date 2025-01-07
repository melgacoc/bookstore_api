# Golang Bookstore API

Este é o CRUD para gerenciar o cadastro de livros em um sistema.

## Funcionalidades

- **Cadastro de um novo livro** Possibilita cadastrar um novo livro utilizando nome, sinopse, autor e categoria;
- **Consultar todos os livros**: Possibilita consultar todos os livros cadastrados;
- **Culsultar livro por ID**: Permite consultar um livro especifico;
- **Atualizar informações de um livro por ID**: Permite atualizar uma ou mais informações de um livro especifico;
- **Deletar um livro por ID**: Permite remover o cadastro de um livro especifico;

## Tecnologias
![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=Swagger&logoColor=white)

## Primeiros Passos

### Pré-requisitos

- Docker
- Docker composer (Opcional)

### Instalação

1. Clone o repositório:

```bash
git@github.com:melgacoc/bookstore_api.git
cd bookstore_api
```

2. Faça o build do container da aplicação e do banco de dados:
```bash
docker-compose build
```

3. Suba os containers:
```bash
docker-compose up
```

### Operações

Consulte a documentação pelo Swagger acessando:
```bash
http://localhost:8080/swagger/index.html
```

Construído usando o modelo de API Restfull possíu as seguintes operações:

### Escolha do DB

A escolha de utilizar um banco relacional para este caso é o fato de ter parametros que podem ser compartilhados por vários livros como os autores e genero.
Sendo assim em situações onde uma informação precisa ser alterada, como por exemplo identificar um erro de grafia em um nome, a correção é simples pois o autor e genero são FK na tabela dos livros.
