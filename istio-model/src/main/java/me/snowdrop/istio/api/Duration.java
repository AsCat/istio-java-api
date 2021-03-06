/*
 * *
 *  * Copyright (C) 2018 Red Hat, Inc.
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *         http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 *
 */
package me.snowdrop.istio.api;

import java.io.IOException;
import java.io.Serializable;
import java.util.HashMap;
import java.util.Map;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.ObjectCodec;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import io.sundr.builder.annotations.Buildable;
import lombok.EqualsAndHashCode;
import lombok.ToString;
import org.joda.time.Period;
import org.joda.time.Seconds;
import org.joda.time.format.PeriodFormatter;
import org.joda.time.format.PeriodFormatterBuilder;

/**
 * @author <a href="claprun@redhat.com">Christophe Laprun</a>
 */
@JsonDeserialize(using = Duration.Deserializer.class)
@JsonSerialize(using = Duration.Serializer.class)
@ToString
@EqualsAndHashCode
public class Duration implements Serializable {
    private Map<String, Object> additionalProperties = new HashMap<>();
    private Integer nanos;
    private Long seconds;
    private final static PeriodFormatter FORMATTER = new PeriodFormatterBuilder()
            .printZeroNever()
            .appendHours()
            .appendSuffix("h")
            .appendMinutes()
            .appendSuffix("m")
            .appendSeconds()
            .appendSuffix("s")
            .appendMillis()
            .appendSuffix("ms")
            .toFormatter();

    /**
     * No args constructor for serialization
     */
    public Duration() {
    }

    @Buildable(generateBuilderPackage = true, builderPackage = "io.fabric8.kubernetes.api.builder", editableEnabled = false, validationEnabled = true)
    public Duration(Integer nanos, Long seconds) {
        this.nanos = nanos;
        this.seconds = seconds;
    }

    public Integer getNanos() {
        return nanos;
    }

    public Long getSeconds() {
        return seconds;
    }

    @JsonAnyGetter
    public Map<String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    @JsonAnySetter
    public void setAdditionalProperty(String name, Object value) {
        this.additionalProperties.put(name, value);
    }

    public static class Deserializer extends JsonDeserializer<Duration> {

        @Override
        public Duration deserialize(JsonParser parser, DeserializationContext ctxt) throws IOException, JsonProcessingException {
            ObjectCodec oc = parser.getCodec();
            JsonNode node = oc.readTree(parser);
            final Period period = FORMATTER.parsePeriod(node.asText());
            return new Duration(0, (long) period.toStandardSeconds().getSeconds());
        }
    }

    public static class Serializer extends JsonSerializer<Duration> {
        @Override
        public void serialize(Duration value, JsonGenerator gen, SerializerProvider serializers) throws IOException {
            gen.writeString(FORMATTER.print(Seconds.seconds(value.seconds.intValue())));
        }
    }
}
