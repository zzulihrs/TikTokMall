module github.com/tiktokmall/backend/app/agent

go 1.23.1

toolchain go1.23.2

replace (
	github.com/tiktokmall/backend/app/frontend => ../frontend
	github.com/tiktokmall/backend/common => ../../common
	github.com/tiktokmall/backend/rpc_gen => ../../rpc_gen
)

require (
	github.com/cloudwego/eino v0.3.7
	github.com/cloudwego/eino-ext/callbacks/langfuse v0.0.0-20250117061805-cd80d1780d76
	github.com/cloudwego/eino-ext/components/model/ark v0.0.0-20250117061805-cd80d1780d76
	github.com/cloudwego/eino-ext/components/retriever/redis v0.0.0-20250117061805-cd80d1780d76
	github.com/cloudwego/eino-ext/components/tool/duckduckgo v0.0.0-20250117061805-cd80d1780d76
	github.com/cloudwego/eino-ext/devops v0.0.0-20250117061805-cd80d1780d76
	github.com/cloudwego/hertz v0.9.6
	github.com/google/uuid v1.6.0
	github.com/hertz-contrib/sse v0.0.6-0.20240617114443-10a844794bf3
	github.com/redis/go-redis/v9 v9.7.0
	github.com/tiktokmall/backend/app/frontend v0.0.0-00010101000000-000000000000
	github.com/tiktokmall/backend/rpc_gen v0.0.0-00010101000000-000000000000
)

require (
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/cloudwego/configmanager v0.2.2 // indirect
	github.com/cloudwego/dynamicgo v0.5.2 // indirect
	github.com/cloudwego/fastpb v0.0.5 // indirect
	github.com/cloudwego/frugal v0.2.3 // indirect
	github.com/cloudwego/gopkg v0.1.3 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cloudwego/kitex v0.12.1 // indirect
	github.com/cloudwego/localsession v0.1.1 // indirect
	github.com/cloudwego/runtimex v0.1.0 // indirect
	github.com/cloudwego/thriftgo v0.3.18 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/hashicorp/consul/api v1.20.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/kitex-contrib/obs-opentelemetry v0.2.9 // indirect
	github.com/kitex-contrib/registry-consul v0.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/tiktokmall/backend/common v0.0.0-00010101000000-000000000000 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
)

require (
	github.com/bytedance/gopkg v0.1.1 // indirect
	github.com/bytedance/sonic v1.12.7 // indirect
	github.com/bytedance/sonic/loader v0.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/eino-ext/components/document/loader/file v0.0.0-20250116071241-3f1eaaafd49c
	github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown v0.0.0-20250116071241-3f1eaaafd49c
	github.com/cloudwego/eino-ext/components/embedding/ark v0.0.0-20250116071241-3f1eaaafd49c
	github.com/cloudwego/eino-ext/components/indexer/redis v0.0.0-20250116071241-3f1eaaafd49c
	github.com/cloudwego/eino-ext/libs/acl/langfuse v0.0.0-20250113033825-eb19b2b6b386 // indirect
	github.com/cloudwego/netpoll v0.6.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/getkin/kin-openapi v0.118.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/goph/emperror v0.17.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nikolalohinski/gonja v1.5.3 // indirect
	github.com/nyaruka/phonenumbers v1.3.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/slongfield/pyfmt v0.0.0-20220222012616-ea85ff4c361f // indirect
	github.com/tidwall/gjson v1.17.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/volcengine/volc-sdk-golang v1.0.23 // indirect
	github.com/volcengine/volcengine-go-sdk v1.0.160 // indirect
	github.com/yargevad/filepathx v1.0.0 // indirect
	golang.org/x/arch v0.12.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/sys v0.30.0 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
