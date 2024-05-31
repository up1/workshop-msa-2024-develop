package com.example.demo;

import io.opentelemetry.api.common.AttributeKey;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.contrib.sampler.RuleBasedRoutingSampler;
import io.opentelemetry.sdk.autoconfigure.spi.AutoConfigurationCustomizerProvider;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class OpenTelemetryConfig {

    @Bean
    public AutoConfigurationCustomizerProvider otelCustomizer() {
        return p ->
                p.addSamplerCustomizer(
                        (fallback, config) ->
                                RuleBasedRoutingSampler.builder(SpanKind.SERVER, fallback)
                                        .drop(AttributeKey.stringKey("url.path"), "^/actuator")
                                        .build());
    }
}
