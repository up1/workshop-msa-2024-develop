# Workshop with centralize config
* [Vault](https://www.vaultproject.io/)

## Step 1 :: Create Vault Server
```
$docker compose up -d vault
$docker compose ps
NAME      IMAGE          COMMAND                  SERVICE   CREATED          STATUS         PORTS
vault     vault:1.13.3   "docker-entrypoint.s…"   vault     57 seconds ago   Up 3 seconds   0.0.0.0:8200->8200/tcp

$docker compose logs --follow
```

Access to Vault Server=http://localhost:8200/
* token=admin

## Step 2 :: Read data from Vault Server with Token
```
$USER_TOKEN=<token> go run read.go
```