# rest-api-2

## Overview
`rest-api-2` é uma API simples em Go que armazena usuários em memória.  
Permite criar usuários e listar todos os usuários através de endpoints HTTP.

---

## Estrutura do Projeto

### cmd/api/
Responsável por iniciar o servidor HTTP na porta `3000`.

### cmd/client/
Cliente de teste que envia uma requisição `POST /users` para a API.

### internal/repositories
Gerencia o armazenamento dos usuários em memória (lista).  
Funções principais:
- Listar usuários
- Adicionar usuário
- Verificar se email já existe

### internal/usecases
Contém a lógica de negócio:
- Validação de email único
- Criação de usuário com UUID

### internal/handlers
Lida com requisições HTTP:
- Recebe JSON
- Chama a lógica do sistema
- Retorna resposta em JSON

### internal/models
Define estruturas de dados:
- User
- Request
- Response
- ErrorResponse

---

## Como Rodar

### 1. Iniciar o servidor
```bash
go run ./cmd/api