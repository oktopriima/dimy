
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
  
  # or you can import the database gank.sql directly into your local database
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

## Authors

- [@oktopriima](https://www.github.com/oktopriima)