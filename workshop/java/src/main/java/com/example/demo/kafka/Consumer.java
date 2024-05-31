package com.example.demo.kafka;

import io.opentelemetry.instrumentation.annotations.WithSpan;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.common.header.Header;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import java.util.Arrays;

@Component
public class Consumer {
    @KafkaListener(topics = "newuser", groupId = "group01")
    @WithSpan(value = "Received-message")
    public void listen(ConsumerRecord<String, String> data) {
        System.out.println("========== Header =============");
        for (Header header : data.headers()) {
            System.out.println(header.key() + " : " + new String(header.value()));
        }
        System.out.println("========== Body =============");
        System.out.println("Received message: " + data.value());
    }
}
