# go-notes

This repository is like CRUD basic operation using golang with Gofiber Framework, database PostgreSQL and completed by Redis for caching, this repository using rate-limiting to prevent blasting api calling at one time

[![golang][Golang]][golang-url][![redis][Redis]][redis-url][![postgresql][Posgtresql]][postgresql-url]

## Acknowledgements

First thing first you must be have 3 main Parts of this. For Golang this repository is using Gofiber, database using PostgreSQL and for caching (Redis) using dependicies Golang.

### .env
Here it is example of .env that used in this repository, you can add the value based on your local setup

```
DB_HOST= ''
DB_NAME= ''
DB_USER= ''
DB_PASSWORD= ''
DB_PORT= 
REDIS_PASSWORD = ''
REDIS_ADDRESS= ''
REDIS_PORT= ''
```

### Runs Program
How the code can run :

```
go mod tidy
go run main.go
```

### Note : This code develop based on Golang 1.21.3

[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[golang-url]: https://go.dev/
[Posgtresql]: https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white
[postgresql-url]: https://www.postgresql.org/
[Redis]: https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white
[redis-url]: https://redis.com/redis-for-dummies/?utm_source=google&utm_medium=cpc&utm_campaign=redis360-brand-uk-17565601660&utm_term=redis&utm_content=redis-for-dummies&gclid=CjwKCAjw3oqoBhAjEiwA_UaLtltfpPKfZtNO6GW-fjxr9yxgeJN9Xa1H8DDeyzpWOVwX1FAAFUCaqBoCpLoQAvD_BwE

