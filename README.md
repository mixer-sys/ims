# ims
Inventory Management Service - Сервис учета товаров

# RUN
```bash
cp env.example .env
docker-compose up
```


###### help
```bash
docker run --name db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db -p 5432:5432 -d postgres:latest


export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://user:password@127.0.0.1:5432/db?sslmode=disable
goose -dir migrations up
export PGPASSWORD="password"
psql -h 127.0.0.1 -p 5432 -U user -d db -c "select * from products;"
```