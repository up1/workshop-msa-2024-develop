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
$docker compose up -d collector
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
* [OpenTelemetry auto-instrumentation and instrumentation libraries for Java](https://github.com/open-telemetry/opentelemetry-java-instrumentation)
  * version = 2.4.0
  * [Broker metric](https://github.com/open-telemetry/opentelemetry-java-instrumentation/blob/main/instrumentation/jmx-metrics/javaagent/kafka-broker.md)
  * [Kafka metrics](https://github.com/open-telemetry/opentelemetry-java-contrib/blob/main/jmx-metrics/README.md)

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

## 4. Working with Apache Kafka UI (development mode only)
* [Kafka UI](https://docs.kafka-ui.provectus.io/)

```
$docker compose up -d kafka-ui
$docker compose ps
$docker compose logs --follow
```

Access to Kafka Admin
* http://localhost:8080
  * user=admin
  * password=admin


## 5. Working with Java and Spring Boot
* [OpenTelemetry with Spring Boot](https://opentelemetry.io/docs/zero-code/java/spring-boot/)

Build jar file
```
$cd java
$mvnw clean package
$cd ..
```

Build docker image
```
$docker compose build java
```

Run container
```
$docker compose up -d java
$docker compose ps
$docker compose logs --follow
```

Call API = http://localhost:8081/api/somkiat