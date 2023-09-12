POC for KYC
==================

Thanks to:

https://blog.logrocket.com/building-simple-app-go-postgresql/

https://www.freecodecamp.org/news/database-migration-golang-migrate/

https://dev.to/andreidascalu/setup-go-with-vscode-in-docker-for-debugging-24ch


## Migrations


With ENV var DATABASE_URL

```
# without docker
export DATABASE_URL="postgresql://kyc:kyc@localhost:7432/kyc?sslmode=disable"
```

- Create migration
```
migrate create -ext sql -dir=db/migrations -seq create_users_table
```

- Apply migration
```
migrate -path db/migrations -database $DATABASE_URL -verbose up
```

- Undo migrations
```
migrate -path db/migrations -database $DATABASE_URL force 1
```
