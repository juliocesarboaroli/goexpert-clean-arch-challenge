# Como rodar

## Banco de dados e RabbitMQ
Utilize o docker para inicializar estes serviços, rodando `docker-compose up -d`.

## Aplicação
Utilize o comando `go run main.go wire_gen.go` e você verá um log com cada porta que pode ser acessada.

## Utilização

**Web:** execute a instrução de `get_orders.http` ou faça uma chamada usando o `curl` abaixo:
```curl
curl --request GET \
  --url http://localhost:8000/orders \
```

**gRPC:**

**GraphQL:** acessar o playground em `http://localhost:8080/` e executar