module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/bigipreceiver

go 1.18

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.70.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.70.0
	github.com/stretchr/testify v1.8.2
	go.opentelemetry.io/collector v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/component v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/confmap v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/consumer v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/pdata v1.0.0-rc7
	go.uber.org/multierr v1.9.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.70.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.8.3 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.70.1-0.20230130215412-26bd7b2bf678 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.38.0 // indirect
	go.opentelemetry.io/otel v1.12.0 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.12.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

retract v0.65.0
