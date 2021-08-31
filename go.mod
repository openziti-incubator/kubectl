// This is a generated file. Do not edit directly.

module github.com/openziti-incubator/kubectl

go 1.16

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/chai2010/gettext-go v1.0.2
	github.com/davecgh/go-spew v1.1.1
	github.com/daviddengcn/go-colortext v1.0.0
	github.com/docker/distribution v2.7.1+incompatible
	github.com/evanphx/json-patch v4.11.0+incompatible
	github.com/exponent-io/jsonpath v0.0.0-20210407135951-1de76d718b3f
	github.com/fatih/camelcase v1.0.0
	github.com/fvbommel/sortorder v1.0.2
	github.com/golangplus/testing v1.0.0 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/googleapis/gnostic v0.5.5
	github.com/jonboulle/clockwork v0.2.2
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de
	github.com/lithammer/dedent v1.1.0
	github.com/mitchellh/go-wordwrap v1.0.1
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/russross/blackfriday v1.6.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/cli-runtime v0.22.1
	k8s.io/client-go v1.5.2
	k8s.io/component-base v0.22.1
	k8s.io/component-helpers v0.22.1
	k8s.io/klog/v2 v2.10.0
	k8s.io/kube-openapi v0.0.0-20210817084001-7fbd8d59e5b8
	k8s.io/metrics v0.22.1
	k8s.io/utils v0.0.0-20210820185131-d34e5cb4466e
	sigs.k8s.io/kustomize/kustomize/v4 v4.3.0
	sigs.k8s.io/kustomize/kyaml v0.11.1
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20210825040442-f20796d02069
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20210825040238-74be3b88bedb
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20210825042947-c992623183f8
	k8s.io/client-go => k8s.io/client-go v0.0.0-20210825040738-3dc80a3333cd
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20210825160035-e7c2dcc7dff4
	k8s.io/component-base => k8s.io/component-base v0.0.0-20210825041339-a33683002a90
	k8s.io/component-helpers => k8s.io/component-helpers v0.0.0-20210825041451-6721137b4907
	k8s.io/metrics => k8s.io/metrics v0.0.0-20210825042829-1468ab25f472
)

replace k8s.io/kubectl => github.com/openziti-incubator/kubectl v0.22.1
