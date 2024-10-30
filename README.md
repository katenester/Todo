# Rest-Api application on golang To-do Lists


## Concepts:
- Development of Web Applications on Go, following the REST API design.
- Working with the framework GIN <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Clean architecture. Dependency injection.
- Docker.Docker compose. Migration db.
- Config <a href="https://github.com/spf13/viper">spf13/viper</a>. 
- Work with env.
- Work with DB(PostgreSql), using library <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Registration and authentication. Work with JWT. Middleware.
- Write SQL query.
- Graceful Shutdown

## Install
1. Clone project `git clone https://github.com/katenester/Todo`
2. Install all dependencies from go.mod `go mod download` `go mod tidy`
3. `make build` `make run`
4. To stop `make stop`

<b>Structure Api</b>

1. Registration `POST http://localhost:8080/auth/sign-up`
2. Authentication `POST http://localhost:8080/auth/sign-in`

Next, you need the Authorization header (JWT token)

3. GET ALL LIST - `GET http://localhost:8080/api/lists`
4. GET ITEMS BY ID - `GET http://localhost:8080/api/lists/{ID_LIST}`
5. CREATE LIST - `POST http://localhost:8080/api/lists`
6. UPDATE LIST- `PUT http://localhost:8080/api/lists/{ID_LIST}`
7. DELETE LIST- `DELETE http://localhost:8080/api/lists/{ID_LIST}`
8. GET ALL ITEMS - `GET http://localhost:8080/api/lists/{ID_LIST}/items`
9. GET ITEMS BY ID - `GET http://localhost:8080/api/items/{ID_ITEM}`
10. CREATE ITEM - `POST http://localhost:8080/api/lists/{ID_LIST}/items`
11. UPDATE ITEM -`PUT http://localhost:8080/api/items/{ID_ITEM}`
12. DELETE ITEM - `DELETE http://localhost:8080/api/items/{ID_ITEM}`

**File: [Postman](https://github.com/katenester/Todo/blob/main/postman/Todo.postman_collection.json)** for import Postman Application  
