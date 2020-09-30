// Code generated for package vfs by go-bindata DO NOT EDIT. (@generated)
// sources:
// charts/README-helm3.md
// charts/README.md
// charts/UPDATING-CHARTS.md
// charts/base/Chart.yaml
// charts/base/NOTES.txt
// charts/base/crds/crd-all.gen.yaml
// charts/base/crds/crd-operator.yaml
// charts/base/files/gen-istio-cluster.yaml
// charts/base/kustomization.yaml
// charts/base/templates/clusterrole.yaml
// charts/base/templates/clusterrolebinding.yaml
// charts/base/templates/crds.yaml
// charts/base/templates/endpoints.yaml
// charts/base/templates/role.yaml
// charts/base/templates/rolebinding.yaml
// charts/base/templates/serviceaccount.yaml
// charts/base/templates/services.yaml
// charts/base/templates/validatingwebhookconfiguration.yaml
// charts/base/values.yaml
// charts/gateways/istio-egress/Chart.yaml
// charts/gateways/istio-egress/NOTES.txt
// charts/gateways/istio-egress/templates/_affinity.tpl
// charts/gateways/istio-egress/templates/_helpers.tpl
// charts/gateways/istio-egress/templates/autoscale.yaml
// charts/gateways/istio-egress/templates/deployment.yaml
// charts/gateways/istio-egress/templates/poddisruptionbudget.yaml
// charts/gateways/istio-egress/templates/preconfigured.yaml
// charts/gateways/istio-egress/templates/role.yaml
// charts/gateways/istio-egress/templates/rolebindings.yaml
// charts/gateways/istio-egress/templates/service.yaml
// charts/gateways/istio-egress/templates/serviceaccount.yaml
// charts/gateways/istio-egress/values.yaml
// charts/gateways/istio-ingress/Chart.yaml
// charts/gateways/istio-ingress/NOTES.txt
// charts/gateways/istio-ingress/templates/_affinity.tpl
// charts/gateways/istio-ingress/templates/autoscale.yaml
// charts/gateways/istio-ingress/templates/deployment.yaml
// charts/gateways/istio-ingress/templates/meshexpansion.yaml
// charts/gateways/istio-ingress/templates/poddisruptionbudget.yaml
// charts/gateways/istio-ingress/templates/preconfigured.yaml
// charts/gateways/istio-ingress/templates/role.yaml
// charts/gateways/istio-ingress/templates/rolebindings.yaml
// charts/gateways/istio-ingress/templates/service.yaml
// charts/gateways/istio-ingress/templates/serviceaccount.yaml
// charts/gateways/istio-ingress/values.yaml
// charts/global.yaml
// charts/istio-cni/Chart.yaml
// charts/istio-cni/templates/clusterrole.yaml
// charts/istio-cni/templates/clusterrolebinding.yaml
// charts/istio-cni/templates/configmap-cni.yaml
// charts/istio-cni/templates/daemonset.yaml
// charts/istio-cni/templates/serviceaccount.yaml
// charts/istio-cni/values.yaml
// charts/istio-control/istio-discovery/Chart.yaml
// charts/istio-control/istio-discovery/NOTES.txt
// charts/istio-control/istio-discovery/files/gen-istio.yaml
// charts/istio-control/istio-discovery/files/injection-template.yaml
// charts/istio-control/istio-discovery/kustomization.yaml
// charts/istio-control/istio-discovery/templates/autoscale.yaml
// charts/istio-control/istio-discovery/templates/configmap-jwks.yaml
// charts/istio-control/istio-discovery/templates/configmap.yaml
// charts/istio-control/istio-discovery/templates/deployment.yaml
// charts/istio-control/istio-discovery/templates/istiod-injector-configmap.yaml
// charts/istio-control/istio-discovery/templates/mutatingwebhook.yaml
// charts/istio-control/istio-discovery/templates/poddisruptionbudget.yaml
// charts/istio-control/istio-discovery/templates/service.yaml
// charts/istio-control/istio-discovery/templates/telemetryv2_1.6.yaml
// charts/istio-control/istio-discovery/templates/telemetryv2_1.7.yaml
// charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml
// charts/istio-control/istio-discovery/values.yaml
// charts/istio-operator/Chart.yaml
// charts/istio-operator/crds/crd-operator.yaml
// charts/istio-operator/files/gen-operator.yaml
// charts/istio-operator/templates/clusterrole.yaml
// charts/istio-operator/templates/clusterrole_binding.yaml
// charts/istio-operator/templates/crds.yaml
// charts/istio-operator/templates/deployment.yaml
// charts/istio-operator/templates/namespace.yaml
// charts/istio-operator/templates/service.yaml
// charts/istio-operator/templates/service_account.yaml
// charts/istio-operator/values.yaml
// charts/istiocoredns/Chart.yaml
// charts/istiocoredns/templates/_affinity.tpl
// charts/istiocoredns/templates/autoscale.yaml
// charts/istiocoredns/templates/clusterrole.yaml
// charts/istiocoredns/templates/clusterrolebinding.yaml
// charts/istiocoredns/templates/configmap.yaml
// charts/istiocoredns/templates/deployment.yaml
// charts/istiocoredns/templates/poddisruptionbudget.yaml
// charts/istiocoredns/templates/service.yaml
// charts/istiocoredns/templates/serviceaccount.yaml
// charts/istiocoredns/values.yaml
// charts/istiod-remote/Chart.yaml
// charts/istiod-remote/NOTES.txt
// charts/istiod-remote/files/gen-istiod-remote.yaml
// charts/istiod-remote/files/injection-template.yaml
// charts/istiod-remote/kustomization.yaml
// charts/istiod-remote/templates/configmap.yaml
// charts/istiod-remote/templates/istiod-injector-configmap.yaml
// charts/istiod-remote/templates/mutatingwebhook.yaml
// charts/istiod-remote/templates/telemetryv2_1.6.yaml
// charts/istiod-remote/templates/telemetryv2_1.7.yaml
// charts/istiod-remote/templates/telemetryv2_1.8.yaml
// charts/istiod-remote/values.yaml
// profiles/default.yaml
// profiles/demo.yaml
// profiles/empty.yaml
// profiles/minimal.yaml
// profiles/preview.yaml
// profiles/remote.yaml
package vfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _chartsReadmeHelm3Md = []byte(`# Helm3 support

## Install

The manifests/ templates support both helm2 and helm3. Please do not introduce helm3-specific changes, many
users are still using helm2 and the operator is currently using the helm2 code to generate.

To install in helm3 you must first create a namespace that you wish to install the following charts to:

`+"`"+``+"`"+``+"`"+`shell script
 kubectl create namespace istio-system
`+"`"+``+"`"+``+"`"+`

We have few charts:

- 'base' creates cluster-wide CRDs, cluster bindings, cluster resources and the istio-system namespace.
  It is possible to customize the namespace, but not recommended.

`+"`"+``+"`"+``+"`"+`shell script
 helm3 install istio-base -n istio-system manifests/charts/base
`+"`"+``+"`"+``+"`"+`

- 'istio-control/istio-discovery' installs a revision of istiod.  You can install it multiple times, with different revisions.
TODO: remove the need to pass -n istio-system

`+"`"+``+"`"+``+"`"+`shell script
 helm3 install -n istio-system istio-17 manifests/charts/istio-control/istio-discovery

 helm3 install -n istio-system istio-canary manifests/charts/istio-control/istio-discovery \
    -f manifests/charts/global.yaml  --set revision=canary --set clusterResources=false

 helm3 install -n istio-system istio-mytest manifests/charts/istio-control/istio-discovery \
    -f manifests/charts/global.yaml  --set revision=mytest --set clusterResources=false
`+"`"+``+"`"+``+"`"+`

- 'ingress' to install a Gateway

Helm3 requires namespaces to be created explicitly, currently we don't support installing multiple gateways in same
namespace - nor is it a good practice. Ingress secrets and access should be separated from control plane.

`+"`"+``+"`"+``+"`"+`shell script
    helm3 install -n istio-system istio-ingress manifests/charts/gateways/istio-ingress -f manifests/charts/global.yaml

    kubectl create ns istio-ingress-canary
    helm3 install -n istio-ingress-canary istio-ingress-canary manifests/charts/gateways/istio-ingress \
      -f manifests/charts/global.yaml --set revision=canary
`+"`"+``+"`"+``+"`"+`

## Namespaces

One of the major changes in helm3 is that the 'release namespace' is no longer created.
That means the first step can't use "-n istio-system" flag, instead we use .global.istioNamespace.
It is possible - but not supported - to install multiple versions of the global, for example in
multi-tenant cases. Each namespace will have a separate root CA, if the built-in CA is used.

TODO: apply the same change for discovery and ingress, so passing -n is no longer needed.

`)

func chartsReadmeHelm3MdBytes() ([]byte, error) {
	return _chartsReadmeHelm3Md, nil
}

func chartsReadmeHelm3Md() (*asset, error) {
	bytes, err := chartsReadmeHelm3MdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/README-helm3.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsReadmeMd = []byte(`# Istio Installer

Note: If making any changes to the charts or values.yaml in this dir, first read [UPDATING-CHARTS.md](UPDATING-CHARTS.md)

Istio installer is a modular, 'a-la-carte' installer for Istio. It is based on a
fork of the Istio helm templates, refactored to increase modularity and isolation.

Goals:
- Improve upgrade experience: users should be able to gradually roll upgrades, with proper
canary deployments for Istio components. It should be possible to deploy a new version while keeping the
stable version in place and gradually migrate apps to the new version.

- More flexibility: the new installer allows multiple 'environments', allowing applications to select
a set of control plane settings and components. While the entire mesh respects the same APIs and config,
apps may target different 'environments' which contain different instances and variants of Istio.

- Better security: separate Istio components reside in different namespaces, allowing different teams or
roles to manage different parts of Istio. For example, a security team would maintain the
root CA and policy, a telemetry team may only have access to Prometheus,
and a different team may maintain the control plane components (which are highly security sensitive).

The install is organized in 'environments' - each environment consists of a set of components
in different namespaces that are configured to work together. Regardless of 'environment',
workloads can talk with each other and obey the Istio configuration resources, but each environment
can use different Istio versions and different configuration defaults.

`+"`"+`istioctl kube-inject`+"`"+` or the automatic sidecar injector are used to select the environment.
In the case of the sidecar injector, the namespace label `+"`"+`istio-env: <NAME_OF_ENV>`+"`"+` is used instead
of the conventional `+"`"+`istio-injected: true`+"`"+`. The name of the environment is defined as the namespace
where the corresponding control plane components (config, discovery, auto-injection) are running.
In the examples below, by default this is the `+"`"+`istio-control`+"`"+` namespace. Pod annotations can also
be used to select a different 'environment'.

## Installing

The new installer is intended to be modular and very explicit about what is installed. It has
far more steps than the Istio installer - but each step is smaller and focused on a specific
feature, and can be performed by different people/teams at different times.

It is strongly recommended that different namespaces are used, with different service accounts.
In particular access to the security-critical production components (root CA, policy, control)
should be locked down and restricted.  The new installer allows multiple instances of
policy/control/telemetry - so testing/staging of new settings and versions can be performed
by a different role than the prod version.

The intended users of this repo are users running Istio in production who want to select, tune
and understand each binary that gets deployed, and select which combination to use.

Note: each component can be installed in parallel with an existing Istio 1.0 or 1.1 install in
`+"`"+`istio-system`+"`"+`. The new components will not interfere with existing apps, but can interoperate
and it is possible to gradually move apps from Istio 1.0/1.1 to the new environments and
across environments ( for example canary -> prod )

Note: there are still some cluster roles that may need to be fixed, most likely cluster permissions
will need to move to the security component.

## Everything is Optional

Each component in the new installer is optional. Users can install the component defined in the new installer,
use the equivalent component in `+"`"+`istio-system`+"`"+`, configured with the official installer, or use a different
version or implementation.

For example you may use your own Prometheus and Grafana installs, or you may use a specialized/custom
certificate provisioning tool, or use components that are centrally managed and running in a different cluster.

This is a work in progress - building on top of the multi-cluster installer.

As an extreme, the goal is to be possible to run Istio workloads in a cluster without installing any Istio component
in that cluster. Currently the minimum we require is the security provider (node agent or citadel).

### Install Istio CRDs

This is the first step of the install. Please do not remove or edit any CRD - config currently requires
all CRDs to be present. On each upgrade it is recommended to reapply the file, to make sure
you get all CRDs.  CRDs are separated by release and by component type in the CRD directory.

Istio has strong integration with certmanager.  Some operators may want to keep their current certmanager
CRDs in place and not have Istio modify them.  In this case, it is necessary to apply CRD files individually.

`+"`"+``+"`"+``+"`"+`bash
kubectl apply -k github.com/istio/installer/base
`+"`"+``+"`"+``+"`"+`

or

`+"`"+``+"`"+``+"`"+`bash
kubectl apply -f base/files
`+"`"+``+"`"+``+"`"+`

### Install Istio-CNI

This is an optional step - CNI must run in a dedicated namespace, it is a 'singleton' and extremely
security sensitive. Access to the CNI namespace must be highly restricted.

**NOTE:** The environment variable `+"`"+`ISTIO_CLUSTER_ISGKE`+"`"+` is assumed to be set to `+"`"+`true`+"`"+` if the cluster
is a GKE cluster.

`+"`"+``+"`"+``+"`"+`bash
ISTIO_CNI_ARGS=
# TODO: What k8s data can we use for this check for whether GKE?
if [[ "${ISTIO_CLUSTER_ISGKE}" == "true" ]]; then
    ISTIO_CNI_ARGS="--set cni.cniBinDir=/home/kubernetes/bin"
fi
iop kube-system istio-cni $IBASE/istio-cni/ ${ISTIO_CNI_ARGS}
`+"`"+``+"`"+``+"`"+`

TODO. It is possible to add Istio-CNI later, and gradually migrate.

### Install Control plane

This can run in any cluster. A mesh should have at least one cluster should run Pilot or equivalent XDS server,
and it is recommended to have Pilot running in each region and in multiple availability zones for multi cluster.

`+"`"+``+"`"+``+"`"+`bash
iop istio-control istio-discovery $IBASE/istio-control/istio-discovery \
            --set global.istioNamespace=istio-system

# Second istio-discovery, using master version of istio
TAG=latest HUB=gcr.io/istio-testing iop istio-master istio-discovery-master $IBASE/istio-control/istio-discovery \
            --set policy.enable=false \
            --set global.istioNamespace=istio-master
`+"`"+``+"`"+``+"`"+`

### Gateways

A cluster may use multiple Gateways, each with a different load balancer IP, domains and certificates.

Since the domain certificates are stored in the gateway namespace, it is recommended to keep each
gateway in a dedicated namespace and restrict access.

For large-scale gateways it is optionally possible to use a dedicated pilot in the gateway namespace.

### Additional test templates

A number of helm test setups are general-purpose and should be installable in any cluster, to confirm
Istio works properly and allow testing the specific install.
`)

func chartsReadmeMdBytes() ([]byte, error) {
	return _chartsReadmeMd, nil
}

func chartsReadmeMd() (*asset, error) {
	bytes, err := chartsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/README.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsUpdatingChartsMd = []byte(`# Upating charts and values.yaml

The charts in the `+"`"+`manifests`+"`"+` directory are used in istioctl to generate an installation manifest. The configuration
settings contained in values.yaml files and passed through the CLI are validated against a
[schema](../../operator/pkg/apis/istio/v1alpha1/values_types.proto).
Whenever making changes in the charts, it's important to follow the below steps.

## Step 0. Check that any schema change really belongs in values.yaml

Is this a new parameter being added? If not, go to the next step.
Dynamic, runtime config that is used to configure Istio components should go into the
[MeshConfig API](https://github.com/istio/api/blob/master/mesh/v1alpha1/config.proto). Values.yaml is being deprecated and adding
to it is discouraged. MeshConfig is the official API which follows API management practices and is dynamic
(does not require component restarts).
Exceptions to this rule are configuration items that affect K8s level settings (resources, mounts etc.)

## Step 1. Make changes in charts and values.yaml in `+"`"+`manifests`+"`"+` directory

## Step 2. Make corresponding values changes in [../profiles/default.yaml](../profiles/default.yaml)

The values.yaml in `+"`"+`manifests`+"`"+` are only used for direct Helm based installations, which is being deprecated.
If any values.yaml changes are being made, the same changes must be made in the `+"`"+`manifests/profiles/default.yaml`+"`"+`
file, which must be in sync with the Helm values in `+"`"+`manifests`+"`"+`.

## Step 3. Update the validation schema

Istioctl uses a [schema](../../operator/pkg/apis/istio/v1alpha1/values_types.proto) to validate the values. Any changes to
the schema must be added here, otherwise istioctl users will see errors.
Once the schema file is updated, run:

`+"`"+``+"`"+``+"`"+`bash
$ make operator-proto
`+"`"+``+"`"+``+"`"+`

This will regenerate the Go structs used for schema validation.

## Step 4. Update the generated manifests

Tests of istioctl use the auto-generated manifests to ensure that the istioctl binary has the correct version of the charts.
These manifests can be found in [gen-istio.yaml](../charts/istio-control/istio-discovery/files/gen-istio.yaml).
To regenerate the manifests, run:

`+"`"+``+"`"+``+"`"+`bash
$ make gen
`+"`"+``+"`"+``+"`"+`

## Step 5. Update golden files

The new charts/values will likely produce different installation manifests. Unit tests that expect a certain command
output will fail for this reason. To update the golden output files, run:

`+"`"+``+"`"+``+"`"+`bash
$ make refresh-goldens
`+"`"+``+"`"+``+"`"+`

This will generate git diffs in the golden output files. Check that the changes are what you expect.

## Step 6. Create a PR using outputs from Steps 1 to 5

Your PR should pass all the checks if you followed these steps.
`)

func chartsUpdatingChartsMdBytes() ([]byte, error) {
	return _chartsUpdatingChartsMd, nil
}

func chartsUpdatingChartsMd() (*asset, error) {
	bytes, err := chartsUpdatingChartsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/UPDATING-CHARTS.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseChartYaml = []byte(`apiVersion: v1
name: base
version: 1.1.0
tillerVersion: ">=2.7.2"
description: Helm chart for deploying Istio cluster resources and CRDs
keywords:
  - istio
sources:
  - http://github.com/istio/istio
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsBaseChartYamlBytes() ([]byte, error) {
	return _chartsBaseChartYaml, nil
}

func chartsBaseChartYaml() (*asset, error) {
	bytes, err := chartsBaseChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseNotesTxt = []byte(`Installs Istio cluster resources: CRDs, cluster bindings and associated service accounts.
`)

func chartsBaseNotesTxtBytes() ([]byte, error) {
	return _chartsBaseNotesTxt, nil
}

func chartsBaseNotesTxt() (*asset, error) {
	bytes, err := chartsBaseNotesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/NOTES.txt", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseCrdsCrdAllGenYaml = []byte(`# DO NOT EDIT - Generated by Cue OpenAPI generator based on Istio APIs.
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: destinationrules.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.host
    description: The name of a service from the service registry
    name: Host
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: DestinationRule
    listKind: DestinationRuleList
    plural: destinationrules
    shortNames:
    - dr
    singular: destinationrule
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting load balancing, outlier detection,
            etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html'
          properties:
            exportTo:
              description: A list of namespaces to which this destination rule is
                exported.
              items:
                format: string
                type: string
              type: array
            host:
              description: The name of a service from the service registry.
              format: string
              type: string
            subsets:
              items:
                properties:
                  labels:
                    additionalProperties:
                      format: string
                      type: string
                    type: object
                  name:
                    description: Name of the subset.
                    format: string
                    type: string
                  trafficPolicy:
                    description: Traffic policies that apply to this subset.
                    properties:
                      connectionPool:
                        properties:
                          http:
                            description: HTTP connection pool settings.
                            properties:
                              h2UpgradePolicy:
                                description: Specify if http1.1 connection should
                                  be upgraded to http2 for the associated destination.
                                enum:
                                - DEFAULT
                                - DO_NOT_UPGRADE
                                - UPGRADE
                                type: string
                              http1MaxPendingRequests:
                                description: Maximum number of pending HTTP requests
                                  to a destination.
                                format: int32
                                type: integer
                              http2MaxRequests:
                                description: Maximum number of requests to a backend.
                                format: int32
                                type: integer
                              idleTimeout:
                                description: The idle timeout for upstream connection
                                  pool connections.
                                type: string
                              maxRequestsPerConnection:
                                description: Maximum number of requests per connection
                                  to a backend.
                                format: int32
                                type: integer
                              maxRetries:
                                format: int32
                                type: integer
                              useClientProtocol:
                                description: If set to true, client protocol will
                                  be preserved while initiating connection to backend.
                                type: boolean
                            type: object
                          tcp:
                            description: Settings common to both HTTP and TCP upstream
                              connections.
                            properties:
                              connectTimeout:
                                description: TCP connection timeout.
                                type: string
                              maxConnections:
                                description: Maximum number of HTTP1 /TCP connections
                                  to a destination host.
                                format: int32
                                type: integer
                              tcpKeepalive:
                                description: If set then set SO_KEEPALIVE on the socket
                                  to enable TCP Keepalives.
                                properties:
                                  interval:
                                    description: The time duration between keep-alive
                                      probes.
                                    type: string
                                  probes:
                                    type: integer
                                  time:
                                    type: string
                                type: object
                            type: object
                        type: object
                      loadBalancer:
                        description: Settings controlling the load balancer algorithms.
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - simple
                            - properties:
                                consistentHash:
                                  oneOf:
                                  - not:
                                      anyOf:
                                      - required:
                                        - httpHeaderName
                                      - required:
                                        - httpCookie
                                      - required:
                                        - useSourceIp
                                      - required:
                                        - httpQueryParameterName
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              required:
                              - consistentHash
                        - required:
                          - simple
                        - properties:
                            consistentHash:
                              oneOf:
                              - not:
                                  anyOf:
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              - required:
                                - httpHeaderName
                              - required:
                                - httpCookie
                              - required:
                                - useSourceIp
                              - required:
                                - httpQueryParameterName
                          required:
                          - consistentHash
                        properties:
                          consistentHash:
                            properties:
                              httpCookie:
                                description: Hash based on HTTP cookie.
                                properties:
                                  name:
                                    description: Name of the cookie.
                                    format: string
                                    type: string
                                  path:
                                    description: Path to set for the cookie.
                                    format: string
                                    type: string
                                  ttl:
                                    description: Lifetime of the cookie.
                                    type: string
                                type: object
                              httpHeaderName:
                                description: Hash based on a specific HTTP header.
                                format: string
                                type: string
                              httpQueryParameterName:
                                description: Hash based on a specific HTTP query parameter.
                                format: string
                                type: string
                              minimumRingSize:
                                type: integer
                              useSourceIp:
                                description: Hash based on the source IP address.
                                type: boolean
                            type: object
                          localityLbSetting:
                            properties:
                              distribute:
                                description: 'Optional: only one of distribute or
                                  failover can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating locality, '/' separated,
                                        e.g.
                                      format: string
                                      type: string
                                    to:
                                      additionalProperties:
                                        type: integer
                                      description: Map of upstream localities to traffic
                                        distribution weights.
                                      type: object
                                  type: object
                                type: array
                              enabled:
                                description: enable locality load balancing, this
                                  is DestinationRule-level and will override mesh
                                  wide settings in entirety.
                                nullable: true
                                type: boolean
                              failover:
                                description: 'Optional: only failover or distribute
                                  can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating region.
                                      format: string
                                      type: string
                                    to:
                                      format: string
                                      type: string
                                  type: object
                                type: array
                            type: object
                          simple:
                            enum:
                            - ROUND_ROBIN
                            - LEAST_CONN
                            - RANDOM
                            - PASSTHROUGH
			    - PREDICTIVE_LEAST_CONN
                            type: string
                        type: object
                      outlierDetection:
                        properties:
                          baseEjectionTime:
                            description: Minimum ejection duration.
                            type: string
                          consecutive5xxErrors:
                            description: Number of 5xx errors before a host is ejected
                              from the connection pool.
                            nullable: true
                            type: integer
                          consecutiveErrors:
                            format: int32
                            type: integer
                          consecutiveGatewayErrors:
                            description: Number of gateway errors before a host is
                              ejected from the connection pool.
                            nullable: true
                            type: integer
                          interval:
                            description: Time interval between ejection sweep analysis.
                            type: string
                          maxEjectionPercent:
                            format: int32
                            type: integer
                          minHealthPercent:
                            format: int32
                            type: integer
                        type: object
                      portLevelSettings:
                        description: Traffic policies specific to individual ports.
                        items:
                          properties:
                            connectionPool:
                              properties:
                                http:
                                  description: HTTP connection pool settings.
                                  properties:
                                    h2UpgradePolicy:
                                      description: Specify if http1.1 connection should
                                        be upgraded to http2 for the associated destination.
                                      enum:
                                      - DEFAULT
                                      - DO_NOT_UPGRADE
                                      - UPGRADE
                                      type: string
                                    http1MaxPendingRequests:
                                      description: Maximum number of pending HTTP
                                        requests to a destination.
                                      format: int32
                                      type: integer
                                    http2MaxRequests:
                                      description: Maximum number of requests to a
                                        backend.
                                      format: int32
                                      type: integer
                                    idleTimeout:
                                      description: The idle timeout for upstream connection
                                        pool connections.
                                      type: string
                                    maxRequestsPerConnection:
                                      description: Maximum number of requests per
                                        connection to a backend.
                                      format: int32
                                      type: integer
                                    maxRetries:
                                      format: int32
                                      type: integer
                                    useClientProtocol:
                                      description: If set to true, client protocol
                                        will be preserved while initiating connection
                                        to backend.
                                      type: boolean
                                  type: object
                                tcp:
                                  description: Settings common to both HTTP and TCP
                                    upstream connections.
                                  properties:
                                    connectTimeout:
                                      description: TCP connection timeout.
                                      type: string
                                    maxConnections:
                                      description: Maximum number of HTTP1 /TCP connections
                                        to a destination host.
                                      format: int32
                                      type: integer
                                    tcpKeepalive:
                                      description: If set then set SO_KEEPALIVE on
                                        the socket to enable TCP Keepalives.
                                      properties:
                                        interval:
                                          description: The time duration between keep-alive
                                            probes.
                                          type: string
                                        probes:
                                          type: integer
                                        time:
                                          type: string
                                      type: object
                                  type: object
                              type: object
                            loadBalancer:
                              description: Settings controlling the load balancer
                                algorithms.
                              oneOf:
                              - not:
                                  anyOf:
                                  - required:
                                    - simple
                                  - properties:
                                      consistentHash:
                                        oneOf:
                                        - not:
                                            anyOf:
                                            - required:
                                              - httpHeaderName
                                            - required:
                                              - httpCookie
                                            - required:
                                              - useSourceIp
                                            - required:
                                              - httpQueryParameterName
                                        - required:
                                          - httpHeaderName
                                        - required:
                                          - httpCookie
                                        - required:
                                          - useSourceIp
                                        - required:
                                          - httpQueryParameterName
                                    required:
                                    - consistentHash
                              - required:
                                - simple
                              - properties:
                                  consistentHash:
                                    oneOf:
                                    - not:
                                        anyOf:
                                        - required:
                                          - httpHeaderName
                                        - required:
                                          - httpCookie
                                        - required:
                                          - useSourceIp
                                        - required:
                                          - httpQueryParameterName
                                    - required:
                                      - httpHeaderName
                                    - required:
                                      - httpCookie
                                    - required:
                                      - useSourceIp
                                    - required:
                                      - httpQueryParameterName
                                required:
                                - consistentHash
                              properties:
                                consistentHash:
                                  properties:
                                    httpCookie:
                                      description: Hash based on HTTP cookie.
                                      properties:
                                        name:
                                          description: Name of the cookie.
                                          format: string
                                          type: string
                                        path:
                                          description: Path to set for the cookie.
                                          format: string
                                          type: string
                                        ttl:
                                          description: Lifetime of the cookie.
                                          type: string
                                      type: object
                                    httpHeaderName:
                                      description: Hash based on a specific HTTP header.
                                      format: string
                                      type: string
                                    httpQueryParameterName:
                                      description: Hash based on a specific HTTP query
                                        parameter.
                                      format: string
                                      type: string
                                    minimumRingSize:
                                      type: integer
                                    useSourceIp:
                                      description: Hash based on the source IP address.
                                      type: boolean
                                  type: object
                                localityLbSetting:
                                  properties:
                                    distribute:
                                      description: 'Optional: only one of distribute
                                        or failover can be set.'
                                      items:
                                        properties:
                                          from:
                                            description: Originating locality, '/'
                                              separated, e.g.
                                            format: string
                                            type: string
                                          to:
                                            additionalProperties:
                                              type: integer
                                            description: Map of upstream localities
                                              to traffic distribution weights.
                                            type: object
                                        type: object
                                      type: array
                                    enabled:
                                      description: enable locality load balancing,
                                        this is DestinationRule-level and will override
                                        mesh wide settings in entirety.
                                      nullable: true
                                      type: boolean
                                    failover:
                                      description: 'Optional: only failover or distribute
                                        can be set.'
                                      items:
                                        properties:
                                          from:
                                            description: Originating region.
                                            format: string
                                            type: string
                                          to:
                                            format: string
                                            type: string
                                        type: object
                                      type: array
                                  type: object
                                simple:
                                  enum:
                                  - ROUND_ROBIN
                                  - LEAST_CONN
                                  - RANDOM
                                  - PASSTHROUGH
				  - PREDICTIVE_LEAST_CONN
                                  type: string
                              type: object
                            outlierDetection:
                              properties:
                                baseEjectionTime:
                                  description: Minimum ejection duration.
                                  type: string
                                consecutive5xxErrors:
                                  description: Number of 5xx errors before a host
                                    is ejected from the connection pool.
                                  nullable: true
                                  type: integer
                                consecutiveErrors:
                                  format: int32
                                  type: integer
                                consecutiveGatewayErrors:
                                  description: Number of gateway errors before a host
                                    is ejected from the connection pool.
                                  nullable: true
                                  type: integer
                                interval:
                                  description: Time interval between ejection sweep
                                    analysis.
                                  type: string
                                maxEjectionPercent:
                                  format: int32
                                  type: integer
                                minHealthPercent:
                                  format: int32
                                  type: integer
                              type: object
                            port:
                              properties:
                                number:
                                  type: integer
                              type: object
                            tls:
                              description: TLS related settings for connections to
                                the upstream service.
                              properties:
                                caCertificates:
                                  format: string
                                  type: string
                                clientCertificate:
                                  description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                                  format: string
                                  type: string
                                credentialName:
                                  format: string
                                  type: string
                                mode:
                                  enum:
                                  - DISABLE
                                  - SIMPLE
                                  - MUTUAL
                                  - ISTIO_MUTUAL
                                  type: string
                                privateKey:
                                  description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                                  format: string
                                  type: string
                                sni:
                                  description: SNI string to present to the server
                                    during TLS handshake.
                                  format: string
                                  type: string
                                subjectAltNames:
                                  items:
                                    format: string
                                    type: string
                                  type: array
                              type: object
                          type: object
                        type: array
                      tls:
                        description: TLS related settings for connections to the upstream
                          service.
                        properties:
                          caCertificates:
                            format: string
                            type: string
                          clientCertificate:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          credentialName:
                            format: string
                            type: string
                          mode:
                            enum:
                            - DISABLE
                            - SIMPLE
                            - MUTUAL
                            - ISTIO_MUTUAL
                            type: string
                          privateKey:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          sni:
                            description: SNI string to present to the server during
                              TLS handshake.
                            format: string
                            type: string
                          subjectAltNames:
                            items:
                              format: string
                              type: string
                            type: array
                        type: object
                    type: object
                type: object
              type: array
            trafficPolicy:
              properties:
                connectionPool:
                  properties:
                    http:
                      description: HTTP connection pool settings.
                      properties:
                        h2UpgradePolicy:
                          description: Specify if http1.1 connection should be upgraded
                            to http2 for the associated destination.
                          enum:
                          - DEFAULT
                          - DO_NOT_UPGRADE
                          - UPGRADE
                          type: string
                        http1MaxPendingRequests:
                          description: Maximum number of pending HTTP requests to
                            a destination.
                          format: int32
                          type: integer
                        http2MaxRequests:
                          description: Maximum number of requests to a backend.
                          format: int32
                          type: integer
                        idleTimeout:
                          description: The idle timeout for upstream connection pool
                            connections.
                          type: string
                        maxRequestsPerConnection:
                          description: Maximum number of requests per connection to
                            a backend.
                          format: int32
                          type: integer
                        maxRetries:
                          format: int32
                          type: integer
                        useClientProtocol:
                          description: If set to true, client protocol will be preserved
                            while initiating connection to backend.
                          type: boolean
                      type: object
                    tcp:
                      description: Settings common to both HTTP and TCP upstream connections.
                      properties:
                        connectTimeout:
                          description: TCP connection timeout.
                          type: string
                        maxConnections:
                          description: Maximum number of HTTP1 /TCP connections to
                            a destination host.
                          format: int32
                          type: integer
                        tcpKeepalive:
                          description: If set then set SO_KEEPALIVE on the socket
                            to enable TCP Keepalives.
                          properties:
                            interval:
                              description: The time duration between keep-alive probes.
                              type: string
                            probes:
                              type: integer
                            time:
                              type: string
                          type: object
                      type: object
                  type: object
                loadBalancer:
                  description: Settings controlling the load balancer algorithms.
                  oneOf:
                  - not:
                      anyOf:
                      - required:
                        - simple
                      - properties:
                          consistentHash:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - httpHeaderName
                                - required:
                                  - httpCookie
                                - required:
                                  - useSourceIp
                                - required:
                                  - httpQueryParameterName
                            - required:
                              - httpHeaderName
                            - required:
                              - httpCookie
                            - required:
                              - useSourceIp
                            - required:
                              - httpQueryParameterName
                        required:
                        - consistentHash
                  - required:
                    - simple
                  - properties:
                      consistentHash:
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - httpHeaderName
                            - required:
                              - httpCookie
                            - required:
                              - useSourceIp
                            - required:
                              - httpQueryParameterName
                        - required:
                          - httpHeaderName
                        - required:
                          - httpCookie
                        - required:
                          - useSourceIp
                        - required:
                          - httpQueryParameterName
                    required:
                    - consistentHash
                  properties:
                    consistentHash:
                      properties:
                        httpCookie:
                          description: Hash based on HTTP cookie.
                          properties:
                            name:
                              description: Name of the cookie.
                              format: string
                              type: string
                            path:
                              description: Path to set for the cookie.
                              format: string
                              type: string
                            ttl:
                              description: Lifetime of the cookie.
                              type: string
                          type: object
                        httpHeaderName:
                          description: Hash based on a specific HTTP header.
                          format: string
                          type: string
                        httpQueryParameterName:
                          description: Hash based on a specific HTTP query parameter.
                          format: string
                          type: string
                        minimumRingSize:
                          type: integer
                        useSourceIp:
                          description: Hash based on the source IP address.
                          type: boolean
                      type: object
                    localityLbSetting:
                      properties:
                        distribute:
                          description: 'Optional: only one of distribute or failover
                            can be set.'
                          items:
                            properties:
                              from:
                                description: Originating locality, '/' separated,
                                  e.g.
                                format: string
                                type: string
                              to:
                                additionalProperties:
                                  type: integer
                                description: Map of upstream localities to traffic
                                  distribution weights.
                                type: object
                            type: object
                          type: array
                        enabled:
                          description: enable locality load balancing, this is DestinationRule-level
                            and will override mesh wide settings in entirety.
                          nullable: true
                          type: boolean
                        failover:
                          description: 'Optional: only failover or distribute can
                            be set.'
                          items:
                            properties:
                              from:
                                description: Originating region.
                                format: string
                                type: string
                              to:
                                format: string
                                type: string
                            type: object
                          type: array
                      type: object
                    simple:
                      enum:
                      - ROUND_ROBIN
                      - LEAST_CONN
                      - RANDOM
                      - PASSTHROUGH
		      - PREDICTIVE_LEAST_CONN
                      type: string
                  type: object
                outlierDetection:
                  properties:
                    baseEjectionTime:
                      description: Minimum ejection duration.
                      type: string
                    consecutive5xxErrors:
                      description: Number of 5xx errors before a host is ejected from
                        the connection pool.
                      nullable: true
                      type: integer
                    consecutiveErrors:
                      format: int32
                      type: integer
                    consecutiveGatewayErrors:
                      description: Number of gateway errors before a host is ejected
                        from the connection pool.
                      nullable: true
                      type: integer
                    interval:
                      description: Time interval between ejection sweep analysis.
                      type: string
                    maxEjectionPercent:
                      format: int32
                      type: integer
                    minHealthPercent:
                      format: int32
                      type: integer
                  type: object
                portLevelSettings:
                  description: Traffic policies specific to individual ports.
                  items:
                    properties:
                      connectionPool:
                        properties:
                          http:
                            description: HTTP connection pool settings.
                            properties:
                              h2UpgradePolicy:
                                description: Specify if http1.1 connection should
                                  be upgraded to http2 for the associated destination.
                                enum:
                                - DEFAULT
                                - DO_NOT_UPGRADE
                                - UPGRADE
                                type: string
                              http1MaxPendingRequests:
                                description: Maximum number of pending HTTP requests
                                  to a destination.
                                format: int32
                                type: integer
                              http2MaxRequests:
                                description: Maximum number of requests to a backend.
                                format: int32
                                type: integer
                              idleTimeout:
                                description: The idle timeout for upstream connection
                                  pool connections.
                                type: string
                              maxRequestsPerConnection:
                                description: Maximum number of requests per connection
                                  to a backend.
                                format: int32
                                type: integer
                              maxRetries:
                                format: int32
                                type: integer
                              useClientProtocol:
                                description: If set to true, client protocol will
                                  be preserved while initiating connection to backend.
                                type: boolean
                            type: object
                          tcp:
                            description: Settings common to both HTTP and TCP upstream
                              connections.
                            properties:
                              connectTimeout:
                                description: TCP connection timeout.
                                type: string
                              maxConnections:
                                description: Maximum number of HTTP1 /TCP connections
                                  to a destination host.
                                format: int32
                                type: integer
                              tcpKeepalive:
                                description: If set then set SO_KEEPALIVE on the socket
                                  to enable TCP Keepalives.
                                properties:
                                  interval:
                                    description: The time duration between keep-alive
                                      probes.
                                    type: string
                                  probes:
                                    type: integer
                                  time:
                                    type: string
                                type: object
                            type: object
                        type: object
                      loadBalancer:
                        description: Settings controlling the load balancer algorithms.
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - simple
                            - properties:
                                consistentHash:
                                  oneOf:
                                  - not:
                                      anyOf:
                                      - required:
                                        - httpHeaderName
                                      - required:
                                        - httpCookie
                                      - required:
                                        - useSourceIp
                                      - required:
                                        - httpQueryParameterName
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              required:
                              - consistentHash
                        - required:
                          - simple
                        - properties:
                            consistentHash:
                              oneOf:
                              - not:
                                  anyOf:
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              - required:
                                - httpHeaderName
                              - required:
                                - httpCookie
                              - required:
                                - useSourceIp
                              - required:
                                - httpQueryParameterName
                          required:
                          - consistentHash
                        properties:
                          consistentHash:
                            properties:
                              httpCookie:
                                description: Hash based on HTTP cookie.
                                properties:
                                  name:
                                    description: Name of the cookie.
                                    format: string
                                    type: string
                                  path:
                                    description: Path to set for the cookie.
                                    format: string
                                    type: string
                                  ttl:
                                    description: Lifetime of the cookie.
                                    type: string
                                type: object
                              httpHeaderName:
                                description: Hash based on a specific HTTP header.
                                format: string
                                type: string
                              httpQueryParameterName:
                                description: Hash based on a specific HTTP query parameter.
                                format: string
                                type: string
                              minimumRingSize:
                                type: integer
                              useSourceIp:
                                description: Hash based on the source IP address.
                                type: boolean
                            type: object
                          localityLbSetting:
                            properties:
                              distribute:
                                description: 'Optional: only one of distribute or
                                  failover can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating locality, '/' separated,
                                        e.g.
                                      format: string
                                      type: string
                                    to:
                                      additionalProperties:
                                        type: integer
                                      description: Map of upstream localities to traffic
                                        distribution weights.
                                      type: object
                                  type: object
                                type: array
                              enabled:
                                description: enable locality load balancing, this
                                  is DestinationRule-level and will override mesh
                                  wide settings in entirety.
                                nullable: true
                                type: boolean
                              failover:
                                description: 'Optional: only failover or distribute
                                  can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating region.
                                      format: string
                                      type: string
                                    to:
                                      format: string
                                      type: string
                                  type: object
                                type: array
                            type: object
                          simple:
                            enum:
                            - ROUND_ROBIN
                            - LEAST_CONN
                            - RANDOM
                            - PASSTHROUGH
			    - PREDICTIVE_LEAST_CONN
                            type: string
                        type: object
                      outlierDetection:
                        properties:
                          baseEjectionTime:
                            description: Minimum ejection duration.
                            type: string
                          consecutive5xxErrors:
                            description: Number of 5xx errors before a host is ejected
                              from the connection pool.
                            nullable: true
                            type: integer
                          consecutiveErrors:
                            format: int32
                            type: integer
                          consecutiveGatewayErrors:
                            description: Number of gateway errors before a host is
                              ejected from the connection pool.
                            nullable: true
                            type: integer
                          interval:
                            description: Time interval between ejection sweep analysis.
                            type: string
                          maxEjectionPercent:
                            format: int32
                            type: integer
                          minHealthPercent:
                            format: int32
                            type: integer
                        type: object
                      port:
                        properties:
                          number:
                            type: integer
                        type: object
                      tls:
                        description: TLS related settings for connections to the upstream
                          service.
                        properties:
                          caCertificates:
                            format: string
                            type: string
                          clientCertificate:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          credentialName:
                            format: string
                            type: string
                          mode:
                            enum:
                            - DISABLE
                            - SIMPLE
                            - MUTUAL
                            - ISTIO_MUTUAL
                            type: string
                          privateKey:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          sni:
                            description: SNI string to present to the server during
                              TLS handshake.
                            format: string
                            type: string
                          subjectAltNames:
                            items:
                              format: string
                              type: string
                            type: array
                        type: object
                    type: object
                  type: array
                tls:
                  description: TLS related settings for connections to the upstream
                    service.
                  properties:
                    caCertificates:
                      format: string
                      type: string
                    clientCertificate:
                      description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                      format: string
                      type: string
                    credentialName:
                      format: string
                      type: string
                    mode:
                      enum:
                      - DISABLE
                      - SIMPLE
                      - MUTUAL
                      - ISTIO_MUTUAL
                      type: string
                    privateKey:
                      description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                      format: string
                      type: string
                    sni:
                      description: SNI string to present to the server during TLS
                        handshake.
                      format: string
                      type: string
                    subjectAltNames:
                      items:
                        format: string
                        type: string
                      type: array
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: envoyfilters.networking.istio.io
spec:
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: EnvoyFilter
    listKind: EnvoyFilterList
    plural: envoyfilters
    singular: envoyfilter
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Customizing Envoy configuration generated by Istio. See more
            details at: https://istio.io/docs/reference/config/networking/envoy-filter.html'
          properties:
            configPatches:
              description: One or more patches with match conditions.
              items:
                properties:
                  applyTo:
                    enum:
                    - INVALID
                    - LISTENER
                    - FILTER_CHAIN
                    - NETWORK_FILTER
                    - HTTP_FILTER
                    - ROUTE_CONFIGURATION
                    - VIRTUAL_HOST
                    - HTTP_ROUTE
                    - CLUSTER
                    type: string
                  match:
                    description: Match on listener/route configuration/cluster.
                    oneOf:
                    - not:
                        anyOf:
                        - required:
                          - listener
                        - required:
                          - routeConfiguration
                        - required:
                          - cluster
                    - required:
                      - listener
                    - required:
                      - routeConfiguration
                    - required:
                      - cluster
                    properties:
                      cluster:
                        description: Match on envoy cluster attributes.
                        properties:
                          name:
                            description: The exact name of the cluster to match.
                            format: string
                            type: string
                          portNumber:
                            description: The service port for which this cluster was
                              generated.
                            type: integer
                          service:
                            description: The fully qualified service name for this
                              cluster.
                            format: string
                            type: string
                          subset:
                            description: The subset associated with the service.
                            format: string
                            type: string
                        type: object
                      context:
                        description: The specific config generation context to match
                          on.
                        enum:
                        - ANY
                        - SIDECAR_INBOUND
                        - SIDECAR_OUTBOUND
                        - GATEWAY
                        type: string
                      listener:
                        description: Match on envoy listener attributes.
                        properties:
                          filterChain:
                            description: Match a specific filter chain in a listener.
                            properties:
                              applicationProtocols:
                                description: Applies only to sidecars.
                                format: string
                                type: string
                              filter:
                                description: The name of a specific filter to apply
                                  the patch to.
                                properties:
                                  name:
                                    description: The filter name to match on.
                                    format: string
                                    type: string
                                  subFilter:
                                    properties:
                                      name:
                                        description: The filter name to match on.
                                        format: string
                                        type: string
                                    type: object
                                type: object
                              name:
                                description: The name assigned to the filter chain.
                                format: string
                                type: string
                              sni:
                                description: The SNI value used by a filter chain's
                                  match condition.
                                format: string
                                type: string
                              transportProtocol:
                                description: Applies only to SIDECAR_INBOUND context.
                                format: string
                                type: string
                            type: object
                          name:
                            description: Match a specific listener by its name.
                            format: string
                            type: string
                          portName:
                            format: string
                            type: string
                          portNumber:
                            type: integer
                        type: object
                      proxy:
                        description: Match on properties associated with a proxy.
                        properties:
                          metadata:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                          proxyVersion:
                            format: string
                            type: string
                        type: object
                      routeConfiguration:
                        description: Match on envoy HTTP route configuration attributes.
                        properties:
                          gateway:
                            format: string
                            type: string
                          name:
                            description: Route configuration name to match on.
                            format: string
                            type: string
                          portName:
                            description: Applicable only for GATEWAY context.
                            format: string
                            type: string
                          portNumber:
                            type: integer
                          vhost:
                            properties:
                              name:
                                format: string
                                type: string
                              route:
                                description: Match a specific route within the virtual
                                  host.
                                properties:
                                  action:
                                    description: Match a route with specific action
                                      type.
                                    enum:
                                    - ANY
                                    - ROUTE
                                    - REDIRECT
                                    - DIRECT_RESPONSE
                                    type: string
                                  name:
                                    format: string
                                    type: string
                                type: object
                            type: object
                        type: object
                    type: object
                  patch:
                    description: The patch to apply along with the operation.
                    properties:
                      filterClass:
                        description: Determines the filter insertion order.
                        enum:
                        - UNSPECIFIED
                        - AUTHN
                        - AUTHZ
                        - STATS
                        type: string
                      operation:
                        description: Determines how the patch should be applied.
                        enum:
                        - INVALID
                        - MERGE
                        - ADD
                        - REMOVE
                        - INSERT_BEFORE
                        - INSERT_AFTER
                        - INSERT_FIRST
                        - REPLACE
                        type: string
                      value:
                        description: The JSON config of the object being patched.
                        type: object
                    type: object
                type: object
              type: array
            workloadSelector:
              properties:
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: gateways.networking.istio.io
spec:
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: Gateway
    listKind: GatewayList
    plural: gateways
    shortNames:
    - gw
    singular: gateway
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting edge load balancer. See more details
            at: https://istio.io/docs/reference/config/networking/gateway.html'
          properties:
            selector:
              additionalProperties:
                format: string
                type: string
              type: object
            servers:
              description: A list of server specifications.
              items:
                properties:
                  bind:
                    format: string
                    type: string
                  defaultEndpoint:
                    format: string
                    type: string
                  hosts:
                    description: One or more hosts exposed by this gateway.
                    items:
                      format: string
                      type: string
                    type: array
                  name:
                    description: An optional name of the server, when set must be
                      unique across all servers.
                    format: string
                    type: string
                  port:
                    properties:
                      name:
                        description: Label assigned to the port.
                        format: string
                        type: string
                      number:
                        description: A valid non-negative integer port number.
                        type: integer
                      protocol:
                        description: The protocol exposed on the port.
                        format: string
                        type: string
                      targetPort:
                        type: integer
                    type: object
                  tls:
                    description: Set of TLS related options that govern the server's
                      behavior.
                    properties:
                      caCertificates:
                        description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                        format: string
                        type: string
                      cipherSuites:
                        description: 'Optional: If specified, only support the specified
                          cipher list.'
                        items:
                          format: string
                          type: string
                        type: array
                      credentialName:
                        format: string
                        type: string
                      httpsRedirect:
                        type: boolean
                      maxProtocolVersion:
                        description: 'Optional: Maximum TLS protocol version.'
                        enum:
                        - TLS_AUTO
                        - TLSV1_0
                        - TLSV1_1
                        - TLSV1_2
                        - TLSV1_3
                        type: string
                      minProtocolVersion:
                        description: 'Optional: Minimum TLS protocol version.'
                        enum:
                        - TLS_AUTO
                        - TLSV1_0
                        - TLSV1_1
                        - TLSV1_2
                        - TLSV1_3
                        type: string
                      mode:
                        enum:
                        - PASSTHROUGH
                        - SIMPLE
                        - MUTUAL
                        - AUTO_PASSTHROUGH
                        - ISTIO_MUTUAL
                        type: string
                      privateKey:
                        description: REQUIRED if mode is `+"`"+`SIMPLE`+"`"+` or `+"`"+`MUTUAL`+"`"+`.
                        format: string
                        type: string
                      serverCertificate:
                        description: REQUIRED if mode is `+"`"+`SIMPLE`+"`"+` or `+"`"+`MUTUAL`+"`"+`.
                        format: string
                        type: string
                      subjectAltNames:
                        items:
                          format: string
                          type: string
                        type: array
                      verifyCertificateHash:
                        items:
                          format: string
                          type: string
                        type: array
                      verifyCertificateSpki:
                        items:
                          format: string
                          type: string
                        type: array
                    type: object
                type: object
              type: array
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: serviceentries.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.hosts
    description: The hosts associated with the ServiceEntry
    name: Hosts
    type: string
  - JSONPath: .spec.location
    description: Whether the service is external to the mesh or part of the mesh (MESH_EXTERNAL
      or MESH_INTERNAL)
    name: Location
    type: string
  - JSONPath: .spec.resolution
    description: Service discovery mode for the hosts (NONE, STATIC, or DNS)
    name: Resolution
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: ServiceEntry
    listKind: ServiceEntryList
    plural: serviceentries
    shortNames:
    - se
    singular: serviceentry
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting service registry. See more details
            at: https://istio.io/docs/reference/config/networking/service-entry.html'
          properties:
            addresses:
              description: The virtual IP addresses associated with the service.
              items:
                format: string
                type: string
              type: array
            endpoints:
              description: One or more endpoints associated with the service.
              items:
                properties:
                  address:
                    format: string
                    type: string
                  labels:
                    additionalProperties:
                      format: string
                      type: string
                    description: One or more labels associated with the endpoint.
                    type: object
                  locality:
                    description: The locality associated with the endpoint.
                    format: string
                    type: string
                  network:
                    format: string
                    type: string
                  ports:
                    additionalProperties:
                      type: integer
                    description: Set of ports associated with the endpoint.
                    type: object
                  serviceAccount:
                    format: string
                    type: string
                  weight:
                    description: The load balancing weight associated with the endpoint.
                    type: integer
                type: object
              type: array
            exportTo:
              description: A list of namespaces to which this service is exported.
              items:
                format: string
                type: string
              type: array
            hosts:
              description: The hosts associated with the ServiceEntry.
              items:
                format: string
                type: string
              type: array
            location:
              enum:
              - MESH_EXTERNAL
              - MESH_INTERNAL
              type: string
            ports:
              description: The ports associated with the external service.
              items:
                properties:
                  name:
                    description: Label assigned to the port.
                    format: string
                    type: string
                  number:
                    description: A valid non-negative integer port number.
                    type: integer
                  protocol:
                    description: The protocol exposed on the port.
                    format: string
                    type: string
                  targetPort:
                    type: integer
                type: object
              type: array
            resolution:
              description: Service discovery mode for the hosts.
              enum:
              - NONE
              - STATIC
              - DNS
              type: string
            subjectAltNames:
              items:
                format: string
                type: string
              type: array
            workloadSelector:
              description: Applicable only for MESH_INTERNAL services.
              properties:
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: sidecars.networking.istio.io
spec:
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: Sidecar
    listKind: SidecarList
    plural: sidecars
    singular: sidecar
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting network reachability of a sidecar.
            See more details at: https://istio.io/docs/reference/config/networking/sidecar.html'
          properties:
            egress:
              items:
                properties:
                  bind:
                    format: string
                    type: string
                  captureMode:
                    enum:
                    - DEFAULT
                    - IPTABLES
                    - NONE
                    type: string
                  hosts:
                    items:
                      format: string
                      type: string
                    type: array
                  port:
                    description: The port associated with the listener.
                    properties:
                      name:
                        description: Label assigned to the port.
                        format: string
                        type: string
                      number:
                        description: A valid non-negative integer port number.
                        type: integer
                      protocol:
                        description: The protocol exposed on the port.
                        format: string
                        type: string
                      targetPort:
                        type: integer
                    type: object
                type: object
              type: array
            ingress:
              items:
                properties:
                  bind:
                    description: The IP to which the listener should be bound.
                    format: string
                    type: string
                  captureMode:
                    enum:
                    - DEFAULT
                    - IPTABLES
                    - NONE
                    type: string
                  defaultEndpoint:
                    format: string
                    type: string
                  port:
                    description: The port associated with the listener.
                    properties:
                      name:
                        description: Label assigned to the port.
                        format: string
                        type: string
                      number:
                        description: A valid non-negative integer port number.
                        type: integer
                      protocol:
                        description: The protocol exposed on the port.
                        format: string
                        type: string
                      targetPort:
                        type: integer
                    type: object
                type: object
              type: array
            outboundTrafficPolicy:
              description: Configuration for the outbound traffic policy.
              properties:
                egressProxy:
                  properties:
                    host:
                      description: The name of a service from the service registry.
                      format: string
                      type: string
                    port:
                      description: Specifies the port on the host that is being addressed.
                      properties:
                        number:
                          type: integer
                      type: object
                    subset:
                      description: The name of a subset within the service.
                      format: string
                      type: string
                  type: object
                mode:
                  enum:
                  - REGISTRY_ONLY
                  - ALLOW_ANY
                  type: string
              type: object
            workloadSelector:
              properties:
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: virtualservices.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.gateways
    description: The names of gateways and sidecars that should apply these routes
    name: Gateways
    type: string
  - JSONPath: .spec.hosts
    description: The destination hosts to which traffic is being sent
    name: Hosts
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: VirtualService
    listKind: VirtualServiceList
    plural: virtualservices
    shortNames:
    - vs
    singular: virtualservice
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting label/content routing, sni routing,
            etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html'
          properties:
            exportTo:
              description: A list of namespaces to which this virtual service is exported.
              items:
                format: string
                type: string
              type: array
            gateways:
              description: The names of gateways and sidecars that should apply these
                routes.
              items:
                format: string
                type: string
              type: array
            hosts:
              description: The destination hosts to which traffic is being sent.
              items:
                format: string
                type: string
              type: array
            http:
              description: An ordered list of route rules for HTTP traffic.
              items:
                properties:
                  corsPolicy:
                    description: Cross-Origin Resource Sharing policy (CORS).
                    properties:
                      allowCredentials:
                        nullable: true
                        type: boolean
                      allowHeaders:
                        items:
                          format: string
                          type: string
                        type: array
                      allowMethods:
                        description: List of HTTP methods allowed to access the resource.
                        items:
                          format: string
                          type: string
                        type: array
                      allowOrigin:
                        description: The list of origins that are allowed to perform
                          CORS requests.
                        items:
                          format: string
                          type: string
                        type: array
                      allowOrigins:
                        description: String patterns that match allowed origins.
                        items:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        type: array
                      exposeHeaders:
                        items:
                          format: string
                          type: string
                        type: array
                      maxAge:
                        type: string
                    type: object
                  delegate:
                    properties:
                      name:
                        description: Name specifies the name of the delegate VirtualService.
                        format: string
                        type: string
                      namespace:
                        description: Namespace specifies the namespace where the delegate
                          VirtualService resides.
                        format: string
                        type: string
                    type: object
                  fault:
                    description: Fault injection policy to apply on HTTP traffic at
                      the client side.
                    properties:
                      abort:
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - httpStatus
                            - required:
                              - grpcStatus
                            - required:
                              - http2Error
                        - required:
                          - httpStatus
                        - required:
                          - grpcStatus
                        - required:
                          - http2Error
                        properties:
                          grpcStatus:
                            format: string
                            type: string
                          http2Error:
                            format: string
                            type: string
                          httpStatus:
                            description: HTTP status code to use to abort the Http
                              request.
                            format: int32
                            type: integer
                          percentage:
                            description: Percentage of requests to be aborted with
                              the error code provided.
                            properties:
                              value:
                                format: double
                                type: number
                            type: object
                        type: object
                      delay:
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - fixedDelay
                            - required:
                              - exponentialDelay
                        - required:
                          - fixedDelay
                        - required:
                          - exponentialDelay
                        properties:
                          exponentialDelay:
                            type: string
                          fixedDelay:
                            description: Add a fixed delay before forwarding the request.
                            type: string
                          percent:
                            description: Percentage of requests on which the delay
                              will be injected (0-100).
                            format: int32
                            type: integer
                          percentage:
                            description: Percentage of requests on which the delay
                              will be injected.
                            properties:
                              value:
                                format: double
                                type: number
                            type: object
                        type: object
                    type: object
                  headers:
                    properties:
                      request:
                        properties:
                          add:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                          remove:
                            items:
                              format: string
                              type: string
                            type: array
                          set:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                        type: object
                      response:
                        properties:
                          add:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                          remove:
                            items:
                              format: string
                              type: string
                            type: array
                          set:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                        type: object
                    type: object
                  match:
                    items:
                      properties:
                        authority:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        gateways:
                          description: Names of gateways where the rule should be
                            applied.
                          items:
                            format: string
                            type: string
                          type: array
                        headers:
                          additionalProperties:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - exact
                                - required:
                                  - prefix
                                - required:
                                  - regex
                            - required:
                              - exact
                            - required:
                              - prefix
                            - required:
                              - regex
                            properties:
                              exact:
                                format: string
                                type: string
                              prefix:
                                format: string
                                type: string
                              regex:
                                description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                                format: string
                                type: string
                            type: object
                          type: object
                        ignoreUriCase:
                          description: Flag to specify whether the URI matching should
                            be case-insensitive.
                          type: boolean
                        method:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        name:
                          description: The name assigned to a match.
                          format: string
                          type: string
                        port:
                          description: Specifies the ports on the host that is being
                            addressed.
                          type: integer
                        queryParams:
                          additionalProperties:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - exact
                                - required:
                                  - prefix
                                - required:
                                  - regex
                            - required:
                              - exact
                            - required:
                              - prefix
                            - required:
                              - regex
                            properties:
                              exact:
                                format: string
                                type: string
                              prefix:
                                format: string
                                type: string
                              regex:
                                description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                                format: string
                                type: string
                            type: object
                          description: Query parameters for matching.
                          type: object
                        scheme:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        sourceLabels:
                          additionalProperties:
                            format: string
                            type: string
                          type: object
                        sourceNamespace:
                          description: Source namespace constraining the applicability
                            of a rule to workloads in that namespace.
                          format: string
                          type: string
                        uri:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        withoutHeaders:
                          additionalProperties:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - exact
                                - required:
                                  - prefix
                                - required:
                                  - regex
                            - required:
                              - exact
                            - required:
                              - prefix
                            - required:
                              - regex
                            properties:
                              exact:
                                format: string
                                type: string
                              prefix:
                                format: string
                                type: string
                              regex:
                                description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                                format: string
                                type: string
                            type: object
                          description: withoutHeader has the same syntax with the
                            header, but has opposite meaning.
                          type: object
                      type: object
                    type: array
                  mirror:
                    properties:
                      host:
                        description: The name of a service from the service registry.
                        format: string
                        type: string
                      port:
                        description: Specifies the port on the host that is being
                          addressed.
                        properties:
                          number:
                            type: integer
                        type: object
                      subset:
                        description: The name of a subset within the service.
                        format: string
                        type: string
                    type: object
                  mirror_percent:
                    description: Percentage of the traffic to be mirrored by the `+"`"+`mirror`+"`"+`
                      field.
                    nullable: true
                    type: integer
                  mirrorPercent:
                    description: Percentage of the traffic to be mirrored by the `+"`"+`mirror`+"`"+`
                      field.
                    nullable: true
                    type: integer
                  mirrorPercentage:
                    description: Percentage of the traffic to be mirrored by the `+"`"+`mirror`+"`"+`
                      field.
                    properties:
                      value:
                        format: double
                        type: number
                    type: object
                  name:
                    description: The name assigned to the route for debugging purposes.
                    format: string
                    type: string
                  redirect:
                    description: A HTTP rule can either redirect or forward (default)
                      traffic.
                    properties:
                      authority:
                        format: string
                        type: string
                      redirectCode:
                        type: integer
                      uri:
                        format: string
                        type: string
                    type: object
                  retries:
                    description: Retry policy for HTTP requests.
                    properties:
                      attempts:
                        description: Number of retries for a given request.
                        format: int32
                        type: integer
                      perTryTimeout:
                        description: Timeout per retry attempt for a given request.
                        type: string
                      retryOn:
                        description: Specifies the conditions under which retry takes
                          place.
                        format: string
                        type: string
                      retryRemoteLocalities:
                        description: Flag to specify whether the retries should retry
                          to other localities.
                        nullable: true
                        type: boolean
                    type: object
                  rewrite:
                    description: Rewrite HTTP URIs and Authority headers.
                    properties:
                      authority:
                        description: rewrite the Authority/Host header with this value.
                        format: string
                        type: string
                      uri:
                        format: string
                        type: string
                    type: object
                  route:
                    description: A HTTP rule can either redirect or forward (default)
                      traffic.
                    items:
                      properties:
                        destination:
                          properties:
                            host:
                              description: The name of a service from the service
                                registry.
                              format: string
                              type: string
                            port:
                              description: Specifies the port on the host that is
                                being addressed.
                              properties:
                                number:
                                  type: integer
                              type: object
                            subset:
                              description: The name of a subset within the service.
                              format: string
                              type: string
                          type: object
                        headers:
                          properties:
                            request:
                              properties:
                                add:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                                remove:
                                  items:
                                    format: string
                                    type: string
                                  type: array
                                set:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                              type: object
                            response:
                              properties:
                                add:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                                remove:
                                  items:
                                    format: string
                                    type: string
                                  type: array
                                set:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                              type: object
                          type: object
                        weight:
                          format: int32
                          type: integer
                      type: object
                    type: array
                  timeout:
                    description: Timeout for HTTP requests, default is disabled.
                    type: string
                type: object
              type: array
            tcp:
              description: An ordered list of route rules for opaque TCP traffic.
              items:
                properties:
                  match:
                    items:
                      properties:
                        destinationSubnets:
                          description: IPv4 or IPv6 ip addresses of destination with
                            optional subnet.
                          items:
                            format: string
                            type: string
                          type: array
                        gateways:
                          description: Names of gateways where the rule should be
                            applied.
                          items:
                            format: string
                            type: string
                          type: array
                        port:
                          description: Specifies the port on the host that is being
                            addressed.
                          type: integer
                        sourceLabels:
                          additionalProperties:
                            format: string
                            type: string
                          type: object
                        sourceNamespace:
                          description: Source namespace constraining the applicability
                            of a rule to workloads in that namespace.
                          format: string
                          type: string
                        sourceSubnet:
                          description: IPv4 or IPv6 ip address of source with optional
                            subnet.
                          format: string
                          type: string
                      type: object
                    type: array
                  route:
                    description: The destination to which the connection should be
                      forwarded to.
                    items:
                      properties:
                        destination:
                          properties:
                            host:
                              description: The name of a service from the service
                                registry.
                              format: string
                              type: string
                            port:
                              description: Specifies the port on the host that is
                                being addressed.
                              properties:
                                number:
                                  type: integer
                              type: object
                            subset:
                              description: The name of a subset within the service.
                              format: string
                              type: string
                          type: object
                        weight:
                          format: int32
                          type: integer
                      type: object
                    type: array
                type: object
              type: array
            tls:
              items:
                properties:
                  match:
                    items:
                      properties:
                        destinationSubnets:
                          description: IPv4 or IPv6 ip addresses of destination with
                            optional subnet.
                          items:
                            format: string
                            type: string
                          type: array
                        gateways:
                          description: Names of gateways where the rule should be
                            applied.
                          items:
                            format: string
                            type: string
                          type: array
                        port:
                          description: Specifies the port on the host that is being
                            addressed.
                          type: integer
                        sniHosts:
                          description: SNI (server name indicator) to match on.
                          items:
                            format: string
                            type: string
                          type: array
                        sourceLabels:
                          additionalProperties:
                            format: string
                            type: string
                          type: object
                        sourceNamespace:
                          description: Source namespace constraining the applicability
                            of a rule to workloads in that namespace.
                          format: string
                          type: string
                      type: object
                    type: array
                  route:
                    description: The destination to which the connection should be
                      forwarded to.
                    items:
                      properties:
                        destination:
                          properties:
                            host:
                              description: The name of a service from the service
                                registry.
                              format: string
                              type: string
                            port:
                              description: Specifies the port on the host that is
                                being addressed.
                              properties:
                                number:
                                  type: integer
                              type: object
                            subset:
                              description: The name of a subset within the service.
                              format: string
                              type: string
                          type: object
                        weight:
                          format: int32
                          type: integer
                      type: object
                    type: array
                type: object
              type: array
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: workloadentries.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  - JSONPath: .spec.address
    description: Address associated with the network endpoint.
    name: Address
    type: string
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: WorkloadEntry
    listKind: WorkloadEntryList
    plural: workloadentries
    shortNames:
    - we
    singular: workloadentry
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting VMs onboarded into the mesh. See more
            details at: https://istio.io/docs/reference/config/networking/workload-entry.html'
          properties:
            address:
              format: string
              type: string
            labels:
              additionalProperties:
                format: string
                type: string
              description: One or more labels associated with the endpoint.
              type: object
            locality:
              description: The locality associated with the endpoint.
              format: string
              type: string
            network:
              format: string
              type: string
            ports:
              additionalProperties:
                type: integer
              description: Set of ports associated with the endpoint.
              type: object
            serviceAccount:
              format: string
              type: string
            weight:
              description: The load balancing weight associated with the endpoint.
              type: integer
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: workloadgroups.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: WorkloadGroup
    listKind: WorkloadGroupList
    plural: workloadgroups
    shortNames:
    - wg
    singular: workloadgroup
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Describes a collection of workload instances. See more details
            at: https://istio.io/docs/reference/config/networking/workload-group.html'
          properties:
            metadata:
              description: Metadata that will be used for all corresponding `+"`"+`WorkloadEntries`+"`"+`.
              properties:
                annotations:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
            probe:
              description: '`+"`"+`ReadinessProbe`+"`"+` describes the configuration the user
                must provide for healthchecking on their workload.'
              oneOf:
              - not:
                  anyOf:
                  - required:
                    - httpGet
                  - required:
                    - tcpSocket
                  - required:
                    - exec
              - required:
                - httpGet
              - required:
                - tcpSocket
              - required:
                - exec
              properties:
                exec:
                  description: health is determined by how the command that is executed
                    exited.
                  properties:
                    command:
                      description: command to run.
                      items:
                        format: string
                        type: string
                      type: array
                  type: object
                failureThreshold:
                  description: Minimum consecutive failures for the probe to be considered
                    failed after having succeeded.
                  format: int32
                  type: integer
                httpGet:
                  properties:
                    host:
                      description: Host name to connect to, defaults to the pod IP.
                      format: string
                      type: string
                    httpHeaders:
                      description: headers the proxy will pass on to make the request.
                      items:
                        properties:
                          name:
                            format: string
                            type: string
                          value:
                            format: string
                            type: string
                        type: object
                      type: array
                    path:
                      description: Path to access on the HTTP server.
                      format: string
                      type: string
                    port:
                      description: port on which the endpoint lives.
                      type: integer
                    scheme:
                      format: string
                      type: string
                  type: object
                initialDelaySeconds:
                  description: Number of seconds after the container has started before
                    readiness probes are initiated.
                  format: int32
                  type: integer
                periodSeconds:
                  description: How often (in seconds) to perform the probe.
                  format: int32
                  type: integer
                successThreshold:
                  description: Minimum consecutive successes for the probe to be considered
                    successful after having failed.
                  format: int32
                  type: integer
                tcpSocket:
                  description: health is determined by if the proxy is able to connect.
                  properties:
                    host:
                      format: string
                      type: string
                    port:
                      type: integer
                  type: object
                timeoutSeconds:
                  description: Number of seconds after which the probe times out.
                  format: int32
                  type: integer
              type: object
            template:
              description: Template to be used for the generation of `+"`"+`WorkloadEntry`+"`"+`
                resources that belong to this `+"`"+`WorkloadGroup`+"`"+`.
              properties:
                address:
                  format: string
                  type: string
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  description: One or more labels associated with the endpoint.
                  type: object
                locality:
                  description: The locality associated with the endpoint.
                  format: string
                  type: string
                network:
                  format: string
                  type: string
                ports:
                  additionalProperties:
                    type: integer
                  description: Set of ports associated with the endpoint.
                  type: object
                serviceAccount:
                  format: string
                  type: string
                weight:
                  description: The load balancing weight associated with the endpoint.
                  type: integer
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    istio: security
    release: istio
  name: authorizationpolicies.security.istio.io
spec:
  group: security.istio.io
  names:
    categories:
    - istio-io
    - security-istio-io
    kind: AuthorizationPolicy
    listKind: AuthorizationPolicyList
    plural: authorizationpolicies
    singular: authorizationpolicy
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration for access control on workloads. See more details
            at: https://istio.io/docs/reference/config/security/authorization-policy.html'
          properties:
            action:
              description: Optional.
              enum:
              - ALLOW
              - DENY
              - AUDIT
              type: string
            rules:
              description: Optional.
              items:
                properties:
                  from:
                    description: Optional.
                    items:
                      properties:
                        source:
                          description: Source specifies the source of a request.
                          properties:
                            ipBlocks:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            namespaces:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notIpBlocks:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notNamespaces:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notPrincipals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notRequestPrincipals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            principals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            requestPrincipals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                          type: object
                      type: object
                    type: array
                  to:
                    description: Optional.
                    items:
                      properties:
                        operation:
                          description: Operation specifies the operation of a request.
                          properties:
                            hosts:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            methods:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notHosts:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notMethods:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notPaths:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notPorts:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            paths:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            ports:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                          type: object
                      type: object
                    type: array
                  when:
                    description: Optional.
                    items:
                      properties:
                        key:
                          description: The name of an Istio attribute.
                          format: string
                          type: string
                        notValues:
                          description: Optional.
                          items:
                            format: string
                            type: string
                          type: array
                        values:
                          description: Optional.
                          items:
                            format: string
                            type: string
                          type: array
                      type: object
                    type: array
                type: object
              type: array
            selector:
              description: Optional.
              properties:
                matchLabels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1beta1
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    istio: security
    release: istio
  name: peerauthentications.security.istio.io
spec:
  group: security.istio.io
  names:
    categories:
    - istio-io
    - security-istio-io
    kind: PeerAuthentication
    listKind: PeerAuthenticationList
    plural: peerauthentications
    shortNames:
    - pa
    singular: peerauthentication
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: PeerAuthentication defines how traffic will be tunneled (or
            not) to the sidecar.
          properties:
            mtls:
              description: Mutual TLS settings for workload.
              properties:
                mode:
                  description: Defines the mTLS mode used for peer authentication.
                  enum:
                  - UNSET
                  - DISABLE
                  - PERMISSIVE
                  - STRICT
                  type: string
              type: object
            portLevelMtls:
              additionalProperties:
                properties:
                  mode:
                    description: Defines the mTLS mode used for peer authentication.
                    enum:
                    - UNSET
                    - DISABLE
                    - PERMISSIVE
                    - STRICT
                    type: string
                type: object
              description: Port specific mutual TLS settings.
              type: object
            selector:
              description: The selector determines the workloads to apply the ChannelAuthentication
                on.
              properties:
                matchLabels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1beta1
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    istio: security
    release: istio
  name: requestauthentications.security.istio.io
spec:
  group: security.istio.io
  names:
    categories:
    - istio-io
    - security-istio-io
    kind: RequestAuthentication
    listKind: RequestAuthenticationList
    plural: requestauthentications
    shortNames:
    - ra
    singular: requestauthentication
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: RequestAuthentication defines what request authentication methods
            are supported by a workload.
          properties:
            jwtRules:
              description: Define the list of JWTs that can be validated at the selected
                workloads' proxy.
              items:
                properties:
                  audiences:
                    items:
                      format: string
                      type: string
                    type: array
                  forwardOriginalToken:
                    description: If set to true, the orginal token will be kept for
                      the ustream request.
                    type: boolean
                  fromHeaders:
                    description: List of header locations from which JWT is expected.
                    items:
                      properties:
                        name:
                          description: The HTTP header name.
                          format: string
                          type: string
                        prefix:
                          description: The prefix that should be stripped before decoding
                            the token.
                          format: string
                          type: string
                      type: object
                    type: array
                  fromParams:
                    description: List of query parameters from which JWT is expected.
                    items:
                      format: string
                      type: string
                    type: array
                  issuer:
                    description: Identifies the issuer that issued the JWT.
                    format: string
                    type: string
                  jwks:
                    description: JSON Web Key Set of public keys to validate signature
                      of the JWT.
                    format: string
                    type: string
                  jwks_uri:
                    format: string
                    type: string
                  jwksUri:
                    format: string
                    type: string
                  outputPayloadToHeader:
                    format: string
                    type: string
                type: object
              type: array
            selector:
              description: The selector determines the workloads to apply the RequestAuthentication
                on.
              properties:
                matchLabels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1beta1
    served: true
    storage: true

---
`)

func chartsBaseCrdsCrdAllGenYamlBytes() ([]byte, error) {
	return _chartsBaseCrdsCrdAllGenYaml, nil
}

func chartsBaseCrdsCrdAllGenYaml() (*asset, error) {
	bytes, err := chartsBaseCrdsCrdAllGenYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/crds/crd-all.gen.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseCrdsCrdOperatorYaml = []byte(`# SYNC WITH manifests/charts/istio-operator/templates
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: istiooperators.install.istio.io
  labels:
    release: istio
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.revision
    description: Istio control plane revision
    name: Revision
    type: string
  - JSONPath: .status.status
    description: IOP current state
    type: string
    name: Status
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: install.istio.io
  names:
    kind: IstioOperator
    plural: istiooperators
    singular: istiooperator
    shortNames:
    - iop
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        spec:
          description: 'Specification of the desired state of the istio control plane resource.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
        status:
          description: 'Status describes each of istio control plane component status at the current time.
            0 means NONE, 1 means UPDATING, 2 means HEALTHY, 3 means ERROR, 4 means RECONCILING.
            More info: https://github.com/istio/api/blob/master/operator/v1alpha1/istio.operator.v1alpha1.pb.html &
            https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
---
`)

func chartsBaseCrdsCrdOperatorYamlBytes() ([]byte, error) {
	return _chartsBaseCrdsCrdOperatorYaml, nil
}

func chartsBaseCrdsCrdOperatorYaml() (*asset, error) {
	bytes, err := chartsBaseCrdsCrdOperatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/crds/crd-operator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseFilesGenIstioClusterYaml = []byte(`---
# Source: crds/crd-all.gen.yaml
# DO NOT EDIT - Generated by Cue OpenAPI generator based on Istio APIs.
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: destinationrules.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.host
    description: The name of a service from the service registry
    name: Host
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: DestinationRule
    listKind: DestinationRuleList
    plural: destinationrules
    shortNames:
    - dr
    singular: destinationrule
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting load balancing, outlier detection,
            etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html'
          properties:
            exportTo:
              description: A list of namespaces to which this destination rule is
                exported.
              items:
                format: string
                type: string
              type: array
            host:
              description: The name of a service from the service registry.
              format: string
              type: string
            subsets:
              items:
                properties:
                  labels:
                    additionalProperties:
                      format: string
                      type: string
                    type: object
                  name:
                    description: Name of the subset.
                    format: string
                    type: string
                  trafficPolicy:
                    description: Traffic policies that apply to this subset.
                    properties:
                      connectionPool:
                        properties:
                          http:
                            description: HTTP connection pool settings.
                            properties:
                              h2UpgradePolicy:
                                description: Specify if http1.1 connection should
                                  be upgraded to http2 for the associated destination.
                                enum:
                                - DEFAULT
                                - DO_NOT_UPGRADE
                                - UPGRADE
                                type: string
                              http1MaxPendingRequests:
                                description: Maximum number of pending HTTP requests
                                  to a destination.
                                format: int32
                                type: integer
                              http2MaxRequests:
                                description: Maximum number of requests to a backend.
                                format: int32
                                type: integer
                              idleTimeout:
                                description: The idle timeout for upstream connection
                                  pool connections.
                                type: string
                              maxRequestsPerConnection:
                                description: Maximum number of requests per connection
                                  to a backend.
                                format: int32
                                type: integer
                              maxRetries:
                                format: int32
                                type: integer
                              useClientProtocol:
                                description: If set to true, client protocol will
                                  be preserved while initiating connection to backend.
                                type: boolean
                            type: object
                          tcp:
                            description: Settings common to both HTTP and TCP upstream
                              connections.
                            properties:
                              connectTimeout:
                                description: TCP connection timeout.
                                type: string
                              maxConnections:
                                description: Maximum number of HTTP1 /TCP connections
                                  to a destination host.
                                format: int32
                                type: integer
                              tcpKeepalive:
                                description: If set then set SO_KEEPALIVE on the socket
                                  to enable TCP Keepalives.
                                properties:
                                  interval:
                                    description: The time duration between keep-alive
                                      probes.
                                    type: string
                                  probes:
                                    type: integer
                                  time:
                                    type: string
                                type: object
                            type: object
                        type: object
                      loadBalancer:
                        description: Settings controlling the load balancer algorithms.
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - simple
                            - properties:
                                consistentHash:
                                  oneOf:
                                  - not:
                                      anyOf:
                                      - required:
                                        - httpHeaderName
                                      - required:
                                        - httpCookie
                                      - required:
                                        - useSourceIp
                                      - required:
                                        - httpQueryParameterName
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              required:
                              - consistentHash
                        - required:
                          - simple
                        - properties:
                            consistentHash:
                              oneOf:
                              - not:
                                  anyOf:
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              - required:
                                - httpHeaderName
                              - required:
                                - httpCookie
                              - required:
                                - useSourceIp
                              - required:
                                - httpQueryParameterName
                          required:
                          - consistentHash
                        properties:
                          consistentHash:
                            properties:
                              httpCookie:
                                description: Hash based on HTTP cookie.
                                properties:
                                  name:
                                    description: Name of the cookie.
                                    format: string
                                    type: string
                                  path:
                                    description: Path to set for the cookie.
                                    format: string
                                    type: string
                                  ttl:
                                    description: Lifetime of the cookie.
                                    type: string
                                type: object
                              httpHeaderName:
                                description: Hash based on a specific HTTP header.
                                format: string
                                type: string
                              httpQueryParameterName:
                                description: Hash based on a specific HTTP query parameter.
                                format: string
                                type: string
                              minimumRingSize:
                                type: integer
                              useSourceIp:
                                description: Hash based on the source IP address.
                                type: boolean
                            type: object
                          localityLbSetting:
                            properties:
                              distribute:
                                description: 'Optional: only one of distribute or
                                  failover can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating locality, '/' separated,
                                        e.g.
                                      format: string
                                      type: string
                                    to:
                                      additionalProperties:
                                        type: integer
                                      description: Map of upstream localities to traffic
                                        distribution weights.
                                      type: object
                                  type: object
                                type: array
                              enabled:
                                description: enable locality load balancing, this
                                  is DestinationRule-level and will override mesh
                                  wide settings in entirety.
                                nullable: true
                                type: boolean
                              failover:
                                description: 'Optional: only failover or distribute
                                  can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating region.
                                      format: string
                                      type: string
                                    to:
                                      format: string
                                      type: string
                                  type: object
                                type: array
                            type: object
                          simple:
                            enum:
                            - ROUND_ROBIN
                            - LEAST_CONN
                            - RANDOM
                            - PASSTHROUGH
			    - PREDICTIVE_LEAST_CONN
                            type: string
                        type: object
                      outlierDetection:
                        properties:
                          baseEjectionTime:
                            description: Minimum ejection duration.
                            type: string
                          consecutive5xxErrors:
                            description: Number of 5xx errors before a host is ejected
                              from the connection pool.
                            nullable: true
                            type: integer
                          consecutiveErrors:
                            format: int32
                            type: integer
                          consecutiveGatewayErrors:
                            description: Number of gateway errors before a host is
                              ejected from the connection pool.
                            nullable: true
                            type: integer
                          interval:
                            description: Time interval between ejection sweep analysis.
                            type: string
                          maxEjectionPercent:
                            format: int32
                            type: integer
                          minHealthPercent:
                            format: int32
                            type: integer
                        type: object
                      portLevelSettings:
                        description: Traffic policies specific to individual ports.
                        items:
                          properties:
                            connectionPool:
                              properties:
                                http:
                                  description: HTTP connection pool settings.
                                  properties:
                                    h2UpgradePolicy:
                                      description: Specify if http1.1 connection should
                                        be upgraded to http2 for the associated destination.
                                      enum:
                                      - DEFAULT
                                      - DO_NOT_UPGRADE
                                      - UPGRADE
                                      type: string
                                    http1MaxPendingRequests:
                                      description: Maximum number of pending HTTP
                                        requests to a destination.
                                      format: int32
                                      type: integer
                                    http2MaxRequests:
                                      description: Maximum number of requests to a
                                        backend.
                                      format: int32
                                      type: integer
                                    idleTimeout:
                                      description: The idle timeout for upstream connection
                                        pool connections.
                                      type: string
                                    maxRequestsPerConnection:
                                      description: Maximum number of requests per
                                        connection to a backend.
                                      format: int32
                                      type: integer
                                    maxRetries:
                                      format: int32
                                      type: integer
                                    useClientProtocol:
                                      description: If set to true, client protocol
                                        will be preserved while initiating connection
                                        to backend.
                                      type: boolean
                                  type: object
                                tcp:
                                  description: Settings common to both HTTP and TCP
                                    upstream connections.
                                  properties:
                                    connectTimeout:
                                      description: TCP connection timeout.
                                      type: string
                                    maxConnections:
                                      description: Maximum number of HTTP1 /TCP connections
                                        to a destination host.
                                      format: int32
                                      type: integer
                                    tcpKeepalive:
                                      description: If set then set SO_KEEPALIVE on
                                        the socket to enable TCP Keepalives.
                                      properties:
                                        interval:
                                          description: The time duration between keep-alive
                                            probes.
                                          type: string
                                        probes:
                                          type: integer
                                        time:
                                          type: string
                                      type: object
                                  type: object
                              type: object
                            loadBalancer:
                              description: Settings controlling the load balancer
                                algorithms.
                              oneOf:
                              - not:
                                  anyOf:
                                  - required:
                                    - simple
                                  - properties:
                                      consistentHash:
                                        oneOf:
                                        - not:
                                            anyOf:
                                            - required:
                                              - httpHeaderName
                                            - required:
                                              - httpCookie
                                            - required:
                                              - useSourceIp
                                            - required:
                                              - httpQueryParameterName
                                        - required:
                                          - httpHeaderName
                                        - required:
                                          - httpCookie
                                        - required:
                                          - useSourceIp
                                        - required:
                                          - httpQueryParameterName
                                    required:
                                    - consistentHash
                              - required:
                                - simple
                              - properties:
                                  consistentHash:
                                    oneOf:
                                    - not:
                                        anyOf:
                                        - required:
                                          - httpHeaderName
                                        - required:
                                          - httpCookie
                                        - required:
                                          - useSourceIp
                                        - required:
                                          - httpQueryParameterName
                                    - required:
                                      - httpHeaderName
                                    - required:
                                      - httpCookie
                                    - required:
                                      - useSourceIp
                                    - required:
                                      - httpQueryParameterName
                                required:
                                - consistentHash
                              properties:
                                consistentHash:
                                  properties:
                                    httpCookie:
                                      description: Hash based on HTTP cookie.
                                      properties:
                                        name:
                                          description: Name of the cookie.
                                          format: string
                                          type: string
                                        path:
                                          description: Path to set for the cookie.
                                          format: string
                                          type: string
                                        ttl:
                                          description: Lifetime of the cookie.
                                          type: string
                                      type: object
                                    httpHeaderName:
                                      description: Hash based on a specific HTTP header.
                                      format: string
                                      type: string
                                    httpQueryParameterName:
                                      description: Hash based on a specific HTTP query
                                        parameter.
                                      format: string
                                      type: string
                                    minimumRingSize:
                                      type: integer
                                    useSourceIp:
                                      description: Hash based on the source IP address.
                                      type: boolean
                                  type: object
                                localityLbSetting:
                                  properties:
                                    distribute:
                                      description: 'Optional: only one of distribute
                                        or failover can be set.'
                                      items:
                                        properties:
                                          from:
                                            description: Originating locality, '/'
                                              separated, e.g.
                                            format: string
                                            type: string
                                          to:
                                            additionalProperties:
                                              type: integer
                                            description: Map of upstream localities
                                              to traffic distribution weights.
                                            type: object
                                        type: object
                                      type: array
                                    enabled:
                                      description: enable locality load balancing,
                                        this is DestinationRule-level and will override
                                        mesh wide settings in entirety.
                                      nullable: true
                                      type: boolean
                                    failover:
                                      description: 'Optional: only failover or distribute
                                        can be set.'
                                      items:
                                        properties:
                                          from:
                                            description: Originating region.
                                            format: string
                                            type: string
                                          to:
                                            format: string
                                            type: string
                                        type: object
                                      type: array
                                  type: object
                                simple:
                                  enum:
                                  - ROUND_ROBIN
                                  - LEAST_CONN
                                  - RANDOM
                                  - PASSTHROUGH
				  - PREDICTIVE_LEAST_CONN
                                  type: string
                              type: object
                            outlierDetection:
                              properties:
                                baseEjectionTime:
                                  description: Minimum ejection duration.
                                  type: string
                                consecutive5xxErrors:
                                  description: Number of 5xx errors before a host
                                    is ejected from the connection pool.
                                  nullable: true
                                  type: integer
                                consecutiveErrors:
                                  format: int32
                                  type: integer
                                consecutiveGatewayErrors:
                                  description: Number of gateway errors before a host
                                    is ejected from the connection pool.
                                  nullable: true
                                  type: integer
                                interval:
                                  description: Time interval between ejection sweep
                                    analysis.
                                  type: string
                                maxEjectionPercent:
                                  format: int32
                                  type: integer
                                minHealthPercent:
                                  format: int32
                                  type: integer
                              type: object
                            port:
                              properties:
                                number:
                                  type: integer
                              type: object
                            tls:
                              description: TLS related settings for connections to
                                the upstream service.
                              properties:
                                caCertificates:
                                  format: string
                                  type: string
                                clientCertificate:
                                  description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                                  format: string
                                  type: string
                                credentialName:
                                  format: string
                                  type: string
                                mode:
                                  enum:
                                  - DISABLE
                                  - SIMPLE
                                  - MUTUAL
                                  - ISTIO_MUTUAL
                                  type: string
                                privateKey:
                                  description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                                  format: string
                                  type: string
                                sni:
                                  description: SNI string to present to the server
                                    during TLS handshake.
                                  format: string
                                  type: string
                                subjectAltNames:
                                  items:
                                    format: string
                                    type: string
                                  type: array
                              type: object
                          type: object
                        type: array
                      tls:
                        description: TLS related settings for connections to the upstream
                          service.
                        properties:
                          caCertificates:
                            format: string
                            type: string
                          clientCertificate:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          credentialName:
                            format: string
                            type: string
                          mode:
                            enum:
                            - DISABLE
                            - SIMPLE
                            - MUTUAL
                            - ISTIO_MUTUAL
                            type: string
                          privateKey:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          sni:
                            description: SNI string to present to the server during
                              TLS handshake.
                            format: string
                            type: string
                          subjectAltNames:
                            items:
                              format: string
                              type: string
                            type: array
                        type: object
                    type: object
                type: object
              type: array
            trafficPolicy:
              properties:
                connectionPool:
                  properties:
                    http:
                      description: HTTP connection pool settings.
                      properties:
                        h2UpgradePolicy:
                          description: Specify if http1.1 connection should be upgraded
                            to http2 for the associated destination.
                          enum:
                          - DEFAULT
                          - DO_NOT_UPGRADE
                          - UPGRADE
                          type: string
                        http1MaxPendingRequests:
                          description: Maximum number of pending HTTP requests to
                            a destination.
                          format: int32
                          type: integer
                        http2MaxRequests:
                          description: Maximum number of requests to a backend.
                          format: int32
                          type: integer
                        idleTimeout:
                          description: The idle timeout for upstream connection pool
                            connections.
                          type: string
                        maxRequestsPerConnection:
                          description: Maximum number of requests per connection to
                            a backend.
                          format: int32
                          type: integer
                        maxRetries:
                          format: int32
                          type: integer
                        useClientProtocol:
                          description: If set to true, client protocol will be preserved
                            while initiating connection to backend.
                          type: boolean
                      type: object
                    tcp:
                      description: Settings common to both HTTP and TCP upstream connections.
                      properties:
                        connectTimeout:
                          description: TCP connection timeout.
                          type: string
                        maxConnections:
                          description: Maximum number of HTTP1 /TCP connections to
                            a destination host.
                          format: int32
                          type: integer
                        tcpKeepalive:
                          description: If set then set SO_KEEPALIVE on the socket
                            to enable TCP Keepalives.
                          properties:
                            interval:
                              description: The time duration between keep-alive probes.
                              type: string
                            probes:
                              type: integer
                            time:
                              type: string
                          type: object
                      type: object
                  type: object
                loadBalancer:
                  description: Settings controlling the load balancer algorithms.
                  oneOf:
                  - not:
                      anyOf:
                      - required:
                        - simple
                      - properties:
                          consistentHash:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - httpHeaderName
                                - required:
                                  - httpCookie
                                - required:
                                  - useSourceIp
                                - required:
                                  - httpQueryParameterName
                            - required:
                              - httpHeaderName
                            - required:
                              - httpCookie
                            - required:
                              - useSourceIp
                            - required:
                              - httpQueryParameterName
                        required:
                        - consistentHash
                  - required:
                    - simple
                  - properties:
                      consistentHash:
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - httpHeaderName
                            - required:
                              - httpCookie
                            - required:
                              - useSourceIp
                            - required:
                              - httpQueryParameterName
                        - required:
                          - httpHeaderName
                        - required:
                          - httpCookie
                        - required:
                          - useSourceIp
                        - required:
                          - httpQueryParameterName
                    required:
                    - consistentHash
                  properties:
                    consistentHash:
                      properties:
                        httpCookie:
                          description: Hash based on HTTP cookie.
                          properties:
                            name:
                              description: Name of the cookie.
                              format: string
                              type: string
                            path:
                              description: Path to set for the cookie.
                              format: string
                              type: string
                            ttl:
                              description: Lifetime of the cookie.
                              type: string
                          type: object
                        httpHeaderName:
                          description: Hash based on a specific HTTP header.
                          format: string
                          type: string
                        httpQueryParameterName:
                          description: Hash based on a specific HTTP query parameter.
                          format: string
                          type: string
                        minimumRingSize:
                          type: integer
                        useSourceIp:
                          description: Hash based on the source IP address.
                          type: boolean
                      type: object
                    localityLbSetting:
                      properties:
                        distribute:
                          description: 'Optional: only one of distribute or failover
                            can be set.'
                          items:
                            properties:
                              from:
                                description: Originating locality, '/' separated,
                                  e.g.
                                format: string
                                type: string
                              to:
                                additionalProperties:
                                  type: integer
                                description: Map of upstream localities to traffic
                                  distribution weights.
                                type: object
                            type: object
                          type: array
                        enabled:
                          description: enable locality load balancing, this is DestinationRule-level
                            and will override mesh wide settings in entirety.
                          nullable: true
                          type: boolean
                        failover:
                          description: 'Optional: only failover or distribute can
                            be set.'
                          items:
                            properties:
                              from:
                                description: Originating region.
                                format: string
                                type: string
                              to:
                                format: string
                                type: string
                            type: object
                          type: array
                      type: object
                    simple:
                      enum:
                      - ROUND_ROBIN
                      - LEAST_CONN
                      - RANDOM
                      - PASSTHROUGH
		      - PREDICTIVE_LEAST_CONN
                      type: string
                  type: object
                outlierDetection:
                  properties:
                    baseEjectionTime:
                      description: Minimum ejection duration.
                      type: string
                    consecutive5xxErrors:
                      description: Number of 5xx errors before a host is ejected from
                        the connection pool.
                      nullable: true
                      type: integer
                    consecutiveErrors:
                      format: int32
                      type: integer
                    consecutiveGatewayErrors:
                      description: Number of gateway errors before a host is ejected
                        from the connection pool.
                      nullable: true
                      type: integer
                    interval:
                      description: Time interval between ejection sweep analysis.
                      type: string
                    maxEjectionPercent:
                      format: int32
                      type: integer
                    minHealthPercent:
                      format: int32
                      type: integer
                  type: object
                portLevelSettings:
                  description: Traffic policies specific to individual ports.
                  items:
                    properties:
                      connectionPool:
                        properties:
                          http:
                            description: HTTP connection pool settings.
                            properties:
                              h2UpgradePolicy:
                                description: Specify if http1.1 connection should
                                  be upgraded to http2 for the associated destination.
                                enum:
                                - DEFAULT
                                - DO_NOT_UPGRADE
                                - UPGRADE
                                type: string
                              http1MaxPendingRequests:
                                description: Maximum number of pending HTTP requests
                                  to a destination.
                                format: int32
                                type: integer
                              http2MaxRequests:
                                description: Maximum number of requests to a backend.
                                format: int32
                                type: integer
                              idleTimeout:
                                description: The idle timeout for upstream connection
                                  pool connections.
                                type: string
                              maxRequestsPerConnection:
                                description: Maximum number of requests per connection
                                  to a backend.
                                format: int32
                                type: integer
                              maxRetries:
                                format: int32
                                type: integer
                              useClientProtocol:
                                description: If set to true, client protocol will
                                  be preserved while initiating connection to backend.
                                type: boolean
                            type: object
                          tcp:
                            description: Settings common to both HTTP and TCP upstream
                              connections.
                            properties:
                              connectTimeout:
                                description: TCP connection timeout.
                                type: string
                              maxConnections:
                                description: Maximum number of HTTP1 /TCP connections
                                  to a destination host.
                                format: int32
                                type: integer
                              tcpKeepalive:
                                description: If set then set SO_KEEPALIVE on the socket
                                  to enable TCP Keepalives.
                                properties:
                                  interval:
                                    description: The time duration between keep-alive
                                      probes.
                                    type: string
                                  probes:
                                    type: integer
                                  time:
                                    type: string
                                type: object
                            type: object
                        type: object
                      loadBalancer:
                        description: Settings controlling the load balancer algorithms.
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - simple
                            - properties:
                                consistentHash:
                                  oneOf:
                                  - not:
                                      anyOf:
                                      - required:
                                        - httpHeaderName
                                      - required:
                                        - httpCookie
                                      - required:
                                        - useSourceIp
                                      - required:
                                        - httpQueryParameterName
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              required:
                              - consistentHash
                        - required:
                          - simple
                        - properties:
                            consistentHash:
                              oneOf:
                              - not:
                                  anyOf:
                                  - required:
                                    - httpHeaderName
                                  - required:
                                    - httpCookie
                                  - required:
                                    - useSourceIp
                                  - required:
                                    - httpQueryParameterName
                              - required:
                                - httpHeaderName
                              - required:
                                - httpCookie
                              - required:
                                - useSourceIp
                              - required:
                                - httpQueryParameterName
                          required:
                          - consistentHash
                        properties:
                          consistentHash:
                            properties:
                              httpCookie:
                                description: Hash based on HTTP cookie.
                                properties:
                                  name:
                                    description: Name of the cookie.
                                    format: string
                                    type: string
                                  path:
                                    description: Path to set for the cookie.
                                    format: string
                                    type: string
                                  ttl:
                                    description: Lifetime of the cookie.
                                    type: string
                                type: object
                              httpHeaderName:
                                description: Hash based on a specific HTTP header.
                                format: string
                                type: string
                              httpQueryParameterName:
                                description: Hash based on a specific HTTP query parameter.
                                format: string
                                type: string
                              minimumRingSize:
                                type: integer
                              useSourceIp:
                                description: Hash based on the source IP address.
                                type: boolean
                            type: object
                          localityLbSetting:
                            properties:
                              distribute:
                                description: 'Optional: only one of distribute or
                                  failover can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating locality, '/' separated,
                                        e.g.
                                      format: string
                                      type: string
                                    to:
                                      additionalProperties:
                                        type: integer
                                      description: Map of upstream localities to traffic
                                        distribution weights.
                                      type: object
                                  type: object
                                type: array
                              enabled:
                                description: enable locality load balancing, this
                                  is DestinationRule-level and will override mesh
                                  wide settings in entirety.
                                nullable: true
                                type: boolean
                              failover:
                                description: 'Optional: only failover or distribute
                                  can be set.'
                                items:
                                  properties:
                                    from:
                                      description: Originating region.
                                      format: string
                                      type: string
                                    to:
                                      format: string
                                      type: string
                                  type: object
                                type: array
                            type: object
                          simple:
                            enum:
                            - ROUND_ROBIN
                            - LEAST_CONN
                            - RANDOM
                            - PASSTHROUGH
			    - PREDICTIVE_LEAST_CONN
                            type: string
                        type: object
                      outlierDetection:
                        properties:
                          baseEjectionTime:
                            description: Minimum ejection duration.
                            type: string
                          consecutive5xxErrors:
                            description: Number of 5xx errors before a host is ejected
                              from the connection pool.
                            nullable: true
                            type: integer
                          consecutiveErrors:
                            format: int32
                            type: integer
                          consecutiveGatewayErrors:
                            description: Number of gateway errors before a host is
                              ejected from the connection pool.
                            nullable: true
                            type: integer
                          interval:
                            description: Time interval between ejection sweep analysis.
                            type: string
                          maxEjectionPercent:
                            format: int32
                            type: integer
                          minHealthPercent:
                            format: int32
                            type: integer
                        type: object
                      port:
                        properties:
                          number:
                            type: integer
                        type: object
                      tls:
                        description: TLS related settings for connections to the upstream
                          service.
                        properties:
                          caCertificates:
                            format: string
                            type: string
                          clientCertificate:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          credentialName:
                            format: string
                            type: string
                          mode:
                            enum:
                            - DISABLE
                            - SIMPLE
                            - MUTUAL
                            - ISTIO_MUTUAL
                            type: string
                          privateKey:
                            description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                            format: string
                            type: string
                          sni:
                            description: SNI string to present to the server during
                              TLS handshake.
                            format: string
                            type: string
                          subjectAltNames:
                            items:
                              format: string
                              type: string
                            type: array
                        type: object
                    type: object
                  type: array
                tls:
                  description: TLS related settings for connections to the upstream
                    service.
                  properties:
                    caCertificates:
                      format: string
                      type: string
                    clientCertificate:
                      description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                      format: string
                      type: string
                    credentialName:
                      format: string
                      type: string
                    mode:
                      enum:
                      - DISABLE
                      - SIMPLE
                      - MUTUAL
                      - ISTIO_MUTUAL
                      type: string
                    privateKey:
                      description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                      format: string
                      type: string
                    sni:
                      description: SNI string to present to the server during TLS
                        handshake.
                      format: string
                      type: string
                    subjectAltNames:
                      items:
                        format: string
                        type: string
                      type: array
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: envoyfilters.networking.istio.io
spec:
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: EnvoyFilter
    listKind: EnvoyFilterList
    plural: envoyfilters
    singular: envoyfilter
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Customizing Envoy configuration generated by Istio. See more
            details at: https://istio.io/docs/reference/config/networking/envoy-filter.html'
          properties:
            configPatches:
              description: One or more patches with match conditions.
              items:
                properties:
                  applyTo:
                    enum:
                    - INVALID
                    - LISTENER
                    - FILTER_CHAIN
                    - NETWORK_FILTER
                    - HTTP_FILTER
                    - ROUTE_CONFIGURATION
                    - VIRTUAL_HOST
                    - HTTP_ROUTE
                    - CLUSTER
                    type: string
                  match:
                    description: Match on listener/route configuration/cluster.
                    oneOf:
                    - not:
                        anyOf:
                        - required:
                          - listener
                        - required:
                          - routeConfiguration
                        - required:
                          - cluster
                    - required:
                      - listener
                    - required:
                      - routeConfiguration
                    - required:
                      - cluster
                    properties:
                      cluster:
                        description: Match on envoy cluster attributes.
                        properties:
                          name:
                            description: The exact name of the cluster to match.
                            format: string
                            type: string
                          portNumber:
                            description: The service port for which this cluster was
                              generated.
                            type: integer
                          service:
                            description: The fully qualified service name for this
                              cluster.
                            format: string
                            type: string
                          subset:
                            description: The subset associated with the service.
                            format: string
                            type: string
                        type: object
                      context:
                        description: The specific config generation context to match
                          on.
                        enum:
                        - ANY
                        - SIDECAR_INBOUND
                        - SIDECAR_OUTBOUND
                        - GATEWAY
                        type: string
                      listener:
                        description: Match on envoy listener attributes.
                        properties:
                          filterChain:
                            description: Match a specific filter chain in a listener.
                            properties:
                              applicationProtocols:
                                description: Applies only to sidecars.
                                format: string
                                type: string
                              filter:
                                description: The name of a specific filter to apply
                                  the patch to.
                                properties:
                                  name:
                                    description: The filter name to match on.
                                    format: string
                                    type: string
                                  subFilter:
                                    properties:
                                      name:
                                        description: The filter name to match on.
                                        format: string
                                        type: string
                                    type: object
                                type: object
                              name:
                                description: The name assigned to the filter chain.
                                format: string
                                type: string
                              sni:
                                description: The SNI value used by a filter chain's
                                  match condition.
                                format: string
                                type: string
                              transportProtocol:
                                description: Applies only to SIDECAR_INBOUND context.
                                format: string
                                type: string
                            type: object
                          name:
                            description: Match a specific listener by its name.
                            format: string
                            type: string
                          portName:
                            format: string
                            type: string
                          portNumber:
                            type: integer
                        type: object
                      proxy:
                        description: Match on properties associated with a proxy.
                        properties:
                          metadata:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                          proxyVersion:
                            format: string
                            type: string
                        type: object
                      routeConfiguration:
                        description: Match on envoy HTTP route configuration attributes.
                        properties:
                          gateway:
                            format: string
                            type: string
                          name:
                            description: Route configuration name to match on.
                            format: string
                            type: string
                          portName:
                            description: Applicable only for GATEWAY context.
                            format: string
                            type: string
                          portNumber:
                            type: integer
                          vhost:
                            properties:
                              name:
                                format: string
                                type: string
                              route:
                                description: Match a specific route within the virtual
                                  host.
                                properties:
                                  action:
                                    description: Match a route with specific action
                                      type.
                                    enum:
                                    - ANY
                                    - ROUTE
                                    - REDIRECT
                                    - DIRECT_RESPONSE
                                    type: string
                                  name:
                                    format: string
                                    type: string
                                type: object
                            type: object
                        type: object
                    type: object
                  patch:
                    description: The patch to apply along with the operation.
                    properties:
                      filterClass:
                        description: Determines the filter insertion order.
                        enum:
                        - UNSPECIFIED
                        - AUTHN
                        - AUTHZ
                        - STATS
                        type: string
                      operation:
                        description: Determines how the patch should be applied.
                        enum:
                        - INVALID
                        - MERGE
                        - ADD
                        - REMOVE
                        - INSERT_BEFORE
                        - INSERT_AFTER
                        - INSERT_FIRST
                        - REPLACE
                        type: string
                      value:
                        description: The JSON config of the object being patched.
                        type: object
                    type: object
                type: object
              type: array
            workloadSelector:
              properties:
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: gateways.networking.istio.io
spec:
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: Gateway
    listKind: GatewayList
    plural: gateways
    shortNames:
    - gw
    singular: gateway
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting edge load balancer. See more details
            at: https://istio.io/docs/reference/config/networking/gateway.html'
          properties:
            selector:
              additionalProperties:
                format: string
                type: string
              type: object
            servers:
              description: A list of server specifications.
              items:
                properties:
                  bind:
                    format: string
                    type: string
                  defaultEndpoint:
                    format: string
                    type: string
                  hosts:
                    description: One or more hosts exposed by this gateway.
                    items:
                      format: string
                      type: string
                    type: array
                  name:
                    description: An optional name of the server, when set must be
                      unique across all servers.
                    format: string
                    type: string
                  port:
                    properties:
                      name:
                        description: Label assigned to the port.
                        format: string
                        type: string
                      number:
                        description: A valid non-negative integer port number.
                        type: integer
                      protocol:
                        description: The protocol exposed on the port.
                        format: string
                        type: string
                      targetPort:
                        type: integer
                    type: object
                  tls:
                    description: Set of TLS related options that govern the server's
                      behavior.
                    properties:
                      caCertificates:
                        description: REQUIRED if mode is `+"`"+`MUTUAL`+"`"+`.
                        format: string
                        type: string
                      cipherSuites:
                        description: 'Optional: If specified, only support the specified
                          cipher list.'
                        items:
                          format: string
                          type: string
                        type: array
                      credentialName:
                        format: string
                        type: string
                      httpsRedirect:
                        type: boolean
                      maxProtocolVersion:
                        description: 'Optional: Maximum TLS protocol version.'
                        enum:
                        - TLS_AUTO
                        - TLSV1_0
                        - TLSV1_1
                        - TLSV1_2
                        - TLSV1_3
                        type: string
                      minProtocolVersion:
                        description: 'Optional: Minimum TLS protocol version.'
                        enum:
                        - TLS_AUTO
                        - TLSV1_0
                        - TLSV1_1
                        - TLSV1_2
                        - TLSV1_3
                        type: string
                      mode:
                        enum:
                        - PASSTHROUGH
                        - SIMPLE
                        - MUTUAL
                        - AUTO_PASSTHROUGH
                        - ISTIO_MUTUAL
                        type: string
                      privateKey:
                        description: REQUIRED if mode is `+"`"+`SIMPLE`+"`"+` or `+"`"+`MUTUAL`+"`"+`.
                        format: string
                        type: string
                      serverCertificate:
                        description: REQUIRED if mode is `+"`"+`SIMPLE`+"`"+` or `+"`"+`MUTUAL`+"`"+`.
                        format: string
                        type: string
                      subjectAltNames:
                        items:
                          format: string
                          type: string
                        type: array
                      verifyCertificateHash:
                        items:
                          format: string
                          type: string
                        type: array
                      verifyCertificateSpki:
                        items:
                          format: string
                          type: string
                        type: array
                    type: object
                type: object
              type: array
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: serviceentries.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.hosts
    description: The hosts associated with the ServiceEntry
    name: Hosts
    type: string
  - JSONPath: .spec.location
    description: Whether the service is external to the mesh or part of the mesh (MESH_EXTERNAL
      or MESH_INTERNAL)
    name: Location
    type: string
  - JSONPath: .spec.resolution
    description: Service discovery mode for the hosts (NONE, STATIC, or DNS)
    name: Resolution
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: ServiceEntry
    listKind: ServiceEntryList
    plural: serviceentries
    shortNames:
    - se
    singular: serviceentry
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting service registry. See more details
            at: https://istio.io/docs/reference/config/networking/service-entry.html'
          properties:
            addresses:
              description: The virtual IP addresses associated with the service.
              items:
                format: string
                type: string
              type: array
            endpoints:
              description: One or more endpoints associated with the service.
              items:
                properties:
                  address:
                    format: string
                    type: string
                  labels:
                    additionalProperties:
                      format: string
                      type: string
                    description: One or more labels associated with the endpoint.
                    type: object
                  locality:
                    description: The locality associated with the endpoint.
                    format: string
                    type: string
                  network:
                    format: string
                    type: string
                  ports:
                    additionalProperties:
                      type: integer
                    description: Set of ports associated with the endpoint.
                    type: object
                  serviceAccount:
                    format: string
                    type: string
                  weight:
                    description: The load balancing weight associated with the endpoint.
                    type: integer
                type: object
              type: array
            exportTo:
              description: A list of namespaces to which this service is exported.
              items:
                format: string
                type: string
              type: array
            hosts:
              description: The hosts associated with the ServiceEntry.
              items:
                format: string
                type: string
              type: array
            location:
              enum:
              - MESH_EXTERNAL
              - MESH_INTERNAL
              type: string
            ports:
              description: The ports associated with the external service.
              items:
                properties:
                  name:
                    description: Label assigned to the port.
                    format: string
                    type: string
                  number:
                    description: A valid non-negative integer port number.
                    type: integer
                  protocol:
                    description: The protocol exposed on the port.
                    format: string
                    type: string
                  targetPort:
                    type: integer
                type: object
              type: array
            resolution:
              description: Service discovery mode for the hosts.
              enum:
              - NONE
              - STATIC
              - DNS
              type: string
            subjectAltNames:
              items:
                format: string
                type: string
              type: array
            workloadSelector:
              description: Applicable only for MESH_INTERNAL services.
              properties:
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: sidecars.networking.istio.io
spec:
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: Sidecar
    listKind: SidecarList
    plural: sidecars
    singular: sidecar
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting network reachability of a sidecar.
            See more details at: https://istio.io/docs/reference/config/networking/sidecar.html'
          properties:
            egress:
              items:
                properties:
                  bind:
                    format: string
                    type: string
                  captureMode:
                    enum:
                    - DEFAULT
                    - IPTABLES
                    - NONE
                    type: string
                  hosts:
                    items:
                      format: string
                      type: string
                    type: array
                  port:
                    description: The port associated with the listener.
                    properties:
                      name:
                        description: Label assigned to the port.
                        format: string
                        type: string
                      number:
                        description: A valid non-negative integer port number.
                        type: integer
                      protocol:
                        description: The protocol exposed on the port.
                        format: string
                        type: string
                      targetPort:
                        type: integer
                    type: object
                type: object
              type: array
            ingress:
              items:
                properties:
                  bind:
                    description: The IP to which the listener should be bound.
                    format: string
                    type: string
                  captureMode:
                    enum:
                    - DEFAULT
                    - IPTABLES
                    - NONE
                    type: string
                  defaultEndpoint:
                    format: string
                    type: string
                  port:
                    description: The port associated with the listener.
                    properties:
                      name:
                        description: Label assigned to the port.
                        format: string
                        type: string
                      number:
                        description: A valid non-negative integer port number.
                        type: integer
                      protocol:
                        description: The protocol exposed on the port.
                        format: string
                        type: string
                      targetPort:
                        type: integer
                    type: object
                type: object
              type: array
            outboundTrafficPolicy:
              description: Configuration for the outbound traffic policy.
              properties:
                egressProxy:
                  properties:
                    host:
                      description: The name of a service from the service registry.
                      format: string
                      type: string
                    port:
                      description: Specifies the port on the host that is being addressed.
                      properties:
                        number:
                          type: integer
                      type: object
                    subset:
                      description: The name of a subset within the service.
                      format: string
                      type: string
                  type: object
                mode:
                  enum:
                  - REGISTRY_ONLY
                  - ALLOW_ANY
                  type: string
              type: object
            workloadSelector:
              properties:
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: virtualservices.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.gateways
    description: The names of gateways and sidecars that should apply these routes
    name: Gateways
    type: string
  - JSONPath: .spec.hosts
    description: The destination hosts to which traffic is being sent
    name: Hosts
    type: string
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: VirtualService
    listKind: VirtualServiceList
    plural: virtualservices
    shortNames:
    - vs
    singular: virtualservice
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting label/content routing, sni routing,
            etc. See more details at: https://istio.io/docs/reference/config/networking/virtual-service.html'
          properties:
            exportTo:
              description: A list of namespaces to which this virtual service is exported.
              items:
                format: string
                type: string
              type: array
            gateways:
              description: The names of gateways and sidecars that should apply these
                routes.
              items:
                format: string
                type: string
              type: array
            hosts:
              description: The destination hosts to which traffic is being sent.
              items:
                format: string
                type: string
              type: array
            http:
              description: An ordered list of route rules for HTTP traffic.
              items:
                properties:
                  corsPolicy:
                    description: Cross-Origin Resource Sharing policy (CORS).
                    properties:
                      allowCredentials:
                        nullable: true
                        type: boolean
                      allowHeaders:
                        items:
                          format: string
                          type: string
                        type: array
                      allowMethods:
                        description: List of HTTP methods allowed to access the resource.
                        items:
                          format: string
                          type: string
                        type: array
                      allowOrigin:
                        description: The list of origins that are allowed to perform
                          CORS requests.
                        items:
                          format: string
                          type: string
                        type: array
                      allowOrigins:
                        description: String patterns that match allowed origins.
                        items:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        type: array
                      exposeHeaders:
                        items:
                          format: string
                          type: string
                        type: array
                      maxAge:
                        type: string
                    type: object
                  delegate:
                    properties:
                      name:
                        description: Name specifies the name of the delegate VirtualService.
                        format: string
                        type: string
                      namespace:
                        description: Namespace specifies the namespace where the delegate
                          VirtualService resides.
                        format: string
                        type: string
                    type: object
                  fault:
                    description: Fault injection policy to apply on HTTP traffic at
                      the client side.
                    properties:
                      abort:
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - httpStatus
                            - required:
                              - grpcStatus
                            - required:
                              - http2Error
                        - required:
                          - httpStatus
                        - required:
                          - grpcStatus
                        - required:
                          - http2Error
                        properties:
                          grpcStatus:
                            format: string
                            type: string
                          http2Error:
                            format: string
                            type: string
                          httpStatus:
                            description: HTTP status code to use to abort the Http
                              request.
                            format: int32
                            type: integer
                          percentage:
                            description: Percentage of requests to be aborted with
                              the error code provided.
                            properties:
                              value:
                                format: double
                                type: number
                            type: object
                        type: object
                      delay:
                        oneOf:
                        - not:
                            anyOf:
                            - required:
                              - fixedDelay
                            - required:
                              - exponentialDelay
                        - required:
                          - fixedDelay
                        - required:
                          - exponentialDelay
                        properties:
                          exponentialDelay:
                            type: string
                          fixedDelay:
                            description: Add a fixed delay before forwarding the request.
                            type: string
                          percent:
                            description: Percentage of requests on which the delay
                              will be injected (0-100).
                            format: int32
                            type: integer
                          percentage:
                            description: Percentage of requests on which the delay
                              will be injected.
                            properties:
                              value:
                                format: double
                                type: number
                            type: object
                        type: object
                    type: object
                  headers:
                    properties:
                      request:
                        properties:
                          add:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                          remove:
                            items:
                              format: string
                              type: string
                            type: array
                          set:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                        type: object
                      response:
                        properties:
                          add:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                          remove:
                            items:
                              format: string
                              type: string
                            type: array
                          set:
                            additionalProperties:
                              format: string
                              type: string
                            type: object
                        type: object
                    type: object
                  match:
                    items:
                      properties:
                        authority:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        gateways:
                          description: Names of gateways where the rule should be
                            applied.
                          items:
                            format: string
                            type: string
                          type: array
                        headers:
                          additionalProperties:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - exact
                                - required:
                                  - prefix
                                - required:
                                  - regex
                            - required:
                              - exact
                            - required:
                              - prefix
                            - required:
                              - regex
                            properties:
                              exact:
                                format: string
                                type: string
                              prefix:
                                format: string
                                type: string
                              regex:
                                description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                                format: string
                                type: string
                            type: object
                          type: object
                        ignoreUriCase:
                          description: Flag to specify whether the URI matching should
                            be case-insensitive.
                          type: boolean
                        method:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        name:
                          description: The name assigned to a match.
                          format: string
                          type: string
                        port:
                          description: Specifies the ports on the host that is being
                            addressed.
                          type: integer
                        queryParams:
                          additionalProperties:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - exact
                                - required:
                                  - prefix
                                - required:
                                  - regex
                            - required:
                              - exact
                            - required:
                              - prefix
                            - required:
                              - regex
                            properties:
                              exact:
                                format: string
                                type: string
                              prefix:
                                format: string
                                type: string
                              regex:
                                description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                                format: string
                                type: string
                            type: object
                          description: Query parameters for matching.
                          type: object
                        scheme:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        sourceLabels:
                          additionalProperties:
                            format: string
                            type: string
                          type: object
                        sourceNamespace:
                          description: Source namespace constraining the applicability
                            of a rule to workloads in that namespace.
                          format: string
                          type: string
                        uri:
                          oneOf:
                          - not:
                              anyOf:
                              - required:
                                - exact
                              - required:
                                - prefix
                              - required:
                                - regex
                          - required:
                            - exact
                          - required:
                            - prefix
                          - required:
                            - regex
                          properties:
                            exact:
                              format: string
                              type: string
                            prefix:
                              format: string
                              type: string
                            regex:
                              description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                              format: string
                              type: string
                          type: object
                        withoutHeaders:
                          additionalProperties:
                            oneOf:
                            - not:
                                anyOf:
                                - required:
                                  - exact
                                - required:
                                  - prefix
                                - required:
                                  - regex
                            - required:
                              - exact
                            - required:
                              - prefix
                            - required:
                              - regex
                            properties:
                              exact:
                                format: string
                                type: string
                              prefix:
                                format: string
                                type: string
                              regex:
                                description: RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
                                format: string
                                type: string
                            type: object
                          description: withoutHeader has the same syntax with the
                            header, but has opposite meaning.
                          type: object
                      type: object
                    type: array
                  mirror:
                    properties:
                      host:
                        description: The name of a service from the service registry.
                        format: string
                        type: string
                      port:
                        description: Specifies the port on the host that is being
                          addressed.
                        properties:
                          number:
                            type: integer
                        type: object
                      subset:
                        description: The name of a subset within the service.
                        format: string
                        type: string
                    type: object
                  mirror_percent:
                    description: Percentage of the traffic to be mirrored by the `+"`"+`mirror`+"`"+`
                      field.
                    nullable: true
                    type: integer
                  mirrorPercent:
                    description: Percentage of the traffic to be mirrored by the `+"`"+`mirror`+"`"+`
                      field.
                    nullable: true
                    type: integer
                  mirrorPercentage:
                    description: Percentage of the traffic to be mirrored by the `+"`"+`mirror`+"`"+`
                      field.
                    properties:
                      value:
                        format: double
                        type: number
                    type: object
                  name:
                    description: The name assigned to the route for debugging purposes.
                    format: string
                    type: string
                  redirect:
                    description: A HTTP rule can either redirect or forward (default)
                      traffic.
                    properties:
                      authority:
                        format: string
                        type: string
                      redirectCode:
                        type: integer
                      uri:
                        format: string
                        type: string
                    type: object
                  retries:
                    description: Retry policy for HTTP requests.
                    properties:
                      attempts:
                        description: Number of retries for a given request.
                        format: int32
                        type: integer
                      perTryTimeout:
                        description: Timeout per retry attempt for a given request.
                        type: string
                      retryOn:
                        description: Specifies the conditions under which retry takes
                          place.
                        format: string
                        type: string
                      retryRemoteLocalities:
                        description: Flag to specify whether the retries should retry
                          to other localities.
                        nullable: true
                        type: boolean
                    type: object
                  rewrite:
                    description: Rewrite HTTP URIs and Authority headers.
                    properties:
                      authority:
                        description: rewrite the Authority/Host header with this value.
                        format: string
                        type: string
                      uri:
                        format: string
                        type: string
                    type: object
                  route:
                    description: A HTTP rule can either redirect or forward (default)
                      traffic.
                    items:
                      properties:
                        destination:
                          properties:
                            host:
                              description: The name of a service from the service
                                registry.
                              format: string
                              type: string
                            port:
                              description: Specifies the port on the host that is
                                being addressed.
                              properties:
                                number:
                                  type: integer
                              type: object
                            subset:
                              description: The name of a subset within the service.
                              format: string
                              type: string
                          type: object
                        headers:
                          properties:
                            request:
                              properties:
                                add:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                                remove:
                                  items:
                                    format: string
                                    type: string
                                  type: array
                                set:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                              type: object
                            response:
                              properties:
                                add:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                                remove:
                                  items:
                                    format: string
                                    type: string
                                  type: array
                                set:
                                  additionalProperties:
                                    format: string
                                    type: string
                                  type: object
                              type: object
                          type: object
                        weight:
                          format: int32
                          type: integer
                      type: object
                    type: array
                  timeout:
                    description: Timeout for HTTP requests, default is disabled.
                    type: string
                type: object
              type: array
            tcp:
              description: An ordered list of route rules for opaque TCP traffic.
              items:
                properties:
                  match:
                    items:
                      properties:
                        destinationSubnets:
                          description: IPv4 or IPv6 ip addresses of destination with
                            optional subnet.
                          items:
                            format: string
                            type: string
                          type: array
                        gateways:
                          description: Names of gateways where the rule should be
                            applied.
                          items:
                            format: string
                            type: string
                          type: array
                        port:
                          description: Specifies the port on the host that is being
                            addressed.
                          type: integer
                        sourceLabels:
                          additionalProperties:
                            format: string
                            type: string
                          type: object
                        sourceNamespace:
                          description: Source namespace constraining the applicability
                            of a rule to workloads in that namespace.
                          format: string
                          type: string
                        sourceSubnet:
                          description: IPv4 or IPv6 ip address of source with optional
                            subnet.
                          format: string
                          type: string
                      type: object
                    type: array
                  route:
                    description: The destination to which the connection should be
                      forwarded to.
                    items:
                      properties:
                        destination:
                          properties:
                            host:
                              description: The name of a service from the service
                                registry.
                              format: string
                              type: string
                            port:
                              description: Specifies the port on the host that is
                                being addressed.
                              properties:
                                number:
                                  type: integer
                              type: object
                            subset:
                              description: The name of a subset within the service.
                              format: string
                              type: string
                          type: object
                        weight:
                          format: int32
                          type: integer
                      type: object
                    type: array
                type: object
              type: array
            tls:
              items:
                properties:
                  match:
                    items:
                      properties:
                        destinationSubnets:
                          description: IPv4 or IPv6 ip addresses of destination with
                            optional subnet.
                          items:
                            format: string
                            type: string
                          type: array
                        gateways:
                          description: Names of gateways where the rule should be
                            applied.
                          items:
                            format: string
                            type: string
                          type: array
                        port:
                          description: Specifies the port on the host that is being
                            addressed.
                          type: integer
                        sniHosts:
                          description: SNI (server name indicator) to match on.
                          items:
                            format: string
                            type: string
                          type: array
                        sourceLabels:
                          additionalProperties:
                            format: string
                            type: string
                          type: object
                        sourceNamespace:
                          description: Source namespace constraining the applicability
                            of a rule to workloads in that namespace.
                          format: string
                          type: string
                      type: object
                    type: array
                  route:
                    description: The destination to which the connection should be
                      forwarded to.
                    items:
                      properties:
                        destination:
                          properties:
                            host:
                              description: The name of a service from the service
                                registry.
                              format: string
                              type: string
                            port:
                              description: Specifies the port on the host that is
                                being addressed.
                              properties:
                                number:
                                  type: integer
                              type: object
                            subset:
                              description: The name of a subset within the service.
                              format: string
                              type: string
                          type: object
                        weight:
                          format: int32
                          type: integer
                      type: object
                    type: array
                type: object
              type: array
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    release: istio
  name: workloadentries.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  - JSONPath: .spec.address
    description: Address associated with the network endpoint.
    name: Address
    type: string
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: WorkloadEntry
    listKind: WorkloadEntryList
    plural: workloadentries
    shortNames:
    - we
    singular: workloadentry
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration affecting VMs onboarded into the mesh. See more
            details at: https://istio.io/docs/reference/config/networking/workload-entry.html'
          properties:
            address:
              format: string
              type: string
            labels:
              additionalProperties:
                format: string
                type: string
              description: One or more labels associated with the endpoint.
              type: object
            locality:
              description: The locality associated with the endpoint.
              format: string
              type: string
            network:
              format: string
              type: string
            ports:
              additionalProperties:
                type: integer
              description: Set of ports associated with the endpoint.
              type: object
            serviceAccount:
              format: string
              type: string
            weight:
              description: The load balancing weight associated with the endpoint.
              type: integer
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: workloadgroups.networking.istio.io
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: networking.istio.io
  names:
    categories:
    - istio-io
    - networking-istio-io
    kind: WorkloadGroup
    listKind: WorkloadGroupList
    plural: workloadgroups
    shortNames:
    - wg
    singular: workloadgroup
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Describes a collection of workload instances. See more details
            at: https://istio.io/docs/reference/config/networking/workload-group.html'
          properties:
            metadata:
              description: Metadata that will be used for all corresponding `+"`"+`WorkloadEntries`+"`"+`.
              properties:
                annotations:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
            probe:
              description: '`+"`"+`ReadinessProbe`+"`"+` describes the configuration the user
                must provide for healthchecking on their workload.'
              oneOf:
              - not:
                  anyOf:
                  - required:
                    - httpGet
                  - required:
                    - tcpSocket
                  - required:
                    - exec
              - required:
                - httpGet
              - required:
                - tcpSocket
              - required:
                - exec
              properties:
                exec:
                  description: health is determined by how the command that is executed
                    exited.
                  properties:
                    command:
                      description: command to run.
                      items:
                        format: string
                        type: string
                      type: array
                  type: object
                failureThreshold:
                  description: Minimum consecutive failures for the probe to be considered
                    failed after having succeeded.
                  format: int32
                  type: integer
                httpGet:
                  properties:
                    host:
                      description: Host name to connect to, defaults to the pod IP.
                      format: string
                      type: string
                    httpHeaders:
                      description: headers the proxy will pass on to make the request.
                      items:
                        properties:
                          name:
                            format: string
                            type: string
                          value:
                            format: string
                            type: string
                        type: object
                      type: array
                    path:
                      description: Path to access on the HTTP server.
                      format: string
                      type: string
                    port:
                      description: port on which the endpoint lives.
                      type: integer
                    scheme:
                      format: string
                      type: string
                  type: object
                initialDelaySeconds:
                  description: Number of seconds after the container has started before
                    readiness probes are initiated.
                  format: int32
                  type: integer
                periodSeconds:
                  description: How often (in seconds) to perform the probe.
                  format: int32
                  type: integer
                successThreshold:
                  description: Minimum consecutive successes for the probe to be considered
                    successful after having failed.
                  format: int32
                  type: integer
                tcpSocket:
                  description: health is determined by if the proxy is able to connect.
                  properties:
                    host:
                      format: string
                      type: string
                    port:
                      type: integer
                  type: object
                timeoutSeconds:
                  description: Number of seconds after which the probe times out.
                  format: int32
                  type: integer
              type: object
            template:
              description: Template to be used for the generation of `+"`"+`WorkloadEntry`+"`"+`
                resources that belong to this `+"`"+`WorkloadGroup`+"`"+`.
              properties:
                address:
                  format: string
                  type: string
                labels:
                  additionalProperties:
                    format: string
                    type: string
                  description: One or more labels associated with the endpoint.
                  type: object
                locality:
                  description: The locality associated with the endpoint.
                  format: string
                  type: string
                network:
                  format: string
                  type: string
                ports:
                  additionalProperties:
                    type: integer
                  description: Set of ports associated with the endpoint.
                  type: object
                serviceAccount:
                  format: string
                  type: string
                weight:
                  description: The load balancing weight associated with the endpoint.
                  type: integer
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1alpha3
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    istio: security
    release: istio
  name: authorizationpolicies.security.istio.io
spec:
  group: security.istio.io
  names:
    categories:
    - istio-io
    - security-istio-io
    kind: AuthorizationPolicy
    listKind: AuthorizationPolicyList
    plural: authorizationpolicies
    singular: authorizationpolicy
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: 'Configuration for access control on workloads. See more details
            at: https://istio.io/docs/reference/config/security/authorization-policy.html'
          properties:
            action:
              description: Optional.
              enum:
              - ALLOW
              - DENY
              - AUDIT
              type: string
            rules:
              description: Optional.
              items:
                properties:
                  from:
                    description: Optional.
                    items:
                      properties:
                        source:
                          description: Source specifies the source of a request.
                          properties:
                            ipBlocks:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            namespaces:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notIpBlocks:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notNamespaces:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notPrincipals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notRequestPrincipals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            principals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            requestPrincipals:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                          type: object
                      type: object
                    type: array
                  to:
                    description: Optional.
                    items:
                      properties:
                        operation:
                          description: Operation specifies the operation of a request.
                          properties:
                            hosts:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            methods:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notHosts:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notMethods:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notPaths:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            notPorts:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            paths:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                            ports:
                              description: Optional.
                              items:
                                format: string
                                type: string
                              type: array
                          type: object
                      type: object
                    type: array
                  when:
                    description: Optional.
                    items:
                      properties:
                        key:
                          description: The name of an Istio attribute.
                          format: string
                          type: string
                        notValues:
                          description: Optional.
                          items:
                            format: string
                            type: string
                          type: array
                        values:
                          description: Optional.
                          items:
                            format: string
                            type: string
                          type: array
                      type: object
                    type: array
                type: object
              type: array
            selector:
              description: Optional.
              properties:
                matchLabels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1beta1
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    istio: security
    release: istio
  name: peerauthentications.security.istio.io
spec:
  group: security.istio.io
  names:
    categories:
    - istio-io
    - security-istio-io
    kind: PeerAuthentication
    listKind: PeerAuthenticationList
    plural: peerauthentications
    shortNames:
    - pa
    singular: peerauthentication
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: PeerAuthentication defines how traffic will be tunneled (or
            not) to the sidecar.
          properties:
            mtls:
              description: Mutual TLS settings for workload.
              properties:
                mode:
                  description: Defines the mTLS mode used for peer authentication.
                  enum:
                  - UNSET
                  - DISABLE
                  - PERMISSIVE
                  - STRICT
                  type: string
              type: object
            portLevelMtls:
              additionalProperties:
                properties:
                  mode:
                    description: Defines the mTLS mode used for peer authentication.
                    enum:
                    - UNSET
                    - DISABLE
                    - PERMISSIVE
                    - STRICT
                    type: string
                type: object
              description: Port specific mutual TLS settings.
              type: object
            selector:
              description: The selector determines the workloads to apply the ChannelAuthentication
                on.
              properties:
                matchLabels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1beta1
    served: true
    storage: true

---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    "helm.sh/resource-policy": keep
  labels:
    app: istio-pilot
    chart: istio
    heritage: Tiller
    istio: security
    release: istio
  name: requestauthentications.security.istio.io
spec:
  group: security.istio.io
  names:
    categories:
    - istio-io
    - security-istio-io
    kind: RequestAuthentication
    listKind: RequestAuthenticationList
    plural: requestauthentications
    shortNames:
    - ra
    singular: requestauthentication
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          description: RequestAuthentication defines what request authentication methods
            are supported by a workload.
          properties:
            jwtRules:
              description: Define the list of JWTs that can be validated at the selected
                workloads' proxy.
              items:
                properties:
                  audiences:
                    items:
                      format: string
                      type: string
                    type: array
                  forwardOriginalToken:
                    description: If set to true, the orginal token will be kept for
                      the ustream request.
                    type: boolean
                  fromHeaders:
                    description: List of header locations from which JWT is expected.
                    items:
                      properties:
                        name:
                          description: The HTTP header name.
                          format: string
                          type: string
                        prefix:
                          description: The prefix that should be stripped before decoding
                            the token.
                          format: string
                          type: string
                      type: object
                    type: array
                  fromParams:
                    description: List of query parameters from which JWT is expected.
                    items:
                      format: string
                      type: string
                    type: array
                  issuer:
                    description: Identifies the issuer that issued the JWT.
                    format: string
                    type: string
                  jwks:
                    description: JSON Web Key Set of public keys to validate signature
                      of the JWT.
                    format: string
                    type: string
                  jwks_uri:
                    format: string
                    type: string
                  jwksUri:
                    format: string
                    type: string
                  outputPayloadToHeader:
                    format: string
                    type: string
                type: object
              type: array
            selector:
              description: The selector determines the workloads to apply the RequestAuthentication
                on.
              properties:
                matchLabels:
                  additionalProperties:
                    format: string
                    type: string
                  type: object
              type: object
          type: object
        status:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      type: object
  versions:
  - name: v1beta1
    served: true
    storage: true

---

---
# Source: crds/crd-operator.yaml
# SYNC WITH manifests/charts/istio-operator/templates
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: istiooperators.install.istio.io
  labels:
    release: istio
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.revision
    description: Istio control plane revision
    name: Revision
    type: string
  - JSONPath: .status.status
    description: IOP current state
    type: string
    name: Status
  - JSONPath: .metadata.creationTimestamp
    description: 'CreationTimestamp is a timestamp representing the server time when
      this object was created. It is not guaranteed to be set in happens-before order
      across separate operations. Clients may not set this value. It is represented
      in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for
      lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata'
    name: Age
    type: date
  group: install.istio.io
  names:
    kind: IstioOperator
    plural: istiooperators
    singular: istiooperator
    shortNames:
    - iop
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        spec:
          description: 'Specification of the desired state of the istio control plane resource.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
        status:
          description: 'Status describes each of istio control plane component status at the current time.
            0 means NONE, 1 means UPDATING, 2 means HEALTHY, 3 means ERROR, 4 means RECONCILING.
            More info: https://github.com/istio/api/blob/master/operator/v1alpha1/istio.operator.v1alpha1.pb.html &
            https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
---

---
# Source: base/templates/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: istio-reader-service-account
  namespace: istio-system
  labels:
    app: istio-reader
    release: istio
---
# Source: base/templates/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: istiod-service-account
  namespace: istio-system
  labels:
    app: istiod
    release: istio
---
# Source: base/templates/clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istiod-istio-system
  labels:
    app: istiod
    release: istio
rules:
  # sidecar injection controller
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["mutatingwebhookconfigurations"]
    verbs: ["get", "list", "watch", "patch"]

  # configuration validation webhook controller
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["validatingwebhookconfigurations"]
    verbs: ["get", "list", "watch", "update"]

  # istio configuration
  - apiGroups: ["config.istio.io", "security.istio.io", "networking.istio.io", "authentication.istio.io"]
    verbs: ["get", "watch", "list"]
    resources: ["*"]

  # auto-detect installed CRD definitions
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["get", "list", "watch"]

  # discovery and routing
  - apiGroups: [""]
    resources: ["pods", "nodes", "services", "namespaces", "endpoints"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["discovery.k8s.io"]
    resources: ["endpointslices"]
    verbs: ["get", "list", "watch"]

  # ingress controller
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses", "ingressclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses/status"]
    verbs: ["*"]

  # required for CA's namespace controller
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["create", "get", "list", "watch", "update"]

  # Istiod and bootstrap.
  - apiGroups: ["certificates.k8s.io"]
    resources:
      - "certificatesigningrequests"
      - "certificatesigningrequests/approval"
      - "certificatesigningrequests/status"
    verbs: ["update", "create", "get", "delete", "watch"]
  - apiGroups: ["certificates.k8s.io"]
    resources:
      - "signers"
    resourceNames:
    - "kubernetes.io/legacy-unknown"
    verbs: ["approve"]

  # Used by Istiod to verify the JWT tokens
  - apiGroups: ["authentication.k8s.io"]
    resources: ["tokenreviews"]
    verbs: ["create"]

  # Used by Istiod to verify gateway SDS
  - apiGroups: ["authorization.k8s.io"]
    resources: ["subjectaccessreviews"]
    verbs: ["create"]

  # Use for Kubernetes Service APIs
  - apiGroups: ["networking.x-k8s.io"]
    resources: ["*"]
    verbs: ["get", "watch", "list"]

  # Needed for multicluster secret reading, possibly ingress certs in the future
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "watch", "list"]
---
# Source: base/templates/clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istio-reader-istio-system
  labels:
    app: istio-reader
    release: istio
rules:
  - apiGroups:
      - "config.istio.io"
      - "security.istio.io"
      - "networking.istio.io"
      - "authentication.istio.io"
    resources: ["*"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["endpoints", "pods", "services", "nodes", "replicationcontrollers", "namespaces"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["apps"]
    resources: ["replicasets"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["authentication.k8s.io"]
    resources: ["tokenreviews"]
    verbs: ["create"]
  - apiGroups: ["authorization.k8s.io"]
    resources: ["subjectaccessreviews"]
    verbs: ["create"]
---
# Source: base/templates/clusterrolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istio-reader-istio-system
  labels:
    app: istio-reader
    release: istio
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istio-reader-istio-system
subjects:
  - kind: ServiceAccount
    name: istio-reader-service-account
    namespace: istio-system
---
# Source: base/templates/clusterrolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istiod-istio-system
  labels:
    app: istiod
    release: istio
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istiod-istio-system
subjects:
  - kind: ServiceAccount
    name: istiod-service-account
    namespace: istio-system
---
# Source: base/templates/role.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: istiod-istio-system
  namespace: istio-system
  labels:
    app: istiod
    release: istio
rules:
# permissions to verify the webhook is ready and rejecting
# invalid config. We use --server-dry-run so no config is persisted.
- apiGroups: ["networking.istio.io"]
  verbs: ["create"]
  resources: ["gateways"]

# For storing CA secret
- apiGroups: [""]
  resources: ["secrets"]
  # TODO lock this down to istio-ca-cert if not using the DNS cert mesh config
  verbs: ["create", "get", "watch", "list", "update", "delete"]
---
# Source: base/templates/rolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: istiod-istio-system
  namespace: istio-system
  labels:
    app: istiod
    release: istio
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: istiod-istio-system
subjects:
  - kind: ServiceAccount
    name: istiod-service-account
    namespace: istio-system
---
# Source: base/templates/validatingwebhookconfiguration.yaml
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: istiod-istio-system
  labels:
    app: istiod
    release: istio
    istio: istiod
webhooks:
  - name: validation.istio.io
    clientConfig:
      service:
        name: istiod
        namespace: istio-system
        path: "/validate"
      caBundle: "" # patched at runtime when the webhook is ready.
    rules:
      - operations:
        - CREATE
        - UPDATE
        apiGroups:
        - config.istio.io
        - security.istio.io
        - authentication.istio.io
        - networking.istio.io
        apiVersions:
        - "*"
        resources:
        - "*"
    # Fail open until the validation webhook is ready. The webhook controller
    # will update this to `+"`"+`Fail`+"`"+` and patch in the `+"`"+`caBundle`+"`"+` when the webhook
    # endpoint is ready.
    failurePolicy: Ignore
    sideEffects: None
    admissionReviewVersions: ["v1beta1", "v1"]
`)

func chartsBaseFilesGenIstioClusterYamlBytes() ([]byte, error) {
	return _chartsBaseFilesGenIstioClusterYaml, nil
}

func chartsBaseFilesGenIstioClusterYaml() (*asset, error) {
	bytes, err := chartsBaseFilesGenIstioClusterYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/files/gen-istio-cluster.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseKustomizationYaml = []byte(`apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - files/gen-istio-cluster.yaml
`)

func chartsBaseKustomizationYamlBytes() ([]byte, error) {
	return _chartsBaseKustomizationYaml, nil
}

func chartsBaseKustomizationYaml() (*asset, error) {
	bytes, err := chartsBaseKustomizationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/kustomization.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istiod-{{ .Values.global.istioNamespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
rules:
  # sidecar injection controller
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["mutatingwebhookconfigurations"]
    verbs: ["get", "list", "watch", "patch"]

  # configuration validation webhook controller
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["validatingwebhookconfigurations"]
    verbs: ["get", "list", "watch", "update"]

  # istio configuration
  - apiGroups: ["config.istio.io", "security.istio.io", "networking.istio.io", "authentication.istio.io"]
    verbs: ["get", "watch", "list"]
    resources: ["*"]
{{- if .Values.global.istiod.enableAnalysis }}
  - apiGroups: ["config.istio.io", "security.istio.io", "networking.istio.io", "authentication.istio.io"]
    verbs: ["update"]
    # TODO: should be on just */status but wildcard is not supported
    resources: ["*"]
{{- end }}

  # auto-detect installed CRD definitions
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["get", "list", "watch"]

  # discovery and routing
  - apiGroups: [""]
    resources: ["pods", "nodes", "services", "namespaces", "endpoints"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["discovery.k8s.io"]
    resources: ["endpointslices"]
    verbs: ["get", "list", "watch"]

  # ingress controller
{{- if .Values.global.istiod.enableAnalysis }}
  - apiGroups: ["extensions", "networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["extensions", "networking.k8s.io"]
    resources: ["ingresses/status"]
    verbs: ["*"]
{{- end}}
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses", "ingressclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["networking.k8s.io"]
    resources: ["ingresses/status"]
    verbs: ["*"]

  # required for CA's namespace controller
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["create", "get", "list", "watch", "update"]

  # Istiod and bootstrap.
  - apiGroups: ["certificates.k8s.io"]
    resources:
      - "certificatesigningrequests"
      - "certificatesigningrequests/approval"
      - "certificatesigningrequests/status"
    verbs: ["update", "create", "get", "delete", "watch"]
  - apiGroups: ["certificates.k8s.io"]
    resources:
      - "signers"
    resourceNames:
    - "kubernetes.io/legacy-unknown"
    verbs: ["approve"]

  # Used by Istiod to verify the JWT tokens
  - apiGroups: ["authentication.k8s.io"]
    resources: ["tokenreviews"]
    verbs: ["create"]

  # Used by Istiod to verify gateway SDS
  - apiGroups: ["authorization.k8s.io"]
    resources: ["subjectaccessreviews"]
    verbs: ["create"]

  # Use for Kubernetes Service APIs
  - apiGroups: ["networking.x-k8s.io"]
    resources: ["*"]
    verbs: ["get", "watch", "list"]

  # Needed for multicluster secret reading, possibly ingress certs in the future
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "watch", "list"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istio-reader-{{ .Values.global.istioNamespace }}
  labels:
    app: istio-reader
    release: {{ .Release.Name }}
rules:
  - apiGroups:
      - "config.istio.io"
      - "security.istio.io"
      - "networking.istio.io"
      - "authentication.istio.io"
    resources: ["*"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["endpoints", "pods", "services", "nodes", "replicationcontrollers", "namespaces"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["apps"]
    resources: ["replicasets"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["authentication.k8s.io"]
    resources: ["tokenreviews"]
    verbs: ["create"]
  - apiGroups: ["authorization.k8s.io"]
    resources: ["subjectaccessreviews"]
    verbs: ["create"]
{{- if .Values.global.centralIstiod }}
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["create", "get", "list", "watch", "update"]
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["mutatingwebhookconfigurations"]
    verbs: ["get", "list", "watch", "patch"]
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["validatingwebhookconfigurations"]
    verbs: ["get", "list", "watch", "update"]
{{- end}}
---
`)

func chartsBaseTemplatesClusterroleYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesClusterroleYaml, nil
}

func chartsBaseTemplatesClusterroleYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istio-reader-{{ .Values.global.istioNamespace }}
  labels:
    app: istio-reader
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istio-reader-{{ .Values.global.istioNamespace }}
subjects:
  - kind: ServiceAccount
    name: istio-reader-service-account
    namespace: {{ .Values.global.istioNamespace }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istiod-{{ .Values.global.istioNamespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istiod-{{ .Values.global.istioNamespace }}
subjects:
  - kind: ServiceAccount
    name: istiod-service-account
    namespace: {{ .Values.global.istioNamespace }}
---
`)

func chartsBaseTemplatesClusterrolebindingYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesClusterrolebindingYaml, nil
}

func chartsBaseTemplatesClusterrolebindingYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesCrdsYaml = []byte(`{{- if .Values.base.enableCRDTemplates }}
{{ .Files.Get "crds/crd-all.gen.yaml" }}
{{ .Files.Get "crds/crd-operator.yaml" }}
{{- end }}
`)

func chartsBaseTemplatesCrdsYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesCrdsYaml, nil
}

func chartsBaseTemplatesCrdsYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesCrdsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/crds.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesEndpointsYaml = []byte(`{{- if .Values.global.remotePilotAddress }}
  {{- if .Values.pilot.enabled }}
apiVersion: v1
kind: Endpoints
metadata:
  name: istiod-remote
  namespace: {{ .Release.Namespace }}
subsets:
- addresses:
  - ip: {{ .Values.global.remotePilotAddress }}
  ports:
  - port: 15012
    name: tcp-istiod
  {{- else if regexMatch "^([0-9]*\\.){3}[0-9]*$" .Values.global.remotePilotAddress }}
apiVersion: v1
kind: Endpoints
metadata:
  name: istiod
  namespace: {{ .Release.Namespace }}
subsets:
- addresses:
  - ip: {{ .Values.global.remotePilotAddress }}
  ports:
  - port: 15012
    name: tcp-istiod
  {{- end }}
---
{{- end }}
`)

func chartsBaseTemplatesEndpointsYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesEndpointsYaml, nil
}

func chartsBaseTemplatesEndpointsYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesEndpointsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/endpoints.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: istiod-{{ .Values.global.istioNamespace }}
  namespace: {{ .Values.global.istioNamespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
rules:
# permissions to verify the webhook is ready and rejecting
# invalid config. We use --server-dry-run so no config is persisted.
- apiGroups: ["networking.istio.io"]
  verbs: ["create"]
  resources: ["gateways"]

# For storing CA secret
- apiGroups: [""]
  resources: ["secrets"]
  # TODO lock this down to istio-ca-cert if not using the DNS cert mesh config
  verbs: ["create", "get", "watch", "list", "update", "delete"]
`)

func chartsBaseTemplatesRoleYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesRoleYaml, nil
}

func chartsBaseTemplatesRoleYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: istiod-{{ .Values.global.istioNamespace }}
  namespace: {{ .Values.global.istioNamespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: istiod-{{ .Values.global.istioNamespace }}
subjects:
  - kind: ServiceAccount
    name: istiod-service-account
    namespace: {{ .Values.global.istioNamespace }}
`)

func chartsBaseTemplatesRolebindingYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesRolebindingYaml, nil
}

func chartsBaseTemplatesRolebindingYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
{{- if .Values.global.imagePullSecrets }}
imagePullSecrets:
{{- range .Values.global.imagePullSecrets }}
  - name: {{ . }}
{{- end }}
{{- end }}
metadata:
  name: istio-reader-service-account
  namespace: {{ .Values.global.istioNamespace }}
  labels:
    app: istio-reader
    release: {{ .Release.Name }}
---
apiVersion: v1
kind: ServiceAccount
  {{- if .Values.global.imagePullSecrets }}
imagePullSecrets:
  {{- range .Values.global.imagePullSecrets }}
  - name: {{ . }}
  {{- end }}
  {{- end }}
metadata:
  name: istiod-service-account
  namespace: {{ .Values.global.istioNamespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
---
`)

func chartsBaseTemplatesServiceaccountYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesServiceaccountYaml, nil
}

func chartsBaseTemplatesServiceaccountYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesServicesYaml = []byte(`{{- if .Values.global.remotePilotAddress }}
  {{- if .Values.pilot.enabled }}
# when istiod is enabled in remote cluster, we can't use istiod service name
apiVersion: v1
kind: Service
metadata:
  name: istiod-remote
  namespace: {{ .Release.Namespace }}
spec:
  ports:
  - port: 15012
    name: tcp-istiod
  clusterIP: None
  {{- else }}
# when istiod isn't enabled in remote cluster, we can use istiod service name
apiVersion: v1
kind: Service
metadata:
  name: istiod
  namespace: {{ .Release.Namespace }}
spec:
  ports:
  - port: 15012
    name: tcp-istiod
  # if the remotePilotAddress is IP addr, we use clusterIP: None.
  # else, we use externalName
  {{- if regexMatch "^([0-9]*\\.){3}[0-9]*$" .Values.global.remotePilotAddress }}
  clusterIP: None
  {{- else }}
  type: ExternalName
  externalName: {{ .Values.global.remotePilotAddress }}
  {{- end }}
  {{- end }}
---
{{- end }}
`)

func chartsBaseTemplatesServicesYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesServicesYaml, nil
}

func chartsBaseTemplatesServicesYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesServicesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/services.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseTemplatesValidatingwebhookconfigurationYaml = []byte(`{{- if .Values.global.configValidation }}
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: istiod-{{ .Values.global.istioNamespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
    istio: istiod
webhooks:
  - name: validation.istio.io
    clientConfig:
      {{- if .Values.base.validationURL }}
      url: {{ .Values.base.validationURL }}
      {{- else }}
      service:
        name: istiod
        namespace: {{ .Values.global.istioNamespace }}
        path: "/validate"
      {{- end }}
      caBundle: "" # patched at runtime when the webhook is ready.
    rules:
      - operations:
        - CREATE
        - UPDATE
        apiGroups:
        - config.istio.io
        - security.istio.io
        - authentication.istio.io
        - networking.istio.io
        apiVersions:
        - "*"
        resources:
        - "*"
    # Fail open until the validation webhook is ready. The webhook controller
    # will update this to `+"`"+`Fail`+"`"+` and patch in the `+"`"+`caBundle`+"`"+` when the webhook
    # endpoint is ready.
    failurePolicy: Ignore
    sideEffects: None
    admissionReviewVersions: ["v1beta1", "v1"]
---
{{- end }}`)

func chartsBaseTemplatesValidatingwebhookconfigurationYamlBytes() ([]byte, error) {
	return _chartsBaseTemplatesValidatingwebhookconfigurationYaml, nil
}

func chartsBaseTemplatesValidatingwebhookconfigurationYaml() (*asset, error) {
	bytes, err := chartsBaseTemplatesValidatingwebhookconfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/templates/validatingwebhookconfiguration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsBaseValuesYaml = []byte(`global:

  # ImagePullSecrets for control plane ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []

  # Used to locate istiod.
  istioNamespace: istio-system

  istiod:
    enableAnalysis: false

  configValidation: true

base:
  # Used for helm2 to add the CRDs to templates.
  enableCRDTemplates: false

  # Validation webhook configuration url
  # For example: https://$remotePilotAddress:15017/validate
  validationURL: ""`)

func chartsBaseValuesYamlBytes() ([]byte, error) {
	return _chartsBaseValuesYaml, nil
}

func chartsBaseValuesYaml() (*asset, error) {
	bytes, err := chartsBaseValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/base/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressChartYaml = []byte(`apiVersion: v1
name: istio-egress
version: 1.1.0
tillerVersion: ">=2.7.2"
description: Helm chart for deploying Istio gateways
keywords:
  - istio
  - egressgateway
  - gateways
sources:
  - http://github.com/istio/istio
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsGatewaysIstioEgressChartYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressChartYaml, nil
}

func chartsGatewaysIstioEgressChartYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressNotesTxt = []byte(`
Changes:
- separate namespace allows:
-- easier reconfig of just the gateway
-- TLS secrets and domain name management is isolated, for better security
-- simplified configuration
-- multiple versions of the ingress can be used, to minize upgrade risks

- the new chart uses the default namespace service account, and doesn't require
additional RBAC permissions.

- simplified label structure. Label change is not supported on upgrade.

- for 'internal load balancer' you should deploy a separate gateway, in a different
namespace.

All ingress gateway have a "app:ingressgateway" label, used to identify it as an
ingress, and an "istio: ingressgateway$SUFFIX" label of Gateway selection.

The Gateways use "istio: ingressgateway$SUFFIX" selectors.


# Multiple gateway versions



# Using different pilot versions



# Migration from istio-system

Istio 1.0 includes the gateways in istio-system. Since the external IP is associated
with the Service and bound to the namespace, it is recommended to:

1. Install the new gateway in a new namespace.
2. Copy any TLS certificate to the new namespace, and configure the domains.
3. Checking the new gateway work - for example by overriding the IP in /etc/hosts
4. Modify the DNS server to add the A record of the new namespace
5. Check traffic
6. Delete the A record corresponding to the gateway in istio-system
7. Upgrade istio-system, disabling the ingressgateway
8. Delete the domain TLS certs from istio-system.

If using certmanager, all Certificate and associated configs must be moved as well.
`)

func chartsGatewaysIstioEgressNotesTxtBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressNotesTxt, nil
}

func chartsGatewaysIstioEgressNotesTxt() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressNotesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/NOTES.txt", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplates_affinityTpl = []byte(`{{/* affinity - https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ */}}

{{- define "nodeaffinity" }}
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
    {{- include "nodeAffinityRequiredDuringScheduling" . }}
    preferredDuringSchedulingIgnoredDuringExecution:
    {{- include "nodeAffinityPreferredDuringScheduling" . }}
{{- end }}

{{- define "nodeAffinityRequiredDuringScheduling" }}
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/arch
          operator: In
          values:
        {{- range $key, $val := .global.arch }}
          {{- if gt ($val | int) 0 }}
          - {{ $key | quote }}
          {{- end }}
        {{- end }}
        {{- $nodeSelector := default .global.defaultNodeSelector .nodeSelector -}}
        {{- range $key, $val := $nodeSelector }}
        - key: {{ $key }}
          operator: In
          values:
          - {{ $val | quote }}
        {{- end }}
{{- end }}

{{- define "nodeAffinityPreferredDuringScheduling" }}
  {{- range $key, $val := .global.arch }}
    {{- if gt ($val | int) 0 }}
    - weight: {{ $val | int }}
      preference:
        matchExpressions:
        - key: kubernetes.io/arch
          operator: In
          values:
          - {{ $key | quote }}
    {{- end }}
  {{- end }}
{{- end }}

{{- define "podAntiAffinity" }}
{{- if or .podAntiAffinityLabelSelector .podAntiAffinityTermLabelSelector}}
  podAntiAffinity:
    {{- if .podAntiAffinityLabelSelector }}
    requiredDuringSchedulingIgnoredDuringExecution:
    {{- include "podAntiAffinityRequiredDuringScheduling" . }}
    {{- end }}
    {{- if .podAntiAffinityTermLabelSelector }}
    preferredDuringSchedulingIgnoredDuringExecution:
    {{- include "podAntiAffinityPreferredDuringScheduling" . }}
    {{- end }}
{{- end }}
{{- end }}

{{- define "podAntiAffinityRequiredDuringScheduling" }}
    {{- range $index, $item := .podAntiAffinityLabelSelector }}
    - labelSelector:
        matchExpressions:
        - key: {{ $item.key }}
          operator: {{ $item.operator }}
          {{- if $item.values }}
          values:
          {{- $vals := split "," $item.values }}
          {{- range $i, $v := $vals }}
          - {{ $v | quote }}
          {{- end }}
          {{- end }}
      topologyKey: {{ $item.topologyKey }}
    {{- end }}
{{- end }}

{{- define "podAntiAffinityPreferredDuringScheduling" }}
    {{- range $index, $item := .podAntiAffinityTermLabelSelector }}
    - podAffinityTerm:
        labelSelector:
          matchExpressions:
          - key: {{ $item.key }}
            operator: {{ $item.operator }}
            {{- if $item.values }}
            values:
            {{- $vals := split "," $item.values }}
            {{- range $i, $v := $vals }}
            - {{ $v | quote }}
            {{- end }}
            {{- end }}
        topologyKey: {{ $item.topologyKey }}
      weight: 100
    {{- end }}
{{- end }}
`)

func chartsGatewaysIstioEgressTemplates_affinityTplBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplates_affinityTpl, nil
}

func chartsGatewaysIstioEgressTemplates_affinityTpl() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplates_affinityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/_affinity.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplates_helpersTpl = []byte(`{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "gateway.name" -}}
{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
{{- default .Chart.Name $gateway.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "gateway.fullname" -}}
{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
{{- if $gateway.fullnameOverride -}}
{{- $gateway.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name $gateway.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "gateway.chart" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
`)

func chartsGatewaysIstioEgressTemplates_helpersTplBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplates_helpersTpl, nil
}

func chartsGatewaysIstioEgressTemplates_helpersTpl() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplates_helpersTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/_helpers.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesAutoscaleYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
{{- if and $gateway.autoscaleEnabled $gateway.autoscaleMin $gateway.autoscaleMax }}
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  maxReplicas: {{ $gateway.autoscaleMax }}
  minReplicas: {{ $gateway.autoscaleMin }}
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{ $gateway.name | default "istio-egressgateway" }}
  metrics:
    - type: Resource
      resource:
        name: cpu
        targetAverageUtilization: {{ $gateway.cpu.targetAverageUtilization }}
---
{{- end }}
`)

func chartsGatewaysIstioEgressTemplatesAutoscaleYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesAutoscaleYaml, nil
}

func chartsGatewaysIstioEgressTemplatesAutoscaleYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesAutoscaleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/autoscale.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesDeploymentYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
{{- if not $gateway.autoscaleEnabled }}
{{- if $gateway.replicaCount }}
  replicas: {{ $gateway.replicaCount }}
{{- end }}
{{- end }}
  selector:
    matchLabels:
{{ $gateway.labels | toYaml | indent 6 }}
  strategy:
    rollingUpdate:
      maxSurge: {{ $gateway.rollingMaxSurge }}
      maxUnavailable: {{ $gateway.rollingMaxUnavailable }}
  template:
    metadata:
      labels:
{{ $gateway.labels | toYaml | indent 8 }}
{{- if eq .Release.Namespace "istio-system"}}
        heritage: Tiller
        release: istio
        chart: gateways
{{- end }}
        service.istio.io/canonical-name: {{ $gateway.name | default "istio-egressgateway" }}
{{- if not (eq .Values.revision "") }}
        service.istio.io/canonical-revision: {{ .Values.revision }}
{{- else}}
        service.istio.io/canonical-revision: latest
{{- end }}
      annotations:
        {{- if .Values.meshConfig.enablePrometheusMerge }}
        prometheus.io/port: "15020"
        prometheus.io/scrape: "true"
        prometheus.io/path: "/stats/prometheus"
        {{- end }}
        sidecar.istio.io/inject: "false"
{{- if $gateway.podAnnotations }}
{{ toYaml $gateway.podAnnotations | indent 8 }}
{{ end }}
    spec:
{{- if not $gateway.runAsRoot }}
      securityContext:
        runAsUser: 1337
        runAsGroup: 1337
        runAsNonRoot: true
        fsGroup: 1337
{{- end }}
      serviceAccountName: {{ $gateway.name | default "istio-egressgateway" }}-service-account
{{- if .Values.global.priorityClassName }}
      priorityClassName: "{{ .Values.global.priorityClassName }}"
{{- end }}
{{- if .Values.global.proxy.enableCoreDump }}
      initContainers:
        - name: enable-core-dump
{{- if contains "/" .Values.global.proxy.image }}
          image: "{{ .Values.global.proxy.image }}"
{{- else }}
          image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image | default "proxyv2" }}:{{ .Values.global.tag }}"
{{- end }}
{{- if .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
          command:
            - /bin/sh
          args:
            - -c
            - sysctl -w kernel.core_pattern=/var/lib/istio/data/core.proxy && ulimit -c unlimited
          securityContext:
            runAsUser: 0
            runAsGroup: 0
            runAsNonRoot: false
            privileged: true
{{- end }}
      containers:
        - name: istio-proxy
{{- if contains "/" .Values.global.proxy.image }}
          image: "{{ .Values.global.proxy.image }}"
{{- else }}
          image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image }}:{{ .Values.global.tag }}"
{{- end }}
{{- if .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
          ports:
            {{- range $key, $val := $gateway.ports }}
            - containerPort: {{ $val.targetPort | default $val.port }}
            {{- end }}
            - containerPort: 15090
              protocol: TCP
              name: http-envoy-prom
          args:
          - proxy
          - router
          - --domain
          - $(POD_NAMESPACE).svc.{{ .Values.global.proxy.clusterDomain }}
        {{- if .Values.global.proxy.logLevel }}
          - --proxyLogLevel={{ .Values.global.proxy.logLevel }}
        {{- end}}
        {{- if .Values.global.proxy.componentLogLevel }}
          - --proxyComponentLogLevel={{ .Values.global.proxy.componentLogLevel }}
        {{- end}}
        {{- if .Values.global.logging.level }}
          - --log_output_level={{ .Values.global.logging.level }}
        {{- end}}
        {{- if .Values.global.logAsJson }}
          - --log_as_json
        {{- end }}
          - --serviceCluster
          - {{ $gateway.name | default "istio-egressgateway" }}
        {{- if .Values.global.sts.servicePort }}
          - --stsPort={{ .Values.global.sts.servicePort }}
        {{- end }}
        {{- if .Values.global.trustDomain }}
          - --trust-domain={{ .Values.global.trustDomain }}
        {{- end }}
        {{- if not $gateway.runAsRoot }}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - ALL
            privileged: false
            readOnlyRootFilesystem: true
        {{- end }}
          readinessProbe:
            failureThreshold: 30
            httpGet:
              path: /healthz/ready
              port: 15021
              scheme: HTTP
            initialDelaySeconds: 1
            periodSeconds: 2
            successThreshold: 1
            timeoutSeconds: 1
          resources:
{{- if $gateway.resources }}
{{ toYaml $gateway.resources | indent 12 }}
{{- else }}
{{ toYaml .Values.global.defaultResources | indent 12 }}
{{- end }}
          env:
          - name: JWT_POLICY
            value: {{ .Values.global.jwtPolicy }}
          - name: PILOT_CERT_PROVIDER
            value: {{ .Values.global.pilotCertProvider }}
          - name: CA_ADDR
          {{- if .Values.global.caAddress }}
            value: {{ .Values.global.caAddress }}
          {{- else }}
            value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Values.global.istioNamespace }}.svc:15012
          {{- end }}
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: INSTANCE_IP
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: status.podIP
          - name: HOST_IP
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: status.hostIP
          - name: SERVICE_ACCOUNT
            valueFrom:
              fieldRef:
                fieldPath: spec.serviceAccountName
          - name: CANONICAL_SERVICE
            valueFrom:
              fieldRef:
                fieldPath: metadata.labels['service.istio.io/canonical-name']
          - name: CANONICAL_REVISION
            valueFrom:
              fieldRef:
                fieldPath: metadata.labels['service.istio.io/canonical-revision']
          - name: ISTIO_META_WORKLOAD_NAME
            value: {{ $gateway.name | default "istio-egressgateway" }}
          - name: ISTIO_META_OWNER
            value: kubernetes://apis/apps/v1/namespaces/{{ .Release.Namespace }}/deployments/{{ $gateway.name | default "istio-egressgateway" }}
          {{- if $.Values.global.meshID }}
          - name: ISTIO_META_MESH_ID
            value: "{{ $.Values.global.meshID }}"
          {{- else if $.Values.global.trustDomain }}
          - name: ISTIO_META_MESH_ID
            value: "{{ $.Values.global.trustDomain }}"
          {{- end }}
          {{- if $gateway.env }}
          {{- range $key, $val := $gateway.env }}
          - name: {{ $key }}
            value: {{ $val }}
          {{- end }}
          {{ end }}
          {{- range $key, $value := .Values.meshConfig.defaultConfig.proxyMetadata }}
          - name: {{ $key }}
            value: "{{ $value }}"
          {{- end }}
{{- if $gateway.podAnnotations }}
          - name: "ISTIO_METAJSON_ANNOTATIONS"
            value: |
{{ toJson $gateway.podAnnotations | indent 16}}
{{ end }}
          - name: ISTIO_META_CLUSTER_ID
            value: "{{ $.Values.global.multiCluster.clusterName | default `+"`"+`Kubernetes`+"`"+` }}"
          volumeMounts:
          - name: istio-envoy
            mountPath: /etc/istio/proxy
          - name: config-volume
            mountPath: /etc/istio/config
          {{- if eq .Values.global.pilotCertProvider "istiod" }}
          - mountPath: /var/run/secrets/istio
            name: istiod-ca-cert
          {{- end }}
          {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
          - name: istio-token
            mountPath: /var/run/secrets/tokens
            readOnly: true
          {{- end }}
          - name: gatewaysdsudspath
            mountPath: /var/run/ingress_gateway
          {{- if .Values.global.mountMtlsCerts }}
          # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
          - name: istio-certs
            mountPath: /etc/certs
            readOnly: true
          {{- end }}
          - mountPath: /var/lib/istio/data
            name: istio-data
          - name: podinfo
            mountPath: /etc/istio/pod
          {{- range $gateway.secretVolumes }}
          - name: {{ .name }}
            mountPath: {{ .mountPath | quote }}
            readOnly: true
          {{- end }}
          {{- range $gateway.configVolumes }}
          {{- if .mountPath }}
          - name: {{ .name }}
            mountPath: {{ .mountPath | quote }}
            readOnly: true
          {{- end }}
          {{- end }}
{{- if $gateway.additionalContainers }}
{{ toYaml $gateway.additionalContainers | indent 8 }}
{{- end }}
      volumes:
      {{- if eq .Values.global.pilotCertProvider "istiod" }}
      - name: istiod-ca-cert
        configMap:
          name: istio-ca-root-cert
      {{- end }}
      - name: podinfo
        downwardAPI:
          items:
            - path: "labels"
              fieldRef:
                fieldPath: metadata.labels
            - path: "annotations"
              fieldRef:
                fieldPath: metadata.annotations
      - name: istio-envoy
        emptyDir: {}
      - name: gatewaysdsudspath
        emptyDir: {}
      - name: istio-data
        emptyDir: {}
{{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
      - name: istio-token
        projected:
          sources:
          - serviceAccountToken:
              path: istio-token
              expirationSeconds: 43200
              audience: {{ .Values.global.sds.token.aud }}
{{- end }}
      {{- if .Values.global.mountMtlsCerts }}
      # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
      - name: istio-certs
        secret:
          secretName: istio.istio-egressgateway-service-account
          optional: true
      {{- end }}
      - name: config-volume
        configMap:
          name: istio{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
          optional: true
      {{- range $gateway.secretVolumes }}
      - name: {{ .name }}
        secret:
          secretName: {{ .secretName | quote }}
          optional: true
      {{- end }}
      {{- range $gateway.configVolumes }}
      - name: {{ .name }}
        configMap:
          name: {{ .configMapName | quote }}
          optional: true
      {{- end }}
      affinity:
      {{- include "nodeaffinity" (dict "global" .Values.global "nodeSelector" $gateway.nodeSelector) | indent 6 }}
      {{- include "podAntiAffinity" $gateway | indent 6 }}
{{- if $gateway.tolerations }}
      tolerations:
{{ toYaml $gateway.tolerations | indent 6 }}
{{- else if .Values.global.defaultTolerations }}
      tolerations:
{{ toYaml .Values.global.defaultTolerations | indent 6 }}
{{- end }}
`)

func chartsGatewaysIstioEgressTemplatesDeploymentYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesDeploymentYaml, nil
}

func chartsGatewaysIstioEgressTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYaml = []byte(`{{- if .Values.global.defaultPodDisruptionBudget.enabled }}
{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | trim | indent 4 }}
    release: {{ .Release.Name }}
spec:
  minAvailable: 1
  selector:
    matchLabels:
{{ $gateway.labels | toYaml | trim | indent 6 }}
{{- end }}
`)

func chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYaml, nil
}

func chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/poddisruptionbudget.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesPreconfiguredYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
{{- if $gateway.zvpn.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: istio-multicluster-egressgateway
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  selector:
    istio: egressgateway
  servers:
  - hosts:
    - "*.{{ $gateway.zvpn.suffix }}"
    port:
      name: tls
      number: 15443
      protocol: TLS
    tls: {}
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: istio-multicluster-egressgateway
  namespace: {{ .Release.Namespace }} 
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  gateways:
  - istio-multicluster-egressgateway
  hosts:
  - "*.{{ $gateway.zvpn.suffix }}"
  tls:
  - match:
    - port: 15443
      sniHosts:
      - "*.{{ $gateway.zvpn.suffix }}"
    route:
    - destination:
        host: non.existent.cluster
        port:
          number: 15443
      weight: 100
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: istio-multicluster-egressgateway
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  workloadLabels:
    istio: egressgateway
  filters:
  - listenerMatch:
      portNumber: 15443
      listenerType: GATEWAY
    filterName: envoy.filters.network.sni_cluster
    filterType: NETWORK
    filterConfig: {}
---
## To ensure all traffic to *.global is using mTLS
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: istio-multicluster-egressgateway
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  host: "*.{{ $gateway.zvpn.suffix }}"
  trafficPolicy:
    tls:
      mode: ISTIO_MUTUAL
---
{{- end }}
`)

func chartsGatewaysIstioEgressTemplatesPreconfiguredYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesPreconfiguredYaml, nil
}

func chartsGatewaysIstioEgressTemplatesPreconfiguredYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesPreconfiguredYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/preconfigured.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesRoleYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}-sds
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "watch", "list"]
---
`)

func chartsGatewaysIstioEgressTemplatesRoleYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesRoleYaml, nil
}

func chartsGatewaysIstioEgressTemplatesRoleYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesRolebindingsYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}-sds
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ $gateway.name | default "istio-egressgateway" }}-sds
subjects:
  - kind: ServiceAccount
    name: {{ $gateway.name | default "istio-egressgateway" }}-service-account
---
`)

func chartsGatewaysIstioEgressTemplatesRolebindingsYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesRolebindingsYaml, nil
}

func chartsGatewaysIstioEgressTemplatesRolebindingsYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesRolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesServiceYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
apiVersion: v1
kind: Service
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}
  namespace: {{ .Release.Namespace }}
  annotations:
    {{- range $key, $val := $gateway.serviceAnnotations }}
    {{ $key }}: {{ $val | quote }}
    {{- end }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  type: ClusterIP
  selector:
{{ $gateway.labels | toYaml | indent 4 }}
  ports:
    {{- range $key, $val := $gateway.ports }}
    -
      {{- range $pkey, $pval := $val }}
      {{ $pkey }}: {{ $pval }}
      {{- end }}
    {{- end }}
---
`)

func chartsGatewaysIstioEgressTemplatesServiceYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesServiceYaml, nil
}

func chartsGatewaysIstioEgressTemplatesServiceYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressTemplatesServiceaccountYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-egressgateway" }}
apiVersion: v1
kind: ServiceAccount
{{- if .Values.global.imagePullSecrets }}
imagePullSecrets:
{{- range .Values.global.imagePullSecrets }}
  - name: {{ . }}
{{- end }}
{{- end }}
metadata:
  name: {{ $gateway.name | default "istio-egressgateway" }}-service-account
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
`)

func chartsGatewaysIstioEgressTemplatesServiceaccountYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressTemplatesServiceaccountYaml, nil
}

func chartsGatewaysIstioEgressTemplatesServiceaccountYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressTemplatesServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/templates/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioEgressValuesYaml = []byte(`# Standalone istio egress gateway.
# Should be installed in a separate namespace, to minimize access to config
gateways:
  istio-egressgateway:
    name: istio-egressgateway
    ports:
    - port: 80
      targetPort: 8080
      name: http2
    - port: 443
      name: https
      targetPort: 8443
      # This is the port where sni routing happens
    - port: 15443
      targetPort: 15443
      name: tls

    # Enable cross-cluster access using SNI matching.
    # Make sure you set suffix if deploying multiple egress gateways.
    zvpn:
      # Must be different for each egress namespace.
      # Used to control the domain name suffix for zero vpn routing.
      # Domain names ending with this suffix will be routed to this egress gateway
      # automatically.
      # This can be a real domain name ( istio.example.com )
      suffix: global
      enabled: false

    labels:
      app: istio-egressgateway
      istio: egressgateway

    # Scalability tuning
    # replicaCount: 1
    rollingMaxSurge: 100%
    rollingMaxUnavailable: 25%
    autoscaleEnabled: true
    autoscaleMin: 1
    autoscaleMax: 5
    resources:
      requests:
        cpu: 100m
        memory: 128Mi
      limits:
        cpu: 2000m
        memory: 1024Mi
    cpu:
      targetAverageUtilization: 80

    serviceAnnotations: {}
    podAnnotations: {}
    type: ClusterIP # change to NodePort or LoadBalancer if need be

    secretVolumes:
    - name: egressgateway-certs
      secretName: istio-egressgateway-certs
      mountPath: /etc/istio/egressgateway-certs
    - name: egressgateway-ca-certs
      secretName: istio-egressgateway-ca-certs
      mountPath: /etc/istio/egressgateway-ca-certs

    configVolumes: []
    additionalContainers: []
    ### Advanced options ############
    # TODO: convert to real options, env should not be exposed
    env:
      # Set this to "external" if and only if you want the egress gateway to
      # act as a transparent SNI gateway that routes mTLS/TLS traffic to
      # external services defined using service entries, where the service
      # entry has resolution set to DNS, has one or more endpoints with
      # network field set to "external". By default its set to "" so that
      # the egress gateway sees the same set of endpoints as the sidecars
      # preserving backward compatibility
      # ISTIO_META_REQUESTED_NETWORK_VIEW: ""
      # A gateway with this mode ensures that pilot generates an additional
      # set of clusters for internal services but without Istio mTLS, to
      # enable cross cluster routing.
      ISTIO_META_ROUTER_MODE: "sni-dnat"

    nodeSelector: {}
    tolerations: []

    # Specify the pod anti-affinity that allows you to constrain which nodes
    # your pod is eligible to be scheduled based on labels on pods that are
    # already running on the node rather than based on labels on nodes.
    # There are currently two types of anti-affinity:
    #    "requiredDuringSchedulingIgnoredDuringExecution"
    #    "preferredDuringSchedulingIgnoredDuringExecution"
    # which denote "hard" vs. "soft" requirements, you can define your values
    # in "podAntiAffinityLabelSelector" and "podAntiAffinityTermLabelSelector"
    # correspondingly.
    # For example:
    # podAntiAffinityLabelSelector:
    # - key: security
    #   operator: In
    #   values: S1,S2
    #   topologyKey: "kubernetes.io/hostname"
    # This pod anti-affinity rule says that the pod requires not to be scheduled
    # onto a node if that node is already running a pod with label having key
    # "security" and value "S1".
    podAntiAffinityLabelSelector: []
    podAntiAffinityTermLabelSelector: []

    # whether to run the gateway in a privileged container
    runAsRoot: false

# Revision is set as 'version' label and part of the resource names when installing multiple control planes.
revision: ""

global:
  # set the default set of namespaces to which services, service entries, virtual services, destination
  # rules should be exported to. Currently only one value can be provided in this list. This value
  # should be one of the following two options:
  # * implies these objects are visible to all namespaces, enabling any sidecar to talk to any other sidecar.
  # . implies these objects are visible to only to sidecars in the same namespace, or if imported as a Sidecar.egress.host
  defaultConfigVisibilitySettings: []

  # enable pod disruption budget for the control plane, which is used to
  # ensure Istio control plane components are gradually upgraded or recovered.
  defaultPodDisruptionBudget:
    enabled: true

  # A minimal set of requested resources to applied to all deployments so that
  # Horizontal Pod Autoscaler will be able to function (if set).
  # Each component can overwrite these default values by adding its own resources
  # block in the relevant section below and setting the desired resources values.
  defaultResources:
    requests:
      cpu: 10m
    #   memory: 128Mi
    # limits:
    #   cpu: 100m
    #   memory: 128Mi

  # Default node tolerations to be applied to all deployments so that all pods can be
  # scheduled to a particular nodes with matching taints. Each component can overwrite
  # these default values by adding its tolerations block in the relevant section below
  # and setting the desired values.
  # Configure this field in case that all pods of Istio control plane are expected to
  # be scheduled to particular nodes with specified taints.
  defaultTolerations: []

  # Default hub for Istio images.
  # Releases are published to docker hub under 'istio' project.
  # Dev builds from prow are on gcr.io
  hub: gcr.io/istio-testing

  # Default tag for Istio images.
  tag: latest

  # Specify image pull policy if default behavior isn't desired.
  # Default behavior: latest images will be Always else IfNotPresent.
  imagePullPolicy: ""

  # ImagePullSecrets for all ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # For components that don't use ServiceAccounts (i.e. grafana, servicegraph, tracing)
  # ImagePullSecrets will be added to the corresponding Deployment(StatefulSet) objects.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []
  # - private-registry-key

  # To output all istio components logs in json format by adding --log_as_json argument to each container argument
  logAsJson: false

  # Comma-separated minimum per-scope logging level of messages to output, in the form of <scope>:<level>,<scope>:<level>
  # The control plane has different scopes depending on component, but can configure default log level across all components
  # If empty, default scope and level will be used as configured in code
  logging:
    level: "default:info"

  # If set to true, the pilot and citadel mtls will be exposed on the
  # ingress gateway
  meshExpansion:
    enabled: false
    # If set to true, the pilot and citadel mtls and the plain text pilot ports
    # will be exposed on an internal gateway
    useILB: false

  # Kubernetes >=v1.11.0 will create two PriorityClass, including system-cluster-critical and
  # system-node-critical, it is better to configure this in order to make sure your Istio pods
  # will not be killed because of low priority class.
  # Refer to https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
  # for more detail.
  priorityClassName: ""

  proxy:
    image: proxyv2

    # CAUTION: It is important to ensure that all Istio helm charts specify the same clusterDomain value
    # cluster domain. Default value is "cluster.local".
    clusterDomain: "cluster.local"

    # Per Component log level for proxy, applies to gateways and sidecars. If a component level is
    # not set, then the global "logLevel" will be used.
    componentLogLevel: "misc:error"

    # If set, newly injected sidecars will have core dumps enabled.
    enableCoreDump: false

    # Log level for proxy, applies to gateways and sidecars.
    # Expected values are: trace|debug|info|warning|error|critical|off
    logLevel: warning

  ##############################################################################################
  # The following values are found in other charts. To effectively modify these values, make   #
  # make sure they are consistent across your Istio helm charts                                #
  ##############################################################################################

  # The customized CA address to retrieve certificates for the pods in the cluster.
  # CSR clients such as the Istio Agent and ingress gateways can use this to specify the CA endpoint.
  caAddress: ""

  # Used to locate istiod.
  istioNamespace: istio-system

  # Configure the policy for validating JWT.
  # Currently, two options are supported: "third-party-jwt" and "first-party-jwt".
  jwtPolicy: "third-party-jwt"

  # Mesh ID means Mesh Identifier. It should be unique within the scope where
  # meshes will interact with each other, but it is not required to be
  # globally/universally unique. For example, if any of the following are true,
  # then two meshes must have different Mesh IDs:
  # - Meshes will have their telemetry aggregated in one place
  # - Meshes will be federated together
  # - Policy will be written referencing one mesh from the other
  #
  # If an administrator expects that any of these conditions may become true in
  # the future, they should ensure their meshes have different Mesh IDs
  # assigned.
  #
  # Within a multicluster mesh, each cluster must be (manually or auto)
  # configured to have the same Mesh ID value. If an existing cluster 'joins' a
  # multicluster mesh, it will need to be migrated to the new mesh ID. Details
  # of migration TBD, and it may be a disruptive operation to change the Mesh
  # ID post-install.
  #
  # If the mesh admin does not specify a value, Istio will use the value of the
  # mesh's Trust Domain. The best practice is to select a proper Trust Domain
  # value.
  meshID: ""

  # Use the user-specified, secret volume mounted key and certs for Pilot and workloads.
  mountMtlsCerts: false

  multiCluster:
    # Set to true to connect two kubernetes clusters via their respective
    # ingressgateway services when pods in each cluster cannot directly
    # talk to one another. All clusters should be using Istio mTLS and must
    # have a shared root CA for this model to work.
    enabled: false
    # Should be set to the name of the cluster this installation will run in. This is required for sidecar injection
    # to properly label proxies
    clusterName: ""

  # Network defines the network this cluster belong to. This name
  # corresponds to the networks in the map of mesh networks.
  network: ""

  # Configure the certificate provider for control plane communication.
  # Currently, two providers are supported: "kubernetes" and "istiod".
  # As some platforms may not have kubernetes signing APIs,
  # Istiod is the default
  pilotCertProvider: istiod

  sds:
    # The JWT token for SDS and the aud field of such JWT. See RFC 7519, section 4.1.3.
    # When a CSR is sent from Citadel Agent to the CA (e.g. Citadel), this aud is to make sure the
    # JWT is intended for the CA.
    token:
      aud: istio-ca

  sts:
    # The service port used by Security Token Service (STS) server to handle token exchange requests.
    # Setting this port to a non-zero value enables STS server.
    servicePort: 0

  # The trust domain corresponds to the trust root of a system
  # Refer to https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#21-trust-domain
  # Indicate the domain used in SPIFFE identity URL
  # The default depends on the environment.
  #   kubernetes: cluster.local
  #   else:  default dns domain
  trustDomain: "cluster.local"

meshConfig:
  enablePrometheusMerge: true
  defaultConfig:
    proxyMetadata: {}
    tracing:
    #      tlsSettings:
    #        mode: DISABLE # DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL
    #        clientCertificate: # example: /etc/istio/tracer/cert-chain.pem
    #        privateKey:        # example: /etc/istio/tracer/key.pem
    #        caCertificates:    # example: /etc/istio/tracer/root-cert.pem
    #        sni:               # example: tracer.somedomain
    #        subjectAltNames: []
    # - tracer.somedomain
`)

func chartsGatewaysIstioEgressValuesYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioEgressValuesYaml, nil
}

func chartsGatewaysIstioEgressValuesYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioEgressValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-egress/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressChartYaml = []byte(`apiVersion: v1
name: istio-ingress
version: 1.1.0
tillerVersion: ">=2.7.2"
description: Helm chart for deploying Istio gateways
keywords:
  - istio
  - ingressgateway
  - gateways
sources:
  - http://github.com/istio/istio
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsGatewaysIstioIngressChartYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressChartYaml, nil
}

func chartsGatewaysIstioIngressChartYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressNotesTxt = []byte(`
Changes:
- separate namespace allows:
-- easier reconfig of just the gateway
-- TLS secrets and domain name management is isolated, for better security
-- simplified configuration
-- multiple versions of the ingress can be used, to minimize upgrade risks

- the new chart uses the default namespace service account, and doesn't require
additional RBAC permissions.

- simplified label and chart structure.
- ability to run a pilot dedicated for the gateway, isolated from the main pilot. This is more robust, safer on upgrades
and allows a bit more flexibility.
- the dedicated pilot-per-ingress is required if the gateway needs to support k8s-style ingress.

# Port and basic host configuration

In order to configure the Service object, the install/upgrade needs to provide a list of all ports.
In the past, this was done when installing/upgrading full istio, and involved some duplication - ports configured
both in upgrade, Gateway and VirtualService.

The new Ingress chart uses a 'values.yaml' (see user-example-ingress), which auto-generates Service ports,
Gateways and basic VirtualService. It is still possible to only configure the ports in Service, and do manual
config for the rest.

All internal services ( telemetry, pilot debug ports, mesh expansion ) can now be configured via the new mechanism.

# Migration from istio-system

Istio 1.0 includes the gateways in istio-system. Since the external IP is associated
with the Service and bound to the namespace, it is recommended to:

1. Install the new gateway in a new namespace.
2. Copy any TLS certificate to the new namespace, and configure the domains.
3. Checking the new gateway work - for example by overriding the IP in /etc/hosts
4. Modify the DNS server to add the A record of the new namespace
5. Check traffic
6. Delete the A record corresponding to the gateway in istio-system
7. Upgrade istio-system, disabling the ingressgateway
8. Delete the domain TLS certs from istio-system.

If using certmanager, all Certificate and associated configs must be moved as well.
`)

func chartsGatewaysIstioIngressNotesTxtBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressNotesTxt, nil
}

func chartsGatewaysIstioIngressNotesTxt() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressNotesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/NOTES.txt", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplates_affinityTpl = []byte(`{{/* affinity - https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ */}}

{{- define "nodeaffinity" }}
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
    {{- include "nodeAffinityRequiredDuringScheduling" . }}
    preferredDuringSchedulingIgnoredDuringExecution:
    {{- include "nodeAffinityPreferredDuringScheduling" . }}
{{- end }}

{{- define "nodeAffinityRequiredDuringScheduling" }}
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/arch
          operator: In
          values:
        {{- range $key, $val := .global.arch }}
          {{- if gt ($val | int) 0 }}
          - {{ $key | quote }}
          {{- end }}
        {{- end }}
        {{- $nodeSelector := default .global.defaultNodeSelector .nodeSelector -}}
        {{- range $key, $val := $nodeSelector }}
        - key: {{ $key }}
          operator: In
          values:
          - {{ $val | quote }}
        {{- end }}
{{- end }}

{{- define "nodeAffinityPreferredDuringScheduling" }}
  {{- range $key, $val := .global.arch }}
    {{- if gt ($val | int) 0 }}
    - weight: {{ $val | int }}
      preference:
        matchExpressions:
        - key: kubernetes.io/arch
          operator: In
          values:
          - {{ $key | quote }}
    {{- end }}
  {{- end }}
{{- end }}

{{- define "podAntiAffinity" }}
{{- if or .podAntiAffinityLabelSelector .podAntiAffinityTermLabelSelector}}
  podAntiAffinity:
    {{- if .podAntiAffinityLabelSelector }}
    requiredDuringSchedulingIgnoredDuringExecution:
    {{- include "podAntiAffinityRequiredDuringScheduling" . }}
    {{- end }}
    {{- if .podAntiAffinityTermLabelSelector }}
    preferredDuringSchedulingIgnoredDuringExecution:
    {{- include "podAntiAffinityPreferredDuringScheduling" . }}
    {{- end }}
{{- end }}
{{- end }}

{{- define "podAntiAffinityRequiredDuringScheduling" }}
    {{- range $index, $item := .podAntiAffinityLabelSelector }}
    - labelSelector:
        matchExpressions:
        - key: {{ $item.key }}
          operator: {{ $item.operator }}
          {{- if $item.values }}
          values:
          {{- $vals := split "," $item.values }}
          {{- range $i, $v := $vals }}
          - {{ $v | quote }}
          {{- end }}
          {{- end }}
      topologyKey: {{ $item.topologyKey }}
    {{- end }}
{{- end }}

{{- define "podAntiAffinityPreferredDuringScheduling" }}
    {{- range $index, $item := .podAntiAffinityTermLabelSelector }}
    - podAffinityTerm:
        labelSelector:
          matchExpressions:
          - key: {{ $item.key }}
            operator: {{ $item.operator }}
            {{- if $item.values }}
            values:
            {{- $vals := split "," $item.values }}
            {{- range $i, $v := $vals }}
            - {{ $v | quote }}
            {{- end }}
            {{- end }}
        topologyKey: {{ $item.topologyKey }}
      weight: 100
    {{- end }}
{{- end }}
`)

func chartsGatewaysIstioIngressTemplates_affinityTplBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplates_affinityTpl, nil
}

func chartsGatewaysIstioIngressTemplates_affinityTpl() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplates_affinityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/_affinity.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesAutoscaleYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
{{- if and $gateway.autoscaleEnabled $gateway.autoscaleMin $gateway.autoscaleMax }}
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  maxReplicas: {{ $gateway.autoscaleMax }}
  minReplicas: {{ $gateway.autoscaleMin }}
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{ $gateway.name | default "istio-ingressgateway" }}
  metrics:
    - type: Resource
      resource:
        name: cpu
        targetAverageUtilization: {{ $gateway.cpu.targetAverageUtilization }}
---
{{- end }}
`)

func chartsGatewaysIstioIngressTemplatesAutoscaleYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesAutoscaleYaml, nil
}

func chartsGatewaysIstioIngressTemplatesAutoscaleYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesAutoscaleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/autoscale.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesDeploymentYaml = []byte(`{{- $gateway := index .Values "gateways" "istio-ingressgateway" }}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
{{- if not $gateway.autoscaleEnabled }}
{{- if $gateway.replicaCount }}
  replicas: {{ $gateway.replicaCount }}
{{- end }}
{{- end }}
  selector:
    matchLabels:
{{ $gateway.labels | toYaml | indent 6 }}
  strategy:
    rollingUpdate:
      maxSurge: {{ $gateway.rollingMaxSurge }}
      maxUnavailable: {{ $gateway.rollingMaxUnavailable }}
  template:
    metadata:
      labels:
{{ $gateway.labels | toYaml | indent 8 }}
{{- if eq .Release.Namespace "istio-system"}}
        heritage: Tiller
        release: istio
        chart: gateways
{{- end }}
        service.istio.io/canonical-name: {{ $gateway.name | default "istio-ingressgateway" }}
        {{- if not (eq .Values.revision "") }}
        service.istio.io/canonical-revision: {{ .Values.revision }}
        {{- else}}
        service.istio.io/canonical-revision: latest
        {{- end }}
      annotations:
        {{- if .Values.meshConfig.enablePrometheusMerge }}
        prometheus.io/port: "15020"
        prometheus.io/scrape: "true"
        prometheus.io/path: "/stats/prometheus"
        {{- end }}
        sidecar.istio.io/inject: "false"
{{- if $gateway.podAnnotations }}
{{ toYaml $gateway.podAnnotations | indent 8 }}
{{ end }}
    spec:
{{- if not $gateway.runAsRoot }}
      securityContext:
        runAsUser: 1337
        runAsGroup: 1337
        runAsNonRoot: true
        fsGroup: 1337
{{- end }}
      serviceAccountName: {{ $gateway.name | default "istio-ingressgateway" }}-service-account
{{- if .Values.global.priorityClassName }}
      priorityClassName: "{{ .Values.global.priorityClassName }}"
{{- end }}
{{- if .Values.global.proxy.enableCoreDump }}
      initContainers:
        - name: enable-core-dump
{{- if contains "/" .Values.global.proxy.image }}
          image: "{{ .Values.global.proxy.image }}"
{{- else }}
          image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image | default "proxyv2" }}:{{ .Values.global.tag }}"
{{- end }}
{{- if .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
          command:
            - /bin/sh
          args:
            - -c
            - sysctl -w kernel.core_pattern=/var/lib/istio/data/core.proxy && ulimit -c unlimited
          securityContext:
            runAsUser: 0
            runAsGroup: 0
            runAsNonRoot: false
            privileged: true
{{- end }}
      containers:
        - name: istio-proxy
{{- if contains "/" .Values.global.proxy.image }}
          image: "{{ .Values.global.proxy.image }}"
{{- else }}
          image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image | default "proxyv2" }}:{{ .Values.global.tag }}"
{{- end }}
{{- if .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
          ports:
            {{- range $key, $val := $gateway.ports }}
            - containerPort: {{ $val.targetPort | default $val.port }}
            {{- end }}
            {{- if $.Values.global.meshExpansion.enabled }}
            {{- range $key, $val := $gateway.meshExpansionPorts }}
            - containerPort: {{ $val.targetPort | default $val.port }}
            {{- end }}
            {{- end }}
            - containerPort: 15090
              protocol: TCP
              name: http-envoy-prom
          args:
          - proxy
          - router
          - --domain
          - $(POD_NAMESPACE).svc.{{ .Values.global.proxy.clusterDomain }}
        {{- if .Values.global.proxy.logLevel }}
          - --proxyLogLevel={{ .Values.global.proxy.logLevel }}
        {{- end}}
        {{- if .Values.global.proxy.componentLogLevel }}
          - --proxyComponentLogLevel={{ .Values.global.proxy.componentLogLevel }}
        {{- end}}
        {{- if .Values.global.logging.level }}
          - --log_output_level={{ .Values.global.logging.level }}
        {{- end}}
        {{- if .Values.global.logAsJson }}
          - --log_as_json
        {{- end }}
          - --serviceCluster
          - {{ $gateway.name | default "istio-ingressgateway" }}
        {{- if .Values.global.sts.servicePort }}
          - --stsPort={{ .Values.global.sts.servicePort }}
        {{- end }}
        {{- if .Values.global.trustDomain }}
          - --trust-domain={{ .Values.global.trustDomain }}
        {{- end }}
        {{- if not $gateway.runAsRoot }}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - ALL
            privileged: false
            readOnlyRootFilesystem: true
        {{- end }}
          readinessProbe:
            failureThreshold: 30
            httpGet:
              path: /healthz/ready
              port: 15021
              scheme: HTTP
            initialDelaySeconds: 1
            periodSeconds: 2
            successThreshold: 1
            timeoutSeconds: 1
          resources:
{{- if $gateway.resources }}
{{ toYaml $gateway.resources | indent 12 }}
{{- else }}
{{ toYaml .Values.global.defaultResources | indent 12 }}
{{- end }}
          env:
          - name: JWT_POLICY
            value: {{ .Values.global.jwtPolicy }}
          - name: PILOT_CERT_PROVIDER
            value: {{ .Values.global.pilotCertProvider }}
          - name: CA_ADDR
          {{- if .Values.global.caAddress }}
            value: {{ .Values.global.caAddress }}
          {{- else }}
            value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Values.global.istioNamespace }}.svc:15012
          {{- end }}
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: INSTANCE_IP
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: status.podIP
          - name: HOST_IP
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: status.hostIP
          - name: SERVICE_ACCOUNT
            valueFrom:
              fieldRef:
                fieldPath: spec.serviceAccountName
          - name: CANONICAL_SERVICE
            valueFrom:
              fieldRef:
                fieldPath: metadata.labels['service.istio.io/canonical-name']
          - name: CANONICAL_REVISION
            valueFrom:
              fieldRef:
                fieldPath: metadata.labels['service.istio.io/canonical-revision']
          - name: ISTIO_META_WORKLOAD_NAME
            value: {{ $gateway.name | default "istio-ingressgateway" }}
          - name: ISTIO_META_OWNER
            value: kubernetes://apis/apps/v1/namespaces/{{ .Release.Namespace }}/deployments/{{ $gateway.name | default "istio-ingressgateway" }}
          {{- if $.Values.global.meshID }}
          - name: ISTIO_META_MESH_ID
            value: "{{ $.Values.global.meshID }}"
          {{- else if $.Values.global.trustDomain }}
          - name: ISTIO_META_MESH_ID
            value: "{{ $.Values.global.trustDomain }}"
          {{- end }}
          {{- range $key, $val := $gateway.env }}
          - name: {{ $key }}
            value: {{ $val }}
          {{- end }}
          {{- range $key, $value := .Values.meshConfig.defaultConfig.proxyMetadata }}
          - name: {{ $key }}
            value: "{{ $value }}"
          {{- end }}
          {{ $network_set := index $gateway.env "ISTIO_META_NETWORK" }}
          {{- if and (not $network_set) .Values.global.network }}
          - name: ISTIO_META_NETWORK
            value: {{ .Values.global.network }}
          {{- end }}
{{- if $gateway.podAnnotations }}
          - name: "ISTIO_METAJSON_ANNOTATIONS"
            value: |
{{ toJson $gateway.podAnnotations | indent 16}}
{{ end }}
          - name: ISTIO_META_CLUSTER_ID
            value: "{{ $.Values.global.multiCluster.clusterName | default `+"`"+`Kubernetes`+"`"+` }}"
          volumeMounts:
          - name: istio-envoy
            mountPath: /etc/istio/proxy
          - name: config-volume
            mountPath: /etc/istio/config
{{- if eq .Values.global.pilotCertProvider "istiod" }}
          - mountPath: /var/run/secrets/istio
            name: istiod-ca-cert
{{- end }}
{{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
          - name: istio-token
            mountPath: /var/run/secrets/tokens
            readOnly: true
{{- end }}
          - name: gatewaysdsudspath
            mountPath: /var/run/ingress_gateway
          {{- if .Values.global.mountMtlsCerts }}
          # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
          - name: istio-certs
            mountPath: /etc/certs
            readOnly: true
          {{- end }}
          - mountPath: /var/lib/istio/data
            name: istio-data
          - name: podinfo
            mountPath: /etc/istio/pod
          {{- range $gateway.secretVolumes }}
          - name: {{ .name }}
            mountPath: {{ .mountPath | quote }}
            readOnly: true
          {{- end }}
          {{- range $gateway.configVolumes }}
          {{- if .mountPath }}
          - name: {{ .name }}
            mountPath: {{ .mountPath | quote }}
            readOnly: true
          {{- end }}
          {{- end }}
{{- if $gateway.additionalContainers }}
{{ toYaml $gateway.additionalContainers | indent 8 }}
{{- end }}
      volumes:
{{- if eq .Values.global.pilotCertProvider "istiod" }}
      - name: istiod-ca-cert
        configMap:
          name: istio-ca-root-cert
{{- end }}
      - name: podinfo
        downwardAPI:
          items:
            - path: "labels"
              fieldRef:
                fieldPath: metadata.labels
            - path: "annotations"
              fieldRef:
                fieldPath: metadata.annotations
      - name: istio-envoy
        emptyDir: {}
      - name: gatewaysdsudspath
        emptyDir: {}
      - name: istio-data
        emptyDir: {}
{{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
      - name: istio-token
        projected:
          sources:
          - serviceAccountToken:
              path: istio-token
              expirationSeconds: 43200
              audience: {{ .Values.global.sds.token.aud }}
{{- end }}
      {{- if .Values.global.mountMtlsCerts }}
      # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
      - name: istio-certs
        secret:
          secretName: istio.istio-ingressgateway-service-account
          optional: true
      {{- end }}
      - name: config-volume
        configMap:
          name: istio{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
          optional: true
      {{- range $gateway.secretVolumes }}
      - name: {{ .name }}
        secret:
          secretName: {{ .secretName | quote }}
          optional: true
      {{- end }}
      {{- range $gateway.configVolumes }}
      - name: {{ .name }}
        configMap:
          name: {{ .configMapName | quote }}
          optional: true
      {{- end }}
      affinity:
      {{- include "nodeaffinity" (dict "global" .Values.global "nodeSelector" $gateway.nodeSelector) | indent 6 }}
      {{- include "podAntiAffinity" $gateway | indent 6 }}
{{- if $gateway.tolerations }}
      tolerations:
{{ toYaml $gateway.tolerations | indent 6 }}
{{- else if .Values.global.defaultTolerations }}
      tolerations:
{{ toYaml .Values.global.defaultTolerations | indent 6 }}
{{- end }}
`)

func chartsGatewaysIstioIngressTemplatesDeploymentYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesDeploymentYaml, nil
}

func chartsGatewaysIstioIngressTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesMeshexpansionYaml = []byte(`{{- if .Values.global.meshExpansion.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: meshexpansion-gateway
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
spec:
  selector:
    istio: ingressgateway
  servers:
    - port:
        number: 15012
        protocol: TCP
        name: tcp-istiod
      hosts:
        - "*"
    - port:
        number: 15017
        protocol: TCP
        name: tcp-istiodwebhook
      hosts:
        - "*"
---

apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: meshexpansion-vs-istiod
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
spec:
  hosts:
  - istiod.{{ .Values.global.istioNamespace }}.svc.{{ .Values.global.proxy.clusterDomain }}
  gateways:
  - meshexpansion-gateway
  tcp:
  - match:
    - port: 15012
    route:
    - destination:
        host: istiod.{{ .Values.global.istioNamespace }}.svc.{{ .Values.global.proxy.clusterDomain }}
        port:
          number: 15012
  - match:
    - port: 15017
    route:
    - destination:
        host: istiod.{{ .Release.Namespace }}.svc.{{ .Values.global.proxy.clusterDomain }}
        port:
          number: 443
---

apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: meshexpansion-dr-istiod
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
spec:
  host: istiod.{{ .Release.Namespace }}.svc.{{ .Values.global.proxy.clusterDomain }}
  trafficPolicy:
    portLevelSettings:
    - port:
        number: 15012
      tls:
        mode: DISABLE
    - port:
        number: 15017
      tls:
        mode: DISABLE

{{- end }}
`)

func chartsGatewaysIstioIngressTemplatesMeshexpansionYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesMeshexpansionYaml, nil
}

func chartsGatewaysIstioIngressTemplatesMeshexpansionYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesMeshexpansionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/meshexpansion.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYaml = []byte(`{{- if .Values.global.defaultPodDisruptionBudget.enabled }}
{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | trim | indent 4 }}
    release: {{ .Release.Name }}
spec:
  minAvailable: 1
  selector:
    matchLabels:
{{ $gateway.labels | toYaml | trim | indent 6 }}
{{- end }}
`)

func chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYaml, nil
}

func chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/poddisruptionbudget.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesPreconfiguredYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
{{- if .Values.global.multiCluster.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: istio-multicluster-ingressgateway
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  selector:
    istio: ingressgateway
  servers:
  - hosts:
    - "*.global"
    port:
      name: tls
      number: 15443
      protocol: TLS
    tls:
      mode: AUTO_PASSTHROUGH
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: istio-multicluster-ingressgateway
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  workloadSelector:
    labels:
      istio: ingressgateway
  configPatches:
  - applyTo: NETWORK_FILTER
    match:
      context: GATEWAY
      listener:
        portNumber: 15443
        filterChain:
          filter:
            name: "envoy.filters.network.sni_cluster"
    patch:
      operation: INSERT_AFTER
      value:
        name: "envoy.filters.network.tcp_cluster_rewrite"
        config:
          cluster_pattern: "\\.global$"
          cluster_replacement: ".svc.{{ .Values.global.proxy.clusterDomain }}"
---
## To ensure all traffic to *.global is using mTLS
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: istio-multicluster-ingressgateway
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
  host: "*.global"
  {{- if .Values.global.defaultConfigVisibilitySettings }}
  exportTo:
  - '*'
  {{- end }}
  trafficPolicy:
    tls:
      mode: ISTIO_MUTUAL
---
{{- end }}
`)

func chartsGatewaysIstioIngressTemplatesPreconfiguredYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesPreconfiguredYaml, nil
}

func chartsGatewaysIstioIngressTemplatesPreconfiguredYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesPreconfiguredYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/preconfigured.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesRoleYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}-sds
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]
---
`)

func chartsGatewaysIstioIngressTemplatesRoleYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesRoleYaml, nil
}

func chartsGatewaysIstioIngressTemplatesRoleYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesRolebindingsYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}-sds
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ $gateway.name | default "istio-ingressgateway" }}-sds
subjects:
- kind: ServiceAccount
  name: {{ $gateway.name | default "istio-ingressgateway" }}-service-account
---
`)

func chartsGatewaysIstioIngressTemplatesRolebindingsYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesRolebindingsYaml, nil
}

func chartsGatewaysIstioIngressTemplatesRolebindingsYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesRolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesServiceYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
{{- if not $gateway.customService }}
apiVersion: v1
kind: Service
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}
  namespace: {{ .Release.Namespace }}
  annotations:
    {{- range $key, $val := $gateway.serviceAnnotations }}
    {{ $key }}: {{ $val | quote }}
    {{- end }}
  labels:
{{ $gateway.labels | toYaml | indent 4 }}
    release: {{ .Release.Name }}
spec:
{{- if $gateway.loadBalancerIP }}
  loadBalancerIP: "{{ $gateway.loadBalancerIP }}"
{{- end }}
{{- if $gateway.loadBalancerSourceRanges }}
  loadBalancerSourceRanges:
{{ toYaml $gateway.loadBalancerSourceRanges | indent 4 }}
{{- end }}
{{- if $gateway.externalTrafficPolicy }}
  externalTrafficPolicy: {{$gateway.externalTrafficPolicy }}
{{- end }}
  type: {{ $gateway.type }}
  selector:
{{ $gateway.labels | toYaml | indent 4 }}
  ports:

    {{- range $key, $val := $gateway.ports }}
    -
      {{- range $pkey, $pval := $val }}
      {{ $pkey}}: {{ $pval }}
      {{- end }}
    {{- end }}

    {{- if $.Values.global.meshExpansion.enabled }}
    {{- range $key, $val := $gateway.meshExpansionPorts }}
    -
      {{- range $pkey, $pval := $val }}
      {{ $pkey}}: {{ $pval }}
      {{- end }}
    {{- end }}
    {{- end }}
  {{ range $app := $gateway.ingressPorts }}
    -
      port: {{ $app.port }}
      name: {{ $app.name }}
  {{- end }}
---
{{ end }}
`)

func chartsGatewaysIstioIngressTemplatesServiceYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesServiceYaml, nil
}

func chartsGatewaysIstioIngressTemplatesServiceYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressTemplatesServiceaccountYaml = []byte(`{{ $gateway := index .Values "gateways" "istio-ingressgateway" }}
apiVersion: v1
kind: ServiceAccount
{{- if .Values.global.imagePullSecrets }}
imagePullSecrets:
{{- range .Values.global.imagePullSecrets }}
  - name: {{ . }}
{{- end }}
{{- end }}
metadata:
  name: {{ $gateway.name | default "istio-ingressgateway" }}-service-account
  namespace: {{ .Release.Namespace }}
  labels:
{{ $gateway.labels | toYaml | trim | indent 4 }}
    release: {{ .Release.Name }}
`)

func chartsGatewaysIstioIngressTemplatesServiceaccountYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressTemplatesServiceaccountYaml, nil
}

func chartsGatewaysIstioIngressTemplatesServiceaccountYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressTemplatesServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/templates/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGatewaysIstioIngressValuesYaml = []byte(`# A-la-carte istio ingress gateway.
# Must be installed in a separate namespace, to minimize access to secrets.

gateways:
  istio-ingressgateway:
    name: istio-ingressgateway
    labels:
      app: istio-ingressgateway
      istio: ingressgateway
    ports:
    ## You can add custom gateway ports in user values overrides, but it must include those ports since helm replaces.
    # Note that AWS ELB will by default perform health checks on the first port
    # on this list. Setting this to the health check port will ensure that health
    # checks always work. https://github.com/istio/istio/issues/12503
    - port: 15021
      targetPort: 15021
      name: status-port
    - port: 80
      targetPort: 8080
      name: http2
    - port: 443
      targetPort: 8443
      name: https
      # This is the port where sni routing happens
    - port: 15443
      targetPort: 15443
      name: tls

    # Scalability tunning
    # replicaCount: 1
    rollingMaxSurge: 100%
    rollingMaxUnavailable: 25%
    autoscaleEnabled: true
    autoscaleMin: 1
    autoscaleMax: 5

    cpu:
      targetAverageUtilization: 80

    resources:
      requests:
        cpu: 100m
        memory: 128Mi
      limits:
        cpu: 2000m
        memory: 1024Mi

    loadBalancerIP: ""
    loadBalancerSourceRanges: []
    serviceAnnotations: {}

    # Enable cross-cluster access using SNI matching
    zvpn:
      enabled: false
      suffix: global

    # To generate an internal load balancer:
    # --set serviceAnnotations.cloud.google.com/load-balancer-type=internal
    #serviceAnnotations:
    #    cloud.google.com/load-balancer-type: "internal"

    podAnnotations: {}
    type: LoadBalancer #change to NodePort, ClusterIP or LoadBalancer if need be

    #### MESH EXPANSION PORTS  ########
    # Pilot and Citadel MTLS ports are enabled in gateway - but will only redirect
    # to pilot/citadel if global.meshExpansion settings are enabled.
    # Delete these ports if mesh expansion is not enabled, to avoid
    # exposing unnecessary ports on the web.
    # You can remove these ports if you are not using mesh expansion
    meshExpansionPorts:
    - port: 15012
      targetPort: 15012
      name: tcp-istiod
    - port: 853
      targetPort: 8853
      name: tcp-dns-tls
    ####### end MESH EXPANSION PORTS ######

    ##############
    secretVolumes:
    - name: ingressgateway-certs
      secretName: istio-ingressgateway-certs
      mountPath: /etc/istio/ingressgateway-certs
    - name: ingressgateway-ca-certs
      secretName: istio-ingressgateway-ca-certs
      mountPath: /etc/istio/ingressgateway-ca-certs

    customService: false
    externalTrafficPolicy: ""

    ingressPorts: []
    additionalContainers: []
    configVolumes: []

    ### Advanced options ############
    env:
      # A gateway with this mode ensures that pilot generates an additional
      # set of clusters for internal services but without Istio mTLS, to
      # enable cross cluster routing.
      ISTIO_META_ROUTER_MODE: "sni-dnat"

    nodeSelector: {}
    tolerations: []

    # Specify the pod anti-affinity that allows you to constrain which nodes
    # your pod is eligible to be scheduled based on labels on pods that are
    # already running on the node rather than based on labels on nodes.
    # There are currently two types of anti-affinity:
    #    "requiredDuringSchedulingIgnoredDuringExecution"
    #    "preferredDuringSchedulingIgnoredDuringExecution"
    # which denote "hard" vs. "soft" requirements, you can define your values
    # in "podAntiAffinityLabelSelector" and "podAntiAffinityTermLabelSelector"
    # correspondingly.
    # For example:
    # podAntiAffinityLabelSelector:
    # - key: security
    #   operator: In
    #   values: S1,S2
    #   topologyKey: "kubernetes.io/hostname"
    # This pod anti-affinity rule says that the pod requires not to be scheduled
    # onto a node if that node is already running a pod with label having key
    # "security" and value "S1".
    podAntiAffinityLabelSelector: []
    podAntiAffinityTermLabelSelector: []

    # whether to run the gateway in a privileged container
    runAsRoot: false

# Revision is set as 'version' label and part of the resource names when installing multiple control planes.
revision: ""

global:
  # set the default set of namespaces to which services, service entries, virtual services, destination
  # rules should be exported to. Currently only one value can be provided in this list. This value
  # should be one of the following two options:
  # * implies these objects are visible to all namespaces, enabling any sidecar to talk to any other sidecar.
  # . implies these objects are visible to only to sidecars in the same namespace, or if imported as a Sidecar.egress.host
  defaultConfigVisibilitySettings: []

  # enable pod disruption budget for the control plane, which is used to
  # ensure Istio control plane components are gradually upgraded or recovered.
  defaultPodDisruptionBudget:
    enabled: true

  # A minimal set of requested resources to applied to all deployments so that
  # Horizontal Pod Autoscaler will be able to function (if set).
  # Each component can overwrite these default values by adding its own resources
  # block in the relevant section below and setting the desired resources values.
  defaultResources:
    requests:
      cpu: 10m
    #   memory: 128Mi
    # limits:
    #   cpu: 100m
    #   memory: 128Mi

  # Default node tolerations to be applied to all deployments so that all pods can be
  # scheduled to a particular nodes with matching taints. Each component can overwrite
  # these default values by adding its tolerations block in the relevant section below
  # and setting the desired values.
  # Configure this field in case that all pods of Istio control plane are expected to
  # be scheduled to particular nodes with specified taints.
  defaultTolerations: []

  # Default hub for Istio images.
  # Releases are published to docker hub under 'istio' project.
  # Dev builds from prow are on gcr.io
  hub: gcr.io/istio-testing

  # Default tag for Istio images.
  tag: latest

  # Specify image pull policy if default behavior isn't desired.
  # Default behavior: latest images will be Always else IfNotPresent.
  imagePullPolicy: ""

  # ImagePullSecrets for all ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # For components that don't use ServiceAccounts (i.e. grafana, servicegraph, tracing)
  # ImagePullSecrets will be added to the corresponding Deployment(StatefulSet) objects.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []
  # - private-registry-key

  # To output all istio components logs in json format by adding --log_as_json argument to each container argument
  logAsJson: false

  # Comma-separated minimum per-scope logging level of messages to output, in the form of <scope>:<level>,<scope>:<level>
  # The control plane has different scopes depending on component, but can configure default log level across all components
  # If empty, default scope and level will be used as configured in code
  logging:
    level: "default:info"

  # If set to true, the pilot and citadel mtls will be exposed on the
  # ingress gateway
  meshExpansion:
    enabled: false
    # If set to true, the pilot and citadel mtls and the plain text pilot ports
    # will be exposed on an internal gateway
    useILB: false

  # Kubernetes >=v1.11.0 will create two PriorityClass, including system-cluster-critical and
  # system-node-critical, it is better to configure this in order to make sure your Istio pods
  # will not be killed because of low priority class.
  # Refer to https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
  # for more detail.
  priorityClassName: ""

  proxy:
    image: proxyv2

    # CAUTION: It is important to ensure that all Istio helm charts specify the same clusterDomain value
    # cluster domain. Default value is "cluster.local".
    clusterDomain: "cluster.local"

    # Per Component log level for proxy, applies to gateways and sidecars. If a component level is
    # not set, then the global "logLevel" will be used.
    componentLogLevel: "misc:error"

    # If set, newly injected sidecars will have core dumps enabled.
    enableCoreDump: false

    # Log level for proxy, applies to gateways and sidecars.
    # Expected values are: trace|debug|info|warning|error|critical|off
    logLevel: warning

  ##############################################################################################
  # The following values are found in other charts. To effectively modify these values, make   #
  # make sure they are consistent across your Istio helm charts                                #
  ##############################################################################################

  # The customized CA address to retrieve certificates for the pods in the cluster.
  # CSR clients such as the Istio Agent and ingress gateways can use this to specify the CA endpoint.
  caAddress: ""

  # Used to locate istiod.
  istioNamespace: istio-system

  # Configure the policy for validating JWT.
  # Currently, two options are supported: "third-party-jwt" and "first-party-jwt".
  jwtPolicy: "third-party-jwt"

  # Mesh ID means Mesh Identifier. It should be unique within the scope where
  # meshes will interact with each other, but it is not required to be
  # globally/universally unique. For example, if any of the following are true,
  # then two meshes must have different Mesh IDs:
  # - Meshes will have their telemetry aggregated in one place
  # - Meshes will be federated together
  # - Policy will be written referencing one mesh from the other
  #
  # If an administrator expects that any of these conditions may become true in
  # the future, they should ensure their meshes have different Mesh IDs
  # assigned.
  #
  # Within a multicluster mesh, each cluster must be (manually or auto)
  # configured to have the same Mesh ID value. If an existing cluster 'joins' a
  # multicluster mesh, it will need to be migrated to the new mesh ID. Details
  # of migration TBD, and it may be a disruptive operation to change the Mesh
  # ID post-install.
  #
  # If the mesh admin does not specify a value, Istio will use the value of the
  # mesh's Trust Domain. The best practice is to select a proper Trust Domain
  # value.
  meshID: ""

  # Use the user-specified, secret volume mounted key and certs for Pilot and workloads.
  mountMtlsCerts: false

  multiCluster:
    # Set to true to connect two kubernetes clusters via their respective
    # ingressgateway services when pods in each cluster cannot directly
    # talk to one another. All clusters should be using Istio mTLS and must
    # have a shared root CA for this model to work.
    enabled: false
    # Should be set to the name of the cluster this installation will run in. This is required for sidecar injection
    # to properly label proxies
    clusterName: ""

  # Network defines the network this cluster belong to. This name
  # corresponds to the networks in the map of mesh networks.
  network: ""

  # Configure the certificate provider for control plane communication.
  # Currently, two providers are supported: "kubernetes" and "istiod".
  # As some platforms may not have kubernetes signing APIs,
  # Istiod is the default
  pilotCertProvider: istiod

  sds:
    # The JWT token for SDS and the aud field of such JWT. See RFC 7519, section 4.1.3.
    # When a CSR is sent from Citadel Agent to the CA (e.g. Citadel), this aud is to make sure the
    # JWT is intended for the CA.
    token:
      aud: istio-ca

  sts:
    # The service port used by Security Token Service (STS) server to handle token exchange requests.
    # Setting this port to a non-zero value enables STS server.
    servicePort: 0

  # The trust domain corresponds to the trust root of a system
  # Refer to https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#21-trust-domain
  # Indicate the domain used in SPIFFE identity URL
  # The default depends on the environment.
  #   kubernetes: cluster.local
  #   else:  default dns domain
  trustDomain: "cluster.local"

meshConfig:
  enablePrometheusMerge: true
  defaultConfig:
    proxyMetadata: {}
    tracing:
    #      tlsSettings:
    #        mode: DISABLE # DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL
    #        clientCertificate: # example: /etc/istio/tracer/cert-chain.pem
    #        privateKey:        # example: /etc/istio/tracer/key.pem
    #        caCertificates:    # example: /etc/istio/tracer/root-cert.pem
    #        sni:               # example: tracer.somedomain
    #        subjectAltNames: []
    # - tracer.somedomain
`)

func chartsGatewaysIstioIngressValuesYamlBytes() ([]byte, error) {
	return _chartsGatewaysIstioIngressValuesYaml, nil
}

func chartsGatewaysIstioIngressValuesYaml() (*asset, error) {
	bytes, err := chartsGatewaysIstioIngressValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/gateways/istio-ingress/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsGlobalYaml = []byte(`# Global and common settings for installing Istio.

# This file is configured for a small scale production cluster.
# Use custom settings to tune up the CPU and scaling.
# Additional values overrides can be used.

# Each individual component will use values from this file, with defaults and 'advanced' settings included in
# its own chart's values.yaml.

# TODO: trim this file to commonly used settings, leave 'advanced' in the individual values.yaml (they can
# still be overridden by users, but won't show in basic documentation.

# This doesn't match istio defaults, which are more geared towards tests and bookinfo.

global:
  # Used to locate istiod.
  istioNamespace: istio-system

  # Default hub for Istio images.
  # Releases are published to docker hub under 'istio' project.
  # Dev builds from prow are on gcr.io
  hub: gcr.io/istio-testing

  # Default tag for Istio images.
  tag: latest

  # Comma-separated minimum per-scope logging level of messages to output, in the form of <scope>:<level>,<scope>:<level>
  # The control plane has different scopes depending on component, but can configure default log level across all components
  # If empty, default scope and level will be used as configured in code
  logging:
    level: "default:info"

  # To output all istio components logs in json format by adding --log_as_json argument to each container argument
  logAsJson: false

  # Enabled by default in master for maximising testing.
  istiod:
    enableAnalysis: false

  # One central istiod controls all remote clusters: disabled by default
  centralIstiod: false

  proxy:
    image: proxyv2

    # cluster domain. Default value is "cluster.local".
    clusterDomain: "cluster.local"

    # Resources for the sidecar.
    resources:
      requests:
        cpu: 100m
        memory: 128Mi
      limits:
        cpu: 2000m
        memory: 1024Mi

    # Log level for proxy, applies to gateways and sidecars.
    # Expected values are: trace|debug|info|warning|error|critical|off
    logLevel: warning

    # Per Component log level for proxy, applies to gateways and sidecars. If a component level is
    # not set, then the global "logLevel" will be used.
    componentLogLevel: "misc:error"

    #If set to true, istio-proxy container will have privileged securityContext
    privileged: false

    # If set, newly injected sidecars will have core dumps enabled.
    enableCoreDump: false

    # Default port for Pilot agent health checks. A value of 0 will disable health checking.
    statusPort: 15020

    # The initial delay for readiness probes in seconds.
    readinessInitialDelaySeconds: 1

    # The period between readiness probes.
    readinessPeriodSeconds: 2

    # The number of successive failed probes before indicating readiness failure.
    readinessFailureThreshold: 30

    # istio egress capture allowlist
    # https://istio.io/docs/tasks/traffic-management/egress.html#calling-external-services-directly
    # example: includeIPRanges: "172.30.0.0/16,172.20.0.0/16"
    # would only capture egress traffic on those two IP Ranges, all other outbound traffic would
    # be allowed by the sidecar
    includeIPRanges: "*"
    excludeIPRanges: ""
    excludeOutboundPorts: ""

    # istio ingress capture allowlist
    # examples:
    #     Redirect only selected ports:            --includeInboundPorts="80,8080"
    excludeInboundPorts: ""

    # This controls the 'policy' in the sidecar injector.
    autoInject: enabled

    # Specify which tracer to use. One of: zipkin, lightstep, datadog, stackdriver.
    # If using stackdriver tracer outside GCP, set env GOOGLE_APPLICATION_CREDENTIALS to the GCP credential file.
    tracer: "zipkin"

  proxy_init:
    # Base name for the proxy_init container, used to configure iptables.
    image: proxyv2
    resources:
      limits:
        cpu: 2000m
        memory: 1024Mi
      requests:
        cpu: 10m
        memory: 10Mi

  # Specify image pull policy if default behavior isn't desired.
  # Default behavior: latest images will be Always else IfNotPresent.
  imagePullPolicy: ""

  # Use the user-specified, secret volume mounted key and certs for Pilot and workloads.
  mountMtlsCerts: false

  # Configuration for each of the supported tracers
  tracer:
    # Configuration for envoy to send trace data to LightStep.
    # Disabled by default.
    # address: the <host>:<port> of the satellite pool
    # accessToken: required for sending data to the pool
    #
    lightstep:
      address: ""                # example: lightstep-satellite:443
      accessToken: ""            # example: abcdefg1234567
    zipkin:
      # Host:Port for reporting trace data in zipkin format. If not specified, will default to
      # zipkin service (port 9411) in the same namespace as the other istio components.
      address: ""
    datadog:
      # Host:Port for submitting traces to the Datadog agent.
      address: "$(HOST_IP):8126"
    stackdriver:
      # enables trace output to stdout.
      debug: false
      # The global default max number of attributes per span.
      maxNumberOfAttributes: 200
      # The global default max number of annotation events per span.
      maxNumberOfAnnotations: 200
      # The global default max number of message events per span.
      maxNumberOfMessageEvents: 200

  # ImagePullSecrets for all ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # For components that don't use ServiceAccounts (i.e. grafana, servicegraph, tracing)
  # ImagePullSecrets will be added to the corresponding Deployment(StatefulSet) objects.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []
    # - private-registry-key

  # Specify pod scheduling arch(amd64, ppc64le, s390x) and weight as follows:
  #   0 - Never scheduled
  #   1 - Least preferred
  #   2 - No preference
  #   3 - Most preferred
  arch:
    amd64: 2
    s390x: 2
    ppc64le: 2

  # Whether to restrict the applications namespace the controller manages;
  # If not set, controller watches all namespaces
  oneNamespace: false

  # Default node selector to be applied to all deployments so that all pods can be
  # constrained to run a particular nodes. Each component can overwrite these default
  # values by adding its node selector block in the relevant section below and setting
  # the desired values.
  defaultNodeSelector: {}

  # Default node tolerations to be applied to all deployments so that all pods can be
  # scheduled to a particular nodes with matching taints. Each component can overwrite
  # these default values by adding its tolerations block in the relevant section below
  # and setting the desired values.
  # Configure this field in case that all pods of Istio control plane are expected to
  # be scheduled to particular nodes with specified taints.
  defaultTolerations: []

  # Whether to perform server-side validation of configuration.
  configValidation: true

  # Custom DNS config for the pod to resolve names of services in other
  # clusters. Use this to add additional search domains, and other settings.
  # see
  # https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#dns-config
  # This does not apply to gateway pods as they typically need a different
  # set of DNS settings than the normal application pods (e.g., in
  # multicluster scenarios).
  # NOTE: If using templates, follow the pattern in the commented example below.
  #podDNSSearchNamespaces:
  #- global
  #- "{{ valueOrDefault .DeploymentMeta.Namespace \"default\" }}.global"

  # If set to true, the pilot and citadel mtls will be exposed on the
  # ingress gateway
  meshExpansion:
    enabled: false
    # If set to true, the pilot and citadel mtls and the plain text pilot ports
    # will be exposed on an internal gateway
    useILB: false

  multiCluster:
    # Set to true to connect two kubernetes clusters via their respective
    # ingressgateway services when pods in each cluster cannot directly
    # talk to one another. All clusters should be using Istio mTLS and must
    # have a shared root CA for this model to work.
    enabled: false
    # Should be set to the name of the cluster this installation will run in. This is required for sidecar injection
    # to properly label proxies
    clusterName: ""

  # A minimal set of requested resources to applied to all deployments so that
  # Horizontal Pod Autoscaler will be able to function (if set).
  # Each component can overwrite these default values by adding its own resources
  # block in the relevant section below and setting the desired resources values.
  defaultResources:
    requests:
      cpu: 10m
    #   memory: 128Mi
    # limits:
    #   cpu: 100m
    #   memory: 128Mi

  # enable pod disruption budget for the control plane, which is used to
  # ensure Istio control plane components are gradually upgraded or recovered.
  defaultPodDisruptionBudget:
    enabled: true
    # The values aren't mutable due to a current PodDisruptionBudget limitation
    # minAvailable: 1

  # Kubernetes >=v1.11.0 will create two PriorityClass, including system-cluster-critical and
  # system-node-critical, it is better to configure this in order to make sure your Istio pods
  # will not be killed because of low priority class.
  # Refer to https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
  # for more detail.
  priorityClassName: ""

  # Use the Mesh Control Protocol (MCP) for configuring Istiod. Requires an MCP source.
  useMCP: false

  # The trust domain corresponds to the trust root of a system
  # Refer to https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#21-trust-domain
  # Indicate the domain used in SPIFFE identity URL
  # The default depends on the environment.
  #   kubernetes: cluster.local
  #   else:  default dns domain
  trustDomain: "cluster.local"

  # Mesh ID means Mesh Identifier. It should be unique within the scope where
  # meshes will interact with each other, but it is not required to be
  # globally/universally unique. For example, if any of the following are true,
  # then two meshes must have different Mesh IDs:
  # - Meshes will have their telemetry aggregated in one place
  # - Meshes will be federated together
  # - Policy will be written referencing one mesh from the other
  #
  # If an administrator expects that any of these conditions may become true in
  # the future, they should ensure their meshes have different Mesh IDs
  # assigned.
  #
  # Within a multicluster mesh, each cluster must be (manually or auto)
  # configured to have the same Mesh ID value. If an existing cluster 'joins' a
  # multicluster mesh, it will need to be migrated to the new mesh ID. Details
  # of migration TBD, and it may be a disruptive operation to change the Mesh
  # ID post-install.
  #
  # If the mesh admin does not specify a value, Istio will use the value of the
  # mesh's Trust Domain. The best practice is to select a proper Trust Domain
  # value.
  meshID: ""

  # The namespace where globally shared configurations should be present.
  # DestinationRules that apply to the entire mesh (e.g., enabling mTLS),
  # default Sidecar configs, etc. should be added to this namespace.
  # configRootNamespace: istio-config

  # set the default set of namespaces to which services, service entries, virtual services, destination
  # rules should be exported to. Currently only one value can be provided in this list. This value
  # should be one of the following two options:
  # * implies these objects are visible to all namespaces, enabling any sidecar to talk to any other sidecar.
  # . implies these objects are visible to only to sidecars in the same namespace, or if imported as a Sidecar.egress.host
  defaultConfigVisibilitySettings: []
#  - '*'
  omitSidecarInjectorConfigMap: false
  sds:
    token:
      aud: istio-ca

  sts:
    # The service port used by Security Token Service (STS) server to handle token exchange requests.
    # Setting this port to a non-zero value enables STS server.
    servicePort: 0

  # The customized CA address to retrieve certificates for the pods in the cluster.
  # CSR clients such as the Istio Agent and ingress gateways can use this to specify the CA endpoint.
  caAddress: ""

  # Configure the mesh networks to be used by the Split Horizon EDS.
  #
  # The following example defines two networks with different endpoints association methods.
  # For `+"`"+`network1`+"`"+` all endpoints that their IP belongs to the provided CIDR range will be
  # mapped to network1. The gateway for this network example is specified by its public IP
  # address and port.
  # The second network, `+"`"+`network2`+"`"+`, in this example is defined differently with all endpoints
  # retrieved through the specified Multi-Cluster registry being mapped to network2. The
  # gateway is also defined differently with the name of the gateway service on the remote
  # cluster. The public IP for the gateway will be determined from that remote service (only
  # LoadBalancer gateway service type is currently supported, for a NodePort type gateway service,
  # it still need to be configured manually).
  #
  # meshNetworks:
  #   network1:
  #     endpoints:
  #     - fromCidr: "192.168.0.1/24"
  #     gateways:
  #     - address: 1.1.1.1
  #       port: 80
  #   network2:
  #     endpoints:
  #     - fromRegistry: reg1
  #     gateways:
  #     - registryServiceName: istio-ingressgateway.istio-system.svc.cluster.local
  #       port: 443
  #
  meshNetworks: {}

  # Network defines the network this cluster belong to. This name
  # corresponds to the networks in the map of mesh networks.
  network: ""

  # Configure whether Operator manages webhook configurations. The current behavior
  # of Istiod is to manage its own webhook configurations.
  # When this option is set as true, Istio Operator, instead of webhooks, manages the
  # webhook configurations. When this option is set as false, webhooks manage their
  # own webhook configurations.
  operatorManageWebhooks: false

  # configure remote pilot and istiod service and endpoint
  remotePilotAddress: ""

  # Configure the certificate provider for control plane communication.
  # Currently, two providers are supported: "kubernetes" and "istiod".
  # As some platforms may not have kubernetes signing APIs,
  # Istiod is the default
  pilotCertProvider: istiod

  # Configure the policy for validating JWT.
  # Currently, two options are supported: "third-party-jwt" and "first-party-jwt".
  jwtPolicy: "third-party-jwt"

# Internal setting - used when generating helm templates for kustomize.
# clusterResources controls the inclusion of cluster-wide resources when generating the charts/installing.
# For backward compat, it is set to 'true', resulting in the old-style installation.
# When set to 'false', all cluster-wide resources will be omitted, and are expected to be installed
# at the same time with the CRDs.
clusterResources: true

meshConfig:
  enablePrometheusMerge: true
  defaultConfig:
    proxyMetadata: {}
    tracing:
#      tlsSettings:
#        mode: DISABLE # DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL
#        clientCertificate: # example: /etc/istio/tracer/cert-chain.pem
#        privateKey:        # example: /etc/istio/tracer/key.pem
#        caCertificates:    # example: /etc/istio/tracer/root-cert.pem
#        sni:               # example: tracer.somedomain
#        subjectAltNames: []
        # - tracer.somedomain
`)

func chartsGlobalYamlBytes() ([]byte, error) {
	return _chartsGlobalYaml, nil
}

func chartsGlobalYaml() (*asset, error) {
	bytes, err := chartsGlobalYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/global.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniChartYaml = []byte(`apiVersion: v1
name: istio-cni
version: 1.1.0
description: Helm chart for istio-cni components
keywords:
  - istio-cni
  - istio
sources:
  - http://github.com/istio/cni
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsIstioCniChartYamlBytes() ([]byte, error) {
	return _chartsIstioCniChartYaml, nil
}

func chartsIstioCniChartYaml() (*asset, error) {
	bytes, err := chartsIstioCniChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniTemplatesClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istio-cni
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
rules:
- apiGroups: [""]
  resources:
  - pods
  - nodes
  verbs:
  - get
---
{{- if .Values.cni.repair.enabled }}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istio-cni-repair-role
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch", "delete", "patch", "update" ]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["get", "list", "watch", "delete", "patch", "update", "create" ]
{{- end }}
---
  {{- if .Values.cni.taint.enabled }}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istio-cni-taint-role
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
rules:
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["get", "list", "watch", "patch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "create", "update"]
  {{- end }}
`)

func chartsIstioCniTemplatesClusterroleYamlBytes() ([]byte, error) {
	return _chartsIstioCniTemplatesClusterroleYaml, nil
}

func chartsIstioCniTemplatesClusterroleYaml() (*asset, error) {
	bytes, err := chartsIstioCniTemplatesClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/templates/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniTemplatesClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istio-cni
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istio-cni
subjects:
- kind: ServiceAccount
  name: istio-cni
  namespace: {{ .Release.Namespace }}
---
{{- if .Values.cni.repair.enabled }}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istio-cni-repair-rolebinding
  namespace: {{ .Release.Namespace}}
  labels:
    k8s-app: istio-cni-repair
subjects:
- kind: ServiceAccount
  name: istio-cni
  namespace: {{ .Release.Namespace}}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istio-cni-repair-role
{{- end }}
---
{{- if ne .Values.cni.psp_cluster_role "" }}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: istio-cni-psp
  namespace: {{ .Release.Namespace }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: {{ .Values.cni.psp_cluster_role }}
subjects:
- kind: ServiceAccount
  name: istio-cni
  namespace: {{ .Release.Namespace }}
{{- end }}
---
  {{- if .Values.cni.taint.enabled }}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istio-cni-taint-rolebinding
  namespace: {{ .Release.Namespace}}
  labels:
    k8s-app: istio-cni-taint
subjects:
  - kind: ServiceAccount
    name: istio-cni
    namespace: {{ .Release.Namespace}}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istio-cni-taint-role
  {{- end }}
`)

func chartsIstioCniTemplatesClusterrolebindingYamlBytes() ([]byte, error) {
	return _chartsIstioCniTemplatesClusterrolebindingYaml, nil
}

func chartsIstioCniTemplatesClusterrolebindingYaml() (*asset, error) {
	bytes, err := chartsIstioCniTemplatesClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/templates/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniTemplatesConfigmapCniYaml = []byte(`kind: ConfigMap
apiVersion: v1
metadata:
  name: istio-cni-config
  namespace: {{ .Release.Namespace }}
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
data:
  # The CNI network configuration to add to the plugin chain on each node.  The special
  # values in this config will be automatically populated.
  cni_network_config: |-
        {
          "cniVersion": "0.3.1",
          "name": "istio-cni",
          "type": "istio-cni",
          "log_level": {{ quote .Values.cni.logLevel }},
          "kubernetes": {
              "kubeconfig": "__KUBECONFIG_FILEPATH__",
              "cni_bin_dir": {{ quote .Values.cni.cniBinDir }},
              "exclude_namespaces": [ {{ range $idx, $ns := .Values.cni.excludeNamespaces }}{{ if $idx }}, {{ end }}{{ quote $ns }}{{ end }} ]
          }
        }
---
  {{- if .Values.cni.taint.enabled }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: "istio-cni-taint-configmap"
  namespace: "kube-system"
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
data:
  config: |
        - name: istio-cni
          selector: k8s-app=istio-cni-node
          namespace: kube-system
  {{- end }}
`)

func chartsIstioCniTemplatesConfigmapCniYamlBytes() ([]byte, error) {
	return _chartsIstioCniTemplatesConfigmapCniYaml, nil
}

func chartsIstioCniTemplatesConfigmapCniYaml() (*asset, error) {
	bytes, err := chartsIstioCniTemplatesConfigmapCniYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/templates/configmap-cni.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniTemplatesDaemonsetYaml = []byte(`# This manifest installs the Istio install-cni container, as well
# as the Istio CNI plugin and config on
# each master and worker node in a Kubernetes cluster.
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: istio-cni-node
  namespace: {{ .Release.Namespace }}
  labels:
    k8s-app: istio-cni-node
    release: {{ .Release.Name }}
spec:
  selector:
    matchLabels:
      k8s-app: istio-cni-node
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  template:
    metadata:
      labels:
        k8s-app: istio-cni-node
      annotations:
        # This, along with the CriticalAddonsOnly toleration below,
        # marks the pod as a critical add-on, ensuring it gets
        # priority scheduling and that its resources are reserved
        # if it ever gets evicted.
        scheduler.alpha.kubernetes.io/critical-pod: ''
        sidecar.istio.io/inject: "false"
        {{- if .Values.cni.podAnnotations }}
{{ toYaml .Values.cni.podAnnotations | indent 8 }}
        {{- end }}
    spec:
      nodeSelector:
        kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
        # Make sure istio-cni-node gets scheduled on all nodes.
        - effect: NoSchedule
          operator: Exists
        # Mark the pod as a critical add-on for rescheduling.
        - key: CriticalAddonsOnly
          operator: Exists
        - effect: NoExecute
          operator: Exists
      priorityClassName: system-cluster-critical
      serviceAccountName: istio-cni
      # Minimize downtime during a rolling upgrade or deletion; tell Kubernetes to do a "force
      # deletion": https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods.
      terminationGracePeriodSeconds: 5
      containers:
        # This container installs the Istio CNI binaries
        # and CNI network config file on each node.
        - name: install-cni
{{- if contains "/" .Values.cni.image }}
          image: "{{ .Values.cni.image }}"
{{- else }}
          image: "{{ .Values.cni.hub | default .Values.global.hub }}/{{ .Values.cni.image | default "install-cni" }}:{{ .Values.cni.tag | default .Values.global.tag }}"
{{- end }}
{{- if or .Values.cni.pullPolicy .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.cni.pullPolicy | default .Values.global.imagePullPolicy }}
{{- end }}
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8000
            initialDelaySeconds: 5
          readinessProbe:
            httpGet:
              path: /readyz
              port: 8000
          command: ["install-cni"]
          env:
{{- if .Values.cni.cniConfFileName }}
            # Name of the CNI config file to create.
            - name: CNI_CONF_NAME
              value: "{{ .Values.cni.cniConfFileName }}"
{{- end }}
            # The CNI network config to install on each node.
            - name: CNI_NETWORK_CONFIG
              valueFrom:
                configMapKeyRef:
                  name: istio-cni-config
                  key: cni_network_config
            - name: CNI_NET_DIR
              value: {{ default "/etc/cni/net.d" .Values.cni.cniConfDir }}
            # Deploy as a standalone CNI plugin or as chained?
            - name: CHAINED_CNI_PLUGIN
              value: "{{ .Values.cni.chained }}"
          volumeMounts:
            - mountPath: /host/opt/cni/bin
              name: cni-bin-dir
            - mountPath: /host/etc/cni/net.d
              name: cni-net-dir
{{- if .Values.cni.repair.enabled }}
        - name: repair-cni
{{- if contains "/" .Values.cni.image }}
          image: "{{ .Values.cni.image }}"
{{- else }}
          image: "{{ .Values.cni.hub | default .Values.global.hub }}/{{ .Values.cni.image | default "install-cni" }}:{{ .Values.cni.tag | default .Values.global.tag }}"
{{- end }}
{{- if or .Values.cni.pullPolicy .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.cni.pullPolicy | default .Values.global.imagePullPolicy }}
{{- end }}

          command: ["/opt/cni/bin/istio-cni-repair"]
          env:
          - name: "REPAIR_NODE-NAME"
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          - name: "REPAIR_LABEL-PODS"
            value: "{{.Values.cni.repair.labelPods}}"
          # Set to true to enable pod deletion
          - name: "REPAIR_DELETE-PODS"
            value: "{{.Values.cni.repair.deletePods}}"
          - name: "REPAIR_RUN-AS-DAEMON"
            value: "true"
          - name: "REPAIR_SIDECAR-ANNOTATION"
            value: "sidecar.istio.io/status"
          - name: "REPAIR_INIT-CONTAINER-NAME"
            value: "{{ .Values.cni.repair.initContainerName }}"
          - name: "REPAIR_BROKEN-POD-LABEL-KEY"
            value: "{{.Values.cni.repair.brokenPodLabelKey}}"
          - name: "REPAIR_BROKEN-POD-LABEL-VALUE"
            value: "{{.Values.cni.repair.brokenPodLabelValue}}"
{{- end }}

{{- if .Values.cni.taint.enabled }}
        - name: taint-controller
{{- if contains "/" .Values.cni.image }}
          image: "{{ .Values.cni.image }}"
{{- else }}
          image: "{{ .Values.cni.hub | default .Values.global.hub }}/{{ .Values.cni.image | default "install-cni" }}:{{ .Values.cni.tag | default .Values.global.tag }}"
{{- end }}
{{- if or .Values.cni.pullPolicy .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.cni.pullPolicy | default .Values.global.imagePullPolicy }}
{{- end }}
          command: ["/opt/cni/bin/istio-cni-taint"]
          env:
          - name: "TAINT_RUN-AS-DAEMON"
            value: "true"
          - name: "TAINT_CONFIGMAP-NAME"
            value: "istio-cni-taint-configmap"
          - name: "TAINT_CONFIGMAP-NAMESPACE"
            value: "kube-system"
{{- end }}
      volumes:
        # Used to install CNI.
        - name: cni-bin-dir
          hostPath:
            path: {{ default "/opt/cni/bin" .Values.cni.cniBinDir }}
        - name: cni-net-dir
          hostPath:
            path: {{ default "/etc/cni/net.d" .Values.cni.cniConfDir }}
`)

func chartsIstioCniTemplatesDaemonsetYamlBytes() ([]byte, error) {
	return _chartsIstioCniTemplatesDaemonsetYaml, nil
}

func chartsIstioCniTemplatesDaemonsetYaml() (*asset, error) {
	bytes, err := chartsIstioCniTemplatesDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/templates/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniTemplatesServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
{{- if .Values.global.imagePullSecrets }}
imagePullSecrets:
{{- range .Values.global.imagePullSecrets }}
  - name: {{ . }}
{{- end }}
{{- end }}
metadata:
  name: istio-cni
  namespace: {{ .Release.Namespace }}
  labels:
    app: istio-cni
    release: {{ .Release.Name }}
`)

func chartsIstioCniTemplatesServiceaccountYamlBytes() ([]byte, error) {
	return _chartsIstioCniTemplatesServiceaccountYaml, nil
}

func chartsIstioCniTemplatesServiceaccountYaml() (*asset, error) {
	bytes, err := chartsIstioCniTemplatesServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/templates/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioCniValuesYaml = []byte(`cni:
  hub: ""
  tag: ""
  image: install-cni
  pullPolicy: ""

  logLevel: info

  # Configuration file to insert istio-cni plugin configuration
  # by default this will be the first file found in the cni-conf-dir
  # Example
  # cniConfFileName: 10-calico.conflist

  # CNI bin and conf dir override settings
  # defaults:
  cniBinDir: /opt/cni/bin
  cniConfDir: /etc/cni/net.d
  cniConfFileName: ""

  excludeNamespaces:
    - istio-system

  # Custom annotations on pod level, if you need them
  podAnnotations: {}

  # If this value is set a RoleBinding will be created
  # in the same namespace as the istio-cni DaemonSet is created.
  # This can be used to bind a preexisting ClusterRole to the istio/cni ServiceAccount
  # e.g. if you use PodSecurityPolicies
  psp_cluster_role: ""

  # Deploy the config files as plugin chain (value "true") or as standalone files in the conf dir (value "false")?
  # Some k8s flavors (e.g. OpenShift) do not support the chain approach, set to false if this is the case
  chained: true

  repair:
    enabled: true
    hub: ""
    tag: ""

    labelPods: true
    deletePods: true

    initContainerName: "istio-validation"

    brokenPodLabelKey: "cni.istio.io/uninitialized"
    brokenPodLabelValue: "true"

  # Experimental taint controller for further race condition mitigation
  taint:
    enabled: false

global:
  # Default hub for Istio images.
  # Releases are published to docker hub under 'istio' project.
  # Dev builds from prow are on gcr.io
  hub: gcr.io/istio-testing

  # Default tag for Istio images.
  tag: latest

  # Specify image pull policy if default behavior isn't desired.
  # Default behavior: latest images will be Always else IfNotPresent.
  imagePullPolicy: ""

  # ImagePullSecrets for all ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # For components that don't use ServiceAccounts (i.e. grafana, servicegraph, tracing)
  # ImagePullSecrets will be added to the corresponding Deployment(StatefulSet) objects.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []
  # - private-registry-key
`)

func chartsIstioCniValuesYamlBytes() ([]byte, error) {
	return _chartsIstioCniValuesYaml, nil
}

func chartsIstioCniValuesYaml() (*asset, error) {
	bytes, err := chartsIstioCniValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-cni/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryChartYaml = []byte(`apiVersion: v1
name: istio-discovery
version: 1.2.0
appVersion: 1.2.0
tillerVersion: ">=2.7.2"
description: Helm chart for istio control plane
keywords:
  - istio
  - istiod
  - istio-discovery
sources:
  - http://github.com/istio/istio
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsIstioControlIstioDiscoveryChartYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryChartYaml, nil
}

func chartsIstioControlIstioDiscoveryChartYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryNotesTxt = []byte(`Minimal control plane for Istio. Pilot and mesh config are included.

MCP and injector should optionally be installed in the same namespace. Alternatively remote
address of an MCP server can be set.

`)

func chartsIstioControlIstioDiscoveryNotesTxtBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryNotesTxt, nil
}

func chartsIstioControlIstioDiscoveryNotesTxt() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryNotesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/NOTES.txt", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryFilesGenIstioYaml = []byte(`---
# Source: istio-discovery/templates/poddisruptionbudget.yaml
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: istiod
  namespace: istio-system
  labels:
    app: istiod
    istio.io/rev: default
    release: istio
    istio: pilot
spec:
  minAvailable: 1
  selector:
    matchLabels:
      app: istiod
      istio: pilot
---
# Source: istio-discovery/templates/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio
  namespace: istio-system
  labels:
    istio.io/rev: default
    release: istio
data:

  # Configuration file for the mesh networks to be used by the Split Horizon EDS.
  meshNetworks: |-
    networks: {}

  mesh: |-
    defaultConfig:
      discoveryAddress: istiod.istio-system.svc:15012
      meshId: cluster.local
      proxyMetadata:
        DNS_AGENT: ""
      tracing:
        zipkin:
          address: zipkin.istio-system:9411
    enablePrometheusMerge: true
    rootNamespace: istio-system
    trustDomain: cluster.local
---
# Source: istio-discovery/templates/istiod-injector-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio-sidecar-injector
  namespace: istio-system
  labels:
    istio.io/rev: default
    release: istio
data:

  values: |-
    {
      "global": {
        "arch": {
          "amd64": 2,
          "ppc64le": 2,
          "s390x": 2
        },
        "caAddress": "",
        "centralIstiod": false,
        "configValidation": true,
        "defaultConfigVisibilitySettings": [],
        "defaultNodeSelector": {},
        "defaultPodDisruptionBudget": {
          "enabled": true
        },
        "defaultResources": {
          "requests": {
            "cpu": "10m"
          }
        },
        "defaultTolerations": [],
        "hub": "gcr.io/istio-testing",
        "imagePullPolicy": "",
        "imagePullSecrets": [],
        "istioNamespace": "istio-system",
        "istiod": {
          "enableAnalysis": false
        },
        "jwtPolicy": "third-party-jwt",
        "logAsJson": false,
        "logging": {
          "level": "default:info"
        },
        "meshExpansion": {
          "enabled": false,
          "useILB": false
        },
        "meshID": "",
        "meshNetworks": {},
        "mountMtlsCerts": false,
        "multiCluster": {
          "clusterName": "",
          "enabled": false
        },
        "network": "",
        "omitSidecarInjectorConfigMap": false,
        "oneNamespace": false,
        "operatorManageWebhooks": false,
        "pilotCertProvider": "istiod",
        "priorityClassName": "",
        "proxy": {
          "autoInject": "enabled",
          "clusterDomain": "cluster.local",
          "componentLogLevel": "misc:error",
          "enableCoreDump": false,
          "excludeIPRanges": "",
          "excludeInboundPorts": "",
          "excludeOutboundPorts": "",
          "holdApplicationUntilProxyStarts": false,
          "image": "proxyv2",
          "includeIPRanges": "*",
          "logLevel": "warning",
          "privileged": false,
          "readinessFailureThreshold": 30,
          "readinessInitialDelaySeconds": 1,
          "readinessPeriodSeconds": 2,
          "resources": {
            "limits": {
              "cpu": "2000m",
              "memory": "1024Mi"
            },
            "requests": {
              "cpu": "100m",
              "memory": "128Mi"
            }
          },
          "statusPort": 15020,
          "tracer": "zipkin"
        },
        "proxy_init": {
          "image": "proxyv2",
          "resources": {
            "limits": {
              "cpu": "2000m",
              "memory": "1024Mi"
            },
            "requests": {
              "cpu": "10m",
              "memory": "10Mi"
            }
          }
        },
        "remotePilotAddress": "",
        "sds": {
          "token": {
            "aud": "istio-ca"
          }
        },
        "sts": {
          "servicePort": 0
        },
        "tag": "latest",
        "tracer": {
          "datadog": {
            "address": "$(HOST_IP):8126"
          },
          "lightstep": {
            "accessToken": "",
            "address": ""
          },
          "stackdriver": {
            "debug": false,
            "maxNumberOfAnnotations": 200,
            "maxNumberOfAttributes": 200,
            "maxNumberOfMessageEvents": 200
          },
          "zipkin": {
            "address": ""
          }
        },
        "trustDomain": "cluster.local",
        "useMCP": false
      },
      "revision": "",
      "sidecarInjectorWebhook": {
        "alwaysInjectSelector": [],
        "enableNamespacesByDefault": false,
        "injectedAnnotations": {},
        "neverInjectSelector": [],
        "objectSelector": {
          "autoInject": true,
          "enabled": false
        },
        "rewriteAppHTTPProbe": true
      }
    }

  # To disable injection: use omitSidecarInjectorConfigMap, which disables the webhook patching
  # and istiod webhook functionality.
  #
  # New fields should not use Values - it is a 'primary' config object, users should be able
  # to fine tune it or use it with kube-inject.
  config: |-
    policy: enabled
    alwaysInjectSelector:
      []
    neverInjectSelector:
      []
    injectedAnnotations:

    template: |
      rewriteAppHTTPProbe: {{ valueOrDefault .Values.sidecarInjectorWebhook.rewriteAppHTTPProbe false }}
      initContainers:
      {{ if ne (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`NONE`+"`"+` }}
      {{ if .Values.istio_cni.enabled -}}
      - name: istio-validation
      {{ else -}}
      - name: istio-init
      {{ end -}}
      {{- if contains "/" .Values.global.proxy_init.image }}
        image: "{{ .Values.global.proxy_init.image }}"
      {{- else }}
        image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
      {{- end }}
        args:
        - istio-iptables
        - "-p"
        - 15001
        - "-z"
        - "15006"
        - "-u"
        - 1337
        - "-m"
        - "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
        - "-i"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
        - "-x"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
        - "-b"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` `+"`"+`*`+"`"+` }}"
        - "-d"
      {{- if excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}
        - "15090,15021,{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
      {{- else }}
        - "15090,15021"
      {{- end }}
        {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") -}}
        - "-q"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
        {{ end -}}
        {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.excludeOutboundPorts "") "") -}}
        - "-o"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
        {{ end -}}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+`) -}}
        - "-k"
        - "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
        {{ end -}}
        {{ if .Values.istio_cni.enabled -}}
        - "--run-validation"
        - "--skip-rule-apply"
        {{ end -}}
        imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
      {{- if .ProxyConfig.ProxyMetadata }}
        env:
        {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
        - name: {{ $key }}
          value: "{{ $value }}"
        {{- end }}
      {{- end }}
        resources:
      {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
          requests:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
            {{ end }}
        {{- end }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
          limits:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
            {{ end }}
        {{- end }}
      {{- else }}
        {{- if .Values.global.proxy.resources }}
          {{ toYaml .Values.global.proxy.resources | indent 4 }}
        {{- end }}
      {{- end }}
        securityContext:
          allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
          privileged: {{ .Values.global.proxy.privileged }}
          capabilities:
        {{- if not .Values.istio_cni.enabled }}
            add:
            - NET_ADMIN
            - NET_RAW
        {{- end }}
            drop:
            - ALL
        {{- if not .Values.istio_cni.enabled }}
          readOnlyRootFilesystem: false
          runAsGroup: 0
          runAsNonRoot: false
          runAsUser: 0
        {{- else }}
          readOnlyRootFilesystem: true
          runAsGroup: 1337
          runAsUser: 1337
          runAsNonRoot: true
        {{- end }}
        restartPolicy: Always
      {{ end -}}
      {{- if eq .Values.global.proxy.enableCoreDump true }}
      - name: enable-core-dump
        args:
        - -c
        - sysctl -w kernel.core_pattern=/var/lib/istio/data/core.proxy && ulimit -c unlimited
        command:
          - /bin/sh
      {{- if contains "/" .Values.global.proxy_init.image }}
        image: "{{ .Values.global.proxy_init.image }}"
      {{- else }}
        image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
      {{- end }}
        imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
        resources: {}
        securityContext:
          allowPrivilegeEscalation: true
          capabilities:
            add:
            - SYS_ADMIN
            drop:
            - ALL
          privileged: true
          readOnlyRootFilesystem: false
          runAsGroup: 0
          runAsNonRoot: false
          runAsUser: 0
      {{ end }}
      containers:
      - name: istio-proxy
      {{- if contains "/" (annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image) }}
        image: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image }}"
      {{- else }}
        image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image }}:{{ .Values.global.tag }}"
      {{- end }}
        ports:
        - containerPort: 15090
          protocol: TCP
          name: http-envoy-prom
        args:
        - proxy
        - sidecar
        - --domain
        - $(POD_NAMESPACE).svc.{{ .Values.global.proxy.clusterDomain }}
        - --serviceCluster
        {{ if ne "" (index .ObjectMeta.Labels "app") -}}
        - "{{ index .ObjectMeta.Labels `+"`"+`app`+"`"+` }}.$(POD_NAMESPACE)"
        {{ else -}}
        - "{{ valueOrDefault .DeploymentMeta.Name `+"`"+`istio-proxy`+"`"+` }}.{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}"
        {{ end -}}
        - --proxyLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/logLevel`+"`"+` .Values.global.proxy.logLevel}}
        - --proxyComponentLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/componentLogLevel`+"`"+` .Values.global.proxy.componentLogLevel}}
      {{- if .Values.global.sts.servicePort }}
        - --stsPort={{ .Values.global.sts.servicePort }}
      {{- end }}
      {{- if .Values.global.trustDomain }}
        - --trust-domain={{ .Values.global.trustDomain }}
      {{- end }}
      {{- if .Values.global.logAsJson }}
        - --log_as_json
      {{- end }}
      {{- if gt .ProxyConfig.Concurrency.GetValue 0 }}
        - --concurrency
        - "{{ .ProxyConfig.Concurrency.GetValue }}"
      {{- end -}}
      {{- if .Values.global.proxy.lifecycle }}
        lifecycle:
          {{ toYaml .Values.global.proxy.lifecycle | indent 4 }}
      {{- else if .Values.global.proxy.holdApplicationUntilProxyStarts}}
        lifecycle:
          postStart:
            exec:
              command:
              - pilot-agent
              - wait
      {{- end }}
        env:
        - name: JWT_POLICY
          value: {{ .Values.global.jwtPolicy }}
        - name: PILOT_CERT_PROVIDER
          value: {{ .Values.global.pilotCertProvider }}
        - name: CA_ADDR
        {{- if .Values.global.caAddress }}
          value: {{ .Values.global.caAddress }}
        {{- else }}
          value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Values.global.istioNamespace }}.svc:15012
        {{- end }}
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: INSTANCE_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: SERVICE_ACCOUNT
          valueFrom:
            fieldRef:
              fieldPath: spec.serviceAccountName
        - name: HOST_IP
          valueFrom:
            fieldRef:
              fieldPath: status.hostIP
        - name: CANONICAL_SERVICE
          valueFrom:
            fieldRef:
              fieldPath: metadata.labels['service.istio.io/canonical-name']
        - name: CANONICAL_REVISION
          valueFrom:
            fieldRef:
              fieldPath: metadata.labels['service.istio.io/canonical-revision']
        - name: PROXY_CONFIG
          value: |
                 {{ protoToJSON .ProxyConfig }}
        - name: ISTIO_META_POD_PORTS
          value: |-
            [
            {{- $first := true }}
            {{- range $index1, $c := .Spec.Containers }}
              {{- range $index2, $p := $c.Ports }}
                {{- if (structToJSON $p) }}
                {{if not $first}},{{end}}{{ structToJSON $p }}
                {{- $first = false }}
                {{- end }}
              {{- end}}
            {{- end}}
            ]
        - name: ISTIO_META_APP_CONTAINERS
          value: "{{- range $index, $container := .Spec.Containers }}{{- if ne $index 0}},{{- end}}{{ $container.Name }}{{- end}}"
        - name: ISTIO_META_CLUSTER_ID
          value: "{{ valueOrDefault .Values.global.multiCluster.clusterName `+"`"+`Kubernetes`+"`"+` }}"
        - name: ISTIO_META_INTERCEPTION_MODE
          value: "{{ or (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/interceptionMode`+"`"+`) .ProxyConfig.InterceptionMode.String }}"
        {{- if .Values.global.network }}
        - name: ISTIO_META_NETWORK
          value: "{{ .Values.global.network }}"
        {{- end }}
        {{ if .ObjectMeta.Annotations }}
        - name: ISTIO_METAJSON_ANNOTATIONS
          value: |
                 {{ toJSON .ObjectMeta.Annotations }}
        {{ end }}
        {{- if .DeploymentMeta.Name }}
        - name: ISTIO_META_WORKLOAD_NAME
          value: "{{ .DeploymentMeta.Name }}"
        {{ end }}
        {{- if and .TypeMeta.APIVersion .DeploymentMeta.Name }}
        - name: ISTIO_META_OWNER
          value: kubernetes://apis/{{ .TypeMeta.APIVersion }}/namespaces/{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}/{{ toLower .TypeMeta.Kind}}s/{{ .DeploymentMeta.Name }}
        {{- end}}
        {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
        - name: ISTIO_BOOTSTRAP_OVERRIDE
          value: "/etc/istio/custom-bootstrap/custom_bootstrap.json"
        {{- end }}
        {{- if .Values.global.meshID }}
        - name: ISTIO_META_MESH_ID
          value: "{{ .Values.global.meshID }}"
        {{- else if .Values.global.trustDomain }}
        - name: ISTIO_META_MESH_ID
          value: "{{ .Values.global.trustDomain }}"
        {{- end }}
        {{- if and (eq .Values.global.proxy.tracer "datadog") (isset .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
        {{- range $key, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
        - name: {{ $key }}
          value: "{{ $value }}"
        {{- end }}
        {{- end }}
        {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
        - name: {{ $key }}
          value: "{{ $value }}"
        {{- end }}
        imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
        {{ if ne (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) `+"`"+`0`+"`"+` }}
        readinessProbe:
          httpGet:
            path: /healthz/ready
            port: 15021
          initialDelaySeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/initialDelaySeconds`+"`"+` .Values.global.proxy.readinessInitialDelaySeconds }}
          periodSeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/periodSeconds`+"`"+` .Values.global.proxy.readinessPeriodSeconds }}
          failureThreshold: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/failureThreshold`+"`"+` .Values.global.proxy.readinessFailureThreshold }}
        {{ end -}}
        securityContext:
          allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
          capabilities:
            {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
            add:
            {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+` -}}
            - NET_ADMIN
            {{- end }}
            {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+` -}}
            - NET_BIND_SERVICE
            {{- end }}
            {{- end }}
            drop:
            - ALL
          privileged: {{ .Values.global.proxy.privileged }}
          readOnlyRootFilesystem: {{ not .Values.global.proxy.enableCoreDump }}
          runAsGroup: 1337
          fsGroup: 1337
          {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
          runAsNonRoot: false
          runAsUser: 0
          {{- else -}}
          runAsNonRoot: true
          runAsUser: 1337
          {{- end }}
        resources:
      {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
          requests:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
            {{ end }}
        {{- end }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
          limits:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
            {{ end }}
        {{- end }}
      {{- else }}
        {{- if .Values.global.proxy.resources }}
          {{ toYaml .Values.global.proxy.resources | indent 4 }}
        {{- end }}
      {{- end }}
        volumeMounts:
        {{- if eq .Values.global.pilotCertProvider "istiod" }}
        - mountPath: /var/run/secrets/istio
          name: istiod-ca-cert
        {{- end }}
        - mountPath: /var/lib/istio/data
          name: istio-data
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
        - mountPath: /etc/istio/custom-bootstrap
          name: custom-bootstrap-volume
        {{- end }}
        # SDS channel between istioagent and Envoy
        - mountPath: /etc/istio/proxy
          name: istio-envoy
        {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
        - mountPath: /var/run/secrets/tokens
          name: istio-token
        {{- end }}
        {{- if .Values.global.mountMtlsCerts }}
        # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
        - mountPath: /etc/certs/
          name: istio-certs
          readOnly: true
        {{- end }}
        - name: istio-podinfo
          mountPath: /etc/istio/pod
         {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
        - mountPath: {{ directory .ProxyConfig.GetTracing.GetTlsSettings.GetCaCertificates }}
          name: lightstep-certs
          readOnly: true
        {{- end }}
          {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+` }}
          {{ range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+`) }}
        - name: "{{  $index }}"
          {{ toYaml $value | indent 4 }}
          {{ end }}
          {{- end }}
      volumes:
      {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
      - name: custom-bootstrap-volume
        configMap:
          name: {{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+` "" }}
      {{- end }}
      # SDS channel between istioagent and Envoy
      - emptyDir:
          medium: Memory
        name: istio-envoy
      - name: istio-data
        emptyDir: {}
      - name: istio-podinfo
        downwardAPI:
          items:
            - path: "labels"
              fieldRef:
                fieldPath: metadata.labels
            - path: "annotations"
              fieldRef:
                fieldPath: metadata.annotations
      {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
      - name: istio-token
        projected:
          sources:
          - serviceAccountToken:
              path: istio-token
              expirationSeconds: 43200
              audience: {{ .Values.global.sds.token.aud }}
      {{- end }}
      {{- if eq .Values.global.pilotCertProvider "istiod" }}
      - name: istiod-ca-cert
        configMap:
          name: istio-ca-root-cert
      {{- end }}
      {{- if .Values.global.mountMtlsCerts }}
      # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
      - name: istio-certs
        secret:
          optional: true
          {{ if eq .Spec.ServiceAccountName "" }}
          secretName: istio.default
          {{ else -}}
          secretName: {{  printf "istio.%s" .Spec.ServiceAccountName }}
          {{  end -}}
      {{- end }}
        {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+` }}
        {{range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+`) }}
      - name: "{{ $index }}"
        {{ toYaml $value | indent 2 }}
        {{ end }}
        {{ end }}
      {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
      - name: lightstep-certs
        secret:
          optional: true
          secretName: lightstep.cacert
      {{- end }}
      podRedirectAnnot:
      {{- if and (.Values.istio_cni.enabled) (not .Values.istio_cni.chained) }}
        k8s.v1.cni.cncf.io/networks: '{{ appendMultusNetwork (index .ObjectMeta.Annotations `+"`"+`k8s.v1.cni.cncf.io/networks`+"`"+`) `+"`"+`istio-cni`+"`"+` }}'
      {{- end }}
        sidecar.istio.io/interceptionMode: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
        traffic.sidecar.istio.io/includeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
        traffic.sidecar.istio.io/excludeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
        traffic.sidecar.istio.io/includeInboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` (includeInboundPorts .Spec.Containers) }}"
        traffic.sidecar.istio.io/excludeInboundPorts: "{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
      {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") }}
        traffic.sidecar.istio.io/includeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
      {{- end }}
      {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne .Values.global.proxy.excludeOutboundPorts "") }}
        traffic.sidecar.istio.io/excludeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
      {{- end }}
        traffic.sidecar.istio.io/kubevirtInterfaces: "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
      {{- if .Values.global.imagePullSecrets }}
      imagePullSecrets:
        {{- range .Values.global.imagePullSecrets }}
        - name: {{ . }}
        {{- end }}
      {{- end }}
---
# Source: istio-discovery/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: istiod
  namespace: istio-system
  labels:
    istio.io/rev: default
    app: istiod
    istio: pilot
    release: istio
spec:
  ports:
    - port: 15010
      name: grpc-xds # plaintext
    - port: 15012
      name: https-dns # mTLS with k8s-signed cert
    - port: 443
      name: https-webhook # validation and injection
      targetPort: 15017
    - port: 15014
      name: http-monitoring # prometheus stats
    - name: dns-tls
      port: 853
      targetPort: 15053
      protocol: TCP
  selector:
    app: istiod
    # Label used by the 'default' service. For versioned deployments we match with app and version.
    # This avoids default deployment picking the canary
    istio: pilot
---
# Source: istio-discovery/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: istiod
  namespace: istio-system
  labels:
    app: istiod
    istio.io/rev: default
    istio: pilot
    release: istio
spec:
  strategy:
    rollingUpdate:
      maxSurge: 100%
      maxUnavailable: 25%
  selector:
    matchLabels:
      istio: pilot
  template:
    metadata:
      labels:
        app: istiod
        istio.io/rev: default
        istio: pilot
      annotations:
        prometheus.io/port: "15014"
        prometheus.io/scrape: "true"
        sidecar.istio.io/inject: "false"
    spec:
      serviceAccountName: istiod-service-account
      securityContext:
        fsGroup: 1337
      containers:
        - name: discovery
          image: "gcr.io/istio-testing/pilot:latest"
          args:
          - "discovery"
          - --monitoringAddr=:15014
          - --log_output_level=default:info
          - --domain
          - cluster.local
          - --trust-domain=cluster.local
          - --keepaliveMaxServerConnectionAge
          - "30m"
          ports:
          - containerPort: 8080
          - containerPort: 15010
          - containerPort: 15017
          - containerPort: 15053
          readinessProbe:
            httpGet:
              path: /ready
              port: 8080
            initialDelaySeconds: 1
            periodSeconds: 3
            timeoutSeconds: 5
          env:
          - name: REVISION
            value: "default"
          - name: JWT_POLICY
            value: third-party-jwt
          - name: PILOT_CERT_PROVIDER
            value: istiod
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: SERVICE_ACCOUNT
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.serviceAccountName
          - name: KUBECONFIG
            value: /var/run/secrets/remote/config
          - name: PILOT_TRACE_SAMPLING
            value: "1"
          - name: PILOT_ENABLE_PROTOCOL_SNIFFING_FOR_OUTBOUND
            value: "true"
          - name: PILOT_ENABLE_PROTOCOL_SNIFFING_FOR_INBOUND
            value: "true"
          - name: INJECTION_WEBHOOK_CONFIG_NAME
            value: istio-sidecar-injector
          - name: ISTIOD_ADDR
            value: istiod.istio-system.svc:15012
          - name: PILOT_ENABLE_ANALYSIS
            value: "false"
          - name: CLUSTER_ID
            value: "Kubernetes"
          - name: CENTRAL_ISTIOD
            value: "false"
          resources:
            requests:
              cpu: 500m
              memory: 2048Mi
          securityContext:
            runAsUser: 1337
            runAsGroup: 1337
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
          volumeMounts:
          - name: config-volume
            mountPath: /etc/istio/config
          - name: istio-token
            mountPath: /var/run/secrets/tokens
            readOnly: true
          - name: local-certs
            mountPath: /var/run/secrets/istio-dns
          - name: cacerts
            mountPath: /etc/cacerts
            readOnly: true
          - name: istio-kubeconfig
            mountPath: /var/run/secrets/remote
            readOnly: true
          - name: inject
            mountPath: /var/lib/istio/inject
            readOnly: true
      volumes:
      # Technically not needed on this pod - but it helps debugging/testing SDS
      # Should be removed after everything works.
      - emptyDir:
          medium: Memory
        name: local-certs
      - name: istio-token
        projected:
          sources:
            - serviceAccountToken:
                audience: istio-ca
                expirationSeconds: 43200
                path: istio-token
      # Optional: user-generated root
      - name: cacerts
        secret:
          secretName: cacerts
          optional: true
      - name: istio-kubeconfig
        secret:
          secretName: istio-kubeconfig
          optional: true
      # Optional - image should have
      - name: inject
        configMap:
          name: istio-sidecar-injector
      - name: config-volume
        configMap:
          name: istio
---
# Source: istio-discovery/templates/autoscale.yaml
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: istiod
  namespace: istio-system
  labels:
    app: istiod
    release: istio
    istio.io/rev: default
spec:
  maxReplicas: 5
  minReplicas: 1
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: istiod
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80
---
# Source: istio-discovery/templates/telemetryv2_1.6.yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.6
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: ANY # inbound, outbound, and gateway
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration: |
                  {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
# Source: istio-discovery/templates/telemetryv2_1.6.yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.6
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Source: istio-discovery/templates/telemetryv2_1.6.yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.6
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {
                    "debug": "false",
                    "stat_prefix": "istio"
                  }
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration: |
                  {
                    "debug": "false",
                    "stat_prefix": "istio"
                  }
                vm_config:
                  vm_id: stats_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {
                    "debug": "false",
                    "stat_prefix": "istio",
                    "disable_host_header_fallback": true
                  }
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Source: istio-discovery/templates/telemetryv2_1.6.yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.6
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration: |
                  {
                    "debug": "false",
                    "stat_prefix": "istio"
                  }
                vm_config:
                  vm_id: tcp_stats_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {
                    "debug": "false",
                    "stat_prefix": "istio"
                  }
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {
                    "debug": "false",
                    "stat_prefix": "istio"
                  }
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---
# Source: istio-discovery/templates/telemetryv2_1.7.yaml
# Note: metadata exchange filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.7
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
# Source: istio-discovery/templates/telemetryv2_1.7.yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.7
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Source: istio-discovery/templates/telemetryv2_1.7.yaml
# Note: http stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.7
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: stats_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio",
                      "disable_host_header_fallback": true
                    }
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Source: istio-discovery/templates/telemetryv2_1.7.yaml
# Note: tcp stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.7
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: tcp_stats_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---
# Source: istio-discovery/templates/telemetryv2_1.8.yaml
# Note: metadata exchange filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.8
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
# Source: istio-discovery/templates/telemetryv2_1.8.yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.8
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Source: istio-discovery/templates/telemetryv2_1.8.yaml
# Note: http stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.8
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: stats_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio",
                      "disable_host_header_fallback": true
                    }
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Source: istio-discovery/templates/telemetryv2_1.8.yaml
# Note: tcp stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.8
  namespace: istio-system
  labels:
    istio.io/rev: default
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: tcp_stats_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "debug": "false",
                      "stat_prefix": "istio"
                    }
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---
# Source: istio-discovery/templates/mutatingwebhook.yaml
# Installed for each revision - not installed for cluster resources ( cluster roles, bindings, crds)
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: istio-sidecar-injector

  labels:
    istio.io/rev: default
    app: sidecar-injector
    release: istio
webhooks:
  - name: sidecar-injector.istio.io
    clientConfig:
      service:
        name: istiod
        namespace: istio-system
        path: "/inject"
      caBundle: ""
    sideEffects: None
    rules:
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    failurePolicy: Fail
    admissionReviewVersions: ["v1beta1", "v1"]
    namespaceSelector:
      matchLabels:
        istio-injection: enabled
`)

func chartsIstioControlIstioDiscoveryFilesGenIstioYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryFilesGenIstioYaml, nil
}

func chartsIstioControlIstioDiscoveryFilesGenIstioYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryFilesGenIstioYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/files/gen-istio.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryFilesInjectionTemplateYaml = []byte(`template: |
  rewriteAppHTTPProbe: {{ valueOrDefault .Values.sidecarInjectorWebhook.rewriteAppHTTPProbe false }}
  initContainers:
  {{ if ne (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`NONE`+"`"+` }}
  {{ if .Values.istio_cni.enabled -}}
  - name: istio-validation
  {{ else -}}
  - name: istio-init
  {{ end -}}
  {{- if contains "/" .Values.global.proxy_init.image }}
    image: "{{ .Values.global.proxy_init.image }}"
  {{- else }}
    image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
  {{- end }}
    args:
    - istio-iptables
    - "-p"
    - 15001
    - "-z"
    - "15006"
    - "-u"
    - 1337
    - "-m"
    - "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
    - "-i"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
    - "-x"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
    - "-b"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` `+"`"+`*`+"`"+` }}"
    - "-d"
  {{- if excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}
    - "15090,15021,{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
  {{- else }}
    - "15090,15021"
  {{- end }}
    {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") -}}
    - "-q"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
    {{ end -}}
    {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.excludeOutboundPorts "") "") -}}
    - "-o"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
    {{ end -}}
    {{ if (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+`) -}}
    - "-k"
    - "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
    {{ end -}}
    {{ if .Values.istio_cni.enabled -}}
    - "--run-validation"
    - "--skip-rule-apply"
    {{ end -}}
    imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
  {{- if .ProxyConfig.ProxyMetadata }}
    env:
    {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
    - name: {{ $key }}
      value: "{{ $value }}"
    {{- end }}
  {{- end }}
    resources:
  {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
      requests:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
        {{ end }}
    {{- end }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
      limits:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
        {{ end }}
    {{- end }}
  {{- else }}
    {{- if .Values.global.proxy.resources }}
      {{ toYaml .Values.global.proxy.resources | indent 4 }}
    {{- end }}
  {{- end }}
    securityContext:
      allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
      privileged: {{ .Values.global.proxy.privileged }}
      capabilities:
    {{- if not .Values.istio_cni.enabled }}
        add:
        - NET_ADMIN
        - NET_RAW
    {{- end }}
        drop:
        - ALL
    {{- if not .Values.istio_cni.enabled }}
      readOnlyRootFilesystem: false
      runAsGroup: 0
      runAsNonRoot: false
      runAsUser: 0
    {{- else }}
      readOnlyRootFilesystem: true
      runAsGroup: 1337
      runAsUser: 1337
      runAsNonRoot: true
    {{- end }}
    restartPolicy: Always
  {{ end -}}
  {{- if eq .Values.global.proxy.enableCoreDump true }}
  - name: enable-core-dump
    args:
    - -c
    - sysctl -w kernel.core_pattern=/var/lib/istio/data/core.proxy && ulimit -c unlimited
    command:
      - /bin/sh
  {{- if contains "/" .Values.global.proxy_init.image }}
    image: "{{ .Values.global.proxy_init.image }}"
  {{- else }}
    image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
  {{- end }}
    imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
    resources: {}
    securityContext:
      allowPrivilegeEscalation: true
      capabilities:
        add:
        - SYS_ADMIN
        drop:
        - ALL
      privileged: true
      readOnlyRootFilesystem: false
      runAsGroup: 0
      runAsNonRoot: false
      runAsUser: 0
  {{ end }}
  containers:
  - name: istio-proxy
  {{- if contains "/" (annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image) }}
    image: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image }}"
  {{- else }}
    image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image }}:{{ .Values.global.tag }}"
  {{- end }}
    ports:
    - containerPort: 15090
      protocol: TCP
      name: http-envoy-prom
    args:
    - proxy
    - sidecar
    - --domain
    - $(POD_NAMESPACE).svc.{{ .Values.global.proxy.clusterDomain }}
    - --serviceCluster
    {{ if ne "" (index .ObjectMeta.Labels "app") -}}
    - "{{ index .ObjectMeta.Labels `+"`"+`app`+"`"+` }}.$(POD_NAMESPACE)"
    {{ else -}}
    - "{{ valueOrDefault .DeploymentMeta.Name `+"`"+`istio-proxy`+"`"+` }}.{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}"
    {{ end -}}
    - --proxyLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/logLevel`+"`"+` .Values.global.proxy.logLevel}}
    - --proxyComponentLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/componentLogLevel`+"`"+` .Values.global.proxy.componentLogLevel}}
  {{- if .Values.global.sts.servicePort }}
    - --stsPort={{ .Values.global.sts.servicePort }}
  {{- end }}
  {{- if .Values.global.trustDomain }}
    - --trust-domain={{ .Values.global.trustDomain }}
  {{- end }}
  {{- if .Values.global.logAsJson }}
    - --log_as_json
  {{- end }}
  {{- if gt .ProxyConfig.Concurrency.GetValue 0 }}
    - --concurrency
    - "{{ .ProxyConfig.Concurrency.GetValue }}"
  {{- end -}}
  {{- if .Values.global.proxy.lifecycle }}
    lifecycle:
      {{ toYaml .Values.global.proxy.lifecycle | indent 4 }}
  {{- else if .Values.global.proxy.holdApplicationUntilProxyStarts}}
    lifecycle:
      postStart:
        exec:
          command:
          - pilot-agent
          - wait
  {{- end }}
    env:
    - name: JWT_POLICY
      value: {{ .Values.global.jwtPolicy }}
    - name: PILOT_CERT_PROVIDER
      value: {{ .Values.global.pilotCertProvider }}
    - name: CA_ADDR
    {{- if .Values.global.caAddress }}
      value: {{ .Values.global.caAddress }}
    {{- else }}
      value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Values.global.istioNamespace }}.svc:15012
    {{- end }}
    - name: POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    - name: POD_NAMESPACE
      valueFrom:
        fieldRef:
          fieldPath: metadata.namespace
    - name: INSTANCE_IP
      valueFrom:
        fieldRef:
          fieldPath: status.podIP
    - name: SERVICE_ACCOUNT
      valueFrom:
        fieldRef:
          fieldPath: spec.serviceAccountName
    - name: HOST_IP
      valueFrom:
        fieldRef:
          fieldPath: status.hostIP
    - name: CANONICAL_SERVICE
      valueFrom:
        fieldRef:
          fieldPath: metadata.labels['service.istio.io/canonical-name']
    - name: CANONICAL_REVISION
      valueFrom:
        fieldRef:
          fieldPath: metadata.labels['service.istio.io/canonical-revision']
    - name: PROXY_CONFIG
      value: |
             {{ protoToJSON .ProxyConfig }}
    - name: ISTIO_META_POD_PORTS
      value: |-
        [
        {{- $first := true }}
        {{- range $index1, $c := .Spec.Containers }}
          {{- range $index2, $p := $c.Ports }}
            {{- if (structToJSON $p) }}
            {{if not $first}},{{end}}{{ structToJSON $p }}
            {{- $first = false }}
            {{- end }}
          {{- end}}
        {{- end}}
        ]
    - name: ISTIO_META_APP_CONTAINERS
      value: "{{- range $index, $container := .Spec.Containers }}{{- if ne $index 0}},{{- end}}{{ $container.Name }}{{- end}}"
    - name: ISTIO_META_CLUSTER_ID
      value: "{{ valueOrDefault .Values.global.multiCluster.clusterName `+"`"+`Kubernetes`+"`"+` }}"
    - name: ISTIO_META_INTERCEPTION_MODE
      value: "{{ or (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/interceptionMode`+"`"+`) .ProxyConfig.InterceptionMode.String }}"
    {{- if .Values.global.network }}
    - name: ISTIO_META_NETWORK
      value: "{{ .Values.global.network }}"
    {{- end }}
    {{ if .ObjectMeta.Annotations }}
    - name: ISTIO_METAJSON_ANNOTATIONS
      value: |
             {{ toJSON .ObjectMeta.Annotations }}
    {{ end }}
    {{- if .DeploymentMeta.Name }}
    - name: ISTIO_META_WORKLOAD_NAME
      value: "{{ .DeploymentMeta.Name }}"
    {{ end }}
    {{- if and .TypeMeta.APIVersion .DeploymentMeta.Name }}
    - name: ISTIO_META_OWNER
      value: kubernetes://apis/{{ .TypeMeta.APIVersion }}/namespaces/{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}/{{ toLower .TypeMeta.Kind}}s/{{ .DeploymentMeta.Name }}
    {{- end}}
    {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
    - name: ISTIO_BOOTSTRAP_OVERRIDE
      value: "/etc/istio/custom-bootstrap/custom_bootstrap.json"
    {{- end }}
    {{- if .Values.global.meshID }}
    - name: ISTIO_META_MESH_ID
      value: "{{ .Values.global.meshID }}"
    {{- else if .Values.global.trustDomain }}
    - name: ISTIO_META_MESH_ID
      value: "{{ .Values.global.trustDomain }}"
    {{- end }}
    {{- if and (eq .Values.global.proxy.tracer "datadog") (isset .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
    {{- range $key, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
    - name: {{ $key }}
      value: "{{ $value }}"
    {{- end }}
    {{- end }}
    {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
    - name: {{ $key }}
      value: "{{ $value }}"
    {{- end }}
    imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
    {{ if ne (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) `+"`"+`0`+"`"+` }}
    readinessProbe:
      httpGet:
        path: /healthz/ready
        port: 15021
      initialDelaySeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/initialDelaySeconds`+"`"+` .Values.global.proxy.readinessInitialDelaySeconds }}
      periodSeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/periodSeconds`+"`"+` .Values.global.proxy.readinessPeriodSeconds }}
      failureThreshold: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/failureThreshold`+"`"+` .Values.global.proxy.readinessFailureThreshold }}
    {{ end -}}
    securityContext:
      allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
      capabilities:
        {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
        add:
        {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+` -}}
        - NET_ADMIN
        {{- end }}
        {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+` -}}
        - NET_BIND_SERVICE
        {{- end }}
        {{- end }}
        drop:
        - ALL
      privileged: {{ .Values.global.proxy.privileged }}
      readOnlyRootFilesystem: {{ not .Values.global.proxy.enableCoreDump }}
      runAsGroup: 1337
      fsGroup: 1337
      {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
      runAsNonRoot: false
      runAsUser: 0
      {{- else -}}
      runAsNonRoot: true
      runAsUser: 1337
      {{- end }}
    resources:
  {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
      requests:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
        {{ end }}
    {{- end }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
      limits:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
        {{ end }}
    {{- end }}
  {{- else }}
    {{- if .Values.global.proxy.resources }}
      {{ toYaml .Values.global.proxy.resources | indent 4 }}
    {{- end }}
  {{- end }}
    volumeMounts:
    {{- if eq .Values.global.pilotCertProvider "istiod" }}
    - mountPath: /var/run/secrets/istio
      name: istiod-ca-cert
    {{- end }}
    - mountPath: /var/lib/istio/data
      name: istio-data
    {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
    - mountPath: /etc/istio/custom-bootstrap
      name: custom-bootstrap-volume
    {{- end }}
    # SDS channel between istioagent and Envoy
    - mountPath: /etc/istio/proxy
      name: istio-envoy
    {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
    - mountPath: /var/run/secrets/tokens
      name: istio-token
    {{- end }}
    {{- if .Values.global.mountMtlsCerts }}
    # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
    - mountPath: /etc/certs/
      name: istio-certs
      readOnly: true
    {{- end }}
    - name: istio-podinfo
      mountPath: /etc/istio/pod
     {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
    - mountPath: {{ directory .ProxyConfig.GetTracing.GetTlsSettings.GetCaCertificates }}
      name: lightstep-certs
      readOnly: true
    {{- end }}
      {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+` }}
      {{ range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+`) }}
    - name: "{{  $index }}"
      {{ toYaml $value | indent 4 }}
      {{ end }}
      {{- end }}
  volumes:
  {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
  - name: custom-bootstrap-volume
    configMap:
      name: {{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+` "" }}
  {{- end }}
  # SDS channel between istioagent and Envoy
  - emptyDir:
      medium: Memory
    name: istio-envoy
  - name: istio-data
    emptyDir: {}
  - name: istio-podinfo
    downwardAPI:
      items:
        - path: "labels"
          fieldRef:
            fieldPath: metadata.labels
        - path: "annotations"
          fieldRef:
            fieldPath: metadata.annotations
  {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
  - name: istio-token
    projected:
      sources:
      - serviceAccountToken:
          path: istio-token
          expirationSeconds: 43200
          audience: {{ .Values.global.sds.token.aud }}
  {{- end }}
  {{- if eq .Values.global.pilotCertProvider "istiod" }}
  - name: istiod-ca-cert
    configMap:
      name: istio-ca-root-cert
  {{- end }}
  {{- if .Values.global.mountMtlsCerts }}
  # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
  - name: istio-certs
    secret:
      optional: true
      {{ if eq .Spec.ServiceAccountName "" }}
      secretName: istio.default
      {{ else -}}
      secretName: {{  printf "istio.%s" .Spec.ServiceAccountName }}
      {{  end -}}
  {{- end }}
    {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+` }}
    {{range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+`) }}
  - name: "{{ $index }}"
    {{ toYaml $value | indent 2 }}
    {{ end }}
    {{ end }}
  {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
  - name: lightstep-certs
    secret:
      optional: true
      secretName: lightstep.cacert
  {{- end }}
  podRedirectAnnot:
  {{- if and (.Values.istio_cni.enabled) (not .Values.istio_cni.chained) }}
    k8s.v1.cni.cncf.io/networks: '{{ appendMultusNetwork (index .ObjectMeta.Annotations `+"`"+`k8s.v1.cni.cncf.io/networks`+"`"+`) `+"`"+`istio-cni`+"`"+` }}'
  {{- end }}
    sidecar.istio.io/interceptionMode: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
    traffic.sidecar.istio.io/includeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
    traffic.sidecar.istio.io/excludeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
    traffic.sidecar.istio.io/includeInboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` (includeInboundPorts .Spec.Containers) }}"
    traffic.sidecar.istio.io/excludeInboundPorts: "{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
  {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") }}
    traffic.sidecar.istio.io/includeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
  {{- end }}
  {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne .Values.global.proxy.excludeOutboundPorts "") }}
    traffic.sidecar.istio.io/excludeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
  {{- end }}
    traffic.sidecar.istio.io/kubevirtInterfaces: "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
  {{- if .Values.global.imagePullSecrets }}
  imagePullSecrets:
    {{- range .Values.global.imagePullSecrets }}
    - name: {{ . }}
    {{- end }}
  {{- end }}
`)

func chartsIstioControlIstioDiscoveryFilesInjectionTemplateYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryFilesInjectionTemplateYaml, nil
}

func chartsIstioControlIstioDiscoveryFilesInjectionTemplateYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryFilesInjectionTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/files/injection-template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryKustomizationYaml = []byte(`apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - files/gen-istio.yaml
`)

func chartsIstioControlIstioDiscoveryKustomizationYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryKustomizationYaml, nil
}

func chartsIstioControlIstioDiscoveryKustomizationYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryKustomizationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/kustomization.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesAutoscaleYaml = []byte(`{{- if and .Values.pilot.autoscaleEnabled .Values.pilot.autoscaleMin .Values.pilot.autoscaleMax }}
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiod
    release: {{ .Release.Name }}
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  maxReplicas: {{ .Values.pilot.autoscaleMax }}
  minReplicas: {{ .Values.pilot.autoscaleMin }}
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: {{ .Values.pilot.cpu.targetAverageUtilization }}
---
{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesAutoscaleYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesAutoscaleYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesAutoscaleYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesAutoscaleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/autoscale.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYaml = []byte(`{{- if .Values.pilot.jwksResolverExtraRootCA }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: pilot-jwks-extra-cacerts{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    release: {{ .Release.Name }}
    istio.io/rev: {{ .Values.revision | default "default" }}
data:
  extra.pem: {{ .Values.pilot.jwksResolverExtraRootCA | quote }}
{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/configmap-jwks.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesConfigmapYaml = []byte(`
{{- define "mesh" }}
    # The trust domain corresponds to the trust root of a system.
    # Refer to https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#21-trust-domain
    trustDomain: {{ .Values.global.trustDomain | quote }}

    defaultConfig:
      {{- if .Values.global.meshID }}
      meshId: {{ .Values.global.meshID }}
      {{- else if .Values.global.trustDomain }}
      meshId: {{ .Values.global.trustDomain }}
      {{- end }}
      tracing:
      {{- if eq .Values.global.proxy.tracer "lightstep" }}
        lightstep:
          # Address of the LightStep Satellite pool
          address: {{ .Values.global.tracer.lightstep.address }}
          # Access Token used to communicate with the Satellite pool
          accessToken: {{ .Values.global.tracer.lightstep.accessToken }}
      {{- else if eq .Values.global.proxy.tracer "zipkin" }}
        zipkin:
          # Address of the Zipkin collector
          address: {{ .Values.global.tracer.zipkin.address | default (print "zipkin." .Values.global.istioNamespace ":9411") }}
      {{- else if eq .Values.global.proxy.tracer "datadog" }}
        datadog:
          # Address of the Datadog Agent
          address: {{ .Values.global.tracer.datadog.address | default "$(HOST_IP):8126" }}
      {{- else if eq .Values.global.proxy.tracer "stackdriver" }}
        stackdriver:
          # enables trace output to stdout.
        {{- if $.Values.global.tracer.stackdriver.debug }}
          debug: {{ $.Values.global.tracer.stackdriver.debug }}
        {{- end }}
        {{- if $.Values.global.tracer.stackdriver.maxNumberOfAttributes }}
          # The global default max number of attributes per span.
          maxNumberOfAttributes: {{ $.Values.global.tracer.stackdriver.maxNumberOfAttributes | default "200" }}
        {{- end }}
        {{- if $.Values.global.tracer.stackdriver.maxNumberOfAnnotations }}
          # The global default max number of annotation events per span.
          maxNumberOfAnnotations: {{ $.Values.global.tracer.stackdriver.maxNumberOfAnnotations | default "200" }}
        {{- end }}
        {{- if $.Values.global.tracer.stackdriver.maxNumberOfMessageEvents }}
          # The global default max number of message events per span.
          maxNumberOfMessageEvents: {{ $.Values.global.tracer.stackdriver.maxNumberOfMessageEvents | default "200" }}
        {{- end }}
      {{- else if eq .Values.global.proxy.tracer "openCensusAgent" }}
      {{- /* Fill in openCensusAgent configuration from meshConfig so it isn't overwritten below */ -}}
        {{ toYaml $.Values.meshConfig.defaultConfig.tracing }}
      {{- end }}

      {{- if .Values.global.remotePilotAddress }}
      discoveryAddress: {{ printf "istiod-remote.%s.svc" .Release.Namespace }}:15012
      {{- else }}
      discoveryAddress: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{.Release.Namespace}}.svc:15012
      {{- end }}
{{- end }}

{{/* We take the mesh config above, defined with individual values.yaml, and merge with .Values.meshConfig */}}
{{/* The intent here is that meshConfig.foo becomes the API, rather than re-inventing the API in values.yaml */}}
{{- $originalMesh := include "mesh" . | fromYaml }}
{{- $mesh := mergeOverwrite $originalMesh .Values.meshConfig }}

{{- if .Values.pilot.configMap }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    release: {{ .Release.Name }}
data:

  # Configuration file for the mesh networks to be used by the Split Horizon EDS.
  meshNetworks: |-
  {{- if .Values.global.meshNetworks }}
    networks:
{{ toYaml .Values.global.meshNetworks | trim | indent 6 }}
  {{- else }}
    networks: {}
  {{- end }}

  mesh: |-
{{- if .Values.meshConfig }}
{{ $mesh | toYaml | indent 4 }}
{{- else }}
{{- include "mesh" . }}
{{- end }}
---
{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesConfigmapYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesConfigmapYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesConfigmapYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiod
    istio.io/rev: {{ .Values.revision | default "default" }}
    istio: pilot
    release: {{ .Release.Name }}
{{- range $key, $val := .Values.pilot.deploymentLabels }}
    {{ $key }}: "{{ $val }}"
{{- end }}
spec:
{{- if not .Values.pilot.autoscaleEnabled }}
{{- if .Values.pilot.replicaCount }}
  replicas: {{ .Values.pilot.replicaCount }}
{{- end }}
{{- end }}
  strategy:
    rollingUpdate:
      maxSurge: {{ .Values.pilot.rollingMaxSurge }}
      maxUnavailable: {{ .Values.pilot.rollingMaxUnavailable }}
  selector:
    matchLabels:
      {{- if ne .Values.revision ""}}
      app: istiod
      istio.io/rev: {{ .Values.revision | default "default" }}
      {{- else }}
      istio: pilot
      {{- end }}
  template:
    metadata:
      labels:
        app: istiod
        istio.io/rev: {{ .Values.revision | default "default" }}
        {{- if eq .Values.revision ""}}
        istio: pilot
        {{- else }}
        istio: istiod
        {{- end }}
      annotations:
        {{- if .Values.meshConfig.enablePrometheusMerge }}
        prometheus.io/port: "15014"
        prometheus.io/scrape: "true"
        {{- end }}
        sidecar.istio.io/inject: "false"
        {{- if .Values.pilot.podAnnotations }}
{{ toYaml .Values.pilot.podAnnotations | indent 8 }}
        {{- end }}
    spec:
      serviceAccountName: istiod-service-account
{{- if .Values.global.priorityClassName }}
      priorityClassName: "{{ .Values.global.priorityClassName }}"
{{- end }}
      securityContext:
        fsGroup: 1337
      containers:
        - name: discovery
{{- if contains "/" .Values.pilot.image }}
          image: "{{ .Values.pilot.image }}"
{{- else }}
          image: "{{ .Values.pilot.hub | default .Values.global.hub }}/{{ .Values.pilot.image | default "pilot" }}:{{ .Values.pilot.tag | default .Values.global.tag }}"
{{- end }}
{{- if .Values.global.imagePullPolicy }}
          imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
          args:
          - "discovery"
          - --monitoringAddr=:15014
{{- if .Values.global.logging.level }}
          - --log_output_level={{ .Values.global.logging.level }}
{{- end}}
{{- if .Values.global.logAsJson }}
          - --log_as_json
{{- end }}
          - --domain
          - {{ .Values.global.proxy.clusterDomain }}
{{- if .Values.global.oneNamespace }}
          - "-a"
          - {{ .Release.Namespace }}
{{- end }}
{{- if .Values.global.trustDomain }}
          - --trust-domain={{ .Values.global.trustDomain }}
{{- end }}
{{- if .Values.pilot.plugins }}
          - --plugins={{ .Values.pilot.plugins }}
{{- end }}
          - --keepaliveMaxServerConnectionAge
          - "{{ .Values.pilot.keepaliveMaxServerConnectionAge }}"
          ports:
          - containerPort: 8080
          - containerPort: 15010
          - containerPort: 15017
          - containerPort: 15053
          readinessProbe:
            httpGet:
              path: /ready
              port: 8080
            initialDelaySeconds: 1
            periodSeconds: 3
            timeoutSeconds: 5
          env:
          - name: REVISION
            value: "{{ .Values.revision | default `+"`"+`default`+"`"+` }}"
          - name: JWT_POLICY
            value: {{ .Values.global.jwtPolicy }}
          - name: PILOT_CERT_PROVIDER
            value: {{ .Values.global.pilotCertProvider }}
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: SERVICE_ACCOUNT
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.serviceAccountName
          - name: KUBECONFIG
            value: /var/run/secrets/remote/config
          {{- if .Values.pilot.env }}
          {{- range $key, $val := .Values.pilot.env }}
          - name: {{ $key }}
            value: "{{ $val }}"
          {{- end }}
          {{- end }}
{{- if .Values.pilot.traceSampling }}
          - name: PILOT_TRACE_SAMPLING
            value: "{{ .Values.pilot.traceSampling }}"
{{- end }}
          - name: PILOT_ENABLE_PROTOCOL_SNIFFING_FOR_OUTBOUND
            value: "{{ .Values.pilot.enableProtocolSniffingForOutbound }}"
          - name: PILOT_ENABLE_PROTOCOL_SNIFFING_FOR_INBOUND
            value: "{{ .Values.pilot.enableProtocolSniffingForInbound }}"
          - name: INJECTION_WEBHOOK_CONFIG_NAME
          {{- if eq .Release.Namespace "istio-system" }}
            value: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
          {{- else }}
            value: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}-{{ .Release.Namespace }}
          {{- end }}
          - name: ISTIOD_ADDR
            value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Release.Namespace }}.svc:15012
          - name: PILOT_ENABLE_ANALYSIS
            value: "{{ .Values.global.istiod.enableAnalysis }}"
          - name: CLUSTER_ID
            value: "{{ $.Values.global.multiCluster.clusterName | default `+"`"+`Kubernetes`+"`"+` }}"
          - name: CENTRAL_ISTIOD
            value: "{{ $.Values.global.centralIstiod | default "false" }}"
          resources:
{{- if .Values.pilot.resources }}
{{ toYaml .Values.pilot.resources | trim | indent 12 }}
{{- else }}
{{ toYaml .Values.global.defaultResources | trim | indent 12 }}
{{- end }}
          securityContext:
            runAsUser: 1337
            runAsGroup: 1337
            runAsNonRoot: true
            capabilities:
              drop:
              - ALL
          volumeMounts:
          - name: config-volume
            mountPath: /etc/istio/config
          {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
          - name: istio-token
            mountPath: /var/run/secrets/tokens
            readOnly: true
          {{- end }}
          - name: local-certs
            mountPath: /var/run/secrets/istio-dns
          - name: cacerts
            mountPath: /etc/cacerts
            readOnly: true
          - name: istio-kubeconfig
            mountPath: /var/run/secrets/remote
            readOnly: true
          - name: inject
            mountPath: /var/lib/istio/inject
            readOnly: true
          {{- if .Values.pilot.jwksResolverExtraRootCA }}
          - name: extracacerts
            mountPath: /cacerts
          {{- end }}
      volumes:
      # Technically not needed on this pod - but it helps debugging/testing SDS
      # Should be removed after everything works.
      - emptyDir:
          medium: Memory
        name: local-certs
      {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
      - name: istio-token
        projected:
          sources:
            - serviceAccountToken:
                audience: {{ .Values.global.sds.token.aud }}
                expirationSeconds: 43200
                path: istio-token
      {{- end }}
      # Optional: user-generated root
      - name: cacerts
        secret:
          secretName: cacerts
          optional: true
      - name: istio-kubeconfig
        secret:
          secretName: istio-kubeconfig
          optional: true
      # Optional - image should have
      - name: inject
        configMap:
          name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
      - name: config-volume
        configMap:
          name: istio{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.pilot.jwksResolverExtraRootCA }}
      - name: extracacerts
        configMap:
          name: pilot-jwks-extra-cacerts{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- end }}
---
`)

func chartsIstioControlIstioDiscoveryTemplatesDeploymentYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesDeploymentYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYaml = []byte(`{{- if not .Values.global.omitSidecarInjectorConfigMap }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    release: {{ .Release.Name }}
data:
{{/* Scope the values to just top level fields used in the template, to reduce the size. */}}
  values: |-
{{ pick .Values "global" "istio_cni" "sidecarInjectorWebhook" "revision" | toPrettyJson | indent 4 }}

  # To disable injection: use omitSidecarInjectorConfigMap, which disables the webhook patching
  # and istiod webhook functionality.
  #
  # New fields should not use Values - it is a 'primary' config object, users should be able
  # to fine tune it or use it with kube-inject.
  config: |-
    policy: {{ .Values.global.proxy.autoInject }}
    alwaysInjectSelector:
{{ toYaml .Values.sidecarInjectorWebhook.alwaysInjectSelector | trim | indent 6 }}
    neverInjectSelector:
{{ toYaml .Values.sidecarInjectorWebhook.neverInjectSelector | trim | indent 6 }}
    injectedAnnotations:
      {{- range $key, $val := .Values.sidecarInjectorWebhook.injectedAnnotations }}
      "{{ $key }}": "{{ $val }}"
      {{- end }}

{{ .Files.Get "files/injection-template.yaml" | trim | indent 4 }}

{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/istiod-injector-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYaml = []byte(`# Installed for each revision - not installed for cluster resources ( cluster roles, bindings, crds)
{{- if not .Values.global.operatorManageWebhooks }}
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
{{- if eq .Release.Namespace "istio-system"}}
  name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
{{ else }}
  name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}-{{ .Release.Namespace }}
{{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    app: sidecar-injector
    release: {{ .Release.Name }}
webhooks:
  - name: sidecar-injector.istio.io
    clientConfig:
      service:
        name: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
        namespace: {{ .Release.Namespace }}
        path: "/inject"
      caBundle: ""
    sideEffects: None
    rules:
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    failurePolicy: Fail
    admissionReviewVersions: ["v1beta1", "v1"]
    namespaceSelector:
{{- if .Values.sidecarInjectorWebhook.enableNamespacesByDefault }}
      matchExpressions:
      - key: name
        operator: NotIn
        values:
        - {{ .Release.Namespace }}
      - key: istio-injection
        operator: NotIn
        values:
        - disabled
      - key: istio-env
        operator: DoesNotExist
      - key: istio.io/rev
        operator: DoesNotExist
{{- else if .Values.revision }}
      matchExpressions:
      - key: istio-injection
        operator: DoesNotExist
      - key: istio.io/rev
        operator: In
        values:
        - {{ .Values.revision }}
{{- else }}
      matchLabels:
        istio-injection: enabled
{{- end }}
{{- if .Values.sidecarInjectorWebhook.objectSelector.enabled }}
    objectSelector:
{{- if .Values.sidecarInjectorWebhook.objectSelector.autoInject }}
      matchExpressions:
      - key: "sidecar.istio.io/inject"
        operator: NotIn
        values:
        - "false"
{{- else if .Values.revision }}
      matchExpressions:
      - key: "sidecar.istio.io/inject"
        operator: DoesNotExist
      - key: istio.io/rev
        operator: In
        values:
        - {{ .Values.revision }}
{{- else }}
      matchLabels:
        "sidecar.istio.io/inject": "true"
{{- end }}
{{- end }}
{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/mutatingwebhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYaml = []byte(`{{- if .Values.global.defaultPodDisruptionBudget.enabled }}
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiod
    istio.io/rev: {{ .Values.revision | default "default" }}
    release: {{ .Release.Name }}
    istio: pilot
spec:
  minAvailable: 1
  selector:
    matchLabels:
      app: istiod
      {{- if ne .Values.revision ""}}
      istio.io/rev: {{ .Values.revision }}
      {{- else }}
      istio: pilot
      {{- end }}
---
{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/poddisruptionbudget.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  name: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    app: istiod
    istio: pilot
    release: {{ .Release.Name }}
spec:
  ports:
    - port: 15010
      name: grpc-xds # plaintext
    - port: 15012
      name: https-dns # mTLS with k8s-signed cert
    - port: 443
      name: https-webhook # validation and injection
      targetPort: 15017
    - port: 15014
      name: http-monitoring # prometheus stats
    - name: dns-tls
      port: 853
      targetPort: 15053
      protocol: TCP
  selector:
    app: istiod
    {{- if ne .Values.revision ""}}
    istio.io/rev: {{ .Values.revision }}
    {{- else }}
    # Label used by the 'default' service. For versioned deployments we match with app and version.
    # This avoids default deployment picking the canary
    istio: pilot
    {{- end }}
---
`)

func chartsIstioControlIstioDiscoveryTemplatesServiceYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesServiceYaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesServiceYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16Yaml = []byte(`{{- if and .Values.telemetry.enabled .Values.telemetry.v2.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: ANY # inbound, outbound, and gateway
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration: |
                  {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
{{- if .Values.telemetry.v2.prometheus.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "destination_cluster": "node.metadata['CLUSTER_ID']",
                          "source_cluster": "downstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio",
                    "disable_host_header_fallback": true{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "destination_cluster": "node.metadata['CLUSTER_ID']",
                          "source_cluster": "downstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: tcp_stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
---

{{- end }}

{{- if .Values.telemetry.v2.stackdriver.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
{{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
{{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_inbound
                configuration: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stackdriver_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s", "disable_host_header_fallback": true}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
---
{{- if .Values.telemetry.v2.accessLogPolicy.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-sampling-accesslog-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "istio.stackdriver"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.access_log
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration: |
                  {
                    "log_window_duration": "{{ .Values.telemetry.v2.accessLogPolicy.logWindowDuration }}"
                  }
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: "envoy.wasm.access_log_policy" }
---
{{- end}}
{{- end}}
{{- end}}
`)

func chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16YamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16Yaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16Yaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/telemetryv2_1.6.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17Yaml = []byte(`{{- if and .Values.telemetry.enabled .Values.telemetry.v2.enabled }}
# Note: metadata exchange filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Note: http stats filter is wasm enabled only in sidecars.
{{- if .Values.telemetry.v2.prometheus.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio",
                      "disable_host_header_fallback": true{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Note: tcp stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---

{{- end }}

{{- if .Values.telemetry.v2.stackdriver.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
{{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
{{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s", "disable_host_header_fallback": true}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
---

apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stackdriver-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
  {{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_OUTBOUND
      proxy:
        proxyVersion: '^1\.7.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
    {{- end }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_INBOUND
      proxy:
        proxyVersion: '^1\.7.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_inbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_inbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
  - applyTo: NETWORK_FILTER
    match:
      context: GATEWAY
      proxy:
        proxyVersion: '^1\.7.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
---
{{- if .Values.telemetry.v2.accessLogPolicy.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-sampling-accesslog-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "istio.stackdriver"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.access_log
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "log_window_duration": "{{ .Values.telemetry.v2.accessLogPolicy.logWindowDuration }}"
                    }
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: "envoy.wasm.access_log_policy" }
---
{{- end}}
{{- end}}
{{- end}}`)

func chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17YamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17Yaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17Yaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/telemetryv2_1.7.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18Yaml = []byte(`{{- if and .Values.telemetry.enabled .Values.telemetry.v2.enabled }}
# Note: metadata exchange filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Note: http stats filter is wasm enabled only in sidecars.
{{- if .Values.telemetry.v2.prometheus.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio",
                      "disable_host_header_fallback": true{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Note: tcp stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---

{{- end }}

{{- if .Values.telemetry.v2.stackdriver.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
{{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}", "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
{{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.inboundAccessLogging }}", "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}", "meshEdgesReportingDuration": "600s", "disable_host_header_fallback": true}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
---

apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stackdriver-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
  {{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_OUTBOUND
      proxy:
        proxyVersion: '^1\.8.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
    {{- end }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_INBOUND
      proxy:
        proxyVersion: '^1\.8.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_inbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.inboundAccessLogging }}"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_inbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
  - applyTo: NETWORK_FILTER
    match:
      context: GATEWAY
      proxy:
        proxyVersion: '^1\.8.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
---
{{- if .Values.telemetry.v2.accessLogPolicy.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-sampling-accesslog-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "istio.stackdriver"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.access_log
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "log_window_duration": "{{ .Values.telemetry.v2.accessLogPolicy.logWindowDuration }}"
                    }
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: "envoy.wasm.access_log_policy" }
---
{{- end }}
{{- end }}
{{- end }}
`)

func chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18YamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18Yaml, nil
}

func chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18Yaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioControlIstioDiscoveryValuesYaml = []byte(`#.Values.pilot for discovery and mesh wide config

## Discovery Settings
pilot:
  autoscaleEnabled: true
  autoscaleMin: 1
  autoscaleMax: 5
  replicaCount: 1
  rollingMaxSurge: 100%
  rollingMaxUnavailable: 25%

  hub: ""
  tag: ""

  # Can be a full hub/image:tag
  image: pilot
  traceSampling: 1.0

  # Resources for a small pilot install
  resources:
    requests:
      cpu: 500m
      memory: 2048Mi

  env: {}

  cpu:
    targetAverageUtilization: 80

  # if protocol sniffing is enabled for outbound
  enableProtocolSniffingForOutbound: true
  # if protocol sniffing is enabled for inbound
  enableProtocolSniffingForInbound: true

  nodeSelector: {}
  podAnnotations: {}

  # You can use jwksResolverExtraRootCA to provide a root certificate
  # in PEM format. This will then be trusted by pilot when resolving
  # JWKS URIs.
  jwksResolverExtraRootCA: ""

  # This is used to set the source of configuration for
  # the associated address in configSource, if nothing is specificed
  # the default MCP is assumed.
  configSource:
    subscribedResources: []

  plugins: []

  # The following is used to limit how long a sidecar can be connected
  # to a pilot. It balances out load across pilot instances at the cost of
  # increasing system churn.
  keepaliveMaxServerConnectionAge: 30m

  # Additional labels to apply to the deployment.
  deploymentLabels: {}


  ## Mesh config settings

  # Install the mesh config map, generated from values.yaml.
  # If false, pilot wil use default values (by default) or user-supplied values.
  configMap: true


sidecarInjectorWebhook:
  # You can use the field called alwaysInjectSelector and neverInjectSelector which will always inject the sidecar or
  # always skip the injection on pods that match that label selector, regardless of the global policy.
  # See https://istio.io/docs/setup/kubernetes/additional-setup/sidecar-injection/#more-control-adding-exceptions
  neverInjectSelector: []
  alwaysInjectSelector: []

  # injectedAnnotations are additional annotations that will be added to the pod spec after injection
  # This is primarily to support PSP annotations. For example, if you defined a PSP with the annotations:
  #
  # annotations:
  #   apparmor.security.beta.kubernetes.io/allowedProfileNames: runtime/default
  #   apparmor.security.beta.kubernetes.io/defaultProfileName: runtime/default
  #
  # The PSP controller would add corresponding annotations to the pod spec for each container. However, this happens before
  # the inject adds additional containers, so we must specify them explicitly here. With the above example, we could specify:
  # injectedAnnotations:
  #   container.apparmor.security.beta.kubernetes.io/istio-init: runtime/default
  #   container.apparmor.security.beta.kubernetes.io/istio-proxy: runtime/default
  injectedAnnotations: {}

  # This enables injection of sidecar in all namespaces,
  # with the exception of namespaces with "istio-injection:disabled" annotation
  # Only one environment should have this enabled.
  enableNamespacesByDefault: false

  # Enable objectSelector to filter out pods with no need for sidecar before calling istio-sidecar-injector.
  # It is disabled by default since this function will only work after k8s v1.15.
  objectSelector:
    enabled: false
    autoInject: true

  rewriteAppHTTPProbe: true

telemetry:
  enabled: true
  v2:
    # For Null VM case now.
    # This also enables metadata exchange.
    enabled: true
    metadataExchange:
      # Indicates whether to enable WebAssembly runtime for metadata exchange filter.
      wasmEnabled: false
    # Indicate if prometheus stats filter is enabled or not
    prometheus:
      enabled: true
      # Indicates whether to enable WebAssembly runtime for stats filter.
      wasmEnabled: false
      # overrides stats EnvoyFilter configuration. 
      configOverride:
        gateway: {}
        inboundSidecar: {}
        outboundSidecar: {}
    # stackdriver filter settings.
    stackdriver:
      enabled: false
      logging: false
      monitoring: false
      topology: false
      disableOutbound: false
      #  configOverride parts give you the ability to override the low level configuration params passed to envoy filter.

      configOverride: {}
      #  e.g.
      #  enable_mesh_edges_reporting: true
      #  disable_server_access_logging: false
      #  meshEdgesReportingDuration: 500s
      #  disable_host_header_fallback: true
    # Access Log Policy Filter Settings. This enables filtering of access logs from stackdriver.
    accessLogPolicy:
      enabled: false
      # To reduce the number of successful logs, default log window duration is
      # set to 12 hours.
      logWindowDuration: "43200s"
# Revision is set as 'version' label and part of the resource names when installing multiple control planes.
revision: ""

# meshConfig defines runtime configuration of components, including Istiod and istio-agent behavior
# See https://istio.io/docs/reference/config/istio.mesh.v1alpha1/ for all available options
meshConfig:

  # Config for the default ProxyConfig.
  # Initially using directly the proxy metadata - can also be activated using annotations
  # on the pod. This is an unsupported low-level API, pending review and decisions on
  # enabling the feature. Enabling the DNS listener is safe - and allows further testing
  # and gradual adoption by setting capture only on specific workloads. It also allows
  # VMs to use other DNS options, like dnsmasq or unbound.
  defaultConfig:
    proxyMetadata:
      # If empty, agent will not start :15013 DNS listener and will not attempt
      # to connect to Istiod DNS-TLS. This will also disable the core dns sidecar in
      # istiod and the dns-over-tls listener.
      # DNS_AGENT: DNS-TLS
      DNS_AGENT: ""

  # The namespace to treat as the administrative root namespace for Istio configuration.
  # When processing a leaf namespace Istio will search for declarations in that namespace first
  # and if none are found it will search in the root namespace. Any matching declaration found in the root namespace
  # is processed as if it were declared in the leaf namespace.
  rootNamespace: "istio-system"

  # TODO: the intent is to eventually have this enabled by default when security is used.
  # It is not clear if user should normally need to configure - the metadata is typically
  # used as an escape and to control testing and rollout, but it is not intended as a long-term
  # stable API.

  # What we may configure in mesh config is the ".global" - and use of other suffixes.
  # No hurry to do this in 1.6, we're trying to prove the code.

global:
  # enable pod disruption budget for the control plane, which is used to
  # ensure Istio control plane components are gradually upgraded or recovered.
  defaultPodDisruptionBudget:
    enabled: true
    # The values aren't mutable due to a current PodDisruptionBudget limitation
    # minAvailable: 1

  # A minimal set of requested resources to applied to all deployments so that
  # Horizontal Pod Autoscaler will be able to function (if set).
  # Each component can overwrite these default values by adding its own resources
  # block in the relevant section below and setting the desired resources values.
  defaultResources:
    requests:
      cpu: 10m
    #   memory: 128Mi
    # limits:
    #   cpu: 100m
    #   memory: 128Mi

  # Default hub for Istio images.
  # Releases are published to docker hub under 'istio' project.
  # Dev builds from prow are on gcr.io
  hub: gcr.io/istio-testing
  # Default tag for Istio images.
  tag: latest

  # Specify image pull policy if default behavior isn't desired.
  # Default behavior: latest images will be Always else IfNotPresent.
  imagePullPolicy: ""

  # ImagePullSecrets for all ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # For components that don't use ServiceAccounts (i.e. grafana, servicegraph, tracing)
  # ImagePullSecrets will be added to the corresponding Deployment(StatefulSet) objects.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []
  # - private-registry-key

  # Enabled by default in master for maximising testing.
  istiod:
    enableAnalysis: false

  # To output all istio components logs in json format by adding --log_as_json argument to each container argument
  logAsJson: false

  # Comma-separated minimum per-scope logging level of messages to output, in the form of <scope>:<level>,<scope>:<level>
  # The control plane has different scopes depending on component, but can configure default log level across all components
  # If empty, default scope and level will be used as configured in code
  logging:
    level: "default:info"

  omitSidecarInjectorConfigMap: false

  # Whether to restrict the applications namespace the controller manages;
  # If not set, controller watches all namespaces
  oneNamespace: false

  # Configure whether Operator manages webhook configurations. The current behavior
  # of Istiod is to manage its own webhook configurations.
  # When this option is set as true, Istio Operator, instead of webhooks, manages the
  # webhook configurations. When this option is set as false, webhooks manage their
  # own webhook configurations.
  operatorManageWebhooks: false

  # Custom DNS config for the pod to resolve names of services in other
  # clusters. Use this to add additional search domains, and other settings.
  # see
  # https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#dns-config
  # This does not apply to gateway pods as they typically need a different
  # set of DNS settings than the normal application pods (e.g., in
  # multicluster scenarios).
  # NOTE: If using templates, follow the pattern in the commented example below.
  #podDNSSearchNamespaces:
  #- global
  #- "{{ valueOrDefault .DeploymentMeta.Namespace \"default\" }}.global"

  # Kubernetes >=v1.11.0 will create two PriorityClass, including system-cluster-critical and
  # system-node-critical, it is better to configure this in order to make sure your Istio pods
  # will not be killed because of low priority class.
  # Refer to https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
  # for more detail.
  priorityClassName: ""

  proxy:
    image: proxyv2

    # This controls the 'policy' in the sidecar injector.
    autoInject: enabled

    # CAUTION: It is important to ensure that all Istio helm charts specify the same clusterDomain value
    # cluster domain. Default value is "cluster.local".
    clusterDomain: "cluster.local"

    # Per Component log level for proxy, applies to gateways and sidecars. If a component level is
    # not set, then the global "logLevel" will be used.
    componentLogLevel: "misc:error"

    # If set, newly injected sidecars will have core dumps enabled.
    enableCoreDump: false

    # istio ingress capture allowlist
    # examples:
    #     Redirect only selected ports:            --includeInboundPorts="80,8080"
    excludeInboundPorts: ""

    # istio egress capture allowlist
    # https://istio.io/docs/tasks/traffic-management/egress.html#calling-external-services-directly
    # example: includeIPRanges: "172.30.0.0/16,172.20.0.0/16"
    # would only capture egress traffic on those two IP Ranges, all other outbound traffic would
    # be allowed by the sidecar
    includeIPRanges: "*"
    excludeIPRanges: ""
    excludeOutboundPorts: ""

    # Log level for proxy, applies to gateways and sidecars.
    # Expected values are: trace|debug|info|warning|error|critical|off
    logLevel: warning

    #If set to true, istio-proxy container will have privileged securityContext
    privileged: false

    # The number of successive failed probes before indicating readiness failure.
    readinessFailureThreshold: 30

    # The initial delay for readiness probes in seconds.
    readinessInitialDelaySeconds: 1

    # The period between readiness probes.
    readinessPeriodSeconds: 2

    # Resources for the sidecar.
    resources:
      requests:
        cpu: 100m
        memory: 128Mi
      limits:
        cpu: 2000m
        memory: 1024Mi

    # Default port for Pilot agent health checks. A value of 0 will disable health checking.
    statusPort: 15020

    # Specify which tracer to use. One of: zipkin, lightstep, datadog, stackdriver.
    # If using stackdriver tracer outside GCP, set env GOOGLE_APPLICATION_CREDENTIALS to the GCP credential file.
    tracer: "zipkin"

    # Controls if sidecar is injected at the front of the container list and blocks the start of the other containers until the proxy is ready
    holdApplicationUntilProxyStarts: false

  proxy_init:
    # Base name for the proxy_init container, used to configure iptables.
    image: proxyv2
    resources:
      limits:
        cpu: 2000m
        memory: 1024Mi
      requests:
        cpu: 10m
        memory: 10Mi

  # configure remote pilot and istiod service and endpoint
  remotePilotAddress: ""

  ##############################################################################################
  # The following values are found in other charts. To effectively modify these values, make   #
  # make sure they are consistent across your Istio helm charts                                #
  ##############################################################################################

  # The customized CA address to retrieve certificates for the pods in the cluster.
  # CSR clients such as the Istio Agent and ingress gateways can use this to specify the CA endpoint.
  caAddress: ""

  # One central istiod controls all remote clusters: disabled by default
  centralIstiod: false

  # Configure the policy for validating JWT.
  # Currently, two options are supported: "third-party-jwt" and "first-party-jwt".
  jwtPolicy: "third-party-jwt"

  # Mesh ID means Mesh Identifier. It should be unique within the scope where
  # meshes will interact with each other, but it is not required to be
  # globally/universally unique. For example, if any of the following are true,
  # then two meshes must have different Mesh IDs:
  # - Meshes will have their telemetry aggregated in one place
  # - Meshes will be federated together
  # - Policy will be written referencing one mesh from the other
  #
  # If an administrator expects that any of these conditions may become true in
  # the future, they should ensure their meshes have different Mesh IDs
  # assigned.
  #
  # Within a multicluster mesh, each cluster must be (manually or auto)
  # configured to have the same Mesh ID value. If an existing cluster 'joins' a
  # multicluster mesh, it will need to be migrated to the new mesh ID. Details
  # of migration TBD, and it may be a disruptive operation to change the Mesh
  # ID post-install.
  #
  # If the mesh admin does not specify a value, Istio will use the value of the
  # mesh's Trust Domain. The best practice is to select a proper Trust Domain
  # value.
  meshID: ""

  # Configure the mesh networks to be used by the Split Horizon EDS.
  #
  # The following example defines two networks with different endpoints association methods.
  # For `+"`"+`network1`+"`"+` all endpoints that their IP belongs to the provided CIDR range will be
  # mapped to network1. The gateway for this network example is specified by its public IP
  # address and port.
  # The second network, `+"`"+`network2`+"`"+`, in this example is defined differently with all endpoints
  # retrieved through the specified Multi-Cluster registry being mapped to network2. The
  # gateway is also defined differently with the name of the gateway service on the remote
  # cluster. The public IP for the gateway will be determined from that remote service (only
  # LoadBalancer gateway service type is currently supported, for a NodePort type gateway service,
  # it still need to be configured manually).
  #
  # meshNetworks:
  #   network1:
  #     endpoints:
  #     - fromCidr: "192.168.0.1/24"
  #     gateways:
  #     - address: 1.1.1.1
  #       port: 80
  #   network2:
  #     endpoints:
  #     - fromRegistry: reg1
  #     gateways:
  #     - registryServiceName: istio-ingressgateway.istio-system.svc.cluster.local
  #       port: 443
  #
  meshNetworks: {}

  # Use the user-specified, secret volume mounted key and certs for Pilot and workloads.
  mountMtlsCerts: false

  multiCluster:
    # Set to true to connect two kubernetes clusters via their respective
    # ingressgateway services when pods in each cluster cannot directly
    # talk to one another. All clusters should be using Istio mTLS and must
    # have a shared root CA for this model to work.
    enabled: false
    # Should be set to the name of the cluster this installation will run in. This is required for sidecar injection
    # to properly label proxies
    clusterName: ""

  # Network defines the network this cluster belong to. This name
  # corresponds to the networks in the map of mesh networks.
  network: ""

  # Configure the certificate provider for control plane communication.
  # Currently, two providers are supported: "kubernetes" and "istiod".
  # As some platforms may not have kubernetes signing APIs,
  # Istiod is the default
  pilotCertProvider: istiod

  sds:
    # The JWT token for SDS and the aud field of such JWT. See RFC 7519, section 4.1.3.
    # When a CSR is sent from Istio Agent to the CA (e.g. Istiod), this aud is to make sure the
    # JWT is intended for the CA.
    token:
      aud: istio-ca

  sts:
    # The service port used by Security Token Service (STS) server to handle token exchange requests.
    # Setting this port to a non-zero value enables STS server.
    servicePort: 0

  # Configuration for each of the supported tracers
  tracer:
    # Configuration for envoy to send trace data to LightStep.
    # Disabled by default.
    # address: the <host>:<port> of the satellite pool
    # accessToken: required for sending data to the pool
    #
    datadog:
      # Host:Port for submitting traces to the Datadog agent.
      address: "$(HOST_IP):8126"
    lightstep:
      address: ""                # example: lightstep-satellite:443
      accessToken: ""            # example: abcdefg1234567
    stackdriver:
      # enables trace output to stdout.
      debug: false
      # The global default max number of message events per span.
      maxNumberOfMessageEvents: 200
      # The global default max number of annotation events per span.
      maxNumberOfAnnotations: 200
      # The global default max number of attributes per span.
      maxNumberOfAttributes: 200
    zipkin:
      # Host:Port for reporting trace data in zipkin format. If not specified, will default to
      # zipkin service (port 9411) in the same namespace as the other istio components.
      address: ""

  # The trust domain corresponds to the trust root of a system
  # Refer to https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#21-trust-domain
  # Indicate the domain used in SPIFFE identity URL
  # The default depends on the environment.
  #   kubernetes: cluster.local
  #   else:  default dns domain
  trustDomain: "cluster.local"

  # Use the Mesh Control Protocol (MCP) for configuring Istiod. Requires an MCP source.
  useMCP: false
`)

func chartsIstioControlIstioDiscoveryValuesYamlBytes() ([]byte, error) {
	return _chartsIstioControlIstioDiscoveryValuesYaml, nil
}

func chartsIstioControlIstioDiscoveryValuesYaml() (*asset, error) {
	bytes, err := chartsIstioControlIstioDiscoveryValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-control/istio-discovery/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorChartYaml = []byte(`apiVersion: v1
name: istio-operator
version: 1.7.0
tillerVersion: ">=2.7.2"
description: Helm chart for deploying Istio operator
keywords:
  - istio
  - operator
sources:
  - https://github.com/istio/istio/tree/master/operator
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsIstioOperatorChartYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorChartYaml, nil
}

func chartsIstioOperatorChartYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorCrdsCrdOperatorYaml = []byte(`# SYNC WITH manifests/charts/base/files
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: istiooperators.install.istio.io
spec:
  group: install.istio.io
  names:
    kind: IstioOperator
    plural: istiooperators
    singular: istiooperator
    shortNames:
    - iop
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        spec:
          description: 'Specification of the desired state of the istio control plane resource.
            More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
        status:
          description: 'Status describes each of istio control plane component status at the current time.
            0 means NONE, 1 means UPDATING, 2 means HEALTHY, 3 means ERROR, 4 means RECONCILING.
            More info: https://github.com/istio/api/blob/master/operator/v1alpha1/istio.operator.v1alpha1.pb.html &
            https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
          type: object
  versions:
  - name: v1alpha1
    served: true
    storage: true
---
`)

func chartsIstioOperatorCrdsCrdOperatorYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorCrdsCrdOperatorYaml, nil
}

func chartsIstioOperatorCrdsCrdOperatorYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorCrdsCrdOperatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/crds/crd-operator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorFilesGenOperatorYaml = []byte(`---
# Source: istio-operator/templates/namespace.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: istio-operator
  labels:
    istio-operator-managed: Reconcile
    istio-injection: disabled
---
# Source: istio-operator/templates/service_account.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: istio-operator
  name: istio-operator
---
# Source: istio-operator/templates/clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: istio-operator
rules:
# istio groups
- apiGroups:
  - authentication.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - config.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - install.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - networking.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - security.istio.io
  resources:
  - '*'
  verbs:
  - '*'
# k8s groups
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  - validatingwebhookconfigurations
  verbs:
  - '*'
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions.apiextensions.k8s.io
  - customresourcedefinitions
  verbs:
  - '*'
- apiGroups:
  - apps
  - extensions
  resources:
  - daemonsets
  - deployments
  - deployments/finalizers
  - ingresses
  - replicasets
  - statefulsets
  verbs:
  - '*'
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
  verbs:
  - '*'
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
  - update
- apiGroups:
  - policy
  resources:
  - poddisruptionbudgets
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterrolebindings
  - clusterroles
  - roles
  - rolebindings
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - configmaps
  - endpoints
  - events
  - namespaces
  - pods
  - persistentvolumeclaims
  - secrets
  - services
  - serviceaccounts
  verbs:
  - '*'
---
# Source: istio-operator/templates/clusterrole_binding.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: istio-operator
subjects:
- kind: ServiceAccount
  name: istio-operator
  namespace: istio-operator
roleRef:
  kind: ClusterRole
  name: istio-operator
  apiGroup: rbac.authorization.k8s.io
---
# Source: istio-operator/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  namespace: istio-operator
  labels:
    name: istio-operator
  name: istio-operator
spec:
  ports:
  - name: http-metrics
    port: 8383
    targetPort: 8383
  selector:
    name: istio-operator
---
# Source: istio-operator/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: istio-operator
  name: istio-operator
spec:
  replicas: 1
  selector:
    matchLabels:
      name: istio-operator
  template:
    metadata:
      labels:
        name: istio-operator
    spec:
      serviceAccountName: istio-operator
      containers:
        - name: istio-operator
          image: gcr.io/istio-testing/operator:1.8-dev
          command:
          - operator
          - server
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - ALL
            privileged: false
            readOnlyRootFilesystem: true
            runAsGroup: 1337
            runAsUser: 1337
            runAsNonRoot: true
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: 200m
              memory: 256Mi
            requests:
              cpu: 50m
              memory: 128Mi
          env:
            - name: WATCH_NAMESPACE
              value: "istio-system"
            - name: LEADER_ELECTION_NAMESPACE
              value: "istio-operator"
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: OPERATOR_NAME
              value: "istio-operator"
            - name: WAIT_FOR_RESOURCES_TIMEOUT
              value: "300s"
            - name: REVISION
              value: ""
`)

func chartsIstioOperatorFilesGenOperatorYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorFilesGenOperatorYaml, nil
}

func chartsIstioOperatorFilesGenOperatorYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorFilesGenOperatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/files/gen-operator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
rules:
# istio groups
- apiGroups:
  - authentication.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - config.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - install.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - networking.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - security.istio.io
  resources:
  - '*'
  verbs:
  - '*'
# k8s groups
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  - validatingwebhookconfigurations
  verbs:
  - '*'
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions.apiextensions.k8s.io
  - customresourcedefinitions
  verbs:
  - '*'
- apiGroups:
  - apps
  - extensions
  resources:
  - daemonsets
  - deployments
  - deployments/finalizers
  - ingresses
  - replicasets
  - statefulsets
  verbs:
  - '*'
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
  verbs:
  - '*'
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
  - update
- apiGroups:
  - policy
  resources:
  - poddisruptionbudgets
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterrolebindings
  - clusterroles
  - roles
  - rolebindings
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - configmaps
  - endpoints
  - events
  - namespaces
  - pods
  - persistentvolumeclaims
  - secrets
  - services
  - serviceaccounts
  verbs:
  - '*'
---
`)

func chartsIstioOperatorTemplatesClusterroleYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesClusterroleYaml, nil
}

func chartsIstioOperatorTemplatesClusterroleYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesClusterrole_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
subjects:
- kind: ServiceAccount
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{.Values.operatorNamespace}}
roleRef:
  kind: ClusterRole
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  apiGroup: rbac.authorization.k8s.io
---
`)

func chartsIstioOperatorTemplatesClusterrole_bindingYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesClusterrole_bindingYaml, nil
}

func chartsIstioOperatorTemplatesClusterrole_bindingYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesClusterrole_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/clusterrole_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesCrdsYaml = []byte(`{{- if .Values.enableCRDTemplates -}}
{{- range $path, $bytes := .Files.Glob "crds/*.yaml" -}}
---
{{ $.Files.Get $path }}
{{- end -}}
{{- end -}}
`)

func chartsIstioOperatorTemplatesCrdsYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesCrdsYaml, nil
}

func chartsIstioOperatorTemplatesCrdsYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesCrdsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/crds.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: {{.Values.operatorNamespace}}
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
spec:
  replicas: 1
  selector:
    matchLabels:
      name: istio-operator
  template:
    metadata:
      labels:
        name: istio-operator
    spec:
      serviceAccountName: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
      containers:
        - name: istio-operator
          image: {{.Values.hub}}/operator:{{.Values.tag}}
          command:
          - operator
          - server
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - ALL
            privileged: false
            readOnlyRootFilesystem: true
            runAsGroup: 1337
            runAsUser: 1337
            runAsNonRoot: true
          imagePullPolicy: IfNotPresent
          resources:
{{ toYaml .Values.operator.resources | trim | indent 12 }}
          env:
            - name: WATCH_NAMESPACE
              value: {{.Values.watchedNamespaces | quote}}
            - name: LEADER_ELECTION_NAMESPACE
              value: {{.Values.operatorNamespace | quote}}
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: OPERATOR_NAME
              value: {{.Values.operatorNamespace | quote}}
            - name: WAIT_FOR_RESOURCES_TIMEOUT
              value: {{.Values.waitForResourcesTimeout | quote}}
            - name: REVISION
              value: {{.Values.revision | quote}}
---
`)

func chartsIstioOperatorTemplatesDeploymentYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesDeploymentYaml, nil
}

func chartsIstioOperatorTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: {{.Values.operatorNamespace}}
  labels:
    istio-operator-managed: Reconcile
    istio-injection: disabled
---
`)

func chartsIstioOperatorTemplatesNamespaceYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesNamespaceYaml, nil
}

func chartsIstioOperatorTemplatesNamespaceYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: {{.Values.operatorNamespace}}
  labels:
    name: istio-operator
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
spec:
  ports:
  - name: http-metrics
    port: 8383
    targetPort: 8383
  selector:
    name: istio-operator
---
`)

func chartsIstioOperatorTemplatesServiceYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesServiceYaml, nil
}

func chartsIstioOperatorTemplatesServiceYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorTemplatesService_accountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.Values.operatorNamespace}}
  name: istio-operator{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
---
`)

func chartsIstioOperatorTemplatesService_accountYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorTemplatesService_accountYaml, nil
}

func chartsIstioOperatorTemplatesService_accountYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorTemplatesService_accountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/templates/service_account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstioOperatorValuesYaml = []byte(`hub: gcr.io/istio-testing
tag: latest

operatorNamespace: istio-operator

# Used to replace istioNamespace to support operator watch multiple namespaces.
watchedNamespaces: istio-system
waitForResourcesTimeout: 300s

# Used for helm2 to add the CRDs to templates.
enableCRDTemplates: false

# revision for the operator resources
revision: ""

# Operator resource defaults
operator:
  resources:
    limits:
      cpu: 200m
      memory: 256Mi
    requests:
      cpu: 50m
      memory: 128Mi

`)

func chartsIstioOperatorValuesYamlBytes() ([]byte, error) {
	return _chartsIstioOperatorValuesYaml, nil
}

func chartsIstioOperatorValuesYaml() (*asset, error) {
	bytes, err := chartsIstioOperatorValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istio-operator/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsChartYaml = []byte(`apiVersion: v1
description: Istio CoreDNS provides DNS resolution for services in multicluster setups.
name: istiocoredns
version: 1.1.0
appVersion: 0.1
tillerVersion: ">=2.7.2"
keywords:
  - istio-addon
  - istiocoredns
sources:
  - https://github.com/istio/istio/
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsIstiocorednsChartYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsChartYaml, nil
}

func chartsIstiocorednsChartYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplates_affinityTpl = []byte(`{{/* affinity - https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ */}}

{{- define "nodeaffinity" }}
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
    {{- include "nodeAffinityRequiredDuringScheduling" . }}
    preferredDuringSchedulingIgnoredDuringExecution:
    {{- include "nodeAffinityPreferredDuringScheduling" . }}
{{- end }}

{{- define "nodeAffinityRequiredDuringScheduling" }}
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/arch
          operator: In
          values:
        {{- range $key, $val := .Values.global.arch }}
          {{- if gt ($val | int) 0 }}
          - {{ $key | quote }}
          {{- end }}
        {{- end }}
        {{- $nodeSelector := default .Values.global.defaultNodeSelector .Values.istiocoredns.nodeSelector -}}
        {{- range $key, $val := $nodeSelector }}
        - key: {{ $key }}
          operator: In
          values:
          - {{ $val | quote }}
        {{- end }}
{{- end }}

{{- define "nodeAffinityPreferredDuringScheduling" }}
  {{- range $key, $val := .Values.global.arch }}
    {{- if gt ($val | int) 0 }}
    - weight: {{ $val | int }}
      preference:
        matchExpressions:
        - key: kubernetes.io/arch
          operator: In
          values:
          - {{ $key | quote }}
    {{- end }}
  {{- end }}
{{- end }}

{{- define "podAntiAffinity" }}
{{- if or .Values.istiocoredns.podAntiAffinityLabelSelector .Values.istiocoredns.podAntiAffinityTermLabelSelector}}
  podAntiAffinity:
    {{- if .Values.istiocoredns.podAntiAffinityLabelSelector }}
    requiredDuringSchedulingIgnoredDuringExecution:
    {{- include "podAntiAffinityRequiredDuringScheduling" . }}
    {{- end }}
    {{- if .Values.istiocoredns.podAntiAffinityTermLabelSelector }}
    preferredDuringSchedulingIgnoredDuringExecution:
    {{- include "podAntiAffinityPreferredDuringScheduling" . }}
    {{- end }}
{{- end }}
{{- end }}

{{- define "podAntiAffinityRequiredDuringScheduling" }}
    {{- range $index, $item := .Values.istiocoredns.podAntiAffinityLabelSelector }}
    - labelSelector:
        matchExpressions:
        - key: {{ $item.key }}
          operator: {{ $item.operator }}
          {{- if $item.values }}
          values:
          {{- $vals := split "," $item.values }}
          {{- range $i, $v := $vals }}
          - {{ $v | quote }}
          {{- end }}
          {{- end }}
      topologyKey: {{ $item.topologyKey }}
    {{- end }}
{{- end }}

{{- define "podAntiAffinityPreferredDuringScheduling" }}
    {{- range $index, $item := .Values.istiocoredns.podAntiAffinityTermLabelSelector }}
    - podAffinityTerm:
        labelSelector:
          matchExpressions:
          - key: {{ $item.key }}
            operator: {{ $item.operator }}
            {{- if $item.values }}
            values:
            {{- $vals := split "," $item.values }}
            {{- range $i, $v := $vals }}
            - {{ $v | quote }}
            {{- end }}
            {{- end }}
        topologyKey: {{ $item.topologyKey }}
      weight: 100
    {{- end }}
{{- end }}
`)

func chartsIstiocorednsTemplates_affinityTplBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplates_affinityTpl, nil
}

func chartsIstiocorednsTemplates_affinityTpl() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplates_affinityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/_affinity.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesAutoscaleYaml = []byte(`{{- if and .Values.istiocoredns.autoscaleEnabled .Values.istiocoredns.autoscaleMin .Values.istiocoredns.autoscaleMax }}
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: istiocoredns
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
spec:
  maxReplicas: {{ .Values.istiocoredns.autoscaleMax }}
  minReplicas: {{ .Values.istiocoredns.autoscaleMin }}
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: istiocoredns
  metrics:
    - type: Resource
      resource:
        name: cpu
        targetAverageUtilization: {{ .Values.istiocoredns.cpu.targetAverageUtilization }}
---
{{- end }}
`)

func chartsIstiocorednsTemplatesAutoscaleYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesAutoscaleYaml, nil
}

func chartsIstiocorednsTemplatesAutoscaleYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesAutoscaleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/autoscale.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: istiocoredns
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
rules:
- apiGroups: ["networking.istio.io"]
  resources: ["*"]
  verbs: ["get", "watch", "list"]
`)

func chartsIstiocorednsTemplatesClusterroleYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesClusterroleYaml, nil
}

func chartsIstiocorednsTemplatesClusterroleYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: istio-istiocoredns-role-binding-{{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: istiocoredns
subjects:
- kind: ServiceAccount
  name: istiocoredns-service-account
  namespace: {{ .Release.Namespace }}
`)

func chartsIstiocorednsTemplatesClusterrolebindingYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesClusterrolebindingYaml, nil
}

func chartsIstiocorednsTemplatesClusterrolebindingYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesConfigmapYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
data:
  Corefile: |
    .:53 {
          errors
          health
          {{ if eq -1 (semver  .Values.istiocoredns.coreDNSTag  | (semver "1.4.0").Compare) }}
          # Removed support for the proxy plugin: https://coredns.io/2019/03/03/coredns-1.4.0-release/
          grpc global 127.0.0.1:8053
          forward . /etc/resolv.conf {
            except global
          }
          {{ else }}
          proxy global 127.0.0.1:8053 {
            protocol grpc insecure
          }
          proxy . /etc/resolv.conf
          {{ end }}
          prometheus :9153
          cache 30
          reload
        }
---
`)

func chartsIstiocorednsTemplatesConfigmapYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesConfigmapYaml, nil
}

func chartsIstiocorednsTemplatesConfigmapYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: istiocoredns
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
spec:
{{- if not .Values.istiocoredns.autoscaleEnabled }}
{{- if .Values.istiocoredns.replicaCount }}
  replicas: {{ .Values.istiocoredns.replicaCount }}
{{- end }}
{{- end }}
  selector:
    matchLabels:
      app: istiocoredns
  strategy:
    rollingUpdate:
      maxSurge: {{ .Values.istiocoredns.rollingMaxSurge }}
      maxUnavailable: {{ .Values.istiocoredns.rollingMaxUnavailable }}
  template:
    metadata:
      name: istiocoredns
      labels:
        app: istiocoredns
        release: {{ .Release.Name }}
      annotations:
        sidecar.istio.io/inject: "false"
        {{- if .Values.istiocoredns.podAnnotations }}
{{ toYaml .Values.istiocoredns.podAnnotations | indent 8 }}
        {{- end }}
    spec:
      serviceAccountName: istiocoredns-service-account
{{- if .Values.global.priorityClassName }}
      priorityClassName: "{{ .Values.global.priorityClassName }}"
{{- end }}
      containers:
      - name: coredns
        image: {{ .Values.istiocoredns.coreDNSImage }}:{{ .Values.istiocoredns.coreDNSTag }}
{{- if .Values.global.imagePullPolicy }}
        imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
        args: [ "-conf", "/etc/coredns/Corefile" ]
        volumeMounts:
        - name: config-volume
          mountPath: /etc/coredns
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        - containerPort: 9153
          name: metrics
          protocol: TCP
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        resources:
{{- if .Values.istiocoredns.resources }}
{{ toYaml .Values.istiocoredns.resources | indent 10 }}
{{- else }}
{{ toYaml .Values.global.defaultResources | indent 10 }}
{{- end }}
      - name: istio-coredns-plugin
        command:
        - /usr/local/bin/plugin
        image: {{ .Values.istiocoredns.coreDNSPluginImage }}
{{- if .Values.global.imagePullPolicy }}
        imagePullPolicy: {{ .Values.global.imagePullPolicy }}
{{- end }}
        ports:
        - containerPort: 8053
          name: dns-grpc
          protocol: TCP
        resources:
{{- if .Values.istiocoredns.resources }}
{{ toYaml .Values.istiocorednsresources | indent 10 }}
{{- else }}
{{ toYaml .Values.global.defaultResources | indent 10 }}
{{- end }}
      dnsPolicy: Default
      volumes:
      - name: config-volume
        configMap:
          name: coredns
          items:
          - key: Corefile
            path: Corefile
      affinity:
      {{- include "nodeaffinity" . | indent 6 }}
      {{- include "podAntiAffinity" . | indent 6 }}
{{- if .Values.istiocoredns.tolerations }}
      tolerations:
{{ toYaml .Values.istiocoredns.tolerations | indent 6 }}
{{- else if .Values.global.defaultTolerations }}
      tolerations:
{{ toYaml .Values.global.defaultTolerations | indent 6 }}
{{- end }}
`)

func chartsIstiocorednsTemplatesDeploymentYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesDeploymentYaml, nil
}

func chartsIstiocorednsTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesPoddisruptionbudgetYaml = []byte(`{{- if .Values.global.defaultPodDisruptionBudget.enabled }}
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  name: istiocoredns
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
spec:
  minAvailable: 1
  selector:
    matchLabels:
      app: istiocoredns
      release: {{ .Release.Name }}
{{- end }}
`)

func chartsIstiocorednsTemplatesPoddisruptionbudgetYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesPoddisruptionbudgetYaml, nil
}

func chartsIstiocorednsTemplatesPoddisruptionbudgetYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesPoddisruptionbudgetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/poddisruptionbudget.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  name: istiocoredns
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
spec:
  selector:
    app: istiocoredns
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
`)

func chartsIstiocorednsTemplatesServiceYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesServiceYaml, nil
}

func chartsIstiocorednsTemplatesServiceYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsTemplatesServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
{{- if .Values.global.imagePullSecrets }}
imagePullSecrets:
{{- range .Values.global.imagePullSecrets }}
  - name: {{ . }}
{{- end }}
{{- end }}
metadata:
  name: istiocoredns-service-account
  namespace: {{ .Release.Namespace }}
  labels:
    app: istiocoredns
    release: {{ .Release.Name }}
`)

func chartsIstiocorednsTemplatesServiceaccountYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsTemplatesServiceaccountYaml, nil
}

func chartsIstiocorednsTemplatesServiceaccountYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsTemplatesServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/templates/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiocorednsValuesYaml = []byte(`#
# addon istiocoredns tracing configuration
#
istiocoredns:
  enabled: false
  autoscaleEnabled: true
  autoscaleMin: 1
  autoscaleMax: 5
  replicaCount: 1
  rollingMaxSurge: 100%
  rollingMaxUnavailable: 25%
  coreDNSImage: coredns/coredns
  coreDNSTag: 1.6.2
  # Source code for the plugin can be found at
  # https://github.com/istio-ecosystem/istio-coredns-plugin
  # The plugin listens for DNS requests from coredns server at 127.0.0.1:8053
  coreDNSPluginImage: istio/coredns-plugin:0.2-istio-1.1
  nodeSelector: {}
  tolerations: []
  podAnnotations: {}
  resources: {}
  # Specify the pod anti-affinity that allows you to constrain which nodes
  # your pod is eligible to be scheduled based on labels on pods that are
  # already running on the node rather than based on labels on nodes.
  # There are currently two types of anti-affinity:
  #    "requiredDuringSchedulingIgnoredDuringExecution"
  #    "preferredDuringSchedulingIgnoredDuringExecution"
  # which denote "hard" vs. "soft" requirements, you can define your values
  # in "podAntiAffinityLabelSelector" and "podAntiAffinityTermLabelSelector"
  # correspondingly.
  # For example:
  # podAntiAffinityLabelSelector:
  # - key: security
  #   operator: In
  #   values: S1,S2
  #   topologyKey: "kubernetes.io/hostname"
  # This pod anti-affinity rule says that the pod requires not to be scheduled
  # onto a node if that node is already running a pod with label having key
  # "security" and value "S1".
  podAntiAffinityLabelSelector: []
  podAntiAffinityTermLabelSelector: []

  cpu:
    targetAverageUtilization: 80

global:
  # Specify pod scheduling arch(amd64, ppc64le, s390x) and weight as follows:
  #   0 - Never scheduled
  #   1 - Least preferred
  #   2 - No preference
  #   3 - Most preferred
  arch:
    amd64: 2
    s390x: 2
    ppc64le: 2

  # Default node selector to be applied to all deployments so that all pods can be
  # constrained to run a particular nodes. Each component can overwrite these default
  # values by adding its node selector block in the relevant section below and setting
  # the desired values.
  defaultNodeSelector: {}

  # enable pod disruption budget for the control plane, which is used to
  # ensure Istio control plane components are gradually upgraded or recovered.
  defaultPodDisruptionBudget:
    enabled: true
    # The values aren't mutable due to a current PodDisruptionBudget limitation
    # minAvailable: 1

  # A minimal set of requested resources to applied to all deployments so that
  # Horizontal Pod Autoscaler will be able to function (if set).
  # Each component can overwrite these default values by adding its own resources
  # block in the relevant section below and setting the desired resources values.
  defaultResources:
    requests:
      cpu: 10m
    #   memory: 128Mi
    # limits:
    #   cpu: 100m
    #   memory: 128Mi

  # Default node tolerations to be applied to all deployments so that all pods can be
  # scheduled to a particular nodes with matching taints. Each component can overwrite
  # these default values by adding its tolerations block in the relevant section below
  # and setting the desired values.
  # Configure this field in case that all pods of Istio control plane are expected to
  # be scheduled to particular nodes with specified taints.
  defaultTolerations: []

  # Specify image pull policy if default behavior isn't desired.
  # Default behavior: latest images will be Always else IfNotPresent.
  imagePullPolicy: ""

  # ImagePullSecrets for all ServiceAccount, list of secrets in the same namespace
  # to use for pulling any images in pods that reference this ServiceAccount.
  # For components that don't use ServiceAccounts (i.e. grafana, servicegraph, tracing)
  # ImagePullSecrets will be added to the corresponding Deployment(StatefulSet) objects.
  # Must be set for any cluster configured with private docker registry.
  imagePullSecrets: []
  # - private-registry-key

  # Kubernetes >=v1.11.0 will create two PriorityClass, including system-cluster-critical and
  # system-node-critical, it is better to configure this in order to make sure your Istio pods
  # will not be killed because of low priority class.
  # Refer to https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
  # for more detail.
  priorityClassName: ""
`)

func chartsIstiocorednsValuesYamlBytes() ([]byte, error) {
	return _chartsIstiocorednsValuesYaml, nil
}

func chartsIstiocorednsValuesYaml() (*asset, error) {
	bytes, err := chartsIstiocorednsValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiocoredns/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteChartYaml = []byte(`apiVersion: v1
name: istiod-remote
version: 1.1.0
appVersion: 1.1.0
tillerVersion: ">=2.7.2"
description: Helm chart for istio control plane configuration for remote clusters
keywords:
  - istio
  - istiod-remote
sources:
  - http://github.com/istio/istio
engine: gotpl
icon: https://istio.io/latest/favicons/android-192x192.png
`)

func chartsIstiodRemoteChartYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteChartYaml, nil
}

func chartsIstiodRemoteChartYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/Chart.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteNotesTxt = []byte(`Minimal control plane configuration for Istio remote data plane. Mutating webhook configuration and sidecar injector configuration are included.


`)

func chartsIstiodRemoteNotesTxtBytes() ([]byte, error) {
	return _chartsIstiodRemoteNotesTxt, nil
}

func chartsIstiodRemoteNotesTxt() (*asset, error) {
	bytes, err := chartsIstiodRemoteNotesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/NOTES.txt", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteFilesGenIstiodRemoteYaml = []byte(`---
# Source: istiod-remote/templates/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio
  namespace: istio-system
  labels:
    istio.io/rev: default
    release: istiod-remote
data:

  # Configuration file for the mesh networks to be used by the Split Horizon EDS.
  meshNetworks: |-
    networks: {}

  mesh: |-
    defaultConfig:
      discoveryAddress: istiod.istio-system.svc:15012
      proxyMetadata:
        DNS_AGENT: ""
      tracing:
        zipkin:
          address: zipkin.istio-system:9411
    enablePrometheusMerge: true
    rootNamespace: istio-system
    trustDomain: cluster.local
---
# Source: istiod-remote/templates/istiod-injector-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio-sidecar-injector
  namespace: istio-system
  labels:
    istio.io/rev: default
    release: istiod-remote
data:

  values: |-
    {
      "global": {
        "arch": {
          "amd64": 2,
          "ppc64le": 2,
          "s390x": 2
        },
        "caAddress": "",
        "centralIstiod": false,
        "configValidation": true,
        "defaultConfigVisibilitySettings": [],
        "defaultNodeSelector": {},
        "defaultPodDisruptionBudget": {
          "enabled": true
        },
        "defaultResources": {
          "requests": {
            "cpu": "10m"
          }
        },
        "defaultTolerations": [],
        "hub": "gcr.io/istio-testing",
        "imagePullPolicy": "",
        "imagePullSecrets": [],
        "istioNamespace": "istio-system",
        "istiod": {
          "enableAnalysis": false
        },
        "jwtPolicy": "third-party-jwt",
        "logAsJson": false,
        "logging": {
          "level": "default:info"
        },
        "meshExpansion": {
          "enabled": false,
          "useILB": false
        },
        "meshID": "",
        "meshNetworks": {},
        "mountMtlsCerts": false,
        "multiCluster": {
          "clusterName": "",
          "enabled": false
        },
        "network": "",
        "omitSidecarInjectorConfigMap": false,
        "oneNamespace": false,
        "operatorManageWebhooks": false,
        "pilotCertProvider": "istiod",
        "priorityClassName": "",
        "proxy": {
          "autoInject": "enabled",
          "clusterDomain": "cluster.local",
          "componentLogLevel": "misc:error",
          "enableCoreDump": false,
          "excludeIPRanges": "",
          "excludeInboundPorts": "",
          "excludeOutboundPorts": "",
          "image": "proxyv2",
          "includeIPRanges": "*",
          "logLevel": "warning",
          "privileged": false,
          "readinessFailureThreshold": 30,
          "readinessInitialDelaySeconds": 1,
          "readinessPeriodSeconds": 2,
          "resources": {
            "limits": {
              "cpu": "2000m",
              "memory": "1024Mi"
            },
            "requests": {
              "cpu": "100m",
              "memory": "128Mi"
            }
          },
          "statusPort": 15020,
          "tracer": "zipkin"
        },
        "proxy_init": {
          "image": "proxyv2",
          "resources": {
            "limits": {
              "cpu": "2000m",
              "memory": "1024Mi"
            },
            "requests": {
              "cpu": "10m",
              "memory": "10Mi"
            }
          }
        },
        "remotePilotAddress": "",
        "sds": {
          "token": {
            "aud": "istio-ca"
          }
        },
        "sts": {
          "servicePort": 0
        },
        "tag": "latest",
        "tracer": {
          "datadog": {
            "address": "$(HOST_IP):8126"
          },
          "lightstep": {
            "accessToken": "",
            "address": ""
          },
          "stackdriver": {
            "debug": false,
            "maxNumberOfAnnotations": 200,
            "maxNumberOfAttributes": 200,
            "maxNumberOfMessageEvents": 200
          },
          "zipkin": {
            "address": ""
          }
        },
        "trustDomain": "cluster.local",
        "useMCP": false
      },
      "revision": "",
      "sidecarInjectorWebhook": {
        "alwaysInjectSelector": [],
        "caBundle": "",
        "enableNamespacesByDefault": false,
        "injectedAnnotations": {},
        "neverInjectSelector": [],
        "objectSelector": {
          "autoInject": true,
          "enabled": false
        }
      }
    }

  # To disable injection: use omitSidecarInjectorConfigMap, which disables the webhook patching
  # and istiod webhook functionality.
  #
  # New fields should not use Values - it is a 'primary' config object, users should be able
  # to fine tune it or use it with kube-inject.
  config: |-
    policy: enabled
    alwaysInjectSelector:
      []
    neverInjectSelector:
      []
    injectedAnnotations:

    template: |
      rewriteAppHTTPProbe: {{ valueOrDefault .Values.sidecarInjectorWebhook.rewriteAppHTTPProbe false }}
      initContainers:
      {{ if ne (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`NONE`+"`"+` }}
      {{ if .Values.istio_cni.enabled -}}
      - name: istio-validation
      {{ else -}}
      - name: istio-init
      {{ end -}}
      {{- if contains "/" .Values.global.proxy_init.image }}
        image: "{{ .Values.global.proxy_init.image }}"
      {{- else }}
        image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
      {{- end }}
        args:
        - istio-iptables
        - "-p"
        - 15001
        - "-z"
        - "15006"
        - "-u"
        - 1337
        - "-m"
        - "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
        - "-i"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
        - "-x"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
        - "-b"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` `+"`"+`*`+"`"+` }}"
        - "-d"
      {{- if excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}
        - "15090,15021,{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
      {{- else }}
        - "15090,15021"
      {{- end }}
        {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") -}}
        - "-q"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
        {{ end -}}
        {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.excludeOutboundPorts "") "") -}}
        - "-o"
        - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
        {{ end -}}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+`) -}}
        - "-k"
        - "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
        {{ end -}}
        {{ if .Values.istio_cni.enabled -}}
        - "--run-validation"
        - "--skip-rule-apply"
        {{ end -}}
        imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
      {{- if .ProxyConfig.ProxyMetadata }}
        env:
        {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
        - name: {{ $key }}
          value: "{{ $value }}"
        {{- end }}
      {{- end }}
        resources:
      {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
          requests:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
            {{ end }}
        {{- end }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
          limits:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
            {{ end }}
        {{- end }}
      {{- else }}
        {{- if .Values.global.proxy.resources }}
          {{ toYaml .Values.global.proxy.resources | indent 4 }}
        {{- end }}
      {{- end }}
        securityContext:
          allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
          privileged: {{ .Values.global.proxy.privileged }}
          capabilities:
        {{- if not .Values.istio_cni.enabled }}
            add:
            - NET_ADMIN
            - NET_RAW
        {{- end }}
            drop:
            - ALL
        {{- if not .Values.istio_cni.enabled }}
          readOnlyRootFilesystem: false
          runAsGroup: 0
          runAsNonRoot: false
          runAsUser: 0
        {{- else }}
          readOnlyRootFilesystem: true
          runAsGroup: 1337
          runAsUser: 1337
          runAsNonRoot: true
        {{- end }}
        restartPolicy: Always
      {{ end -}}
      {{- if eq .Values.global.proxy.enableCoreDump true }}
      - name: enable-core-dump
        args:
        - -c
        - sysctl -w kernel.core_pattern=/var/lib/istio/data/core.proxy && ulimit -c unlimited
        command:
          - /bin/sh
      {{- if contains "/" .Values.global.proxy_init.image }}
        image: "{{ .Values.global.proxy_init.image }}"
      {{- else }}
        image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
      {{- end }}
        imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
        resources: {}
        securityContext:
          allowPrivilegeEscalation: true
          capabilities:
            add:
            - SYS_ADMIN
            drop:
            - ALL
          privileged: true
          readOnlyRootFilesystem: false
          runAsGroup: 0
          runAsNonRoot: false
          runAsUser: 0
      {{ end }}
      containers:
      - name: istio-proxy
      {{- if contains "/" (annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image) }}
        image: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image }}"
      {{- else }}
        image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image }}:{{ .Values.global.tag }}"
      {{- end }}
        ports:
        - containerPort: 15090
          protocol: TCP
          name: http-envoy-prom
        args:
        - proxy
        - sidecar
        - --domain
        - $(POD_NAMESPACE).svc.{{ .Values.global.proxy.clusterDomain }}
        - --serviceCluster
        {{ if ne "" (index .ObjectMeta.Labels "app") -}}
        - "{{ index .ObjectMeta.Labels `+"`"+`app`+"`"+` }}.$(POD_NAMESPACE)"
        {{ else -}}
        - "{{ valueOrDefault .DeploymentMeta.Name `+"`"+`istio-proxy`+"`"+` }}.{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}"
        {{ end -}}
        - --proxyLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/logLevel`+"`"+` .Values.global.proxy.logLevel}}
        - --proxyComponentLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/componentLogLevel`+"`"+` .Values.global.proxy.componentLogLevel}}
      {{- if .Values.global.sts.servicePort }}
        - --stsPort={{ .Values.global.sts.servicePort }}
      {{- end }}
      {{- if .Values.global.trustDomain }}
        - --trust-domain={{ .Values.global.trustDomain }}
      {{- end }}
      {{- if .Values.global.logAsJson }}
        - --log_as_json
      {{- end }}
      {{- if gt .ProxyConfig.Concurrency.GetValue 0 }}
        - --concurrency
        - "{{ .ProxyConfig.Concurrency.GetValue }}"
      {{- end -}}
      {{- if .Values.global.proxy.lifecycle }}
        lifecycle:
          {{ toYaml .Values.global.proxy.lifecycle | indent 4 }}
      {{- else if .Values.global.proxy.holdApplicationUntilProxyStarts}}
        lifecycle:
          postStart:
            exec:
              command:
              - pilot-agent
              - wait
      {{- end }}
        env:
        - name: JWT_POLICY
          value: {{ .Values.global.jwtPolicy }}
        - name: PILOT_CERT_PROVIDER
          value: {{ .Values.global.pilotCertProvider }}
        - name: CA_ADDR
        {{- if .Values.global.caAddress }}
          value: {{ .Values.global.caAddress }}
        {{- else }}
          value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Values.global.istioNamespace }}.svc:15012
        {{- end }}
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: INSTANCE_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: SERVICE_ACCOUNT
          valueFrom:
            fieldRef:
              fieldPath: spec.serviceAccountName
        - name: HOST_IP
          valueFrom:
            fieldRef:
              fieldPath: status.hostIP
        - name: CANONICAL_SERVICE
          valueFrom:
            fieldRef:
              fieldPath: metadata.labels['service.istio.io/canonical-name']
        - name: CANONICAL_REVISION
          valueFrom:
            fieldRef:
              fieldPath: metadata.labels['service.istio.io/canonical-revision']
        - name: PROXY_CONFIG
          value: |
                 {{ protoToJSON .ProxyConfig }}
        - name: ISTIO_META_POD_PORTS
          value: |-
            [
            {{- $first := true }}
            {{- range $index1, $c := .Spec.Containers }}
              {{- range $index2, $p := $c.Ports }}
                {{- if (structToJSON $p) }}
                {{if not $first}},{{end}}{{ structToJSON $p }}
                {{- $first = false }}
                {{- end }}
              {{- end}}
            {{- end}}
            ]
        - name: ISTIO_META_APP_CONTAINERS
          value: "{{- range $index, $container := .Spec.Containers }}{{- if ne $index 0}},{{- end}}{{ $container.Name }}{{- end}}"
        - name: ISTIO_META_CLUSTER_ID
          value: "{{ valueOrDefault .Values.global.multiCluster.clusterName `+"`"+`Kubernetes`+"`"+` }}"
        - name: ISTIO_META_INTERCEPTION_MODE
          value: "{{ or (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/interceptionMode`+"`"+`) .ProxyConfig.InterceptionMode.String }}"
        {{- if .Values.global.network }}
        - name: ISTIO_META_NETWORK
          value: "{{ .Values.global.network }}"
        {{- end }}
        {{ if .ObjectMeta.Annotations }}
        - name: ISTIO_METAJSON_ANNOTATIONS
          value: |
                 {{ toJSON .ObjectMeta.Annotations }}
        {{ end }}
        {{- if .DeploymentMeta.Name }}
        - name: ISTIO_META_WORKLOAD_NAME
          value: "{{ .DeploymentMeta.Name }}"
        {{ end }}
        {{- if and .TypeMeta.APIVersion .DeploymentMeta.Name }}
        - name: ISTIO_META_OWNER
          value: kubernetes://apis/{{ .TypeMeta.APIVersion }}/namespaces/{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}/{{ toLower .TypeMeta.Kind}}s/{{ .DeploymentMeta.Name }}
        {{- end}}
        {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
        - name: ISTIO_BOOTSTRAP_OVERRIDE
          value: "/etc/istio/custom-bootstrap/custom_bootstrap.json"
        {{- end }}
        {{- if .Values.global.meshID }}
        - name: ISTIO_META_MESH_ID
          value: "{{ .Values.global.meshID }}"
        {{- else if .Values.global.trustDomain }}
        - name: ISTIO_META_MESH_ID
          value: "{{ .Values.global.trustDomain }}"
        {{- end }}
        {{- if and (eq .Values.global.proxy.tracer "datadog") (isset .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
        {{- range $key, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
        - name: {{ $key }}
          value: "{{ $value }}"
        {{- end }}
        {{- end }}
        {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
        - name: {{ $key }}
          value: "{{ $value }}"
        {{- end }}
        imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
        {{ if ne (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) `+"`"+`0`+"`"+` }}
        readinessProbe:
          httpGet:
            path: /healthz/ready
            port: 15021
          initialDelaySeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/initialDelaySeconds`+"`"+` .Values.global.proxy.readinessInitialDelaySeconds }}
          periodSeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/periodSeconds`+"`"+` .Values.global.proxy.readinessPeriodSeconds }}
          failureThreshold: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/failureThreshold`+"`"+` .Values.global.proxy.readinessFailureThreshold }}
        {{ end -}}
        securityContext:
          allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
          capabilities:
            {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
            add:
            {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+` -}}
            - NET_ADMIN
            {{- end }}
            {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+` -}}
            - NET_BIND_SERVICE
            {{- end }}
            {{- end }}
            drop:
            - ALL
          privileged: {{ .Values.global.proxy.privileged }}
          readOnlyRootFilesystem: {{ not .Values.global.proxy.enableCoreDump }}
          runAsGroup: 1337
          fsGroup: 1337
          {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
          runAsNonRoot: false
          runAsUser: 0
          {{- else -}}
          runAsNonRoot: true
          runAsUser: 1337
          {{- end }}
        resources:
      {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
          requests:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
            {{ end }}
        {{- end }}
        {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
          limits:
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
            cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
            {{ end }}
            {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
            memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
            {{ end }}
        {{- end }}
      {{- else }}
        {{- if .Values.global.proxy.resources }}
          {{ toYaml .Values.global.proxy.resources | indent 4 }}
        {{- end }}
      {{- end }}
        volumeMounts:
        {{- if eq .Values.global.pilotCertProvider "istiod" }}
        - mountPath: /var/run/secrets/istio
          name: istiod-ca-cert
        {{- end }}
        - mountPath: /var/lib/istio/data
          name: istio-data
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
        - mountPath: /etc/istio/custom-bootstrap
          name: custom-bootstrap-volume
        {{- end }}
        # SDS channel between istioagent and Envoy
        - mountPath: /etc/istio/proxy
          name: istio-envoy
        {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
        - mountPath: /var/run/secrets/tokens
          name: istio-token
        {{- end }}
        {{- if .Values.global.mountMtlsCerts }}
        # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
        - mountPath: /etc/certs/
          name: istio-certs
          readOnly: true
        {{- end }}
        - name: istio-podinfo
          mountPath: /etc/istio/pod
         {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
        - mountPath: {{ directory .ProxyConfig.GetTracing.GetTlsSettings.GetCaCertificates }}
          name: lightstep-certs
          readOnly: true
        {{- end }}
          {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+` }}
          {{ range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+`) }}
        - name: "{{  $index }}"
          {{ toYaml $value | indent 4 }}
          {{ end }}
          {{- end }}
      volumes:
      {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
      - name: custom-bootstrap-volume
        configMap:
          name: {{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+` "" }}
      {{- end }}
      # SDS channel between istioagent and Envoy
      - emptyDir:
          medium: Memory
        name: istio-envoy
      - name: istio-data
        emptyDir: {}
      - name: istio-podinfo
        downwardAPI:
          items:
            - path: "labels"
              fieldRef:
                fieldPath: metadata.labels
            - path: "annotations"
              fieldRef:
                fieldPath: metadata.annotations
      {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
      - name: istio-token
        projected:
          sources:
          - serviceAccountToken:
              path: istio-token
              expirationSeconds: 43200
              audience: {{ .Values.global.sds.token.aud }}
      {{- end }}
      {{- if eq .Values.global.pilotCertProvider "istiod" }}
      - name: istiod-ca-cert
        configMap:
          name: istio-ca-root-cert
      {{- end }}
      {{- if .Values.global.mountMtlsCerts }}
      # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
      - name: istio-certs
        secret:
          optional: true
          {{ if eq .Spec.ServiceAccountName "" }}
          secretName: istio.default
          {{ else -}}
          secretName: {{  printf "istio.%s" .Spec.ServiceAccountName }}
          {{  end -}}
      {{- end }}
        {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+` }}
        {{range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+`) }}
      - name: "{{ $index }}"
        {{ toYaml $value | indent 2 }}
        {{ end }}
        {{ end }}
      {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
      - name: lightstep-certs
        secret:
          optional: true
          secretName: lightstep.cacert
      {{- end }}
      podRedirectAnnot:
      {{- if and (.Values.istio_cni.enabled) (not .Values.istio_cni.chained) }}
        k8s.v1.cni.cncf.io/networks: '{{ appendMultusNetwork (index .ObjectMeta.Annotations `+"`"+`k8s.v1.cni.cncf.io/networks`+"`"+`) `+"`"+`istio-cni`+"`"+` }}'
      {{- end }}
        sidecar.istio.io/interceptionMode: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
        traffic.sidecar.istio.io/includeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
        traffic.sidecar.istio.io/excludeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
        traffic.sidecar.istio.io/includeInboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` (includeInboundPorts .Spec.Containers) }}"
        traffic.sidecar.istio.io/excludeInboundPorts: "{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
      {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") }}
        traffic.sidecar.istio.io/includeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
      {{- end }}
      {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne .Values.global.proxy.excludeOutboundPorts "") }}
        traffic.sidecar.istio.io/excludeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
      {{- end }}
        traffic.sidecar.istio.io/kubevirtInterfaces: "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
      {{- if .Values.global.imagePullSecrets }}
      imagePullSecrets:
        {{- range .Values.global.imagePullSecrets }}
        - name: {{ . }}
        {{- end }}
      {{- end }}
---
# Source: istiod-remote/templates/mutatingwebhook.yaml
# Installed for each revision - not installed for cluster resources ( cluster roles, bindings, crds)
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: istio-sidecar-injector

  labels:
    istio.io/rev: default
    app: sidecar-injector
    release: istiod-remote
webhooks:
  - name: sidecar-injector.istio.io
    clientConfig:
      service:
        name: istiod
        namespace: istio-system
        path: "/inject"
      caBundle: 
    rules:
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    failurePolicy: Fail
    admissionReviewVersions: ["v1beta1", "v1"]
    sideEffects: None
    namespaceSelector:
      matchLabels:
        istio-injection: enabled
`)

func chartsIstiodRemoteFilesGenIstiodRemoteYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteFilesGenIstiodRemoteYaml, nil
}

func chartsIstiodRemoteFilesGenIstiodRemoteYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteFilesGenIstiodRemoteYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/files/gen-istiod-remote.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteFilesInjectionTemplateYaml = []byte(`template: |
  rewriteAppHTTPProbe: {{ valueOrDefault .Values.sidecarInjectorWebhook.rewriteAppHTTPProbe false }}
  initContainers:
  {{ if ne (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`NONE`+"`"+` }}
  {{ if .Values.istio_cni.enabled -}}
  - name: istio-validation
  {{ else -}}
  - name: istio-init
  {{ end -}}
  {{- if contains "/" .Values.global.proxy_init.image }}
    image: "{{ .Values.global.proxy_init.image }}"
  {{- else }}
    image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
  {{- end }}
    args:
    - istio-iptables
    - "-p"
    - 15001
    - "-z"
    - "15006"
    - "-u"
    - 1337
    - "-m"
    - "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
    - "-i"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
    - "-x"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
    - "-b"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` `+"`"+`*`+"`"+` }}"
    - "-d"
  {{- if excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}
    - "15090,15021,{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
  {{- else }}
    - "15090,15021"
  {{- end }}
    {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") -}}
    - "-q"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
    {{ end -}}
    {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.excludeOutboundPorts "") "") -}}
    - "-o"
    - "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
    {{ end -}}
    {{ if (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+`) -}}
    - "-k"
    - "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
    {{ end -}}
    {{ if .Values.istio_cni.enabled -}}
    - "--run-validation"
    - "--skip-rule-apply"
    {{ end -}}
    imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
  {{- if .ProxyConfig.ProxyMetadata }}
    env:
    {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
    - name: {{ $key }}
      value: "{{ $value }}"
    {{- end }}
  {{- end }}
    resources:
  {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
      requests:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
        {{ end }}
    {{- end }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
      limits:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
        {{ end }}
    {{- end }}
  {{- else }}
    {{- if .Values.global.proxy.resources }}
      {{ toYaml .Values.global.proxy.resources | indent 4 }}
    {{- end }}
  {{- end }}
    securityContext:
      allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
      privileged: {{ .Values.global.proxy.privileged }}
      capabilities:
    {{- if not .Values.istio_cni.enabled }}
        add:
        - NET_ADMIN
        - NET_RAW
    {{- end }}
        drop:
        - ALL
    {{- if not .Values.istio_cni.enabled }}
      readOnlyRootFilesystem: false
      runAsGroup: 0
      runAsNonRoot: false
      runAsUser: 0
    {{- else }}
      readOnlyRootFilesystem: true
      runAsGroup: 1337
      runAsUser: 1337
      runAsNonRoot: true
    {{- end }}
    restartPolicy: Always
  {{ end -}}
  {{- if eq .Values.global.proxy.enableCoreDump true }}
  - name: enable-core-dump
    args:
    - -c
    - sysctl -w kernel.core_pattern=/var/lib/istio/data/core.proxy && ulimit -c unlimited
    command:
      - /bin/sh
  {{- if contains "/" .Values.global.proxy_init.image }}
    image: "{{ .Values.global.proxy_init.image }}"
  {{- else }}
    image: "{{ .Values.global.hub }}/{{ .Values.global.proxy_init.image }}:{{ .Values.global.tag }}"
  {{- end }}
    imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
    resources: {}
    securityContext:
      allowPrivilegeEscalation: true
      capabilities:
        add:
        - SYS_ADMIN
        drop:
        - ALL
      privileged: true
      readOnlyRootFilesystem: false
      runAsGroup: 0
      runAsNonRoot: false
      runAsUser: 0
  {{ end }}
  containers:
  - name: istio-proxy
  {{- if contains "/" (annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image) }}
    image: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/proxyImage`+"`"+` .Values.global.proxy.image }}"
  {{- else }}
    image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image }}:{{ .Values.global.tag }}"
  {{- end }}
    ports:
    - containerPort: 15090
      protocol: TCP
      name: http-envoy-prom
    args:
    - proxy
    - sidecar
    - --domain
    - $(POD_NAMESPACE).svc.{{ .Values.global.proxy.clusterDomain }}
    - --serviceCluster
    {{ if ne "" (index .ObjectMeta.Labels "app") -}}
    - "{{ index .ObjectMeta.Labels `+"`"+`app`+"`"+` }}.$(POD_NAMESPACE)"
    {{ else -}}
    - "{{ valueOrDefault .DeploymentMeta.Name `+"`"+`istio-proxy`+"`"+` }}.{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}"
    {{ end -}}
    - --proxyLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/logLevel`+"`"+` .Values.global.proxy.logLevel}}
    - --proxyComponentLogLevel={{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/componentLogLevel`+"`"+` .Values.global.proxy.componentLogLevel}}
  {{- if .Values.global.sts.servicePort }}
    - --stsPort={{ .Values.global.sts.servicePort }}
  {{- end }}
  {{- if .Values.global.trustDomain }}
    - --trust-domain={{ .Values.global.trustDomain }}
  {{- end }}
  {{- if .Values.global.logAsJson }}
    - --log_as_json
  {{- end }}
  {{- if gt .ProxyConfig.Concurrency.GetValue 0 }}
    - --concurrency
    - "{{ .ProxyConfig.Concurrency.GetValue }}"
  {{- end -}}
  {{- if .Values.global.proxy.lifecycle }}
    lifecycle:
      {{ toYaml .Values.global.proxy.lifecycle | indent 4 }}
  {{- else if .Values.global.proxy.holdApplicationUntilProxyStarts}}
    lifecycle:
      postStart:
        exec:
          command:
          - pilot-agent
          - wait
  {{- end }}
    env:
    - name: JWT_POLICY
      value: {{ .Values.global.jwtPolicy }}
    - name: PILOT_CERT_PROVIDER
      value: {{ .Values.global.pilotCertProvider }}
    - name: CA_ADDR
    {{- if .Values.global.caAddress }}
      value: {{ .Values.global.caAddress }}
    {{- else }}
      value: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{ .Values.global.istioNamespace }}.svc:15012
    {{- end }}
    - name: POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    - name: POD_NAMESPACE
      valueFrom:
        fieldRef:
          fieldPath: metadata.namespace
    - name: INSTANCE_IP
      valueFrom:
        fieldRef:
          fieldPath: status.podIP
    - name: SERVICE_ACCOUNT
      valueFrom:
        fieldRef:
          fieldPath: spec.serviceAccountName
    - name: HOST_IP
      valueFrom:
        fieldRef:
          fieldPath: status.hostIP
    - name: CANONICAL_SERVICE
      valueFrom:
        fieldRef:
          fieldPath: metadata.labels['service.istio.io/canonical-name']
    - name: CANONICAL_REVISION
      valueFrom:
        fieldRef:
          fieldPath: metadata.labels['service.istio.io/canonical-revision']
    - name: PROXY_CONFIG
      value: |
             {{ protoToJSON .ProxyConfig }}
    - name: ISTIO_META_POD_PORTS
      value: |-
        [
        {{- $first := true }}
        {{- range $index1, $c := .Spec.Containers }}
          {{- range $index2, $p := $c.Ports }}
            {{- if (structToJSON $p) }}
            {{if not $first}},{{end}}{{ structToJSON $p }}
            {{- $first = false }}
            {{- end }}
          {{- end}}
        {{- end}}
        ]
    - name: ISTIO_META_APP_CONTAINERS
      value: "{{- range $index, $container := .Spec.Containers }}{{- if ne $index 0}},{{- end}}{{ $container.Name }}{{- end}}"
    - name: ISTIO_META_CLUSTER_ID
      value: "{{ valueOrDefault .Values.global.multiCluster.clusterName `+"`"+`Kubernetes`+"`"+` }}"
    - name: ISTIO_META_INTERCEPTION_MODE
      value: "{{ or (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/interceptionMode`+"`"+`) .ProxyConfig.InterceptionMode.String }}"
    {{- if .Values.global.network }}
    - name: ISTIO_META_NETWORK
      value: "{{ .Values.global.network }}"
    {{- end }}
    {{ if .ObjectMeta.Annotations }}
    - name: ISTIO_METAJSON_ANNOTATIONS
      value: |
             {{ toJSON .ObjectMeta.Annotations }}
    {{ end }}
    {{- if .DeploymentMeta.Name }}
    - name: ISTIO_META_WORKLOAD_NAME
      value: "{{ .DeploymentMeta.Name }}"
    {{ end }}
    {{- if and .TypeMeta.APIVersion .DeploymentMeta.Name }}
    - name: ISTIO_META_OWNER
      value: kubernetes://apis/{{ .TypeMeta.APIVersion }}/namespaces/{{ valueOrDefault .DeploymentMeta.Namespace `+"`"+`default`+"`"+` }}/{{ toLower .TypeMeta.Kind}}s/{{ .DeploymentMeta.Name }}
    {{- end}}
    {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
    - name: ISTIO_BOOTSTRAP_OVERRIDE
      value: "/etc/istio/custom-bootstrap/custom_bootstrap.json"
    {{- end }}
    {{- if .Values.global.meshID }}
    - name: ISTIO_META_MESH_ID
      value: "{{ .Values.global.meshID }}"
    {{- else if .Values.global.trustDomain }}
    - name: ISTIO_META_MESH_ID
      value: "{{ .Values.global.trustDomain }}"
    {{- end }}
    {{- if and (eq .Values.global.proxy.tracer "datadog") (isset .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
    {{- range $key, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`apm.datadoghq.com/env`+"`"+`) }}
    - name: {{ $key }}
      value: "{{ $value }}"
    {{- end }}
    {{- end }}
    {{- range $key, $value := .ProxyConfig.ProxyMetadata }}
    - name: {{ $key }}
      value: "{{ $value }}"
    {{- end }}
    imagePullPolicy: "{{ valueOrDefault .Values.global.imagePullPolicy `+"`"+`Always`+"`"+` }}"
    {{ if ne (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) `+"`"+`0`+"`"+` }}
    readinessProbe:
      httpGet:
        path: /healthz/ready
        port: 15021
      initialDelaySeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/initialDelaySeconds`+"`"+` .Values.global.proxy.readinessInitialDelaySeconds }}
      periodSeconds: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/periodSeconds`+"`"+` .Values.global.proxy.readinessPeriodSeconds }}
      failureThreshold: {{ annotation .ObjectMeta `+"`"+`readiness.status.sidecar.istio.io/failureThreshold`+"`"+` .Values.global.proxy.readinessFailureThreshold }}
    {{ end -}}
    securityContext:
      allowPrivilegeEscalation: {{ .Values.global.proxy.privileged }}
      capabilities:
        {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
        add:
        {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+` -}}
        - NET_ADMIN
        {{- end }}
        {{ if eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+` -}}
        - NET_BIND_SERVICE
        {{- end }}
        {{- end }}
        drop:
        - ALL
      privileged: {{ .Values.global.proxy.privileged }}
      readOnlyRootFilesystem: {{ not .Values.global.proxy.enableCoreDump }}
      runAsGroup: 1337
      fsGroup: 1337
      {{ if or (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode) `+"`"+`TPROXY`+"`"+`) (eq (annotation .ObjectMeta `+"`"+`sidecar.istio.io/capNetBindService`+"`"+` .Values.global.proxy.capNetBindService) `+"`"+`true`+"`"+`) -}}
      runAsNonRoot: false
      runAsUser: 0
      {{- else -}}
      runAsNonRoot: true
      runAsUser: 1337
      {{- end }}
    resources:
  {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) }}
      requests:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPU`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemory`+"`"+` }}"
        {{ end }}
    {{- end }}
    {{- if or (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) }}
      limits:
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+`) -}}
        cpu: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyCPULimit`+"`"+` }}"
        {{ end }}
        {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+`) -}}
        memory: "{{ index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/proxyMemoryLimit`+"`"+` }}"
        {{ end }}
    {{- end }}
  {{- else }}
    {{- if .Values.global.proxy.resources }}
      {{ toYaml .Values.global.proxy.resources | indent 4 }}
    {{- end }}
  {{- end }}
    volumeMounts:
    {{- if eq .Values.global.pilotCertProvider "istiod" }}
    - mountPath: /var/run/secrets/istio
      name: istiod-ca-cert
    {{- end }}
    - mountPath: /var/lib/istio/data
      name: istio-data
    {{ if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
    - mountPath: /etc/istio/custom-bootstrap
      name: custom-bootstrap-volume
    {{- end }}
    # SDS channel between istioagent and Envoy
    - mountPath: /etc/istio/proxy
      name: istio-envoy
    {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
    - mountPath: /var/run/secrets/tokens
      name: istio-token
    {{- end }}
    {{- if .Values.global.mountMtlsCerts }}
    # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
    - mountPath: /etc/certs/
      name: istio-certs
      readOnly: true
    {{- end }}
    - name: istio-podinfo
      mountPath: /etc/istio/pod
     {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
    - mountPath: {{ directory .ProxyConfig.GetTracing.GetTlsSettings.GetCaCertificates }}
      name: lightstep-certs
      readOnly: true
    {{- end }}
      {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+` }}
      {{ range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolumeMount`+"`"+`) }}
    - name: "{{  $index }}"
      {{ toYaml $value | indent 4 }}
      {{ end }}
      {{- end }}
  volumes:
  {{- if (isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+`) }}
  - name: custom-bootstrap-volume
    configMap:
      name: {{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/bootstrapOverride`+"`"+` "" }}
  {{- end }}
  # SDS channel between istioagent and Envoy
  - emptyDir:
      medium: Memory
    name: istio-envoy
  - name: istio-data
    emptyDir: {}
  - name: istio-podinfo
    downwardAPI:
      items:
        - path: "labels"
          fieldRef:
            fieldPath: metadata.labels
        - path: "annotations"
          fieldRef:
            fieldPath: metadata.annotations
  {{- if eq .Values.global.jwtPolicy "third-party-jwt" }}
  - name: istio-token
    projected:
      sources:
      - serviceAccountToken:
          path: istio-token
          expirationSeconds: 43200
          audience: {{ .Values.global.sds.token.aud }}
  {{- end }}
  {{- if eq .Values.global.pilotCertProvider "istiod" }}
  - name: istiod-ca-cert
    configMap:
      name: istio-ca-root-cert
  {{- end }}
  {{- if .Values.global.mountMtlsCerts }}
  # Use the key and cert mounted to /etc/certs/ for the in-cluster mTLS communications.
  - name: istio-certs
    secret:
      optional: true
      {{ if eq .Spec.ServiceAccountName "" }}
      secretName: istio.default
      {{ else -}}
      secretName: {{  printf "istio.%s" .Spec.ServiceAccountName }}
      {{  end -}}
  {{- end }}
    {{- if isset .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+` }}
    {{range $index, $value := fromJSON (index .ObjectMeta.Annotations `+"`"+`sidecar.istio.io/userVolume`+"`"+`) }}
  - name: "{{ $index }}"
    {{ toYaml $value | indent 2 }}
    {{ end }}
    {{ end }}
  {{- if and (eq .Values.global.proxy.tracer "lightstep") .ProxyConfig.GetTracing.GetTlsSettings }}
  - name: lightstep-certs
    secret:
      optional: true
      secretName: lightstep.cacert
  {{- end }}
  podRedirectAnnot:
  {{- if and (.Values.istio_cni.enabled) (not .Values.istio_cni.chained) }}
    k8s.v1.cni.cncf.io/networks: '{{ appendMultusNetwork (index .ObjectMeta.Annotations `+"`"+`k8s.v1.cni.cncf.io/networks`+"`"+`) `+"`"+`istio-cni`+"`"+` }}'
  {{- end }}
    sidecar.istio.io/interceptionMode: "{{ annotation .ObjectMeta `+"`"+`sidecar.istio.io/interceptionMode`+"`"+` .ProxyConfig.InterceptionMode }}"
    traffic.sidecar.istio.io/includeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundIPRanges`+"`"+` .Values.global.proxy.includeIPRanges }}"
    traffic.sidecar.istio.io/excludeOutboundIPRanges: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundIPRanges`+"`"+` .Values.global.proxy.excludeIPRanges }}"
    traffic.sidecar.istio.io/includeInboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeInboundPorts`+"`"+` (includeInboundPorts .Spec.Containers) }}"
    traffic.sidecar.istio.io/excludeInboundPorts: "{{ excludeInboundPort (annotation .ObjectMeta `+"`"+`status.sidecar.istio.io/port`+"`"+` .Values.global.proxy.statusPort) (annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeInboundPorts`+"`"+` .Values.global.proxy.excludeInboundPorts) }}"
  {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+`) (ne (valueOrDefault .Values.global.proxy.includeOutboundPorts "") "") }}
    traffic.sidecar.istio.io/includeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/includeOutboundPorts`+"`"+` .Values.global.proxy.includeOutboundPorts }}"
  {{- end }}
  {{ if or (isset .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+`) (ne .Values.global.proxy.excludeOutboundPorts "") }}
    traffic.sidecar.istio.io/excludeOutboundPorts: "{{ annotation .ObjectMeta `+"`"+`traffic.sidecar.istio.io/excludeOutboundPorts`+"`"+` .Values.global.proxy.excludeOutboundPorts }}"
  {{- end }}
    traffic.sidecar.istio.io/kubevirtInterfaces: "{{ index .ObjectMeta.Annotations `+"`"+`traffic.sidecar.istio.io/kubevirtInterfaces`+"`"+` }}"
  {{- if .Values.global.imagePullSecrets }}
  imagePullSecrets:
    {{- range .Values.global.imagePullSecrets }}
    - name: {{ . }}
    {{- end }}
  {{- end }}
`)

func chartsIstiodRemoteFilesInjectionTemplateYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteFilesInjectionTemplateYaml, nil
}

func chartsIstiodRemoteFilesInjectionTemplateYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteFilesInjectionTemplateYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/files/injection-template.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteKustomizationYaml = []byte(`apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - files/gen-istiod-remote.yaml
`)

func chartsIstiodRemoteKustomizationYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteKustomizationYaml, nil
}

func chartsIstiodRemoteKustomizationYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteKustomizationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/kustomization.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteTemplatesConfigmapYaml = []byte(`
{{- define "mesh" }}
    # The trust domain corresponds to the trust root of a system.
    # Refer to https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#21-trust-domain
    trustDomain: {{ .Values.global.trustDomain | quote }}

    defaultConfig:
      tracing:
      {{- if eq .Values.global.proxy.tracer "lightstep" }}
        lightstep:
          # Address of the LightStep Satellite pool
          address: {{ .Values.global.tracer.lightstep.address }}
          # Access Token used to communicate with the Satellite pool
          accessToken: {{ .Values.global.tracer.lightstep.accessToken }}
      {{- else if eq .Values.global.proxy.tracer "zipkin" }}
        zipkin:
          # Address of the Zipkin collector
          address: {{ .Values.global.tracer.zipkin.address | default (print "zipkin." .Values.global.istioNamespace ":9411") }}
      {{- else if eq .Values.global.proxy.tracer "datadog" }}
        datadog:
          # Address of the Datadog Agent
          address: {{ .Values.global.tracer.datadog.address | default "$(HOST_IP):8126" }}
      {{- else if eq .Values.global.proxy.tracer "stackdriver" }}
        stackdriver:
          # enables trace output to stdout.
        {{- if $.Values.global.tracer.stackdriver.debug }}
          debug: {{ $.Values.global.tracer.stackdriver.debug }}
        {{- end }}
        {{- if $.Values.global.tracer.stackdriver.maxNumberOfAttributes }}
          # The global default max number of attributes per span.
          maxNumberOfAttributes: {{ $.Values.global.tracer.stackdriver.maxNumberOfAttributes | default "200" }}
        {{- end }}
        {{- if $.Values.global.tracer.stackdriver.maxNumberOfAnnotations }}
          # The global default max number of annotation events per span.
          maxNumberOfAnnotations: {{ $.Values.global.tracer.stackdriver.maxNumberOfAnnotations | default "200" }}
        {{- end }}
        {{- if $.Values.global.tracer.stackdriver.maxNumberOfMessageEvents }}
          # The global default max number of message events per span.
          maxNumberOfMessageEvents: {{ $.Values.global.tracer.stackdriver.maxNumberOfMessageEvents | default "200" }}
        {{- end }}
      {{- else if eq .Values.global.proxy.tracer "openCensusAgent" }}
      {{- /* Fill in openCensusAgent configuration from meshConfig so it isn't overwritten below */ -}}
        {{ toYaml $.Values.meshConfig.defaultConfig.tracing }}
      {{- end }}

      {{- if .Values.global.remotePilotAddress }}
      discoveryAddress: {{ printf "istiod-remote.%s.svc" .Release.Namespace }}:15012
      {{- else }}
      discoveryAddress: istiod{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}.{{.Release.Namespace}}.svc:15012
      {{- end }}
{{- end }}

{{/* We take the mesh config above, defined with individual values.yaml, and merge with .Values.meshConfig */}}
{{/* The intent here is that meshConfig.foo becomes the API, rather than re-inventing the API in values.yaml */}}
{{- $originalMesh := include "mesh" . | fromYaml }}
{{- $mesh := mergeOverwrite $originalMesh .Values.meshConfig }}

{{- if not .Values.global.centralIstiod }}
{{- if .Values.pilot.configMap }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    release: {{ .Release.Name }}
data:

  # Configuration file for the mesh networks to be used by the Split Horizon EDS.
  meshNetworks: |-
  {{- if .Values.global.meshNetworks }}
    networks:
{{ toYaml .Values.global.meshNetworks | trim | indent 6 }}
  {{- else }}
    networks: {}
  {{- end }}

  mesh: |-
{{- if .Values.meshConfig }}
{{ $mesh | toYaml | indent 4 }}
{{- else }}
{{- include "mesh" . }}
{{- end }}
---
{{- end }}
{{- end }}
`)

func chartsIstiodRemoteTemplatesConfigmapYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteTemplatesConfigmapYaml, nil
}

func chartsIstiodRemoteTemplatesConfigmapYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteTemplatesConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/templates/configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYaml = []byte(`{{- if not .Values.global.omitSidecarInjectorConfigMap }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  namespace: {{ .Release.Namespace }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    release: {{ .Release.Name }}
data:
{{/* Scope the values to just top level fields used in the template, to reduce the size. */}}
  values: |-
{{ pick .Values "global" "istio_cni" "sidecarInjectorWebhook" "revision" | toPrettyJson | indent 4 }}

  # To disable injection: use omitSidecarInjectorConfigMap, which disables the webhook patching
  # and istiod webhook functionality.
  #
  # New fields should not use Values - it is a 'primary' config object, users should be able
  # to fine tune it or use it with kube-inject.
  config: |-
    policy: {{ .Values.global.proxy.autoInject }}
    alwaysInjectSelector:
{{ toYaml .Values.sidecarInjectorWebhook.alwaysInjectSelector | trim | indent 6 }}
    neverInjectSelector:
{{ toYaml .Values.sidecarInjectorWebhook.neverInjectSelector | trim | indent 6 }}
    injectedAnnotations:
      {{- range $key, $val := .Values.sidecarInjectorWebhook.injectedAnnotations }}
      "{{ $key }}": "{{ $val }}"
      {{- end }}

{{ .Files.Get "files/injection-template.yaml" | trim | indent 4 }}

{{- end }}
`)

func chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYaml, nil
}

func chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/templates/istiod-injector-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteTemplatesMutatingwebhookYaml = []byte(`# Installed for each revision - not installed for cluster resources ( cluster roles, bindings, crds)
{{- if not .Values.global.operatorManageWebhooks }}
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
{{- if eq .Release.Namespace "istio-system"}}
  name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
{{ else }}
  name: istio-sidecar-injector{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}-{{ .Release.Namespace }}
{{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
    app: sidecar-injector
    release: {{ .Release.Name }}
webhooks:
  - name: sidecar-injector.istio.io
    clientConfig:
      {{- if .Values.istiodRemote.injectionURL }}
      # the url doesn't take query params, so use inject/cluster/x/net/y format
      url: {{ .Values.istiodRemote.injectionURL }}/cluster/{{ .Values.global.multiCluster.clusterName | default "remote0" }}/net/{{ .Values.global.network | default "network0" }}
      {{- else }}
      service:
        name: istiod
        namespace: {{ .Values.global.istioNamespace }}
        path: "/inject"
      {{- end }}
      caBundle: {{ .Values.sidecarInjectorWebhook.caBundle }}
    rules:
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    failurePolicy: Fail
    admissionReviewVersions: ["v1beta1", "v1"]
    sideEffects: None
    namespaceSelector:
{{- if .Values.sidecarInjectorWebhook.enableNamespacesByDefault }}
      matchExpressions:
      - key: name
        operator: NotIn
        values:
        - {{ .Release.Namespace }}
      - key: istio-injection
        operator: NotIn
        values:
        - disabled
      - key: istio-env
        operator: DoesNotExist
      - key: istio.io/rev
        operator: DoesNotExist
{{- else if .Values.revision }}
      matchExpressions:
      - key: istio-injection
        operator: DoesNotExist
      - key: istio.io/rev
        operator: In
        values:
        - {{ .Values.revision }}
{{- else }}
      matchLabels:
        istio-injection: enabled
{{- end }}
{{- if .Values.sidecarInjectorWebhook.objectSelector.enabled }}
    objectSelector:
{{- if .Values.sidecarInjectorWebhook.objectSelector.autoInject }}
      matchExpressions:
      - key: "sidecar.istio.io/inject"
        operator: NotIn
        values:
        - "false"
{{- else if .Values.revision }}
      matchExpressions:
      - key: "sidecar.istio.io/inject"
        operator: DoesNotExist
      - key: istio.io/rev
        operator: In
        values:
        - {{ .Values.revision }}
{{- else }}
      matchLabels:
        "sidecar.istio.io/inject": "true"
{{- end }}
{{- end }}
{{- end }}
`)

func chartsIstiodRemoteTemplatesMutatingwebhookYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteTemplatesMutatingwebhookYaml, nil
}

func chartsIstiodRemoteTemplatesMutatingwebhookYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteTemplatesMutatingwebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/templates/mutatingwebhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteTemplatesTelemetryv2_16Yaml = []byte(`{{- if and .Values.telemetry.enabled .Values.telemetry.v2.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: ANY # inbound, outbound, and gateway
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration: |
                  {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
{{- if .Values.telemetry.v2.prometheus.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "destination_cluster": "node.metadata['CLUSTER_ID']",
                          "source_cluster": "downstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio",
                    "disable_host_header_fallback": true{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "destination_cluster": "node.metadata['CLUSTER_ID']",
                          "source_cluster": "downstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: tcp_stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                  {
                    "debug": "false",
                    "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                    "metrics": [
                      {
                        "dimensions": {
                          "source_cluster": "node.metadata['CLUSTER_ID']",
                          "destination_cluster": "upstream_peer.cluster_id"
                        }
                      }
                    ]
                    {{- end }}
                  }
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
---

{{- end }}

{{- if .Values.telemetry.v2.stackdriver.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
{{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
{{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_inbound
                configuration: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stackdriver_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s", "disable_host_header_fallback": true}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
---
{{- if .Values.telemetry.v2.accessLogPolicy.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-sampling-accesslog-filter-1.6{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '1\.6.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "istio.stackdriver"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.access_log
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration: |
                  {
                    "log_window_duration": "{{ .Values.telemetry.v2.accessLogPolicy.logWindowDuration }}"
                  }
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: "envoy.wasm.access_log_policy" }
---
{{- end}}
{{- end}}
{{- end}}
`)

func chartsIstiodRemoteTemplatesTelemetryv2_16YamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteTemplatesTelemetryv2_16Yaml, nil
}

func chartsIstiodRemoteTemplatesTelemetryv2_16Yaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteTemplatesTelemetryv2_16YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/templates/telemetryv2_1.6.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteTemplatesTelemetryv2_17Yaml = []byte(`{{- if and .Values.telemetry.enabled .Values.telemetry.v2.enabled }}
# Note: metadata exchange filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Note: http stats filter is wasm enabled only in sidecars.
{{- if .Values.telemetry.v2.prometheus.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio",
                      "disable_host_header_fallback": true{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Note: tcp stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---

{{- end }}

{{- if .Values.telemetry.v2.stackdriver.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
{{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
{{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "meshEdgesReportingDuration": "600s", "disable_host_header_fallback": true}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
---

apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stackdriver-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
  {{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_OUTBOUND
      proxy:
        proxyVersion: '^1\.7.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
    {{- end }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_INBOUND
      proxy:
        proxyVersion: '^1\.7.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_inbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_inbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
  - applyTo: NETWORK_FILTER
    match:
      context: GATEWAY
      proxy:
        proxyVersion: '^1\.7.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
---
{{- if .Values.telemetry.v2.accessLogPolicy.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-sampling-accesslog-filter-1.7{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '1\.7.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "istio.stackdriver"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.access_log
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "log_window_duration": "{{ .Values.telemetry.v2.accessLogPolicy.logWindowDuration }}"
                    }
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: "envoy.wasm.access_log_policy" }
---
{{- end}}
{{- end}}
{{- end}}`)

func chartsIstiodRemoteTemplatesTelemetryv2_17YamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteTemplatesTelemetryv2_17Yaml, nil
}

func chartsIstiodRemoteTemplatesTelemetryv2_17Yaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteTemplatesTelemetryv2_17YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/templates/telemetryv2_1.7.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteTemplatesTelemetryv2_18Yaml = []byte(`{{- if and .Values.telemetry.enabled .Values.telemetry.v2.enabled }}
# Note: metadata exchange filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: metadata-exchange-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  {{- if .Values.telemetry.v2.metadataExchange.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/metadata-exchange-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {}
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.metadata_exchange
---
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-metadata-exchange-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener: {}
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.metadata_exchange
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
            value:
              protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
    - applyTo: CLUSTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        cluster: {}
      patch:
        operation: MERGE
        value:
          filters:
          - name: istio.metadata_exchange
            typed_config:
              "@type": type.googleapis.com/udpa.type.v1.TypedStruct
              type_url: type.googleapis.com/envoy.tcp.metadataexchange.config.MetadataExchange
              value:
                protocol: istio-peer-exchange
---
# Note: http stats filter is wasm enabled only in sidecars.
{{- if .Values.telemetry.v2.prometheus.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stats-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
                  {{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio",
                      "disable_host_header_fallback": true{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: envoy.wasm.stats
---
# Note: tcp stats filter is wasm enabled only in sidecars.
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stats-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.inboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "destination_cluster": "node.metadata['CLUSTER_ID']",
                            "source_cluster": "downstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.inboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_inbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.outboundSidecar }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.outboundSidecar | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  {{- if .Values.telemetry.v2.prometheus.wasmEnabled }}
                  runtime: envoy.wasm.runtime.v8
                  allow_precompiled: true
                  code:
                    local:
                      filename: /etc/istio/extensions/stats-filter.compiled.wasm
                  {{- else }}
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
                  {{- end }}
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.tcp_proxy"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stats
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
            value:
              config:
                root_id: stats_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.prometheus.configOverride.gateway }}
                    {
                      "debug": "false",
                      "stat_prefix": "istio"{{- if .Values.global.multiCluster.clusterName }},
                      "metrics": [
                        {
                          "dimensions": {
                            "source_cluster": "node.metadata['CLUSTER_ID']",
                            "destination_cluster": "upstream_peer.cluster_id"
                          }
                        }
                      ]
                      {{- end }}
                    }
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.prometheus.configOverride.gateway | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: tcp_stats_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local:
                      inline_string: "envoy.wasm.stats"
---

{{- end }}

{{- if .Values.telemetry.v2.stackdriver.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
{{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_OUTBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}", "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
{{- end }}
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_inbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.inboundAccessLogging }}", "meshEdgesReportingDuration": "600s"}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_inbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        proxy:
          proxyVersion: '^1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "envoy.router"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.stackdriver
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                root_id: stackdriver_outbound
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                    {"enable_mesh_edges_reporting": {{ .Values.telemetry.v2.stackdriver.topology }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}", "meshEdgesReportingDuration": "600s", "disable_host_header_fallback": true}
                    {{- else }}
                    {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                    {{- end }}
                vm_config:
                  vm_id: stackdriver_outbound
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: envoy.wasm.null.stackdriver }
---

apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: tcp-stackdriver-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
  {{- if not .Values.telemetry.v2.stackdriver.disableOutbound }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_OUTBOUND
      proxy:
        proxyVersion: '^1\.8.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
    {{- end }}
  - applyTo: NETWORK_FILTER
    match:
      context: SIDECAR_INBOUND
      proxy:
        proxyVersion: '^1\.8.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_inbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"disable_server_access_logging": {{ not .Values.telemetry.v2.stackdriver.logging }}, "access_logging": "{{ .Values.telemetry.v2.stackdriver.inboundAccessLogging }}"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_inbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
  - applyTo: NETWORK_FILTER
    match:
      context: GATEWAY
      proxy:
        proxyVersion: '^1\.8.*'
      listener:
        filterChain:
          filter:
            name: "envoy.tcp_proxy"
    patch:
      operation: INSERT_BEFORE
      value:
        name: istio.stackdriver
        typed_config:
          "@type": type.googleapis.com/udpa.type.v1.TypedStruct
          type_url: type.googleapis.com/envoy.extensions.filters.network.wasm.v3.Wasm
          value:
            config:
              root_id: stackdriver_outbound
              configuration:
                "@type": "type.googleapis.com/google.protobuf.StringValue"
                value: |
                  {{- if not .Values.telemetry.v2.stackdriver.configOverride }}
                  {"access_logging": "{{ .Values.telemetry.v2.stackdriver.outboundAccessLogging }}"}
                  {{- else }}
                  {{ toJson .Values.telemetry.v2.stackdriver.configOverride | indent 18 }}
                  {{- end }}
              vm_config:
                vm_id: stackdriver_outbound
                runtime: envoy.wasm.runtime.null
                code:
                  local: { inline_string: envoy.wasm.null.stackdriver }
---
{{- if .Values.telemetry.v2.accessLogPolicy.enabled }}
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: stackdriver-sampling-accesslog-filter-1.8{{- if not (eq .Values.revision "") }}-{{ .Values.revision }}{{- end }}
  {{- if .Values.meshConfig.rootNamespace }}
  namespace: {{ .Values.meshConfig.rootNamespace }}
  {{- else }}
  namespace: {{ .Release.Namespace }}
  {{- end }}
  labels:
    istio.io/rev: {{ .Values.revision | default "default" }}
spec:
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: SIDECAR_INBOUND
        proxy:
          proxyVersion: '1\.8.*'
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
              subFilter:
                name: "istio.stackdriver"
      patch:
        operation: INSERT_BEFORE
        value:
          name: istio.access_log
          typed_config:
            "@type": type.googleapis.com/udpa.type.v1.TypedStruct
            type_url: type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
            value:
              config:
                configuration:
                  "@type": "type.googleapis.com/google.protobuf.StringValue"
                  value: |
                    {
                      "log_window_duration": "{{ .Values.telemetry.v2.accessLogPolicy.logWindowDuration }}"
                    }
                vm_config:
                  runtime: envoy.wasm.runtime.null
                  code:
                    local: { inline_string: "envoy.wasm.access_log_policy" }
---
{{- end }}
{{- end }}
{{- end }}
`)

func chartsIstiodRemoteTemplatesTelemetryv2_18YamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteTemplatesTelemetryv2_18Yaml, nil
}

func chartsIstiodRemoteTemplatesTelemetryv2_18Yaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteTemplatesTelemetryv2_18YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/templates/telemetryv2_1.8.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chartsIstiodRemoteValuesYaml = []byte(`meshConfig:

  # Config for the default ProxyConfig.
  # Initially using directly the proxy metadata - can also be activated using annotations
  # on the pod. This is an unsupported low-level API, pending review and decisions on
  # enabling the feature. Enabling the DNS listener is safe - and allows further testing
  # and gradual adoption by setting capture only on specific workloads. It also allows
  # VMs to use other DNS options, like dnsmasq or unbound.
  defaultConfig:
    proxyMetadata:
      # If empty, agent will not start :15013 DNS listener and will not attempt
      # to connect to Istiod DNS-TLS. This will also disable the core dns sidecar in
      # istiod and the dns-over-tls listener.
      # DNS_AGENT: DNS-TLS
      DNS_AGENT: ""

  # The namespace to treat as the administrative root namespace for Istio configuration.
  # When processing a leaf namespace Istio will search for declarations in that namespace first
  # and if none are found it will search in the root namespace. Any matching declaration found in the root namespace
  # is processed as if it were declared in the leaf namespace.
  rootNamespace: "istio-system"

sidecarInjectorWebhook:
  # You can use the field called alwaysInjectSelector and neverInjectSelector which will always inject the sidecar or
  # always skip the injection on pods that match that label selector, regardless of the global policy.
  # See https://istio.io/docs/setup/kubernetes/additional-setup/sidecar-injection/#more-control-adding-exceptions
  neverInjectSelector: []
  alwaysInjectSelector: []

  # injectedAnnotations are additional annotations that will be added to the pod spec after injection
  # This is primarily to support PSP annotations. For example, if you defined a PSP with the annotations:
  #
  # annotations:
  #   apparmor.security.beta.kubernetes.io/allowedProfileNames: runtime/default
  #   apparmor.security.beta.kubernetes.io/defaultProfileName: runtime/default
  #
  # The PSP controller would add corresponding annotations to the pod spec for each container. However, this happens before
  # the inject adds additional containers, so we must specify them explicitly here. With the above example, we could specify:
  # injectedAnnotations:
  #   container.apparmor.security.beta.kubernetes.io/istio-init: runtime/default
  #   container.apparmor.security.beta.kubernetes.io/istio-proxy: runtime/default
  injectedAnnotations: {}

  # This enables injection of sidecar in all namespaces,
  # with the exception of namespaces with "istio-injection:disabled" annotation
  # Only one environment should have this enabled.
  enableNamespacesByDefault: false

  # Enable objectSelector to filter out pods with no need for sidecar before calling istio-sidecar-injector.
  # It is disabled by default since this function will only work after k8s v1.15.
  objectSelector:
    enabled: false
    autoInject: true

  # caBundle to be patched at runtime or user can specify it manually
  caBundle: ""

istiodRemote:
  # Sidecar injector mutating webhook configuration url
  # For example: https://$remotePilotAddress:15017/inject
  injectionURL: ""

# Revision is set as 'version' label and part of the resource names when installing multiple control planes.
revision: ""

telemetry:
  enabled: false
  v2:
    # This also enables metadata exchange.
    enabled: true
    metadataExchange:
      # Indicates whether to enable WebAssembly runtime for metadata exchange filter.
      wasmEnabled: false
    # Indicate if prometheus stats filter is enabled or not
    prometheus:
      enabled: true
      # Indicates whether to enable WebAssembly runtime for stats filter.
      wasmEnabled: false
      # overrides stats EnvoyFilter configuration. 
      configOverride:
        gateway: {}
        inboundSidecar: {}
        outboundSidecar: {}
    # stackdriver filter settings.
    stackdriver:
      enabled: false
      logging: false
      monitoring: false
      topology: false
      disableOutbound: false
      #  configOverride parts give you the ability to override the low level configuration params passed to envoy filter.

      configOverride: {}
      #  e.g.
      #  enable_mesh_edges_reporting: true
      #  disable_server_access_logging: false
      #  meshEdgesReportingDuration: 500s
      #  disable_host_header_fallback: true
    # Access Log Policy Filter Settings. This enables filtering of access logs from stackdriver.
    accessLogPolicy:
      enabled: false
      # To reduce the number of successful logs, default log window duration is
      # set to 12 hours.
      logWindowDuration: "43200s"
pilot:
  # Install the mesh config map, generated from values.yaml.
  # If false, pilot wil use default values (by default) or user-supplied values.
  configMap: true
global:

  # One central istiod controls all remote clusters: disabled by default
  centralIstiod: false
`)

func chartsIstiodRemoteValuesYamlBytes() ([]byte, error) {
	return _chartsIstiodRemoteValuesYaml, nil
}

func chartsIstiodRemoteValuesYaml() (*asset, error) {
	bytes, err := chartsIstiodRemoteValuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "charts/istiod-remote/values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesDefaultYaml = []byte(`apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
spec:
  hub: gcr.io/istio-testing
  tag: latest

  # You may override parts of meshconfig by uncommenting the following lines.
  meshConfig:
    defaultConfig:
      proxyMetadata: {}
    enablePrometheusMerge: true
    # Opt-out of global http2 upgrades.
    # Destination rule is used to opt-in.
    # h2_upgrade_policy: DO_NOT_UPGRADE

  # Traffic management feature
  components:
    base:
      enabled: true
    pilot:
      enabled: true
      k8s:
        env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 1
          periodSeconds: 3
          timeoutSeconds: 5
        strategy:
          rollingUpdate:
            maxSurge: "100%"
            maxUnavailable: "25%"

    # Istio Gateway feature
    ingressGateways:
    - name: istio-ingressgateway
      enabled: true
      k8s:
        env:
          - name: ISTIO_META_ROUTER_MODE
            value: "sni-dnat"
        service:
          ports:
            - port: 15021
              targetPort: 15021
              name: status-port
            - port: 80
              targetPort: 8080
              name: http2
            - port: 443
              targetPort: 8443
              name: https
            - port: 15012
              targetPort: 15012
              name: tcp-istiod
            - port: 15443
              targetPort: 15443
              name: tls
        hpaSpec:
          maxReplicas: 5
          minReplicas: 1
          scaleTargetRef:
            apiVersion: apps/v1
            kind: Deployment
            name: istio-ingressgateway
          metrics:
            - type: Resource
              resource:
                name: cpu
                targetAverageUtilization: 80
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 2000m
            memory: 1024Mi
        strategy:
          rollingUpdate:
            maxSurge: "100%"
            maxUnavailable: "25%"

    egressGateways:
    - name: istio-egressgateway
      enabled: false
      k8s:
        env:
          - name: ISTIO_META_ROUTER_MODE
            value: "sni-dnat"
        service:
          ports:
            - port: 80
              name: http2
              targetPort: 8080
            - port: 443
              name: https
              targetPort: 8443
            - port: 15443
              targetPort: 15443
              name: tls
        hpaSpec:
          maxReplicas: 5
          minReplicas: 1
          scaleTargetRef:
            apiVersion: apps/v1
            kind: Deployment
            name: istio-egressgateway
          metrics:
            - type: Resource
              resource:
                name: cpu
                targetAverageUtilization: 80
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 2000m
            memory: 1024Mi
        strategy:
          rollingUpdate:
            maxSurge: "100%"
            maxUnavailable: "25%"
    # Istio CNI feature
    cni:
      enabled: false
    
    # istiod remote configuration wwhen istiod isn't installed on the cluster
    istiodRemote:
      enabled: false

  addonComponents:
    istiocoredns:
      enabled: false

  # Global values passed through to helm global.yaml.
  # Please keep this in sync with manifests/charts/global.yaml
  values:
    global:
      istioNamespace: istio-system
      istiod:
        enableAnalysis: false
      logging:
        level: "default:info"
      logAsJson: false
      pilotCertProvider: istiod
      jwtPolicy: third-party-jwt
      proxy:
        image: proxyv2
        clusterDomain: "cluster.local"
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 2000m
            memory: 1024Mi
        logLevel: warning
        componentLogLevel: "misc:error"
        privileged: false
        enableCoreDump: false
        statusPort: 15020
        readinessInitialDelaySeconds: 1
        readinessPeriodSeconds: 2
        readinessFailureThreshold: 30
        includeIPRanges: "*"
        excludeIPRanges: ""
        excludeOutboundPorts: ""
        excludeInboundPorts: ""
        autoInject: enabled
        tracer: "zipkin"
      proxy_init:
        image: proxyv2
        resources:
          limits:
            cpu: 2000m
            memory: 1024Mi
          requests:
            cpu: 10m
            memory: 10Mi
      # Specify image pull policy if default behavior isn't desired.
      # Default behavior: latest images will be Always else IfNotPresent.
      imagePullPolicy: ""
      operatorManageWebhooks: false
      tracer:
        lightstep: {}
        zipkin: {}
        datadog: {}
        stackdriver: {}
      imagePullSecrets: []
      arch:
        amd64: 2
        s390x: 2
        ppc64le: 2
      oneNamespace: false
      defaultNodeSelector: {}
      configValidation: true
      meshExpansion:
        enabled: false
        useILB: false
      multiCluster:
        enabled: false
        clusterName: ""
      omitSidecarInjectorConfigMap: false
      network: ""
      defaultResources:
        requests:
          cpu: 10m
      defaultPodDisruptionBudget:
        enabled: true
      priorityClassName: ""
      useMCP: false
      trustDomain: "cluster.local"
      sds:
        token:
          aud: istio-ca
      sts:
        servicePort: 0
      meshNetworks: {}
      mountMtlsCerts: false
    base:
      enableCRDTemplates: false
      validationURL: ""
    pilot:
      autoscaleEnabled: true
      autoscaleMin: 1
      autoscaleMax: 5
      replicaCount: 1
      image: pilot
      traceSampling: 1.0
      env: {}
      cpu:
        targetAverageUtilization: 80
      nodeSelector: {}
      keepaliveMaxServerConnectionAge: 30m
      enableProtocolSniffingForOutbound: true
      enableProtocolSniffingForInbound: true
      deploymentLabels:
      configMap: true

    telemetry:
      enabled: true
      v2:
        enabled: true
        metadataExchange:
          wasmEnabled: false
        prometheus:
          wasmEnabled: false
          enabled: true
        stackdriver:
          enabled: false
          logging: false
          monitoring: false
          topology: false
          configOverride: {}

    istiodRemote:
      injectionURL: ""
      
    gateways:
      istio-egressgateway:
        zvpn: {}
        env: {}
        autoscaleEnabled: true
        type: ClusterIP
        name: istio-egressgateway
        secretVolumes:
          - name: egressgateway-certs
            secretName: istio-egressgateway-certs
            mountPath: /etc/istio/egressgateway-certs
          - name: egressgateway-ca-certs
            secretName: istio-egressgateway-ca-certs
            mountPath: /etc/istio/egressgateway-ca-certs

      istio-ingressgateway:
        autoscaleEnabled: true
        type: LoadBalancer
        name: istio-ingressgateway
        zvpn: {}
        env: {}
        secretVolumes:
          - name: ingressgateway-certs
            secretName: istio-ingressgateway-certs
            mountPath: /etc/istio/ingressgateway-certs
          - name: ingressgateway-ca-certs
            secretName: istio-ingressgateway-ca-certs
            mountPath: /etc/istio/ingressgateway-ca-certs

    sidecarInjectorWebhook:
      enableNamespacesByDefault: false
      rewriteAppHTTPProbe: true
      objectSelector:
        enabled: false
        autoInject: true

    istiocoredns:
      coreDNSImage: coredns/coredns
      coreDNSTag: 1.6.2
      coreDNSPluginImage: istio/coredns-plugin:0.2-istio-1.1

    clusterResources: true
`)

func profilesDefaultYamlBytes() ([]byte, error) {
	return _profilesDefaultYaml, nil
}

func profilesDefaultYaml() (*asset, error) {
	bytes, err := profilesDefaultYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles/default.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesDemoYaml = []byte(`apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  meshConfig:
    accessLogFile: /dev/stdout
  components:
    egressGateways:
    - name: istio-egressgateway
      enabled: true
      k8s:
        resources:
          requests:
            cpu: 10m
            memory: 40Mi

    ingressGateways:
    - name: istio-ingressgateway
      enabled: true
      k8s:
        resources:
          requests:
            cpu: 10m
            memory: 40Mi
        service:
          ports:
            ## You can add custom gateway ports in user values overrides, but it must include those ports since helm replaces.
            # Note that AWS ELB will by default perform health checks on the first port
            # on this list. Setting this to the health check port will ensure that health
            # checks always work. https://github.com/istio/istio/issues/12503
            - port: 15021
              targetPort: 15021
              name: status-port
            - port: 80
              targetPort: 8080
              name: http2
            - port: 443
              targetPort: 8443
              name: https
            - port: 31400
              targetPort: 31400
              name: tcp
              # This is the port where sni routing happens
            - port: 15443
              targetPort: 15443
              name: tls

    pilot:
      k8s:
        env:
          - name: PILOT_TRACE_SAMPLING
            value: "100"
        resources:
          requests:
            cpu: 10m
            memory: 100Mi

  values:
    global:
      proxy:
        resources:
          requests:
            cpu: 10m
            memory: 40Mi

    pilot:
      autoscaleEnabled: false

    gateways:
      istio-egressgateway:
        autoscaleEnabled: false
      istio-ingressgateway:
        autoscaleEnabled: false
`)

func profilesDemoYamlBytes() ([]byte, error) {
	return _profilesDemoYaml, nil
}

func profilesDemoYaml() (*asset, error) {
	bytes, err := profilesDemoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles/demo.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesEmptyYaml = []byte(`# The empty profile has everything disabled
# This is useful as a base for custom user configuration
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  components:
    base:
      enabled: false
    pilot:
      enabled: false
    ingressGateways:
    - name: istio-ingressgateway
      enabled: false
`)

func profilesEmptyYamlBytes() ([]byte, error) {
	return _profilesEmptyYaml, nil
}

func profilesEmptyYaml() (*asset, error) {
	bytes, err := profilesEmptyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles/empty.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesMinimalYaml = []byte(`# The minimal profile will install just the core control plane
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  components:
    ingressGateways:
    - name: istio-ingressgateway
      enabled: false
`)

func profilesMinimalYamlBytes() ([]byte, error) {
	return _profilesMinimalYaml, nil
}

func profilesMinimalYaml() (*asset, error) {
	bytes, err := profilesMinimalYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles/minimal.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesPreviewYaml = []byte(`# The preview profile contains features that are experimental.
# This is intended to explore new features coming to Istio.
# Stability, security, and performance are not guaranteed - use at your own risk.
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  meshConfig:
    defaultConfig:
      proxyMetadata:
        # Enable Istio agent to handle DNS requests for known hosts
        # Unknown hosts will automatically be resolved using upstream dns servers in resolv.conf
        ISTIO_META_DNS_CAPTURE: "ALL"
        # proxy envoy xds via istio agent (pre-req for dns)
        ISTIO_META_PROXY_XDS_VIA_AGENT: "enable"
  values:
    telemetry:
      v2:
        metadataExchange:
          wasmEnabled: true
        prometheus:
          wasmEnabled: true
`)

func profilesPreviewYamlBytes() ([]byte, error) {
	return _profilesPreviewYaml, nil
}

func profilesPreviewYaml() (*asset, error) {
	bytes, err := profilesPreviewYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles/preview.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesRemoteYaml = []byte(`apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec: {}
`)

func profilesRemoteYamlBytes() ([]byte, error) {
	return _profilesRemoteYaml, nil
}

func profilesRemoteYaml() (*asset, error) {
	bytes, err := profilesRemoteYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles/remote.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"charts/README-helm3.md":                                                        chartsReadmeHelm3Md,
	"charts/README.md":                                                              chartsReadmeMd,
	"charts/UPDATING-CHARTS.md":                                                     chartsUpdatingChartsMd,
	"charts/base/Chart.yaml":                                                        chartsBaseChartYaml,
	"charts/base/NOTES.txt":                                                         chartsBaseNotesTxt,
	"charts/base/crds/crd-all.gen.yaml":                                             chartsBaseCrdsCrdAllGenYaml,
	"charts/base/crds/crd-operator.yaml":                                            chartsBaseCrdsCrdOperatorYaml,
	"charts/base/files/gen-istio-cluster.yaml":                                      chartsBaseFilesGenIstioClusterYaml,
	"charts/base/kustomization.yaml":                                                chartsBaseKustomizationYaml,
	"charts/base/templates/clusterrole.yaml":                                        chartsBaseTemplatesClusterroleYaml,
	"charts/base/templates/clusterrolebinding.yaml":                                 chartsBaseTemplatesClusterrolebindingYaml,
	"charts/base/templates/crds.yaml":                                               chartsBaseTemplatesCrdsYaml,
	"charts/base/templates/endpoints.yaml":                                          chartsBaseTemplatesEndpointsYaml,
	"charts/base/templates/role.yaml":                                               chartsBaseTemplatesRoleYaml,
	"charts/base/templates/rolebinding.yaml":                                        chartsBaseTemplatesRolebindingYaml,
	"charts/base/templates/serviceaccount.yaml":                                     chartsBaseTemplatesServiceaccountYaml,
	"charts/base/templates/services.yaml":                                           chartsBaseTemplatesServicesYaml,
	"charts/base/templates/validatingwebhookconfiguration.yaml":                     chartsBaseTemplatesValidatingwebhookconfigurationYaml,
	"charts/base/values.yaml":                                                       chartsBaseValuesYaml,
	"charts/gateways/istio-egress/Chart.yaml":                                       chartsGatewaysIstioEgressChartYaml,
	"charts/gateways/istio-egress/NOTES.txt":                                        chartsGatewaysIstioEgressNotesTxt,
	"charts/gateways/istio-egress/templates/_affinity.tpl":                          chartsGatewaysIstioEgressTemplates_affinityTpl,
	"charts/gateways/istio-egress/templates/_helpers.tpl":                           chartsGatewaysIstioEgressTemplates_helpersTpl,
	"charts/gateways/istio-egress/templates/autoscale.yaml":                         chartsGatewaysIstioEgressTemplatesAutoscaleYaml,
	"charts/gateways/istio-egress/templates/deployment.yaml":                        chartsGatewaysIstioEgressTemplatesDeploymentYaml,
	"charts/gateways/istio-egress/templates/poddisruptionbudget.yaml":               chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYaml,
	"charts/gateways/istio-egress/templates/preconfigured.yaml":                     chartsGatewaysIstioEgressTemplatesPreconfiguredYaml,
	"charts/gateways/istio-egress/templates/role.yaml":                              chartsGatewaysIstioEgressTemplatesRoleYaml,
	"charts/gateways/istio-egress/templates/rolebindings.yaml":                      chartsGatewaysIstioEgressTemplatesRolebindingsYaml,
	"charts/gateways/istio-egress/templates/service.yaml":                           chartsGatewaysIstioEgressTemplatesServiceYaml,
	"charts/gateways/istio-egress/templates/serviceaccount.yaml":                    chartsGatewaysIstioEgressTemplatesServiceaccountYaml,
	"charts/gateways/istio-egress/values.yaml":                                      chartsGatewaysIstioEgressValuesYaml,
	"charts/gateways/istio-ingress/Chart.yaml":                                      chartsGatewaysIstioIngressChartYaml,
	"charts/gateways/istio-ingress/NOTES.txt":                                       chartsGatewaysIstioIngressNotesTxt,
	"charts/gateways/istio-ingress/templates/_affinity.tpl":                         chartsGatewaysIstioIngressTemplates_affinityTpl,
	"charts/gateways/istio-ingress/templates/autoscale.yaml":                        chartsGatewaysIstioIngressTemplatesAutoscaleYaml,
	"charts/gateways/istio-ingress/templates/deployment.yaml":                       chartsGatewaysIstioIngressTemplatesDeploymentYaml,
	"charts/gateways/istio-ingress/templates/meshexpansion.yaml":                    chartsGatewaysIstioIngressTemplatesMeshexpansionYaml,
	"charts/gateways/istio-ingress/templates/poddisruptionbudget.yaml":              chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYaml,
	"charts/gateways/istio-ingress/templates/preconfigured.yaml":                    chartsGatewaysIstioIngressTemplatesPreconfiguredYaml,
	"charts/gateways/istio-ingress/templates/role.yaml":                             chartsGatewaysIstioIngressTemplatesRoleYaml,
	"charts/gateways/istio-ingress/templates/rolebindings.yaml":                     chartsGatewaysIstioIngressTemplatesRolebindingsYaml,
	"charts/gateways/istio-ingress/templates/service.yaml":                          chartsGatewaysIstioIngressTemplatesServiceYaml,
	"charts/gateways/istio-ingress/templates/serviceaccount.yaml":                   chartsGatewaysIstioIngressTemplatesServiceaccountYaml,
	"charts/gateways/istio-ingress/values.yaml":                                     chartsGatewaysIstioIngressValuesYaml,
	"charts/global.yaml":                                                            chartsGlobalYaml,
	"charts/istio-cni/Chart.yaml":                                                   chartsIstioCniChartYaml,
	"charts/istio-cni/templates/clusterrole.yaml":                                   chartsIstioCniTemplatesClusterroleYaml,
	"charts/istio-cni/templates/clusterrolebinding.yaml":                            chartsIstioCniTemplatesClusterrolebindingYaml,
	"charts/istio-cni/templates/configmap-cni.yaml":                                 chartsIstioCniTemplatesConfigmapCniYaml,
	"charts/istio-cni/templates/daemonset.yaml":                                     chartsIstioCniTemplatesDaemonsetYaml,
	"charts/istio-cni/templates/serviceaccount.yaml":                                chartsIstioCniTemplatesServiceaccountYaml,
	"charts/istio-cni/values.yaml":                                                  chartsIstioCniValuesYaml,
	"charts/istio-control/istio-discovery/Chart.yaml":                               chartsIstioControlIstioDiscoveryChartYaml,
	"charts/istio-control/istio-discovery/NOTES.txt":                                chartsIstioControlIstioDiscoveryNotesTxt,
	"charts/istio-control/istio-discovery/files/gen-istio.yaml":                     chartsIstioControlIstioDiscoveryFilesGenIstioYaml,
	"charts/istio-control/istio-discovery/files/injection-template.yaml":            chartsIstioControlIstioDiscoveryFilesInjectionTemplateYaml,
	"charts/istio-control/istio-discovery/kustomization.yaml":                       chartsIstioControlIstioDiscoveryKustomizationYaml,
	"charts/istio-control/istio-discovery/templates/autoscale.yaml":                 chartsIstioControlIstioDiscoveryTemplatesAutoscaleYaml,
	"charts/istio-control/istio-discovery/templates/configmap-jwks.yaml":            chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYaml,
	"charts/istio-control/istio-discovery/templates/configmap.yaml":                 chartsIstioControlIstioDiscoveryTemplatesConfigmapYaml,
	"charts/istio-control/istio-discovery/templates/deployment.yaml":                chartsIstioControlIstioDiscoveryTemplatesDeploymentYaml,
	"charts/istio-control/istio-discovery/templates/istiod-injector-configmap.yaml": chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYaml,
	"charts/istio-control/istio-discovery/templates/mutatingwebhook.yaml":           chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYaml,
	"charts/istio-control/istio-discovery/templates/poddisruptionbudget.yaml":       chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYaml,
	"charts/istio-control/istio-discovery/templates/service.yaml":                   chartsIstioControlIstioDiscoveryTemplatesServiceYaml,
	"charts/istio-control/istio-discovery/templates/telemetryv2_1.6.yaml":           chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16Yaml,
	"charts/istio-control/istio-discovery/templates/telemetryv2_1.7.yaml":           chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17Yaml,
	"charts/istio-control/istio-discovery/templates/telemetryv2_1.8.yaml":           chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18Yaml,
	"charts/istio-control/istio-discovery/values.yaml":                              chartsIstioControlIstioDiscoveryValuesYaml,
	"charts/istio-operator/Chart.yaml":                                              chartsIstioOperatorChartYaml,
	"charts/istio-operator/crds/crd-operator.yaml":                                  chartsIstioOperatorCrdsCrdOperatorYaml,
	"charts/istio-operator/files/gen-operator.yaml":                                 chartsIstioOperatorFilesGenOperatorYaml,
	"charts/istio-operator/templates/clusterrole.yaml":                              chartsIstioOperatorTemplatesClusterroleYaml,
	"charts/istio-operator/templates/clusterrole_binding.yaml":                      chartsIstioOperatorTemplatesClusterrole_bindingYaml,
	"charts/istio-operator/templates/crds.yaml":                                     chartsIstioOperatorTemplatesCrdsYaml,
	"charts/istio-operator/templates/deployment.yaml":                               chartsIstioOperatorTemplatesDeploymentYaml,
	"charts/istio-operator/templates/namespace.yaml":                                chartsIstioOperatorTemplatesNamespaceYaml,
	"charts/istio-operator/templates/service.yaml":                                  chartsIstioOperatorTemplatesServiceYaml,
	"charts/istio-operator/templates/service_account.yaml":                          chartsIstioOperatorTemplatesService_accountYaml,
	"charts/istio-operator/values.yaml":                                             chartsIstioOperatorValuesYaml,
	"charts/istiocoredns/Chart.yaml":                                                chartsIstiocorednsChartYaml,
	"charts/istiocoredns/templates/_affinity.tpl":                                   chartsIstiocorednsTemplates_affinityTpl,
	"charts/istiocoredns/templates/autoscale.yaml":                                  chartsIstiocorednsTemplatesAutoscaleYaml,
	"charts/istiocoredns/templates/clusterrole.yaml":                                chartsIstiocorednsTemplatesClusterroleYaml,
	"charts/istiocoredns/templates/clusterrolebinding.yaml":                         chartsIstiocorednsTemplatesClusterrolebindingYaml,
	"charts/istiocoredns/templates/configmap.yaml":                                  chartsIstiocorednsTemplatesConfigmapYaml,
	"charts/istiocoredns/templates/deployment.yaml":                                 chartsIstiocorednsTemplatesDeploymentYaml,
	"charts/istiocoredns/templates/poddisruptionbudget.yaml":                        chartsIstiocorednsTemplatesPoddisruptionbudgetYaml,
	"charts/istiocoredns/templates/service.yaml":                                    chartsIstiocorednsTemplatesServiceYaml,
	"charts/istiocoredns/templates/serviceaccount.yaml":                             chartsIstiocorednsTemplatesServiceaccountYaml,
	"charts/istiocoredns/values.yaml":                                               chartsIstiocorednsValuesYaml,
	"charts/istiod-remote/Chart.yaml":                                               chartsIstiodRemoteChartYaml,
	"charts/istiod-remote/NOTES.txt":                                                chartsIstiodRemoteNotesTxt,
	"charts/istiod-remote/files/gen-istiod-remote.yaml":                             chartsIstiodRemoteFilesGenIstiodRemoteYaml,
	"charts/istiod-remote/files/injection-template.yaml":                            chartsIstiodRemoteFilesInjectionTemplateYaml,
	"charts/istiod-remote/kustomization.yaml":                                       chartsIstiodRemoteKustomizationYaml,
	"charts/istiod-remote/templates/configmap.yaml":                                 chartsIstiodRemoteTemplatesConfigmapYaml,
	"charts/istiod-remote/templates/istiod-injector-configmap.yaml":                 chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYaml,
	"charts/istiod-remote/templates/mutatingwebhook.yaml":                           chartsIstiodRemoteTemplatesMutatingwebhookYaml,
	"charts/istiod-remote/templates/telemetryv2_1.6.yaml":                           chartsIstiodRemoteTemplatesTelemetryv2_16Yaml,
	"charts/istiod-remote/templates/telemetryv2_1.7.yaml":                           chartsIstiodRemoteTemplatesTelemetryv2_17Yaml,
	"charts/istiod-remote/templates/telemetryv2_1.8.yaml":                           chartsIstiodRemoteTemplatesTelemetryv2_18Yaml,
	"charts/istiod-remote/values.yaml":                                              chartsIstiodRemoteValuesYaml,
	"profiles/default.yaml":                                                         profilesDefaultYaml,
	"profiles/demo.yaml":                                                            profilesDemoYaml,
	"profiles/empty.yaml":                                                           profilesEmptyYaml,
	"profiles/minimal.yaml":                                                         profilesMinimalYaml,
	"profiles/preview.yaml":                                                         profilesPreviewYaml,
	"profiles/remote.yaml":                                                          profilesRemoteYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"charts": &bintree{nil, map[string]*bintree{
		"README-helm3.md":    &bintree{chartsReadmeHelm3Md, map[string]*bintree{}},
		"README.md":          &bintree{chartsReadmeMd, map[string]*bintree{}},
		"UPDATING-CHARTS.md": &bintree{chartsUpdatingChartsMd, map[string]*bintree{}},
		"base": &bintree{nil, map[string]*bintree{
			"Chart.yaml": &bintree{chartsBaseChartYaml, map[string]*bintree{}},
			"NOTES.txt":  &bintree{chartsBaseNotesTxt, map[string]*bintree{}},
			"crds": &bintree{nil, map[string]*bintree{
				"crd-all.gen.yaml":  &bintree{chartsBaseCrdsCrdAllGenYaml, map[string]*bintree{}},
				"crd-operator.yaml": &bintree{chartsBaseCrdsCrdOperatorYaml, map[string]*bintree{}},
			}},
			"files": &bintree{nil, map[string]*bintree{
				"gen-istio-cluster.yaml": &bintree{chartsBaseFilesGenIstioClusterYaml, map[string]*bintree{}},
			}},
			"kustomization.yaml": &bintree{chartsBaseKustomizationYaml, map[string]*bintree{}},
			"templates": &bintree{nil, map[string]*bintree{
				"clusterrole.yaml":                    &bintree{chartsBaseTemplatesClusterroleYaml, map[string]*bintree{}},
				"clusterrolebinding.yaml":             &bintree{chartsBaseTemplatesClusterrolebindingYaml, map[string]*bintree{}},
				"crds.yaml":                           &bintree{chartsBaseTemplatesCrdsYaml, map[string]*bintree{}},
				"endpoints.yaml":                      &bintree{chartsBaseTemplatesEndpointsYaml, map[string]*bintree{}},
				"role.yaml":                           &bintree{chartsBaseTemplatesRoleYaml, map[string]*bintree{}},
				"rolebinding.yaml":                    &bintree{chartsBaseTemplatesRolebindingYaml, map[string]*bintree{}},
				"serviceaccount.yaml":                 &bintree{chartsBaseTemplatesServiceaccountYaml, map[string]*bintree{}},
				"services.yaml":                       &bintree{chartsBaseTemplatesServicesYaml, map[string]*bintree{}},
				"validatingwebhookconfiguration.yaml": &bintree{chartsBaseTemplatesValidatingwebhookconfigurationYaml, map[string]*bintree{}},
			}},
			"values.yaml": &bintree{chartsBaseValuesYaml, map[string]*bintree{}},
		}},
		"gateways": &bintree{nil, map[string]*bintree{
			"istio-egress": &bintree{nil, map[string]*bintree{
				"Chart.yaml": &bintree{chartsGatewaysIstioEgressChartYaml, map[string]*bintree{}},
				"NOTES.txt":  &bintree{chartsGatewaysIstioEgressNotesTxt, map[string]*bintree{}},
				"templates": &bintree{nil, map[string]*bintree{
					"_affinity.tpl":            &bintree{chartsGatewaysIstioEgressTemplates_affinityTpl, map[string]*bintree{}},
					"_helpers.tpl":             &bintree{chartsGatewaysIstioEgressTemplates_helpersTpl, map[string]*bintree{}},
					"autoscale.yaml":           &bintree{chartsGatewaysIstioEgressTemplatesAutoscaleYaml, map[string]*bintree{}},
					"deployment.yaml":          &bintree{chartsGatewaysIstioEgressTemplatesDeploymentYaml, map[string]*bintree{}},
					"poddisruptionbudget.yaml": &bintree{chartsGatewaysIstioEgressTemplatesPoddisruptionbudgetYaml, map[string]*bintree{}},
					"preconfigured.yaml":       &bintree{chartsGatewaysIstioEgressTemplatesPreconfiguredYaml, map[string]*bintree{}},
					"role.yaml":                &bintree{chartsGatewaysIstioEgressTemplatesRoleYaml, map[string]*bintree{}},
					"rolebindings.yaml":        &bintree{chartsGatewaysIstioEgressTemplatesRolebindingsYaml, map[string]*bintree{}},
					"service.yaml":             &bintree{chartsGatewaysIstioEgressTemplatesServiceYaml, map[string]*bintree{}},
					"serviceaccount.yaml":      &bintree{chartsGatewaysIstioEgressTemplatesServiceaccountYaml, map[string]*bintree{}},
				}},
				"values.yaml": &bintree{chartsGatewaysIstioEgressValuesYaml, map[string]*bintree{}},
			}},
			"istio-ingress": &bintree{nil, map[string]*bintree{
				"Chart.yaml": &bintree{chartsGatewaysIstioIngressChartYaml, map[string]*bintree{}},
				"NOTES.txt":  &bintree{chartsGatewaysIstioIngressNotesTxt, map[string]*bintree{}},
				"templates": &bintree{nil, map[string]*bintree{
					"_affinity.tpl":            &bintree{chartsGatewaysIstioIngressTemplates_affinityTpl, map[string]*bintree{}},
					"autoscale.yaml":           &bintree{chartsGatewaysIstioIngressTemplatesAutoscaleYaml, map[string]*bintree{}},
					"deployment.yaml":          &bintree{chartsGatewaysIstioIngressTemplatesDeploymentYaml, map[string]*bintree{}},
					"meshexpansion.yaml":       &bintree{chartsGatewaysIstioIngressTemplatesMeshexpansionYaml, map[string]*bintree{}},
					"poddisruptionbudget.yaml": &bintree{chartsGatewaysIstioIngressTemplatesPoddisruptionbudgetYaml, map[string]*bintree{}},
					"preconfigured.yaml":       &bintree{chartsGatewaysIstioIngressTemplatesPreconfiguredYaml, map[string]*bintree{}},
					"role.yaml":                &bintree{chartsGatewaysIstioIngressTemplatesRoleYaml, map[string]*bintree{}},
					"rolebindings.yaml":        &bintree{chartsGatewaysIstioIngressTemplatesRolebindingsYaml, map[string]*bintree{}},
					"service.yaml":             &bintree{chartsGatewaysIstioIngressTemplatesServiceYaml, map[string]*bintree{}},
					"serviceaccount.yaml":      &bintree{chartsGatewaysIstioIngressTemplatesServiceaccountYaml, map[string]*bintree{}},
				}},
				"values.yaml": &bintree{chartsGatewaysIstioIngressValuesYaml, map[string]*bintree{}},
			}},
		}},
		"global.yaml": &bintree{chartsGlobalYaml, map[string]*bintree{}},
		"istio-cni": &bintree{nil, map[string]*bintree{
			"Chart.yaml": &bintree{chartsIstioCniChartYaml, map[string]*bintree{}},
			"templates": &bintree{nil, map[string]*bintree{
				"clusterrole.yaml":        &bintree{chartsIstioCniTemplatesClusterroleYaml, map[string]*bintree{}},
				"clusterrolebinding.yaml": &bintree{chartsIstioCniTemplatesClusterrolebindingYaml, map[string]*bintree{}},
				"configmap-cni.yaml":      &bintree{chartsIstioCniTemplatesConfigmapCniYaml, map[string]*bintree{}},
				"daemonset.yaml":          &bintree{chartsIstioCniTemplatesDaemonsetYaml, map[string]*bintree{}},
				"serviceaccount.yaml":     &bintree{chartsIstioCniTemplatesServiceaccountYaml, map[string]*bintree{}},
			}},
			"values.yaml": &bintree{chartsIstioCniValuesYaml, map[string]*bintree{}},
		}},
		"istio-control": &bintree{nil, map[string]*bintree{
			"istio-discovery": &bintree{nil, map[string]*bintree{
				"Chart.yaml": &bintree{chartsIstioControlIstioDiscoveryChartYaml, map[string]*bintree{}},
				"NOTES.txt":  &bintree{chartsIstioControlIstioDiscoveryNotesTxt, map[string]*bintree{}},
				"files": &bintree{nil, map[string]*bintree{
					"gen-istio.yaml":          &bintree{chartsIstioControlIstioDiscoveryFilesGenIstioYaml, map[string]*bintree{}},
					"injection-template.yaml": &bintree{chartsIstioControlIstioDiscoveryFilesInjectionTemplateYaml, map[string]*bintree{}},
				}},
				"kustomization.yaml": &bintree{chartsIstioControlIstioDiscoveryKustomizationYaml, map[string]*bintree{}},
				"templates": &bintree{nil, map[string]*bintree{
					"autoscale.yaml":                 &bintree{chartsIstioControlIstioDiscoveryTemplatesAutoscaleYaml, map[string]*bintree{}},
					"configmap-jwks.yaml":            &bintree{chartsIstioControlIstioDiscoveryTemplatesConfigmapJwksYaml, map[string]*bintree{}},
					"configmap.yaml":                 &bintree{chartsIstioControlIstioDiscoveryTemplatesConfigmapYaml, map[string]*bintree{}},
					"deployment.yaml":                &bintree{chartsIstioControlIstioDiscoveryTemplatesDeploymentYaml, map[string]*bintree{}},
					"istiod-injector-configmap.yaml": &bintree{chartsIstioControlIstioDiscoveryTemplatesIstiodInjectorConfigmapYaml, map[string]*bintree{}},
					"mutatingwebhook.yaml":           &bintree{chartsIstioControlIstioDiscoveryTemplatesMutatingwebhookYaml, map[string]*bintree{}},
					"poddisruptionbudget.yaml":       &bintree{chartsIstioControlIstioDiscoveryTemplatesPoddisruptionbudgetYaml, map[string]*bintree{}},
					"service.yaml":                   &bintree{chartsIstioControlIstioDiscoveryTemplatesServiceYaml, map[string]*bintree{}},
					"telemetryv2_1.6.yaml":           &bintree{chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_16Yaml, map[string]*bintree{}},
					"telemetryv2_1.7.yaml":           &bintree{chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_17Yaml, map[string]*bintree{}},
					"telemetryv2_1.8.yaml":           &bintree{chartsIstioControlIstioDiscoveryTemplatesTelemetryv2_18Yaml, map[string]*bintree{}},
				}},
				"values.yaml": &bintree{chartsIstioControlIstioDiscoveryValuesYaml, map[string]*bintree{}},
			}},
		}},
		"istio-operator": &bintree{nil, map[string]*bintree{
			"Chart.yaml": &bintree{chartsIstioOperatorChartYaml, map[string]*bintree{}},
			"crds": &bintree{nil, map[string]*bintree{
				"crd-operator.yaml": &bintree{chartsIstioOperatorCrdsCrdOperatorYaml, map[string]*bintree{}},
			}},
			"files": &bintree{nil, map[string]*bintree{
				"gen-operator.yaml": &bintree{chartsIstioOperatorFilesGenOperatorYaml, map[string]*bintree{}},
			}},
			"templates": &bintree{nil, map[string]*bintree{
				"clusterrole.yaml":         &bintree{chartsIstioOperatorTemplatesClusterroleYaml, map[string]*bintree{}},
				"clusterrole_binding.yaml": &bintree{chartsIstioOperatorTemplatesClusterrole_bindingYaml, map[string]*bintree{}},
				"crds.yaml":                &bintree{chartsIstioOperatorTemplatesCrdsYaml, map[string]*bintree{}},
				"deployment.yaml":          &bintree{chartsIstioOperatorTemplatesDeploymentYaml, map[string]*bintree{}},
				"namespace.yaml":           &bintree{chartsIstioOperatorTemplatesNamespaceYaml, map[string]*bintree{}},
				"service.yaml":             &bintree{chartsIstioOperatorTemplatesServiceYaml, map[string]*bintree{}},
				"service_account.yaml":     &bintree{chartsIstioOperatorTemplatesService_accountYaml, map[string]*bintree{}},
			}},
			"values.yaml": &bintree{chartsIstioOperatorValuesYaml, map[string]*bintree{}},
		}},
		"istiocoredns": &bintree{nil, map[string]*bintree{
			"Chart.yaml": &bintree{chartsIstiocorednsChartYaml, map[string]*bintree{}},
			"templates": &bintree{nil, map[string]*bintree{
				"_affinity.tpl":            &bintree{chartsIstiocorednsTemplates_affinityTpl, map[string]*bintree{}},
				"autoscale.yaml":           &bintree{chartsIstiocorednsTemplatesAutoscaleYaml, map[string]*bintree{}},
				"clusterrole.yaml":         &bintree{chartsIstiocorednsTemplatesClusterroleYaml, map[string]*bintree{}},
				"clusterrolebinding.yaml":  &bintree{chartsIstiocorednsTemplatesClusterrolebindingYaml, map[string]*bintree{}},
				"configmap.yaml":           &bintree{chartsIstiocorednsTemplatesConfigmapYaml, map[string]*bintree{}},
				"deployment.yaml":          &bintree{chartsIstiocorednsTemplatesDeploymentYaml, map[string]*bintree{}},
				"poddisruptionbudget.yaml": &bintree{chartsIstiocorednsTemplatesPoddisruptionbudgetYaml, map[string]*bintree{}},
				"service.yaml":             &bintree{chartsIstiocorednsTemplatesServiceYaml, map[string]*bintree{}},
				"serviceaccount.yaml":      &bintree{chartsIstiocorednsTemplatesServiceaccountYaml, map[string]*bintree{}},
			}},
			"values.yaml": &bintree{chartsIstiocorednsValuesYaml, map[string]*bintree{}},
		}},
		"istiod-remote": &bintree{nil, map[string]*bintree{
			"Chart.yaml": &bintree{chartsIstiodRemoteChartYaml, map[string]*bintree{}},
			"NOTES.txt":  &bintree{chartsIstiodRemoteNotesTxt, map[string]*bintree{}},
			"files": &bintree{nil, map[string]*bintree{
				"gen-istiod-remote.yaml":  &bintree{chartsIstiodRemoteFilesGenIstiodRemoteYaml, map[string]*bintree{}},
				"injection-template.yaml": &bintree{chartsIstiodRemoteFilesInjectionTemplateYaml, map[string]*bintree{}},
			}},
			"kustomization.yaml": &bintree{chartsIstiodRemoteKustomizationYaml, map[string]*bintree{}},
			"templates": &bintree{nil, map[string]*bintree{
				"configmap.yaml":                 &bintree{chartsIstiodRemoteTemplatesConfigmapYaml, map[string]*bintree{}},
				"istiod-injector-configmap.yaml": &bintree{chartsIstiodRemoteTemplatesIstiodInjectorConfigmapYaml, map[string]*bintree{}},
				"mutatingwebhook.yaml":           &bintree{chartsIstiodRemoteTemplatesMutatingwebhookYaml, map[string]*bintree{}},
				"telemetryv2_1.6.yaml":           &bintree{chartsIstiodRemoteTemplatesTelemetryv2_16Yaml, map[string]*bintree{}},
				"telemetryv2_1.7.yaml":           &bintree{chartsIstiodRemoteTemplatesTelemetryv2_17Yaml, map[string]*bintree{}},
				"telemetryv2_1.8.yaml":           &bintree{chartsIstiodRemoteTemplatesTelemetryv2_18Yaml, map[string]*bintree{}},
			}},
			"values.yaml": &bintree{chartsIstiodRemoteValuesYaml, map[string]*bintree{}},
		}},
	}},
	"profiles": &bintree{nil, map[string]*bintree{
		"default.yaml": &bintree{profilesDefaultYaml, map[string]*bintree{}},
		"demo.yaml":    &bintree{profilesDemoYaml, map[string]*bintree{}},
		"empty.yaml":   &bintree{profilesEmptyYaml, map[string]*bintree{}},
		"minimal.yaml": &bintree{profilesMinimalYaml, map[string]*bintree{}},
		"preview.yaml": &bintree{profilesPreviewYaml, map[string]*bintree{}},
		"remote.yaml":  &bintree{profilesRemoteYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
