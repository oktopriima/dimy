
# GOLANG SIMPLE ORDERS

Take home exercise
- Language : Golang 1.20
- Framework : Echo
- ORM : gorm
- Database : MySQL

## Run Locally

Clone the project

```bash
  git clone https://github.com/oktopriima/dimy
```

Go to the project directory

```bash
  cd $GOPATH/src/gitthub.com/oktopriima/dimy
```

Setup env
```bash
  cp env.yaml.example env.yaml
  # adjust value based on your local setup
```

Migrate database
```bash
  # install rubenv migration tools
  go get -v github.com/rubenv/sql-migrate/... 

  # only for the migration
  # because I use different tools for migration
  cp dbconfig-example.yml dbconfig.yml
  # adjust the database connection

  sql-migrate up --env=dev
  
  # or you can import the database dimy.sql directly into your local database
  
```

Install dependencies

```bash
  go mod tidy -compact=1.20
  go mod vendor
```

Start the server

```bash
  go build
  ./dimy
```

Docker setup
```bash
# update env.yaml change the mysql part
# build the docker images
docker build -t dimy .

# run your application
docker run -d dimy

# check docker ip for golang container
# change HOST on postman environment from localhost to docker IP
```

## AVAILABLE ENDPOINT
Create Orders
```bash
curl --location 'http://localhost:8080/api/orders' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id" : 2,
    "address_id" : 2,
    "products" : [
        {
            "id" : 1,
            "quantity": 10
        },
        {
            "id" : 2,
            "quantity": 10
        }
    ]
}'
```

Pay Orders
```bash
curl --location 'http://localhost:8080/api/orders/payment' \
--header 'Content-Type: application/json' \
--data '{
    "order_id" : 3,
    "payments" : [
        {
            "id" : 1,
            "value" : 100000
        },
        {
            "id" : 2,
            "value" : 50000
        }
    ]
}'
```

Detail Customer
```bash
curl --location 'http://localhost:8080/api/customers/1'
```
## Authors

- [@oktopriima](https://www.github.com/oktopriima)