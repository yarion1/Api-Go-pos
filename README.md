# Go Products API

Esta é uma API de gerenciamento de produtos desenvolvida em **Go** utilizando o framework **Fiber**.

A API oferece operações de **CRUD** (Create, Read, Update e Delete) para produtos, além de funcionalidades de **autenticação e autorização** com **JWT (JSON Web Tokens)**.

O banco de dados utilizado é o **SQLite3**, e o ORM **GORM** facilita a interação com o banco de dados.

A documentação da API é gerada automaticamente com **Swagger**. Para testes, utilizamos o **Testify**, uma biblioteca poderosa para testes unitários e de integração.

## 🚀 Funcionalidades

### 🔐 Autenticação e Autorização
- Cadastro de usuários: `POST /users`
- Geração de token JWT para autenticação: `POST /users/generate_token`
- Proteção de rotas com **middleware JWT**

### 📦 CRUD de Produtos
- **Criar produto**: `POST /products`
- **Buscar produto por ID**: `GET /products/:id`
- **Listar todos os produtos**: `GET /products`
- **Atualizar produto**: `PUT /products/:id`
- **Excluir produto**: `DELETE /products/:id`

## 🛠 Tecnologias Utilizadas
- **Golang** + **Fiber** (framework web)
- **SQLite3** + **GORM** (ORM para banco de dados)
- **JWT** (JSON Web Tokens para autenticação)
- **Swagger** (documentação automática da API)
- **Testify** (testes unitários e de integração)

## 📖 Instalação e Configuração

### 1️⃣ Pré-requisitos
Certifique-se de ter instalado:
- [Go](https://go.dev/dl/) (1.18 ou superior)
- [SQLite3](https://www.sqlite.org/download.html)

### 2️⃣ Clonando o repositório
```sh
git clone https://github.com/seu-usuario/go-products-api.git
cd go-products-api
