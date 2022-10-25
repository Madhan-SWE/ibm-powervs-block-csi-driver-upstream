module sigs.k8s.io/ibm-powervs-block-csi-driver

go 1.18

require (
	github.com/IBM-Cloud/power-go-client v1.2.1
	github.com/IBM/go-sdk-core/v5 v5.10.2
	github.com/IBM/platform-services-go-sdk v0.28.7
	github.com/container-storage-interface/spec v1.5.0
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/glog v1.0.0
	github.com/golang/mock v1.6.0
	github.com/kubernetes-csi/csi-test v2.2.0+incompatible
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.22.1
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f
	google.golang.org/grpc v1.50.0
	k8s.io/api v0.24.6
	k8s.io/apimachinery v0.24.6
	k8s.io/client-go v1.24.6
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.80.1
	k8s.io/kubernetes v1.24.6
	k8s.io/mount-utils v0.24.6
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
	sigs.k8s.io/controller-runtime v0.11.2
)

require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.13 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/aws/aws-sdk-go v1.38.49 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/go-logr/zapr v1.2.0 // indirect
	github.com/go-openapi/analysis v0.21.2 // indirect
	github.com/go-openapi/errors v0.20.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/runtime v0.23.0 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-openapi/validate v0.20.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.mongodb.org/mongo-driver v1.10.0 // indirect
	go.opentelemetry.io/contrib v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0 // indirect
	go.opentelemetry.io/otel v0.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp v0.20.0 // indirect
	go.opentelemetry.io/otel/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/trace v0.20.0 // indirect
	go.opentelemetry.io/proto/otlp v0.7.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.24.6 // indirect
	k8s.io/apiserver v0.24.6 // indirect
	k8s.io/component-base v0.24.6 // indirect
	k8s.io/component-helpers v0.24.6 // indirect
	k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42 // indirect
	k8s.io/kubectl v0.0.0 // indirect
	k8s.io/kubelet v0.0.0 // indirect
	k8s.io/pod-security-admission v0.0.0 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.30 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.24.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.24.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.24.6
	k8s.io/apiserver => k8s.io/apiserver v0.24.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.24.6
	k8s.io/client-go => k8s.io/client-go v0.24.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.24.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.24.6
	k8s.io/code-generator => k8s.io/code-generator v0.24.6
	k8s.io/component-base => k8s.io/component-base v0.24.6
	k8s.io/component-helpers => k8s.io/component-helpers v0.24.6
	k8s.io/controller-manager => k8s.io/controller-manager v0.24.6
	k8s.io/cri-api => k8s.io/cri-api v0.24.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.24.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.24.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.24.6
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.24.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.24.6
	k8s.io/kubectl => k8s.io/kubectl v0.24.6
	k8s.io/kubelet => k8s.io/kubelet v0.24.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.24.6
	k8s.io/metrics => k8s.io/metrics v0.24.6
	k8s.io/mount-utils => k8s.io/mount-utils v0.24.6
	k8s.io/node-api => k8s.io/node-api v0.24.6
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.24.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.24.6
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.24.6
	k8s.io/sample-controller => k8s.io/sample-controller v0.24.6
	vbom.ml/util => github.com/fvbommel/util v0.0.0-20180919145318-efcd4e0f9787
)
