module github.com/broadinstitute/sherlock/sherlock-webhook-proxy

go 1.19

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.7.4
	github.com/broadinstitute/sherlock/go-shared v0.0.0
	github.com/broadinstitute/sherlock/sherlock-go-client v0.0.0
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-playground/webhooks/v6 v6.3.0
)

replace github.com/broadinstitute/sherlock/go-shared => ../go-shared

replace github.com/broadinstitute/sherlock/sherlock-go-client => ../sherlock-go-client

require (
	cloud.google.com/go/functions v1.13.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/cloudevents/sdk-go/v2 v2.14.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/spec v0.20.8 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-openapi/validate v0.22.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
