
# Fruits API 🍉🍊🍍
 
REST API built from scratch with a MySQL database

## Configuration 🔧

1.) Add and configure the `.env` file in the root directory :

``` bash
scratch-api
    ├── api
    ├── bin
    ├── scripts
    ├── storage
    ├── types
    ├── .env          <--- create .env file
    ├── .gitignore
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── Makefile
```
```env
    # .env
    DB_PASSWORD="YourDatabasePassword123"
```

2.) Create and run mysql server
- ensure docker is installed on your machine
```bash
    docker-compose up --build -d
```

## Run API ▶️
- ensure make is installed on your system:
```bash
    // linux
    sudo apt-get install build-essential

    // windows (assuming you already have choco installed)
    choco install make
``` 
- Run API
```bash
    make run
```
