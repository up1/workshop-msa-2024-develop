package com.example.demo;

import io.micrometer.core.instrument.MeterRegistry;
import io.micrometer.core.instrument.Metrics;
import io.opentelemetry.api.common.AttributeKey;
import io.opentelemetry.api.trace.SpanKind;
import io.opentelemetry.contrib.sampler.RuleBasedRoutingSampler;
import io.opentelemetry.sdk.autoconfigure.spi.AutoConfigurationCustomizerProvider;
import org.springframework.boot.autoconfigure.condition.ConditionalOnClass;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.Optional;

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

    @Bean
    @ConditionalOnClass(name = "io.opentelemetry.javaagent.OpenTelemetryAgent")
    public MeterRegistry otelRegistry() {
        Optional<MeterRegistry> otelRegistry = Metrics.globalRegistry.getRegistries().stream()
                .filter(r -> r.getClass().getName().contains("OpenTelemetryMeterRegistry"))
                .findAny();
        otelRegistry.ifPresent(Metrics.globalRegistry::remove);
        return otelRegistry.orElse(null);
    }
}
