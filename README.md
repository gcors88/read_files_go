<h1 align="center"> Read Files - GOLANG Project </h1>

## Índice

- [Sobre](#Sobre)
- [Tecnologias Utilizadas](#Tecnologias)
- [Iniciando o projeto](#Iniciando)

<hr>


<!-- About -->

# Sobre

<p align="left"> Este projeto foi feito com o intuito de percorrer uma pasta contendo arquivos txt onde cada campo é separado por um circunflexo "^".
Apos percorrer cada arquivo e fazer a leitura dos mesmos é feito um upsert do dado do usuário em um banco de dados postgres. </p>

<!-- TECHNOLOGIES -->

# Tecnologias

- **Tecnologias**
  - [Go](https://go.dev/)
  - [Docker](https://www.docker.com/)
  - [Gorm](https://gorm.io/index.html)
  - [Postgres](https://www.postgresql.org/)

<hr>


<!-- TECHNOLOGIES -->

# Iniciando

### Pré-requisitos

- Go lang

  ```sh
  https://go.dev/
  ```

- Docker

  ```sh
  https://www.docker.com/
  ```

<hr>


### Instalação

1. Após clonar o repositório em um diretório local intalaremos os pacotes com o comando:

   ```sh
   go get ./...
   ```

2. Inicializar o banco de dados:

   ```sh
   docker-compose up kafka
   ```

### Uso

1 - Inicialize o script executando o comando abaixo:
   ```sh
   go run main.go
   ```


2 - Serão exibidos os logs dos dados que foram criado/atualizados no banco


3 - Para visualização dos dados recomendo a instalação do dbeaver ou algum outro client de banco de dados de sua preferência.




Glauber Oliveira - [Linkedin](https://www.linkedin.com/in/gcolliveira/) - glauber17230@gmail.com 

