// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/core/0000_00_cluster-version-operator_00_namespace.yaml
// assets/core/0000_50_cluster-openshift-controller-manager_00_namespace.yaml
// assets/core/0000_50_service-ca-operator_01_namespace.yaml
// assets/core/0000_50_service-ca-operator_02_service.yaml
// assets/core/0000_50_service-ca-operator_03_cm.yaml
// assets/core/0000_50_service-ca-operator_04_sa.yaml
// assets/core/0000_60_service-ca_01_namespace.yaml
// assets/core/0000_60_service-ca_04_sa.yaml
// assets/core/0000_70_dns-operator_00-namespace.yaml
// assets/core/0000_70_dns-operator_01-service-account.yaml
// assets/core/0000_70_dns-operator_01-service.yaml
// assets/core/0000_70_dns_00-namespace.yaml
// assets/core/0000_70_dns_01-configmap.yaml
// assets/core/0000_70_dns_01-service-account.yaml
// assets/core/0000_70_dns_01-service.yaml
// assets/core/0000_80_hostpath-provisioner-namespace.yaml
// assets/core/0000_80_hostpath-provisioner-serviceaccount.yaml
// assets/core/0000_80_openshift-router-cm.yaml
// assets/core/0000_80_openshift-router-namespace.yaml
// assets/core/0000_80_openshift-router-service-account.yaml
// assets/core/0000_80_openshift-router-service.yaml
// assets/core/0001_00_cluster-version-operator_03_service.yaml
// assets/core/openshift-sdn-cm.yaml
package assets

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

var _assetsCore0000_00_clusterVersionOperator_00_namespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-cluster-version
  annotations:
    openshift.io/node-selector: ""
  labels:
    name: openshift-cluster-version
    openshift.io/run-level: "1"
    openshift.io/cluster-monitoring: "true"
`)

func assetsCore0000_00_clusterVersionOperator_00_namespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_00_clusterVersionOperator_00_namespaceYaml, nil
}

func assetsCore0000_00_clusterVersionOperator_00_namespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_00_clusterVersionOperator_00_namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_00_cluster-version-operator_00_namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  annotations:
    include.release.openshift.io/self-managed-high-availability: "true"
    openshift.io/node-selector: ""
  labels:
    openshift.io/cluster-monitoring: "true"
  name: openshift-controller-manager
`)

func assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYaml, nil
}

func assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_50_cluster-openshift-controller-manager_00_namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_50_serviceCaOperator_01_namespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  labels:
    openshift.io/run-level: "1"
    openshift.io/cluster-monitoring: "true"
  name: openshift-service-ca-operator
  annotations:
    openshift.io/node-selector: ""
`)

func assetsCore0000_50_serviceCaOperator_01_namespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_50_serviceCaOperator_01_namespaceYaml, nil
}

func assetsCore0000_50_serviceCaOperator_01_namespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_50_serviceCaOperator_01_namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_50_service-ca-operator_01_namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_50_serviceCaOperator_02_serviceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  annotations:
    service.beta.openshift.io/serving-cert-secret-name: serving-cert
  labels:
    app: service-ca-operator
  name: metrics
  namespace: openshift-service-ca-operator
spec:
  ports:
  - name: https
    port: 443
    protocol: TCP
    targetPort: 8443
  selector:
    app: service-ca-operator
  sessionAffinity: None
  type: ClusterIP
`)

func assetsCore0000_50_serviceCaOperator_02_serviceYamlBytes() ([]byte, error) {
	return _assetsCore0000_50_serviceCaOperator_02_serviceYaml, nil
}

func assetsCore0000_50_serviceCaOperator_02_serviceYaml() (*asset, error) {
	bytes, err := assetsCore0000_50_serviceCaOperator_02_serviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_50_service-ca-operator_02_service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_50_serviceCaOperator_03_cmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-ca-operator
  name: service-ca-operator-config
data:
  operator-config.yaml: |
    apiVersion: operator.openshift.io/v1alpha1
    kind: GenericOperatorConfig
`)

func assetsCore0000_50_serviceCaOperator_03_cmYamlBytes() ([]byte, error) {
	return _assetsCore0000_50_serviceCaOperator_03_cmYaml, nil
}

func assetsCore0000_50_serviceCaOperator_03_cmYaml() (*asset, error) {
	bytes, err := assetsCore0000_50_serviceCaOperator_03_cmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_50_service-ca-operator_03_cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_50_serviceCaOperator_04_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-service-ca-operator
  name: service-ca-operator
  labels:
    app: service-ca-operator
`)

func assetsCore0000_50_serviceCaOperator_04_saYamlBytes() ([]byte, error) {
	return _assetsCore0000_50_serviceCaOperator_04_saYaml, nil
}

func assetsCore0000_50_serviceCaOperator_04_saYaml() (*asset, error) {
	bytes, err := assetsCore0000_50_serviceCaOperator_04_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_50_service-ca-operator_04_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_60_serviceCa_01_namespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  labels:
    openshift.io/run-level: "1"
    openshift.io/cluster-monitoring: "true"
  name: openshift-service-ca
`)

func assetsCore0000_60_serviceCa_01_namespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_60_serviceCa_01_namespaceYaml, nil
}

func assetsCore0000_60_serviceCa_01_namespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_60_serviceCa_01_namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_60_service-ca_01_namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_60_serviceCa_04_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-service-ca
  name: service-ca
  labels:
    app: service-ca
`)

func assetsCore0000_60_serviceCa_04_saYamlBytes() ([]byte, error) {
	return _assetsCore0000_60_serviceCa_04_saYaml, nil
}

func assetsCore0000_60_serviceCa_04_saYaml() (*asset, error) {
	bytes, err := assetsCore0000_60_serviceCa_04_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_60_service-ca_04_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dnsOperator_00NamespaceYaml = []byte(`kind: Namespace
apiVersion: v1
metadata:
  annotations:
    openshift.io/node-selector: ""
  name: openshift-dns-operator
  labels:
    # set value to avoid depending on kube admission that depends on openshift apis
    openshift.io/run-level: "0"
    # allow openshift-monitoring to look for ServiceMonitor objects in this namespace
    openshift.io/cluster-monitoring: "true"
`)

func assetsCore0000_70_dnsOperator_00NamespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dnsOperator_00NamespaceYaml, nil
}

func assetsCore0000_70_dnsOperator_00NamespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dnsOperator_00NamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns-operator_00-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dnsOperator_01ServiceAccountYaml = []byte(`# Account for the operator itself. It should require namespace scoped
# permissions.
kind: ServiceAccount
apiVersion: v1
metadata:
  name: dns-operator
  namespace: openshift-dns-operator
`)

func assetsCore0000_70_dnsOperator_01ServiceAccountYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dnsOperator_01ServiceAccountYaml, nil
}

func assetsCore0000_70_dnsOperator_01ServiceAccountYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dnsOperator_01ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns-operator_01-service-account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dnsOperator_01ServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  annotations:
    service.beta.openshift.io/serving-cert-secret-name: metrics-tls
  labels:
    name: dns-operator
  name: metrics
  namespace: openshift-dns-operator
spec:
  ports:
  - name: metrics
    port: 9393
    targetPort: metrics
  selector:
    name: dns-operator
  type: ClusterIP
`)

func assetsCore0000_70_dnsOperator_01ServiceYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dnsOperator_01ServiceYaml, nil
}

func assetsCore0000_70_dnsOperator_01ServiceYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dnsOperator_01ServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns-operator_01-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dns_00NamespaceYaml = []byte(`kind: Namespace
apiVersion: v1
metadata:
  annotations:
    openshift.io/node-selector: ""
  name: openshift-dns
  labels:
    # set value to avoid depending on kube admission that depends on openshift apis
    openshift.io/run-level: "0"
    # allow openshift-monitoring to look for ServiceMonitor objects in this namespace
    openshift.io/cluster-monitoring: "true"
`)

func assetsCore0000_70_dns_00NamespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dns_00NamespaceYaml, nil
}

func assetsCore0000_70_dns_00NamespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dns_00NamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns_00-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dns_01ConfigmapYaml = []byte(`apiVersion: v1
data:
  Corefile: |
    .:5353 {
        errors
        health
        kubernetes cluster.local in-addr.arpa ip6.arpa {
            pods insecure
            upstream
            fallthrough in-addr.arpa ip6.arpa
        }
        prometheus :9153
        forward . /etc/resolv.conf {
            policy sequential
        }
        cache 30
        reload
    }
kind: ConfigMap
metadata:
  labels:
    dns.operator.openshift.io/owning-dns: default
  name: dns-default
  namespace: openshift-dns
`)

func assetsCore0000_70_dns_01ConfigmapYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dns_01ConfigmapYaml, nil
}

func assetsCore0000_70_dns_01ConfigmapYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dns_01ConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns_01-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dns_01ServiceAccountYaml = []byte(`kind: ServiceAccount
apiVersion: v1
metadata:
  name: dns
  namespace: openshift-dns
`)

func assetsCore0000_70_dns_01ServiceAccountYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dns_01ServiceAccountYaml, nil
}

func assetsCore0000_70_dns_01ServiceAccountYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dns_01ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns_01-service-account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_70_dns_01ServiceYaml = []byte(`kind: Service
apiVersion: v1
metadata:
  annotations:
    service.beta.openshift.io/serving-cert-secret-name: dns-default-metrics-tls
  labels:
      dns.operator.openshift.io/owning-dns: default
  name: dns-default
  namespace: openshift-dns          
# name, namespace,labels and annotations are set at runtime
spec:
  clusterIP: {{.ClusterIP}}
  ports:
  - name: dns
    port: 53
    targetPort: dns
    protocol: UDP
  - name: dns-tcp
    port: 53
    targetPort: dns-tcp
    protocol: TCP
  - name: metrics
    port: 9154
    targetPort: metrics
    protocol: TCP
  selector:
    dns.operator.openshift.io/daemonset-dns: default    
`)

func assetsCore0000_70_dns_01ServiceYamlBytes() ([]byte, error) {
	return _assetsCore0000_70_dns_01ServiceYaml, nil
}

func assetsCore0000_70_dns_01ServiceYaml() (*asset, error) {
	bytes, err := assetsCore0000_70_dns_01ServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_70_dns_01-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_80_hostpathProvisionerNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: kubevirt-hostpath-provisioner`)

func assetsCore0000_80_hostpathProvisionerNamespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_80_hostpathProvisionerNamespaceYaml, nil
}

func assetsCore0000_80_hostpathProvisionerNamespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_80_hostpathProvisionerNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_80_hostpath-provisioner-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_80_hostpathProvisionerServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: kubevirt-hostpath-provisioner-admin
  namespace: kubevirt-hostpath-provisioner`)

func assetsCore0000_80_hostpathProvisionerServiceaccountYamlBytes() ([]byte, error) {
	return _assetsCore0000_80_hostpathProvisionerServiceaccountYaml, nil
}

func assetsCore0000_80_hostpathProvisionerServiceaccountYaml() (*asset, error) {
	bytes, err := assetsCore0000_80_hostpathProvisionerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_80_hostpath-provisioner-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_80_openshiftRouterCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-ingress
  name: service-ca-bundle 
  annotations:
    service.beta.openshift.io/inject-cabundle: "true"
`)

func assetsCore0000_80_openshiftRouterCmYamlBytes() ([]byte, error) {
	return _assetsCore0000_80_openshiftRouterCmYaml, nil
}

func assetsCore0000_80_openshiftRouterCmYaml() (*asset, error) {
	bytes, err := assetsCore0000_80_openshiftRouterCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_80_openshift-router-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_80_openshiftRouterNamespaceYaml = []byte(`kind: Namespace
apiVersion: v1
metadata:
  name: openshift-ingress
  annotations:
    openshift.io/node-selector: ""
  labels:
    # allow openshift-monitoring to look for ServiceMonitor objects in this namespace
    openshift.io/cluster-monitoring: "true"
    name: openshift-ingress
    network.openshift.io/policy-group: ingress
`)

func assetsCore0000_80_openshiftRouterNamespaceYamlBytes() ([]byte, error) {
	return _assetsCore0000_80_openshiftRouterNamespaceYaml, nil
}

func assetsCore0000_80_openshiftRouterNamespaceYaml() (*asset, error) {
	bytes, err := assetsCore0000_80_openshiftRouterNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_80_openshift-router-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_80_openshiftRouterServiceAccountYaml = []byte(`# Account for routers created by the operator. It will require cluster scoped
# permissions related to Route processing.
kind: ServiceAccount
apiVersion: v1
metadata:
  name: router
  namespace: openshift-ingress
`)

func assetsCore0000_80_openshiftRouterServiceAccountYamlBytes() ([]byte, error) {
	return _assetsCore0000_80_openshiftRouterServiceAccountYaml, nil
}

func assetsCore0000_80_openshiftRouterServiceAccountYaml() (*asset, error) {
	bytes, err := assetsCore0000_80_openshiftRouterServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_80_openshift-router-service-account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0000_80_openshiftRouterServiceYaml = []byte(`kind: Service
apiVersion: v1
metadata:
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: router-certs-default
  labels:
    ingresscontroller.operator.openshift.io/deployment-ingresscontroller: default
  name: router-internal-default
  namespace: openshift-ingress     
spec:
  selector:
    ingresscontroller.operator.openshift.io/deployment-ingresscontroller: default
  type: ClusterIP
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: http
  - name: https
    port: 443
    protocol: TCP
    targetPort: https
  - name: metrics
    port: 1936
    protocol: TCP
    targetPort: 1936
`)

func assetsCore0000_80_openshiftRouterServiceYamlBytes() ([]byte, error) {
	return _assetsCore0000_80_openshiftRouterServiceYaml, nil
}

func assetsCore0000_80_openshiftRouterServiceYaml() (*asset, error) {
	bytes, err := assetsCore0000_80_openshiftRouterServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0000_80_openshift-router-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCore0001_00_clusterVersionOperator_03_serviceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  name: cluster-version-operator
  namespace: openshift-cluster-version
  labels:
    k8s-app: cluster-version-operator
  annotations:
    service.beta.openshift.io/serving-cert-secret-name: cluster-version-operator-serving-cert
    exclude.release.openshift.io/internal-openshift-hosted: "true"
spec:
  type: ClusterIP
  selector:
    k8s-app: cluster-version-operator
  ports:
  - name: metrics
    port: 9099 # chosen to be in the internal open range
`)

func assetsCore0001_00_clusterVersionOperator_03_serviceYamlBytes() ([]byte, error) {
	return _assetsCore0001_00_clusterVersionOperator_03_serviceYaml, nil
}

func assetsCore0001_00_clusterVersionOperator_03_serviceYaml() (*asset, error) {
	bytes, err := assetsCore0001_00_clusterVersionOperator_03_serviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/0001_00_cluster-version-operator_03_service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsCoreOpenshiftSdnCmYaml = []byte(`apiVersion: v1
data:
  kube-proxy-config.yaml: |-
    apiVersion: kubeproxy.config.k8s.io/v1alpha1
    bindAddress: 0.0.0.0
    #bindAddressHardFail: false
    clientConnection:
      acceptContentTypes: ""
      burst: 0
      contentType: ""
      kubeconfig: ""
      qps: 0
    clusterCIDR: 10.42.0.0/16
    configSyncPeriod: 0s
    conntrack:
      maxPerCore: null
      min: null
      tcpCloseWaitTimeout: null
      tcpEstablishedTimeout: null
    detectLocalMode: ""
    enableProfiling: false
    featureGates:
      EndpointSlice: false
      EndpointSliceProxying: false
    healthzBindAddress: 0.0.0.0:10256
    hostnameOverride: ""
    iptables:
      masqueradeAll: false
      masqueradeBit: 0
      minSyncPeriod: 0s
      syncPeriod: 0s
    ipvs:
      excludeCIDRs: null
      minSyncPeriod: 0s
      scheduler: ""
      strictARP: false
      syncPeriod: 0s
      tcpFinTimeout: 0s
      tcpTimeout: 0s
      udpTimeout: 0s
    kind: KubeProxyConfiguration
    metricsBindAddress: 0.0.0.0:29101
    mode: unidling+iptables
    nodePortAddresses: null
    oomScoreAdj: null
    portRange: ""
    showHiddenMetricsForVersion: ""
    udpIdleTimeout: 0s
    winkernel:
      enableDSR: false
      networkName: ""
      sourceVip: ""
kind: ConfigMap
metadata:
  name: sdn-config
  namespace: openshift-sdn
`)

func assetsCoreOpenshiftSdnCmYamlBytes() ([]byte, error) {
	return _assetsCoreOpenshiftSdnCmYaml, nil
}

func assetsCoreOpenshiftSdnCmYaml() (*asset, error) {
	bytes, err := assetsCoreOpenshiftSdnCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/core/openshift-sdn-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/core/0000_00_cluster-version-operator_00_namespace.yaml":             assetsCore0000_00_clusterVersionOperator_00_namespaceYaml,
	"assets/core/0000_50_cluster-openshift-controller-manager_00_namespace.yaml": assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYaml,
	"assets/core/0000_50_service-ca-operator_01_namespace.yaml":                  assetsCore0000_50_serviceCaOperator_01_namespaceYaml,
	"assets/core/0000_50_service-ca-operator_02_service.yaml":                    assetsCore0000_50_serviceCaOperator_02_serviceYaml,
	"assets/core/0000_50_service-ca-operator_03_cm.yaml":                         assetsCore0000_50_serviceCaOperator_03_cmYaml,
	"assets/core/0000_50_service-ca-operator_04_sa.yaml":                         assetsCore0000_50_serviceCaOperator_04_saYaml,
	"assets/core/0000_60_service-ca_01_namespace.yaml":                           assetsCore0000_60_serviceCa_01_namespaceYaml,
	"assets/core/0000_60_service-ca_04_sa.yaml":                                  assetsCore0000_60_serviceCa_04_saYaml,
	"assets/core/0000_70_dns-operator_00-namespace.yaml":                         assetsCore0000_70_dnsOperator_00NamespaceYaml,
	"assets/core/0000_70_dns-operator_01-service-account.yaml":                   assetsCore0000_70_dnsOperator_01ServiceAccountYaml,
	"assets/core/0000_70_dns-operator_01-service.yaml":                           assetsCore0000_70_dnsOperator_01ServiceYaml,
	"assets/core/0000_70_dns_00-namespace.yaml":                                  assetsCore0000_70_dns_00NamespaceYaml,
	"assets/core/0000_70_dns_01-configmap.yaml":                                  assetsCore0000_70_dns_01ConfigmapYaml,
	"assets/core/0000_70_dns_01-service-account.yaml":                            assetsCore0000_70_dns_01ServiceAccountYaml,
	"assets/core/0000_70_dns_01-service.yaml":                                    assetsCore0000_70_dns_01ServiceYaml,
	"assets/core/0000_80_hostpath-provisioner-namespace.yaml":                    assetsCore0000_80_hostpathProvisionerNamespaceYaml,
	"assets/core/0000_80_hostpath-provisioner-serviceaccount.yaml":               assetsCore0000_80_hostpathProvisionerServiceaccountYaml,
	"assets/core/0000_80_openshift-router-cm.yaml":                               assetsCore0000_80_openshiftRouterCmYaml,
	"assets/core/0000_80_openshift-router-namespace.yaml":                        assetsCore0000_80_openshiftRouterNamespaceYaml,
	"assets/core/0000_80_openshift-router-service-account.yaml":                  assetsCore0000_80_openshiftRouterServiceAccountYaml,
	"assets/core/0000_80_openshift-router-service.yaml":                          assetsCore0000_80_openshiftRouterServiceYaml,
	"assets/core/0001_00_cluster-version-operator_03_service.yaml":               assetsCore0001_00_clusterVersionOperator_03_serviceYaml,
	"assets/core/openshift-sdn-cm.yaml":                                          assetsCoreOpenshiftSdnCmYaml,
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
	"assets": {nil, map[string]*bintree{
		"core": {nil, map[string]*bintree{
			"0000_00_cluster-version-operator_00_namespace.yaml":             {assetsCore0000_00_clusterVersionOperator_00_namespaceYaml, map[string]*bintree{}},
			"0000_50_cluster-openshift-controller-manager_00_namespace.yaml": {assetsCore0000_50_clusterOpenshiftControllerManager_00_namespaceYaml, map[string]*bintree{}},
			"0000_50_service-ca-operator_01_namespace.yaml":                  {assetsCore0000_50_serviceCaOperator_01_namespaceYaml, map[string]*bintree{}},
			"0000_50_service-ca-operator_02_service.yaml":                    {assetsCore0000_50_serviceCaOperator_02_serviceYaml, map[string]*bintree{}},
			"0000_50_service-ca-operator_03_cm.yaml":                         {assetsCore0000_50_serviceCaOperator_03_cmYaml, map[string]*bintree{}},
			"0000_50_service-ca-operator_04_sa.yaml":                         {assetsCore0000_50_serviceCaOperator_04_saYaml, map[string]*bintree{}},
			"0000_60_service-ca_01_namespace.yaml":                           {assetsCore0000_60_serviceCa_01_namespaceYaml, map[string]*bintree{}},
			"0000_60_service-ca_04_sa.yaml":                                  {assetsCore0000_60_serviceCa_04_saYaml, map[string]*bintree{}},
			"0000_70_dns-operator_00-namespace.yaml":                         {assetsCore0000_70_dnsOperator_00NamespaceYaml, map[string]*bintree{}},
			"0000_70_dns-operator_01-service-account.yaml":                   {assetsCore0000_70_dnsOperator_01ServiceAccountYaml, map[string]*bintree{}},
			"0000_70_dns-operator_01-service.yaml":                           {assetsCore0000_70_dnsOperator_01ServiceYaml, map[string]*bintree{}},
			"0000_70_dns_00-namespace.yaml":                                  {assetsCore0000_70_dns_00NamespaceYaml, map[string]*bintree{}},
			"0000_70_dns_01-configmap.yaml":                                  {assetsCore0000_70_dns_01ConfigmapYaml, map[string]*bintree{}},
			"0000_70_dns_01-service-account.yaml":                            {assetsCore0000_70_dns_01ServiceAccountYaml, map[string]*bintree{}},
			"0000_70_dns_01-service.yaml":                                    {assetsCore0000_70_dns_01ServiceYaml, map[string]*bintree{}},
			"0000_80_hostpath-provisioner-namespace.yaml":                    {assetsCore0000_80_hostpathProvisionerNamespaceYaml, map[string]*bintree{}},
			"0000_80_hostpath-provisioner-serviceaccount.yaml":               {assetsCore0000_80_hostpathProvisionerServiceaccountYaml, map[string]*bintree{}},
			"0000_80_openshift-router-cm.yaml":                               {assetsCore0000_80_openshiftRouterCmYaml, map[string]*bintree{}},
			"0000_80_openshift-router-namespace.yaml":                        {assetsCore0000_80_openshiftRouterNamespaceYaml, map[string]*bintree{}},
			"0000_80_openshift-router-service-account.yaml":                  {assetsCore0000_80_openshiftRouterServiceAccountYaml, map[string]*bintree{}},
			"0000_80_openshift-router-service.yaml":                          {assetsCore0000_80_openshiftRouterServiceYaml, map[string]*bintree{}},
			"0001_00_cluster-version-operator_03_service.yaml":               {assetsCore0001_00_clusterVersionOperator_03_serviceYaml, map[string]*bintree{}},
			"openshift-sdn-cm.yaml":                                          {assetsCoreOpenshiftSdnCmYaml, map[string]*bintree{}},
		}},
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
