module github.com/broadinstitute/sherlock/sherlock

go 1.22.0

toolchain go1.23.0

require (
	cloud.google.com/go/bigquery v1.63.1
	cloud.google.com/go/cloudsqlconn v1.13.0
	cloud.google.com/go/iam v1.2.2
	cloud.google.com/go/kms v1.20.1
	cloud.google.com/go/resourcemanager v1.10.2
	contrib.go.opencensus.io/exporter/prometheus v0.4.2
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.0
	github.com/PagerDuty/go-pagerduty v1.8.0
	github.com/avast/retry-go/v4 v4.6.0
	github.com/broadinstitute/sherlock/go-shared v0.0.0
	github.com/coreos/go-oidc/v3 v3.11.0
	github.com/creasty/defaults v1.8.0
	github.com/dustinkirkland/golang-petname v0.0.0-20240428194347-eebcea082ee0
	github.com/gin-contrib/cors v1.7.2
	github.com/gin-contrib/secure v1.1.0
	github.com/gin-gonic/gin v1.10.0
	github.com/go-jose/go-jose/v4 v4.0.4
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/google/go-cmp v0.6.0
	github.com/google/go-github/v58 v58.0.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.7.1
	github.com/jinzhu/copier v0.4.0
	github.com/knadh/koanf v1.5.0
	github.com/microsoft/kiota-abstractions-go v1.7.0
	github.com/microsoftgraph/msgraph-sdk-go v1.51.0
	github.com/microsoftgraph/msgraph-sdk-go-core v1.2.1
	github.com/pact-foundation/pact-go/v2 v2.0.8
	github.com/rs/zerolog v1.33.0
	github.com/sanity-io/litter v1.5.5
	github.com/sethvargo/go-password v0.3.1
	github.com/slack-go/slack v0.15.0
	github.com/stretchr/testify v1.9.0
	github.com/swaggo/files/v2 v2.0.2-0.20240712141554-0590c09f83eb
	github.com/swaggo/gin-swagger v1.6.1-0.20231210095754-aa92a0ac3f26
	github.com/swaggo/swag v1.16.3
	github.com/zitadel/oidc/v3 v3.31.0
	go.opencensus.io v0.24.0
	golang.org/x/crypto v0.28.0
	golang.org/x/exp v0.0.0-20240823005443-9b4947da3948
	golang.org/x/net v0.30.0
	golang.org/x/oauth2 v0.23.0
	golang.org/x/text v0.19.0
	google.golang.org/api v0.204.0
	gorm.io/datatypes v1.2.4
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
)

replace github.com/broadinstitute/sherlock/go-shared => ../go-shared

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.10.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.5 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	cloud.google.com/go/longrunning v0.6.1 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.15.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.2 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/apache/arrow/go/v15 v15.0.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.7.1 // indirect
	github.com/bytedance/sonic v1.12.2 // indirect
	github.com/bytedance/sonic/loader v0.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cjlapao/common-go v0.0.40 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.5 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-chi/chi/v5 v5.1.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/microsoft/kiota-authentication-azure-go v1.1.0 // indirect
	github.com/microsoft/kiota-http-go v1.4.4 // indirect
	github.com/microsoft/kiota-serialization-form-go v1.0.0 // indirect
	github.com/microsoft/kiota-serialization-json-go v1.0.8 // indirect
	github.com/microsoft/kiota-serialization-multipart-go v1.0.0 // indirect
	github.com/microsoft/kiota-serialization-text-go v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/muhlemmer/gu v0.3.1 // indirect
	github.com/muhlemmer/httpforwarded v0.1.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.20.2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/prometheus/statsd_exporter v0.27.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/std-uritemplate/std-uritemplate/go v1.0.5 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	github.com/zitadel/logging v0.6.1 // indirect
	github.com/zitadel/schema v1.3.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/arch v0.9.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	google.golang.org/genproto v0.0.0-20241021214115-324edc3d5d38 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241021214115-324edc3d5d38 // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
)
