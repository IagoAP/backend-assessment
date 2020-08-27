# Processo Seletivo da T10

## Como rodar o projeto

O projeto foi criado em golang utilizando postgres, para que o projeto funcione temos apenas dois requisitos:

* golang
  >[Instale o golang na maquina.](https://golang.org/dl/)
* docker-compose
  >[Instale o docker compose para configurar o ambiente necessario na maquina.](https://docs.docker.com/compose/install/)

A partir do momento que esses dois requisitos sao cumpridos temos que rodar apenas dois comandos da pasta root do projeto:

```ShellScript
docker-compose up
go run main.go
```

## Base de dados

A base de dados foi estruturada de acordo com a arquitetura cqrs e, entao temos duas bases com suas respectivas tabelas:

  >pst10:
  >
  >* entities
  >* tokens
  >* product
  >
  >pst102:
  >
  >* model

A base pst10 e utilizada como base confiavel e a base pst102 e utilizada para retornar os dados ao usuario.

PS: os arquivos .sql estao disponiveis na pasta resources, mas nao sao necessarios para subir o ambiente.

### entities

| id | user_type | name | username | password |
|----|-----------|------|----------|----------|

A tabela de entidades produz um id sequencial, em user_type e armazenado a informacao sobre qual user_type aquelas credenciais pertencem. O nome e apenas opicional.

### tokens

| id_entities | token | expiration_time |
|-------------|-------|-----------------|

Essa tabela representa a lista de tokens relaciado com a entidade que o gerou, sendo eles validos ou nao, para registro dos dados.

### product

| id | description | customer_mid | customer_email | externalApp_id | assuperuser_id | activated |
|----|-------------|--------------|----------------|---------------|----------------|-----------|

Essa tabela representa os produtos cadastrados por algum ExternalApp e tabel representa se o produto foi aprovado, nao aprovado ou nao analisado por algum SuperUser.

### model

| id_externalApp | id_product | id_superUser | description | customer_mid | customer_email | activated |
|----------------|------------|--------------|------------|---------------|----------------|-----------|

Essa tabela representa todas as informacoes consideradas pertinentes a qualquer uma das entidades e ela e preenchida conforme as tabelas acima sao preenchidas.

## API

A API desenvolvida e composta de 4 requisicoes do tipo POST e 1 requisicao do tipo GET

### [POST] /RequestToken

Entidades: ExternalApp, SuperUser

body:

```javascript
{
  "Username": "User1",
  "Password": "User1"
}

```

Exemplo de resposta:

```javascript
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIwLTA4LTI3VDA3OjIwOjMyLTAzOjAwIiwidXNlcl9pZCI6MX0.qmr31P3Oxs6Pl3ZGJ5mzDhWycJW8slvuSX0cWCLXUZI"
}

```

### [POST] /IssueProductActivation

Entidades: ExternalApp

haders:

```javascript
"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIwLTA4LTI2VDIxOjIyOjQ5LTAzOjAwIiwidXNlcl9pZCI6MX0.qtFfHZEWiIq5MCRyAbFgnrxqX_ULEhdyqkRvbkuXX2U"
```

body:

```javascript
{
  "Description": "Teste",
  "CustomerMid": 2,
  "CustomerEmail": "iagoaph@gmail.com"
}

```

Exemplo de resposta:

```javascript
{
  "ActivationID": "83d2fca5f4284d53a41f8b60ac36c491",
  "message": "ok"
}

```

### [POST] /ApproveActivation

Entidades: SuperUser

haders:

```javascript
"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIwLTA4LTI2VDIxOjIyOjQ5LTAzOjAwIiwidXNlcl9pZCI6MX0.qtFfHZEWiIq5MCRyAbFgnrxqX_ULEhdyqkRvbkuXX2U"
```

body:

```javascript
{
  "ID": "9bfbb31013044313a959b363e1cb6b84"
}

```

Exemplo de resposta:

```javascript
{
  "message": "ok"
}

```

### [POST] /RejectActivation

Entidades: SuperUser

haders:

```javascript
"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIwLTA4LTI2VDIxOjIyOjQ5LTAzOjAwIiwidXNlcl9pZCI6MX0.qtFfHZEWiIq5MCRyAbFgnrxqX_ULEhdyqkRvbkuXX2U"
```

body:

```javascript
{
  "ID": "9bfbb31013044313a959b363e1cb6b84"
}

```

Exemplo de resposta:

```javascript
{
  "message": "ok"
}

```

### [GET] /ActivationRequests

Entidades: ExternalApp, SuperUser

haders:

```javascript
"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIwLTA4LTI2VDIxOjIyOjQ5LTAzOjAwIiwidXNlcl9pZCI6MX0.qtFfHZEWiIq5MCRyAbFgnrxqX_ULEhdyqkRvbkuXX2U"
```

Exemplo de resposta:

```javascript
[
  {
    "ExternalAppID": "1",
    "ID": "c69415ce8e344f52976478048deb20bf",
    "SuperUserID": "Produto nao avaliado.",
    "Description": "Teste",
    "CustomerMid": "2",
    "CustomerEmail": "iagoaph@gmail.com",
    "Activated": "Produto nao avaliado"
  },
  {
    "ExternalAppID": "1",
    "ID": "c69415ce8e344f52976478048deb20bf",
    "SuperUserID": "Produto nao avaliado.",
    "Description": "Teste",
    "CustomerMid": "2",
    "CustomerEmail": "iagoaph@gmail.com",
    "Activated": "Produto nao avaliado"
  },
  {
    "ExternalAppID": "1",
    "ID": "c69415ce8e344f52976478048deb20bf",
    "SuperUserID": "Produto nao avaliado.",
    "Description": "Teste",
    "CustomerMid": "2",
    "CustomerEmail": "iagoaph@gmail.com",
    "Activated": "Produto nao avaliado"
  }
]

```
