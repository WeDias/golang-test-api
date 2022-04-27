[![License](http://img.shields.io/github/license/BureauTech/Cadastrol-Server)](https://github.com/WeDias/golang-test-api/blob/main/LICENSE)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/WeDias/golang-test-api/)
[![Go CI](https://github.com/WeDias/golang-test-api/actions/workflows/go.yml/badge.svg)](https://github.com/WeDias/golang-test-api/actions/workflows/go.yml)
[![Docker Image CI](https://github.com/WeDias/golang-test-api/actions/workflows/docker-image.yml/badge.svg)](https://github.com/WeDias/golang-test-api/actions/workflows/docker-image.yml)

# golang-test-api
  A simple CRUD API made with Go, Postgres, FIber, Gorm and Docker.

- ## Cloning the repository
  To clone the repository run the following command:
  ```bash
  $ git clone https://github.com/WeDias/golang-test-api.git
  ```
  
- ## How to install and run (docker)
  Inside the newly cloned project folder run the following command:
  ```bash
  $ docker compose up
  # or
  $ docker-compose up
  ```

- ## How to install and run (manually)

- ### Installing the dependencies and setup env
  Inside the newly cloned project folder run the following command to install the dependencies:
  ```bash
  $ go mod download
  ```
  To configure the environment run this command:
  ```bash
  $ bash setup-env.sh
  ```
  It will generate a .env file with the variables below, you can edit to adapt with your database settings.
  ```bash
  PG_HOST=localhost
  PG_PASS=postgres
  PG_USER=postgres
  PG_DBNM=postgres
  PG_PORT=5432
  
  API_PORT=:3000
  ```
  After that, run the ddl-database.sql file in your database which is inside the resources folder.

- ### Running the application
  Inside the project, run the following command to run the application:
  ```bash
  $ go run main.go
  ```

- ## API
  #### Create a new product:
  ```bash
  # POST /api/v1/product

  $ curl --request POST \
  --url http://localhost:3000/api/v1/product \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "product1",
	"price": 10.99,
	"stock": 10
  }'

  # > {"id":1,"name":"product1","price":10.99,"stock":10}
  ```

  #### Get all products:
  ```bash
  # GET /api/v1/product

  $ curl --request GET \
  --url http://localhost:3000/api/v1/product

  # > [{"id":1,"name":"product1","price":10.99,"stock":10}]
  ```

  #### Get a product by id:
  ```bash
  # GET /api/v1/product/1

  $ curl --request GET \
  --url http://localhost:3000/api/v1/product/1

  # > {"id":1,"name":"product1","price":10.99,"stock":10}
  ```
  
  #### Update a product by id:
  ```bash
  # PUT /api/v1/product/1
 
  $ curl --request PUT \
  --url http://localhost:3000/api/v1/product/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "a product",
	"price": 8.99,
	"stock": 5
  }'

  # > {"id":1,"name":"a product","price":8.99,"stock":5}
  ```

  #### Delete a product by id:
  ```bash
  # DELETE /api/v1/product/1

  $ curl --request DELETE \
  --url http://localhost:3000/api/v1/product/1

  # > {"id":1,"name":"a product","price":8.99,"stock":5}
  ```
