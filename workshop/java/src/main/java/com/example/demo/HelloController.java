package com.example.demo;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.metrics.Meter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class HelloController {

    private final Logger logger = LoggerFactory.getLogger(HelloController.class);

    @Autowired
    private OpenTelemetry openTelemetry;

    @Autowired
    HelloService helloService;

    @GetMapping("/{name}")
    public String sayHi(@PathVariable String name) {
        // Log
        logger.debug("Hello with: {}", name);

        // Metric
        Meter meter = openTelemetry.getMeter("demo");
        meter.counterBuilder("say_hi").setDescription("Count called data").build();

        // Trace
        // Create new span
        helloService.newProcess(name);


        return "Hello with " + name;
    }

}
