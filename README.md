## Super hero API

Rest API que usa a [SuperHeroAPI](https://superheroapi.com/) como fonte de dados

### Info

- Versão do Go: 1.14
- Bibliotecas:
  - [Fiber](github.com/gofiber/fiber)
  - [gorm](github.com/jinzhu/gorm)
  - [go.uuid](github.com/satori/go.uuid)
  - [godotenv](github.com/joho/godotenv)
  - [testify](github.com/stretchr/testify)

### Como Usar

Crie um arquivo **.env** na raiz do projeto seguindo o arquivo **.env.example**

E, com o docker instalado, execute o comando:

```bash
docker-compose up
```

### Testes

Primeiro configure o PostgreSQL

```bash
docker run --rm --name pgsql -d -p 5432:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_DB=levpay_test postgres:12-alpine
```

Depois:

```bash
go test ./...
```

### Endpoints

| MÉTODO | ENDPOINT             | DESCRIÇÃO                          | BODY                         |
| ------ | -------------------- | ---------------------------------- | ---------------------------- |
| GET    | /super               | Listar todos os supers registrados |
| GET    | /super/heros         | Listar todos os herois registrados |
| GET    | /super/villains      | Listar todos os vilões registrados |
| GET    | /search?name=XXXXXXX | Buscar super por nome              |
| GET    | /super/:id           | Buscar super por ID                |
| POST   | /super/              | Criar um super                     | { "name": "CHARACTER NAME" } |
| DELETE | /super/:id           | Remover um super                   |
