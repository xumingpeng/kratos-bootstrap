module github.com/xumingpeng/kratos-bootstrap/bootstrap

go 1.24.0

toolchain go1.24.3

replace (
	github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.4.1

	github.com/tx7do/kratos-bootstrap/api => ../api

	github.com/tx7do/kratos-bootstrap/config => ../config
	github.com/tx7do/kratos-bootstrap/logger => ../logger
	github.com/tx7do/kratos-bootstrap/registry => ../registry
	github.com/tx7do/kratos-bootstrap/tracer => ../tracer
	github.com/tx7do/kratos-bootstrap/utils => ../utils
)

require (
	auroraride.com/aurservd v0.0.0-20250512035008-91d265103e6d
	auroraride.com/auth v0.0.0-20250512101944-2f1fd53c5dba
	github.com/go-kratos/kratos/v2 v2.8.4
	github.com/google/subcommands v1.2.0
	github.com/olekukonko/tablewriter v1.0.2
	github.com/redis/go-redis/v9 v9.8.0
	github.com/spf13/cobra v1.9.1
	github.com/tx7do/kratos-bootstrap/api v0.0.13
	github.com/tx7do/kratos-bootstrap/config v0.0.9
	github.com/tx7do/kratos-bootstrap/logger v0.0.9
	github.com/tx7do/kratos-bootstrap/registry v0.0.9
	github.com/tx7do/kratos-bootstrap/tracer v0.0.9
	github.com/tx7do/kratos-bootstrap/utils v0.1.3
	golang.org/x/tools v0.33.0
)

require (
	dario.cat/mergo v1.0.2 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.63.107 // indirect
	github.com/apolloconfig/agollo/v4 v4.4.0 // indirect
	github.com/armon/go-metrics v0.5.4 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dromara/carbon/v2 v2.6.4 // indirect
	github.com/emicklei/go-restful/v3 v3.12.2 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/fxamacker/cbor/v2 v2.8.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.9 // indirect
	github.com/go-chassis/cari v0.9.0 // indirect
	github.com/go-chassis/foundation v0.4.0 // indirect
	github.com/go-chassis/openlog v1.1.3 // indirect
	github.com/go-chassis/sc-client v0.7.0 // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/go-json-experiment/json v0.0.0-20250417205406-170dfdcf87d1 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/config/kubernetes/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/log/fluent/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/servicecomb/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20250429074618-c82f7957223f // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.26.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/go-zookeeper/zk v1.0.4 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.9 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	github.com/hashicorp/consul/api v1.32.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.4 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/serf v0.10.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.13-0.20220915233716-71ac16282d12 // indirect
	github.com/karlseguin/ccache/v2 v2.0.8 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/labstack/echo/v4 v4.13.3 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nacos-group/nacos-sdk-go v1.1.5 // indirect
	github.com/olekukonko/errors v1.1.0 // indirect
	github.com/olekukonko/ll v0.0.6-0.20250511102614-9564773e9d27 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.4 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sagikazarmark/locafero v0.9.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sony/sonyflake v1.2.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.8.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/spf13/viper v1.20.1 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.11 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xuri/efp v0.0.1 // indirect
	github.com/xuri/excelize/v2 v2.9.0 // indirect
	github.com/xuri/nfp v0.0.1 // indirect
	go.etcd.io/bbolt v1.4.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.21 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.21 // indirect
	go.etcd.io/etcd/client/v3 v3.5.21 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.35.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	go.opentelemetry.io/proto/otlp v1.6.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/exp v0.0.0-20250506013437-ce4c2cf36ca6 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250505200425-f936aa4a68b2 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250505200425-f936aa4a68b2 // indirect
	google.golang.org/grpc v1.72.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.33.0 // indirect
	k8s.io/apimachinery v0.33.0 // indirect
	k8s.io/client-go v0.33.0 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff // indirect
	k8s.io/utils v0.0.0-20250502105355-0f33e8f1c979 // indirect
	resty.dev/v3 v3.0.0-beta.2 // indirect
	sigs.k8s.io/json v0.0.0-20241014173422-cfa47c3a1cc8 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.7.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
