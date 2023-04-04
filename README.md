# Instructions to run project from scratch
## Prerequisites
- Docker
- Golang

## Frontend
Under `vue-app/`, run
```
docker-compose up -d
```

## Backend
Under `vue-api/`

1. Start database
```
docker-compose up -d
```
2. Start backend service
```
make start
```

In the browser, visit `localhost`, and use login credential:
```
user: admin1@admin.com
password: password
```
