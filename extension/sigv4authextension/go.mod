module github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension

go 1.18

require (
	github.com/aws/aws-sdk-go-v2 v1.17.8
	github.com/aws/aws-sdk-go-v2/config v1.18.20
	github.com/aws/aws-sdk-go-v2/credentials v1.13.19
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.8
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/collector v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/component v0.70.1-0.20230130215412-26bd7b2bf678
	go.opentelemetry.io/collector/confmap v0.70.1-0.20230130215412-26bd7b2bf678
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.52.3
)

require (
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.26 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.7 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/featuregate v0.70.1-0.20230130215412-26bd7b2bf678 // indirect
	go.opentelemetry.io/otel v1.12.0 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.12.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v0.65.0
