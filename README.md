# Link Shortener
## A small project to demonstrate how to build a webapp based on containerized microservices
---
**⚠️ Note: This is not production ready. Do not expose to an unsecure environment.**

This project is structured in 3 containers/services:

    - Frontend
        - Nginx: Serves the static frontend and reverse-proxies the backend routines that access and manage the database
    
    - Database
        - CouchDB: A document-oriented NoSQL database with a REST API

    - Backend
        - Go routines: Some simple Go routines that manage http requests from the frontend to generate new links in the database and redirects users that have already registered in the database

## How to run

First **clone** this repository and **navigate** to it. 

Compile the go backend:

```bash
cd Backend
go build genLink.go
cd .. # Go back to the root folder
```

Then you just need to run:

```bash
docker-compose up
```

You are **done**! Navigate to http://localhost:8080/ to test the application

## How to uninstall/remove
Run:
```
docker-compose down --rmi all
```

[Optional] To eliminate all unused docker images on your system do:

*⚠️ Careful with this command* 
```bash
docker image prune --all
```

**Done!** Remember to delete the cloned repository as well

