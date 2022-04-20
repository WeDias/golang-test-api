# golang-test-api

  A simple CRUD API made with Go, Postgres, FIber, Gorm and Docker.

- ## Cloning the repository
  To clone the repository run the following command:
  ```
  $ git clone https://github.com/WeDias/golang-test-api.git
  ```
  
- ## How to install and run (docker)
	Inside the newly cloned project folder run the following command:
	```
	$ docker compose up
	```

- ## How to install and run (manually)

- ### Installing the dependencies
  Inside the newly cloned project folder run the following command to install the dependencies:
  ```
  $ go mod download
  ```

- ### Running the application
  Inside the project, run the following command to run the application:
  ```
  $ go run main.go
  ```

- ## API
  ```
	# POST /api/v1/product

	$ curl --request POST \
	--url http://localhost:3000/api/v1/product \
	--header 'Content-Type: application/json' \
	--data '{
			"name": "product1",
			"price": 10.99,
			"stock": 10
	}'

	> {"id":1,"name":"product1","price":10.99,"stock":10}
	```

	```
	# GET /api/v1/product

	$ curl --request GET \
	--url http://localhost:3000/api/v1/product

	> [{"id":1,"name":"product1","price":10.99,"stock":10}]
	```

	```
	# GET /api/v1/product/1

	$ curl --request GET \
	--url http://localhost:3000/api/v1/product/1

	> {"id":1,"name":"product1","price":10.99,"stock":10}
	```

	```
	# PUT /api/v1/product/1

	$ curl --request PUT \
	--url http://localhost:3000/api/v1/product/1 \
	--header 'Content-Type: application/json' \
	--data '{
			"name": "a product",
			"price": 8.99,
			"stock": 5
	}'

	> {"id":1,"name":"a product","price":8.99,"stock":5}
	```

	```
	# DELETE /api/v1/product/1

	$ curl --request DELETE \
	--url http://localhost:3000/api/v1/product/1

	> {"id":1,"name":"a product","price":8.99,"stock":5}
  ```
