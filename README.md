# fold

This project demonstrates sync of postgres data to elastic search.
Also it demonstrates how we can use grpc and http for scalable solutions.

### folder structure
```
├── DockerfileApi
├── DockerfileGrpc
├── README.md
├── cmd
│   └── docker
│       ├── api
│       │   └── main.go
│       └── grpc
│           └── main.go
├── controller
│   ├── hashtagHandler.go
│   ├── projectHandler.go
│   ├── resp.go
│   └── userHandler.go
├── db
│   └── migrations
│       ├── 000001_create_user_table.down.sql
│       ├── 000001_create_user_table.up.sql
│       ├── 000002_create_project_table.down.sql
│       └── 000002_create_project_table.up.sql
├── docker-compose.yml
├── go.mod
├── go.sum
├── main
├── pgsync
│   ├── Dockerfile
│   ├── runserver.sh
│   ├── schema.json
│   └── wait-for-it.sh
├── protobuf
│   ├── Makefile.golang
│   ├── golang
│   │   ├── api
│   │   │   ├── hashtag
│   │   │   │   └── hashtag.pb.go
│   │   │   ├── project
│   │   │   │   └── project.pb.go
│   │   │   └── user
│   │   │       └── user.pb.go
│   │   ├── error.pb.go
│   │   └── grpc
│   │       ├── hashtag
│   │       │   ├── hashtag.pb.go
│   │       │   └── hashtag_grpc.pb.go
│   │       ├── project
│   │       │   ├── project.pb.go
│   │       │   └── project_grpc.pb.go
│   │       └── user
│   │           ├── user.pb.go
│   │           └── user_grpc.pb.go
│   └── proto
│       ├── api
│       │   ├── hashtag
│       │   │   └── hashtag.proto
│       │   ├── project
│       │   │   └── project.proto
│       │   └── user
│       │       └── user.proto
│       ├── error.proto
│       └── grpc
│           ├── hashtag
│           │   └── hashtag.proto
│           ├── project
│           │   └── project.proto
│           └── user
│               └── user.proto
└── service
    ├── hashtagService.go
    ├── model
    │   ├── hashtag.go
    │   ├── project.go
    │   ├── project_hashtags.go
    │   ├── user.go
    │   └── user_project.go
    ├── projectService.go
    └── userService.go



```

from the above structure:
1. `cmd/docker/api` - this is for external api endpoints http layer
2. `cmd/docker/grpc` - this is internal service with grpc
3. `protobuf` - this holds all the proto file and its compilations.
4. `pgsync`- to sync data from postgres to elastic we are using pgsync.
5. `service` - this actually implements grpc services
6. `controller`- this hold http endpoints handler.
7. `.env` - contains the environment for postgres sql - (you can change it)
### Pre requisites
1. `docker`, `golang`, `golang-migrate`
2. clone this repo. 

### Installation
1. install `docker`
2. install `golang`
3. install `golang-migrate` 
 [click here](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
   
4. run the migrate
   ```
    migrate -database "postgres://folduser:folduser_password@localhost:5432/fold_db?sslmode=disable" -path db/migrations/ up
   ```
5. run docker-compose
   ```
   docker-compose up
   ```
6. you are all set. 

### API curls
1. create user
   ```
   curl --location 'localhost:8000/v1/user' \
   --header 'Content-Type: application/json' \
   --data '{
     "name":"Ganesh shastri"
   }'
   ```
2. create project
   ```
   curl --location 'localhost:8000/v1/project' \
   --header 'Content-Type: application/json' \
   --data '{
     "name":"Ganesh'\''s project",
    "description":"test description project"
   }'
   ```

3. link user to project
   ```
   curl --location --request PUT 'localhost:8000/v1/user/2/linkproject' \
   --header 'Content-Type: application/json' \
   --data '{
   "projectIds":[2]
   }'
   ```
4. create hashtag
   ```
   curl --location 'localhost:8000/v1/hashtag' \
   --header 'Content-Type: application/json' \
   --data '{
   "name":"Ganesh hashtag"
   }'
   ```
5. link hastag to project
   ```
   curl --location --request PUT 'localhost:8000/v1/project/4/linkhashtag' \
   --header 'Content-Type: application/json' \
   --data '{
   "hashtagIds":[3]
   }'
   ```
6. Query from `elastic search`
   params: 
     `username`,`userid`,`hashtag`,`slug`,`desc`
   Curl:
     ```
    curl --location 'localhost:8000/v1/userproject?username=test'
     ```
   `slug` & `desc` follows fuzzy search in elasticsearch.

#### Advantages of this architecture:
   1. Highly scalable.
2. We can scale individual services like api service or grpc service.
3. Separation of logic. All business logic will be in grpc service.
4. Internal calls can be managed with grpc avoiding http latency in internal-service-calls.

Note: Not using any cloud deployments, all the required services are dockerized. We can simply leverage it to any cloud.