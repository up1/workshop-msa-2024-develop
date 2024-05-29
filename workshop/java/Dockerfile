FROM azul/zulu-openjdk-alpine:17-jre
WORKDIR /app

COPY target/demo.jar app.jar
COPY otel/opentelemetry-javaagent.jar opentelemetry-javaagent.jar

EXPOSE 8080
ENTRYPOINT java -jar -javaagent:opentelemetry-javaagent.jar app.jar