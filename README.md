The application is implemented in order to consolidate knowledge about the communication of the application and databases, as well as the implementation of a Clean Architecture.

`/users*` enpoints interact to the PostgreSQL.

`/authors*` enpoints interact to the MongoDB.

### Build & Run (Locally)
1. Rename `.env example` file to `.env`
2. Run command for build and up application:
```
make drun
```
3. Run command for initialize Postrges database:
```
make migrate
```

### REST API
Endpoint | Method | Response codes | Description
--- | --- | --- | ---
*/users* | GET | 200, 500 | Get list of users
*/users/:id* | GET | 200, 404, 500 | Get user by id
*/users* | POST | 201, 422, 500 | Create user
*/users/:id* | PUT | 204, 404, 400, 422, 500 | Update user
*/users/:id* | DELETE | 204, 404, 500 | Delete user by id
 |  |  | 
*/authors* | GET | 200, 500 | Get list of authors
*/authors/:id* | GET | 200, 404, 500 | Get author by id
*/authors* | POST | 201, 422, 500 | Create author
*/authors/:id* | PUT | 204, 404, 400, 422, 500 | Update author
*/authors/:id* | DELETE | 204, 404, 500 | Delete author by id
