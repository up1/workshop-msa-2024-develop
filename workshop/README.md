# Microservices workshop
* [The Twelve-Factor App](https://12factor.net/)
* Properties of service
  * Observability


## 1. Build and run go service (port=9000)
```
$docker compose build go
$docker compose up -d go
$docker compose ps      
NAME            IMAGE         COMMAND      SERVICE   CREATED         STATUS         PORTS
workshop-go-1   workshop-go   "/app/api"   go        3 seconds ago   Up 2 seconds   0.0.0.0:9000->9000/tcp

$docker compose logs --follow
```

## 2. Observable service with [LGTM stack from Grafana](https://grafana.com)
* Log
* Grafana (port=3000)
* Trace
* Metric

Start LGTM stack
```
$docker compose up -d lgtm
$docker compose ps
NAME              IMAGE                     COMMAND                  SERVICE   CREATED          STATUS          PORTS
workshop-lgtm-1   grafana/otel-lgtm:0.6.0   "/bin/sh -c ./run-al…"   lgtm      37 seconds ago   Up 36 seconds   0.0.0.0:3000->3000/tcp, 0.0.0.0:4317-4318->4317-4318/tcp

$docker compose logs --follow
```

Access to grafana
* http://localhost:3000/explore
  * user=admin
  * password=admin

## 3. Working with Apache Kafka
* Producer = Go
* Consumer = Java + Spring Boot 3
* Use [Custom Kafka Docker image from Binami](https://hub.docker.com/r/bitnami/kafka)

Start Kafka
```
$docker compose up -d
$docker compose ps
NAME                   IMAGE                     COMMAND                  SERVICE     CREATED          STATUS                    PORTS
kafka                  bitnami/kafka:3.6         "/opt/bitnami/script…"   kafka       13 seconds ago   Up 12 seconds (healthy)   9092/tcp, 0.0.0.0:29092->29092/tcp
workshop-collector-1   grafana/otel-lgtm:0.6.0   "/bin/sh -c ./run-al…"   collector   13 seconds ago   Up 12 seconds             0.0.0.0:3000->3000/tcp, 0.0.0.0:4317-4318->4317-4318/tcp
workshop-go-1          workshop-go               "/app/api"               go          13 seconds ago   Up 5 seconds              0.0.0.0:9000->9000/tcp

$docker compose logs --follow
```
