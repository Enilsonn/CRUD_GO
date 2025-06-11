# 📘 CRUD\_GO

Uma API REST simples desenvolvida em Go para gerenciar usuários em memória. Este projeto implementa as operações básicas de **Create, Read, Update e Delete (CRUD)** utilizando a biblioteca [Chi](https://github.com/go-chi/chi) para roteamento.

---

## 🚀 Funcionalidades

- ✅ Criar um novo usuário
- 🔍 Buscar um usuário por ID
- 📋 Listar todos os usuários
- ✏️ Atualizar usuário por ID
- ❌ Deletar usuário por ID

---

## 🏗️ Tecnologias Utilizadas

- **Go (Golang)**
- **Chi Router**

---

## 📦 Instalação

```bash
# Clone o repositório
git clone https://github.com/Enilsonn/CRUD_GO.git
cd CRUD_GO

# Instale as dependências (caso utilize go.mod)
go mod tidy

# Execute a API
go run main.go
```

A API estará disponível em:\
📍 `http://localhost:8080`

---

## 📬 Endpoints

| Método | Rota    | Descrição                | Corpo da Requisição                        |
| ------ | ------- | ------------------------ | ------------------------------------------ |
| POST   | `/`     | Cria um novo usuário     | `{ "name": "", "email": "", "phone": "" }` |
| GET    | `/`     | Lista todos os usuários  | —                                          |
| GET    | `/{id}` | Busca um usuário por ID  | —                                          |
| PUT    | `/{id}` | Atualiza um usuário      | `{ "name": "", "email": "", "phone": "" }` |
| DELETE | `/{id}` | Deleta um usuário por ID | —                                          |

---

## 🧪 Exemplo de Uso (via curl)

```bash
# Criar um usuário
curl -X POST http://localhost:8080/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Fabio","email":"fabiolima@email.com","phone":"123456"}'

# Buscar todos os usuários
curl http://localhost:8080/

# Buscar usuário por ID
curl http://localhost:8080/1

# Atualizar usuário
curl -X PUT http://localhost:8080/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Fabio Lima","email":"fabiolima@email.com","phone":"99999"}'

# Deletar usuário
curl -X DELETE http://localhost:8080/1
```

---

## 🗃️ Estrutura Interna

- `main.go`: Inicia o servidor e configura dependências
- `apiMux.go`: Define rotas e middlewares
- `handlers/`: Implementação dos métodos de CRUD
- `middlewares/`: Middleware para setar `Content-Type: application/json`
- `tools/`: Funções utilitárias como serialização de JSON

---

## 📌 Observações

- Os dados são armazenados apenas em memória (`map[int]User`), ou seja, reiniciar o servidor apaga os dados.
- Essa API é ideal para fins de aprendizado e experimentação com Go.

---

## 👤 Autor

**Enilson Lima da Costa Junior**\
🔗 [LinkedIn](https://www.linkedin.com/in/enilson-lima-944532140/)\
👨‍💻 [GitHub](https://github.com/Enilsonn)

# CRUD_GO
