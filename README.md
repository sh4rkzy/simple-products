
# **Product Management API**

## **Descrição**
A **Product Management API** é uma aplicação construída em **Go (v1.23)** que fornece endpoints para gerenciar produtos, permitindo criar e listar produtos de maneira eficiente. O projeto utiliza MongoDB como banco de dados e é estruturado com boas práticas de arquitetura modular, utilizando o framework **Gin** para criar APIs RESTful.

---

## **Requisitos**

- **Go** versão 1.23 ou superior.
- **MongoDB** em execução local ou remoto.
- **Docker** e **Docker Compose** (opcional para ambiente de desenvolvimento).

---

## **Funcionalidades**

### Endpoints Implementados

#### 1. **Health Check**
- **Método:** `GET`
- **Rota:** `/api/v1/health`
- **Descrição:** Verifica se o servidor está ativo.
- **Resposta de Exemplo:**
  ```json
  {
    "application": "OK",
    "databases": "OK",
    "status_code": 200,
    "transaction": {
        "timestamp": "2024-11-24 23:17:12",
        "transaction_id": "fe68942c-9b8f-49b7-8d76-4ab719dae67b"
  }

  ```

#### 2. **Listar Produtos**
- **Método:** `GET`
- **Rota:** `/api/v1/products`
- **Descrição:** Retorna todos os produtos cadastrados.
- **Resposta de Exemplo:**
  ```json
  {
    "status_code": 200,
    "message": "OK",
    "products": [
      {
        "product_id": "1a2b3c4d",
        "name": "Smartphone XYZ",
        "price": 1299.99,
        "dt_created": "2024-11-24 10:00:00",
        "dt_updated": "2024-11-24 10:00:00"
      }
    ],
    "transaction": {
      "transaction_id": "9i8u7y6t5r4e3w2q1",
      "timestamp": "2024-11-24 12:00:00"
    }
  }
  ```

#### 3. **Criar Produto**
- **Método:** `POST`
- **Rota:** `/api/v1/products`
- **Descrição:** Cria um novo produto com os dados fornecidos.
- **Body de Exemplo:**
  ```json
  {
    "name": "Smartphone XYZ",
    "price": 1299.99
  }
  ```
- **Resposta de Exemplo:**
  ```json
  {
    "status_code": 201,
    "message": "Produto criado com sucesso",
    "product_id": "1a2b3c4d",
    "transaction": {
      "transaction_id": "8h7g6f5e4d3c2b1a",
      "timestamp": "2024-11-24 12:00:00"
    }
  }
  ```

---

## **Como Iniciar o Projeto**

### **Executar com Docker Compose**
Para rodar o projeto em um ambiente isolado, você pode usar o Docker Compose. Isso já configura o fluxo com MongoDB e Nginx como proxy reverso.

1. Certifique-se de que o **Docker** e **Docker Compose** estão instalados.
2. Execute o comando abaixo para iniciar os containers:
   ```bash
   docker-compose up --build
   ```
3. O servidor estará disponível em `http://localhost`.

### **Executar Manualmente (Localmente)**
Se preferir rodar localmente em modo de desenvolvimento, siga as etapas abaixo:

1. Configure o banco de dados MongoDB (pode rodar com o Docker Compose somente para o MongoDB).
2. Adicione a variável de ambiente `MONGO_URI` no arquivo `.env`:
   ```
   MONGO_URI=mongodb://localhost:27017
   ```
3. Instale as dependências:
   ```bash
   go mod tidy
   ```
4. Execute o servidor:
   ```bash
   go run main.go
   ```
5. O servidor estará acessível em `http://localhost:8080`.

---

## **Configuração do Nginx**
O Nginx é usado como proxy reverso para redirecionar o tráfego para o serviço da API. A configuração do Nginx está no arquivo `nginx.conf`.

Exemplo de configuração:
```nginx
server {
    listen 80;

    server_name localhost;

    location / {
        proxy_pass http://api:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root /usr/share/nginx/html;
    }
}
```

---

## **Docker Compose**
Aqui está o exemplo de um `docker-compose.yml` atualizado para o projeto:
```yaml
version: "3.9"
services:
  api:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: api
    ports:
      - "3000:3000"
    environment:
      - MONGO_URI=mongodb://mongo:27017
    depends_on:
      - mongo
    networks:
      - backend

  mongo:
    image: mongo:6.0
    container_name: mongodb
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: root
    volumes:
      - mongo-data:/data/db
    networks:
      - backend

  nginx:
    image: nginx:latest
    container_name: nginx
    ports:
      - "80:80"
    volumes:
      - ./nginx.conf:/etc/nginx/conf.d/default.conf:ro
    depends_on:
      - api
    networks:
      - backend

networks:
  backend:
    driver: bridge

volumes:
  mongo-data:
```

---

## **Tecnologias Utilizadas**

- **Go (1.23)**: Linguagem principal da aplicação.
- **Gin**: Framework para construção de APIs REST.
- **MongoDB**: Banco de dados para persistência de dados.
- **Docker**: Para criar e executar ambientes isolados.
- **Docker Compose**: Para gerenciar múltiplos containers no ambiente de desenvolvimento.
- **Nginx**: Proxy reverso para gerenciar o tráfego de requisições.

---

## **Licença**
Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
