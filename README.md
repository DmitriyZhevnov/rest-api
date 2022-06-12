# rest-api
Application that interacts with SQL and NoSQL Dbs. 

# user-service

# REST API

GET /users  -- list of users -- 200, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users -- create user -- 201, 422, 500
PUT / users/:id -- update user -- 204, 404, 400, 422, 500
DELETE /users/:id -- delete user by id -- 204, 404, 500