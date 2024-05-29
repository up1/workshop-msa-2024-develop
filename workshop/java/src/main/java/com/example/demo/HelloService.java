package com.example.demo;


import io.opentelemetry.api.trace.Span;
import io.opentelemetry.instrumentation.annotations.WithSpan;
import org.springframework.stereotype.Service;

@Service
public class HelloService {

    @WithSpan(value = "new-process")
    public void newProcess(String name) {
        Span currentSpan = Span.current();
        currentSpan.addEvent("ADD EVENT TO newProcess SPAN");
        currentSpan.setAttribute("name", name);
    }

}
