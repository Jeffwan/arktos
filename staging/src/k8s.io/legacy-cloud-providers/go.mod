// This is a generated file. Do not edit directly.

module k8s.io/legacy-cloud-providers

go 1.13

require (
	cloud.google.com/go v0.38.0
	github.com/Azure/azure-sdk-for-go v35.0.0+incompatible

	github.com/Azure/go-autorest/autorest v0.9.0
	github.com/Azure/go-autorest/autorest/adal v0.5.0
	github.com/Azure/go-autorest/autorest/to v0.2.0
	github.com/Azure/go-autorest/autorest/validation v0.1.0 // indirect
	github.com/GoogleCloudPlatform/k8s-cloud-provider v0.0.0-20190822182118-27a4ced34534
	github.com/aws/aws-sdk-go v1.28.2
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.4.1 // indirect
	github.com/rubiojr/go-vhd v0.0.0-20160810183302-0bfd3b39853c
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/vmware/govmomi v0.20.3
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.8.0
	gopkg.in/gcfg.v1 v1.2.0
	gopkg.in/warnings.v0 v0.1.1 // indirect
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/cloud-provider v0.0.0
	k8s.io/csi-translation-lib v0.0.0
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
	sigs.k8s.io/yaml v1.2.0
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.34.0
	github.com/GoogleCloudPlatform/k8s-cloud-provider => github.com/GoogleCloudPlatform/k8s-cloud-provider v0.0.0-20181220005116-f8e995905100
	github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.16.26
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go v0.0.0-20160705203006-01aeca54ebda
	github.com/google/gofuzz => github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d
	github.com/hashicorp/golang-lru => github.com/hashicorp/golang-lru v0.5.0
	github.com/json-iterator/go => github.com/json-iterator/go v0.0.0-20180701071628-ab8a2e0c74be
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	github.com/prometheus/common => github.com/prometheus/common v0.0.0-20181126121408-4724e9255275
	github.com/vmware/govmomi => github.com/vmware/govmomi v0.20.1
	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200622214017-ed371f2e16b4
	golang.org/x/text => golang.org/x/text v0.3.1-0.20181227161524-e6919f6577db
	google.golang.org/api => google.golang.org/api v0.0.0-20181220000619-583d854617af
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/client-go => ../client-go
	k8s.io/cloud-provider => ../cloud-provider
	k8s.io/csi-translation-lib => ../csi-translation-lib
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/legacy-cloud-providers => ../legacy-cloud-providers
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
	sigs.k8s.io/yaml => sigs.k8s.io/yaml v1.1.0
)
