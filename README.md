## Super hero API

Rest API que usa a [SuperHeroAPI](https://superheroapi.com/) como fonte de dados

Crie um arquivo **.env** na raiz do projeto seguindo o arquivo **.env.example**

### Endpoints

| MÉTODO | ENDPOINT             | DESCRIÇÃO                          |
| ------ | -------------------- | ---------------------------------- |
| GET    | /super               | Listar todos os supers registrados |
| GET    | /super/heros         | Listar todos os herois registrados |
| GET    | /super/villains      | Listar todos os vilões registrados |
| GET    | /search?name=XXXXXXX | Buscar super por nome              |
| GET    | /super/:id           | Buscar super por ID                |
| POST   | /super/              | Criar um super                     |
| DELETE | /super/:id           | Remover um super                   |
