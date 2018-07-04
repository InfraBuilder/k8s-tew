package config

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/darxkies/k8s-tew/utils"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type InternalConfig struct {
	BaseDirectory string
	Name          string
	Node          *Node
	Config        *Config
}

func (config *InternalConfig) GetTemplateAssetFilename(name string) string {
	return fmt.Sprintf(`{{asset_file "%s"}}`, name)
}

func (config *InternalConfig) GetTemplateAssetDirectory(name string) string {
	return fmt.Sprintf(`{{asset_directory "%s"}}`, name)
}

func (config *InternalConfig) GetFullTargetAssetFilename(name string) string {
	return path.Join(config.Config.DeploymentDirectory, config.GetRelativeAssetFilename(name))
}

func (config *InternalConfig) GetFullLocalAssetFilename(name string) string {
	return path.Join(config.BaseDirectory, config.GetRelativeAssetFilename(name))
}

func (config *InternalConfig) GetRelativeAssetFilename(name string) string {
	var result *AssetFile
	var ok bool
	var directory *AssetDirectory

	if result, ok = config.Config.Assets.Files[name]; !ok {
		log.WithFields(log.Fields{"name": name}).Fatal("missing asset file")
	}

	if directory, ok = config.Config.Assets.Directories[result.Directory]; !ok {
		log.WithFields(log.Fields{"name": name}).Fatal("missing asset directory")
	}

	filename := path.Join(directory.Directory, name)

	resultFilename, error := config.ApplyTemplate("asset-file", filename)
	if error != nil {
		log.WithFields(log.Fields{"name": name, "error": error}).Fatal("asset file expansion")
	}

	return path.Join("/", resultFilename)
}

func (config *InternalConfig) GetFullLocalAssetDirectory(name string) string {
	return path.Join(config.BaseDirectory, config.GetRelativeAssetDirectory(name))
}

func (config *InternalConfig) GetFullTargetAssetDirectory(name string) string {
	return path.Join(config.Config.DeploymentDirectory, config.GetRelativeAssetDirectory(name))
}

func (config *InternalConfig) GetRelativeAssetDirectory(name string) string {
	var result *AssetDirectory
	var ok bool

	if result, ok = config.Config.Assets.Directories[name]; !ok {
		log.WithFields(log.Fields{"name": name}).Fatal("missing asset directory")
	}

	return result.Directory
}

func (config *InternalConfig) SetNode(nodeName string, node *Node) {
	config.Name = nodeName
	config.Node = node
}

func NewInternalConfig(baseDirectory string) *InternalConfig {
	config := &InternalConfig{}
	config.BaseDirectory = baseDirectory

	config.Config = NewConfig()

	return config
}

func (config *InternalConfig) registerAssetDirectories() {
	// Config
	config.addAssetDirectory(utils.CONFIG_DIRECTORY, Labels{}, config.getRelativeConfigDirectory())
	config.addAssetDirectory(utils.CERTIFICATES_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.CONFIG_DIRECTORY), utils.CERTIFICATES_SUBDIRECTORY))
	config.addAssetDirectory(utils.CNI_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.CONFIG_DIRECTORY), utils.CNI_SUBDIRECTORY))
	config.addAssetDirectory(utils.CRI_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.CONFIG_DIRECTORY), utils.CRI_SUBDIRECTORY))

	// K8S Config
	config.addAssetDirectory(utils.K8S_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.CONFIG_DIRECTORY), utils.K8S_SUBDIRECTORY))
	config.addAssetDirectory(utils.K8S_KUBE_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.K8S_CONFIG_DIRECTORY), utils.KUBECONFIG_SUBDIRECTORY))
	config.addAssetDirectory(utils.K8S_SECURITY_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.K8S_CONFIG_DIRECTORY), utils.SECURITY_SUBDIRECTORY))
	config.addAssetDirectory(utils.K8S_SETUP_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.K8S_CONFIG_DIRECTORY), utils.SETUP_SUBDIRECTORY))
	config.addAssetDirectory(utils.K8S_MANIFESTS_DIRECTORY, Labels{utils.NODE_WORKER}, path.Join(config.GetRelativeAssetDirectory(utils.K8S_CONFIG_DIRECTORY), utils.MANIFESTS_SUBDIRECTORY))

	// Binaries
	config.addAssetDirectory(utils.BINARIES_DIRECTORY, Labels{}, path.Join(utils.OPTIONAL_SUBDIRECTORY, utils.K8S_TEW_SUBDIRECTORY, utils.BINARY_SUBDIRECTORY))
	config.addAssetDirectory(utils.K8S_BINARIES_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.BINARIES_DIRECTORY), utils.K8S_SUBDIRECTORY))
	config.addAssetDirectory(utils.ETCD_BINARIES_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.BINARIES_DIRECTORY), utils.ETCD_SUBDIRECTORY))
	config.addAssetDirectory(utils.CRI_BINARIES_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.BINARIES_DIRECTORY), utils.CRI_SUBDIRECTORY))
	config.addAssetDirectory(utils.CNI_BINARIES_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.BINARIES_DIRECTORY), utils.CNI_SUBDIRECTORY))
	config.addAssetDirectory(utils.GOBETWEEN_BINARIES_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.BINARIES_DIRECTORY), utils.LOAD_BALANCER_SUBDIRECTORY))

	// Misc
	config.addAssetDirectory(utils.GOBETWEEN_CONFIG_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.CONFIG_DIRECTORY), utils.LOAD_BALANCER_SUBDIRECTORY))
	config.addAssetDirectory(utils.DYNAMIC_DATA_DIRECTORY, Labels{}, path.Join(utils.VARIABLE_SUBDIRECTORY, utils.LIBRARY_SUBDIRECTORY, utils.K8S_TEW_SUBDIRECTORY))
	config.addAssetDirectory(utils.ETCD_DATA_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.DYNAMIC_DATA_DIRECTORY), utils.ETCD_SUBDIRECTORY))
	config.addAssetDirectory(utils.CONTAINERD_DATA_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.DYNAMIC_DATA_DIRECTORY), utils.CONTAINERD_SUBDIRECTORY))
	config.addAssetDirectory(utils.KUBELET_DATA_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.DYNAMIC_DATA_DIRECTORY), utils.KUBELET_SUBDIRECTORY))
	config.addAssetDirectory(utils.LOGGING_DIRECTORY, Labels{}, path.Join(utils.VARIABLE_SUBDIRECTORY, utils.LOGGING_SUBDIRECTORY, utils.K8S_TEW_SUBDIRECTORY))
	config.addAssetDirectory(utils.SERVICE_DIRECTORY, Labels{}, path.Join(utils.CONFIG_SUBDIRECTORY, utils.SYSTEMD_SUBDIRECTORY, utils.SYSTEM_SUBDIRECTORY))
	config.addAssetDirectory(utils.CONTAINERD_STATE_DIRECTORY, Labels{}, path.Join(utils.VARIABLE_SUBDIRECTORY, utils.RUN_SUBDIRECTORY, utils.K8S_TEW_SUBDIRECTORY, utils.CONTAINERD_SUBDIRECTORY))
	config.addAssetDirectory(utils.PROFILE_DIRECTORY, Labels{}, path.Join(utils.CONFIG_SUBDIRECTORY, utils.PROFILE_D_SUBDIRECTORY))
	config.addAssetDirectory(utils.HELM_DATA_DIRECTORY, Labels{}, path.Join(config.GetRelativeAssetDirectory(utils.DYNAMIC_DATA_DIRECTORY), utils.HELM_SUBDIRECTORY))
	config.addAssetDirectory(utils.TEMPORARY_DIRECTORY, Labels{}, path.Join(utils.TEMPORARY_SUBDIRECTORY))
}

func (config *InternalConfig) registerAssetFiles() {
	// Config
	config.addAssetFile(utils.CONFIG_FILENAME, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.CONFIG_DIRECTORY)

	// Binaries
	config.addAssetFile(utils.K8S_TEW_BINARY, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.BINARIES_DIRECTORY)

	// CNI Binaries
	config.addAssetFile(utils.BRIDGE_BINARY, Labels{utils.NODE_WORKER}, utils.CNI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.FLANNEL_BINARY, Labels{utils.NODE_WORKER}, utils.CNI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.LOOPBACK_BINARY, Labels{utils.NODE_WORKER}, utils.CNI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.HOST_LOCAL_BINARY, Labels{utils.NODE_WORKER}, utils.CNI_BINARIES_DIRECTORY)

	// ContainerD Binaries
	config.addAssetFile(utils.CONTAINERD_BINARY, Labels{utils.NODE_WORKER}, utils.CRI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.CONTAINERD_SHIM_BINARY, Labels{utils.NODE_WORKER}, utils.CRI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.CTR_BINARY, Labels{utils.NODE_WORKER}, utils.CRI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.RUNC_BINARY, Labels{utils.NODE_WORKER}, utils.CRI_BINARIES_DIRECTORY)
	config.addAssetFile(utils.CRICTL_BINARY, Labels{utils.NODE_WORKER}, utils.CRI_BINARIES_DIRECTORY)

	// Etcd Binaries
	config.addAssetFile(utils.ETCD_BINARY, Labels{utils.NODE_CONTROLLER}, utils.ETCD_BINARIES_DIRECTORY)
	config.addAssetFile(utils.ETCDCTL_BINARY, Labels{utils.NODE_CONTROLLER}, utils.ETCD_BINARIES_DIRECTORY)
	config.addAssetFile(utils.FLANNELD_BINARY, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.ETCD_BINARIES_DIRECTORY)

	// K8S Binaries
	config.addAssetFile(utils.KUBECTL_BINARY, Labels{utils.NODE_CONTROLLER}, utils.K8S_BINARIES_DIRECTORY)
	config.addAssetFile(utils.KUBE_APISERVER_BINARY, Labels{utils.NODE_CONTROLLER}, utils.K8S_BINARIES_DIRECTORY)
	config.addAssetFile(utils.KUBE_CONTROLLER_MANAGER_BINARY, Labels{utils.NODE_CONTROLLER}, utils.K8S_BINARIES_DIRECTORY)
	config.addAssetFile(utils.KUBELET_BINARY, Labels{utils.NODE_WORKER}, utils.K8S_BINARIES_DIRECTORY)
	config.addAssetFile(utils.KUBE_PROXY_BINARY, Labels{utils.NODE_WORKER}, utils.K8S_BINARIES_DIRECTORY)
	config.addAssetFile(utils.KUBE_SCHEDULER_BINARY, Labels{utils.NODE_CONTROLLER}, utils.K8S_BINARIES_DIRECTORY)

	// Helm Binary
	config.addAssetFile(utils.HELM_BINARY, Labels{}, utils.K8S_BINARIES_DIRECTORY)

	// Gobetween Binary
	config.addAssetFile(utils.GOBETWEEN_BINARY, Labels{utils.NODE_CONTROLLER}, utils.GOBETWEEN_BINARIES_DIRECTORY)

	// Certificates
	config.addAssetFile(utils.CA_PEM, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.CA_KEY_PEM, Labels{utils.NODE_CONTROLLER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.VIRTUAL_IP_PEM, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.VIRTUAL_IP_KEY_PEM, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.FLANNELD_PEM, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.FLANNELD_KEY_PEM, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.KUBERNETES_PEM, Labels{utils.NODE_CONTROLLER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.KUBERNETES_KEY_PEM, Labels{utils.NODE_CONTROLLER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.SERVICE_ACCOUNT_PEM, Labels{utils.NODE_CONTROLLER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.SERVICE_ACCOUNT_KEY_PEM, Labels{utils.NODE_CONTROLLER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.ADMIN_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.ADMIN_KEY_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.CONTROLLER_MANAGER_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.CONTROLLER_MANAGER_KEY_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.SCHEDULER_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.SCHEDULER_KEY_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.PROXY_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.PROXY_KEY_PEM, Labels{}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.KUBELET_PEM, Labels{utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)
	config.addAssetFile(utils.KUBELET_KEY_PEM, Labels{utils.NODE_WORKER}, utils.CERTIFICATES_DIRECTORY)

	// Kubeconfig
	config.addAssetFile(utils.ADMIN_KUBECONFIG, Labels{}, utils.K8S_KUBE_CONFIG_DIRECTORY)
	config.addAssetFile(utils.CONTROLLER_MANAGER_KUBECONFIG, Labels{utils.NODE_CONTROLLER}, utils.K8S_KUBE_CONFIG_DIRECTORY)
	config.addAssetFile(utils.SCHEDULER_KUBECONFIG, Labels{utils.NODE_CONTROLLER}, utils.K8S_KUBE_CONFIG_DIRECTORY)
	config.addAssetFile(utils.PROXY_KUBECONFIG, Labels{utils.NODE_WORKER}, utils.K8S_KUBE_CONFIG_DIRECTORY)
	config.addAssetFile(utils.KUBELET_KUBECONFIG, Labels{utils.NODE_WORKER}, utils.K8S_KUBE_CONFIG_DIRECTORY)

	// Security
	config.addAssetFile(utils.ENCRYPTION_CONFIG, Labels{utils.NODE_CONTROLLER}, utils.K8S_SECURITY_CONFIG_DIRECTORY)

	// CNI
	config.addAssetFile(utils.NET_CONFIG, Labels{utils.NODE_WORKER}, utils.CNI_CONFIG_DIRECTORY)
	config.addAssetFile(utils.CNI_CONFIG, Labels{utils.NODE_WORKER}, utils.CNI_CONFIG_DIRECTORY)

	// CRI
	config.addAssetFile(utils.CONTAINERD_CONFIG, Labels{utils.NODE_WORKER}, utils.CRI_CONFIG_DIRECTORY)
	config.addAssetFile(utils.CONTAINERD_SOCK, Labels{}, utils.CONTAINERD_STATE_DIRECTORY)

	// Service
	config.addAssetFile(utils.SERVICE_CONFIG, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.SERVICE_DIRECTORY)

	// K8S Setup
	config.addAssetFile(utils.K8S_KUBELET_SETUP, Labels{}, utils.K8S_SETUP_CONFIG_DIRECTORY)
	config.addAssetFile(utils.K8S_ADMIN_USER_SETUP, Labels{}, utils.K8S_SETUP_CONFIG_DIRECTORY)
	config.addAssetFile(utils.K8S_HELM_USER_SETUP, Labels{}, utils.K8S_SETUP_CONFIG_DIRECTORY)

	// K8S Config
	config.addAssetFile(utils.K8S_KUBE_SCHEDULER_CONFIG, Labels{utils.NODE_CONTROLLER}, utils.K8S_CONFIG_DIRECTORY)
	config.addAssetFile(utils.K8S_KUBELET_CONFIG, Labels{utils.NODE_WORKER}, utils.K8S_CONFIG_DIRECTORY)

	// Profile
	config.addAssetFile(utils.K8S_TEW_PROFILE, Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, utils.PROFILE_DIRECTORY)

	// Gobetween
	config.addAssetFile(utils.GOBETWEEN_CONFIG, Labels{utils.NODE_CONTROLLER}, utils.GOBETWEEN_CONFIG_DIRECTORY)
}

func (config *InternalConfig) registerServers() {
	// Servers
	config.addServer("etcd", Labels{utils.NODE_CONTROLLER}, config.GetTemplateAssetFilename(utils.ETCD_BINARY), map[string]string{
		"name":                        "{{.Name}}",
		"cert-file":                   config.GetTemplateAssetFilename(utils.KUBERNETES_PEM),
		"key-file":                    config.GetTemplateAssetFilename(utils.KUBERNETES_KEY_PEM),
		"peer-cert-file":              config.GetTemplateAssetFilename(utils.KUBERNETES_PEM),
		"peer-key-file":               config.GetTemplateAssetFilename(utils.KUBERNETES_KEY_PEM),
		"trusted-ca-file":             config.GetTemplateAssetFilename(utils.CA_PEM),
		"peer-trusted-ca-file":        config.GetTemplateAssetFilename(utils.CA_PEM),
		"peer-client-cert-auth":       "",
		"client-cert-auth":            "",
		"initial-advertise-peer-urls": "https://{{.Node.IP}}:2380",
		"listen-peer-urls":            "https://{{.Node.IP}}:2380",
		"listen-client-urls":          "https://{{.Node.IP}}:2379",
		"advertise-client-urls":       "https://{{.Node.IP}}:2379",
		"initial-cluster-token":       "etcd-cluster",
		"initial-cluster":             "{{etcd_cluster}}",
		"initial-cluster-state":       "new",
		"data-dir":                    config.GetTemplateAssetDirectory(utils.ETCD_DATA_DIRECTORY),
	})

	config.addServer("flanneld", Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, config.GetTemplateAssetFilename(utils.FLANNELD_BINARY), map[string]string{
		"etcd-endpoints": "{{etcd_servers}}",
		"etcd-cafile":    config.GetTemplateAssetFilename(utils.CA_PEM),
		"etcd-certfile":  config.GetTemplateAssetFilename(utils.FLANNELD_PEM),
		"etcd-keyfile":   config.GetTemplateAssetFilename(utils.FLANNELD_KEY_PEM),
		"iface-regex":    "{{.Node.IP}}",
		"v":              "0",
	})

	config.addServer("containerd", Labels{utils.NODE_WORKER}, config.GetTemplateAssetFilename(utils.CONTAINERD_BINARY), map[string]string{
		"config": config.GetTemplateAssetFilename(utils.CONTAINERD_CONFIG),
	})

	config.addServer("gobetween", Labels{utils.NODE_CONTROLLER}, config.GetTemplateAssetFilename(utils.GOBETWEEN_BINARY), map[string]string{
		"config": config.GetTemplateAssetFilename(utils.GOBETWEEN_CONFIG),
	})

	config.addServer("kube-apiserver", Labels{utils.NODE_CONTROLLER}, config.GetTemplateAssetFilename(utils.KUBE_APISERVER_BINARY), map[string]string{
		"allow-privileged":                        "true",
		"advertise-address":                       "{{.Node.IP}}",
		"apiserver-count":                         "{{controllers_count}}",
		"audit-log-maxage":                        "30",
		"audit-log-maxbackup":                     "3",
		"audit-log-maxsize":                       "100",
		"audit-log-path":                          path.Join(config.GetTemplateAssetDirectory(utils.LOGGING_DIRECTORY), utils.AUDIT_LOG),
		"authorization-mode":                      "Node,RBAC",
		"bind-address":                            "0.0.0.0",
		"secure-port":                             "{{.Config.APIServerPort}}",
		"client-ca-file":                          config.GetTemplateAssetFilename(utils.CA_PEM),
		"enable-admission-plugins":                "Initializers,NamespaceLifecycle,NodeRestriction,LimitRanger,ServiceAccount,DefaultStorageClass,ResourceQuota",
		"enable-swagger-ui":                       "true",
		"etcd-cafile":                             config.GetTemplateAssetFilename(utils.CA_PEM),
		"etcd-certfile":                           config.GetTemplateAssetFilename(utils.KUBERNETES_PEM),
		"etcd-keyfile":                            config.GetTemplateAssetFilename(utils.KUBERNETES_KEY_PEM),
		"etcd-servers":                            "{{etcd_servers}}",
		"event-ttl":                               "1h",
		"experimental-encryption-provider-config": config.GetTemplateAssetFilename(utils.ENCRYPTION_CONFIG),
		"kubelet-certificate-authority":           config.GetTemplateAssetFilename(utils.CA_PEM),
		"kubelet-client-certificate":              config.GetTemplateAssetFilename(utils.KUBERNETES_PEM),
		"kubelet-client-key":                      config.GetTemplateAssetFilename(utils.KUBERNETES_KEY_PEM),
		"kubelet-https":                           "true",
		"runtime-config":                          "api/all",
		"service-account-key-file":                config.GetTemplateAssetFilename(utils.SERVICE_ACCOUNT_PEM),
		"service-cluster-ip-range":                "{{.Config.ClusterIPRange}}",
		"service-node-port-range":                 "30000-32767",
		"tls-cert-file":                           config.GetTemplateAssetFilename(utils.KUBERNETES_PEM),
		"tls-private-key-file":                    config.GetTemplateAssetFilename(utils.KUBERNETES_KEY_PEM),
		"v": "0",
	})

	config.addServer("kube-controller-manager", Labels{utils.NODE_CONTROLLER}, config.GetTemplateAssetFilename(utils.KUBE_CONTROLLER_MANAGER_BINARY), map[string]string{
		"address":                          "0.0.0.0",
		"cluster-cidr":                     "{{.Config.ClusterCIDR}}",
		"cluster-name":                     "kubernetes",
		"cluster-signing-cert-file":        config.GetTemplateAssetFilename(utils.CA_PEM),
		"cluster-signing-key-file":         config.GetTemplateAssetFilename(utils.CA_KEY_PEM),
		"kubeconfig":                       config.GetTemplateAssetFilename(utils.CONTROLLER_MANAGER_KUBECONFIG),
		"leader-elect":                     "true",
		"root-ca-file":                     config.GetTemplateAssetFilename(utils.CA_PEM),
		"service-account-private-key-file": config.GetTemplateAssetFilename(utils.SERVICE_ACCOUNT_KEY_PEM),
		"service-cluster-ip-range":         "{{.Config.ClusterIPRange}}",
		"use-service-account-credentials":  "true",
		"v": "0",
	})

	config.addServer("kube-scheduler", Labels{utils.NODE_CONTROLLER}, config.GetTemplateAssetFilename(utils.KUBE_SCHEDULER_BINARY), map[string]string{
		"config": config.GetTemplateAssetFilename(utils.K8S_KUBE_SCHEDULER_CONFIG),
		"v":      "0",
	})

	config.addServer("kube-proxy", Labels{utils.NODE_WORKER}, config.GetTemplateAssetFilename(utils.KUBE_PROXY_BINARY), map[string]string{
		"cluster-cidr": "{{.Config.ClusterCIDR}}",
		"kubeconfig":   config.GetTemplateAssetFilename(utils.PROXY_KUBECONFIG),
		"proxy-mode":   "iptables",
		"v":            "0",
	})

	config.addServer("kubelet", Labels{utils.NODE_WORKER}, config.GetTemplateAssetFilename(utils.KUBELET_BINARY), map[string]string{
		"config":                       config.GetTemplateAssetFilename(utils.K8S_KUBELET_CONFIG),
		"container-runtime":            "remote",
		"container-runtime-endpoint":   "unix://" + config.GetTemplateAssetFilename(utils.CONTAINERD_SOCK),
		"image-pull-progress-deadline": "2m",
		"kubeconfig":                   config.GetTemplateAssetFilename(utils.KUBELET_KUBECONFIG),
		"network-plugin":               "cni",
		"register-node":                "true",
		"allow-privileged":             "true",
		"resolv-conf":                  "{{.Config.ResolvConf}}",
		"root-dir":                     config.GetTemplateAssetDirectory(utils.KUBELET_DATA_DIRECTORY),
		"v":                            "0",
	})
}

func (config *InternalConfig) registerCommands() {
	kubectlCommand := fmt.Sprintf("%s --kubeconfig %s", config.GetFullLocalAssetFilename(utils.KUBECTL_BINARY), config.GetFullLocalAssetFilename(utils.ADMIN_KUBECONFIG))
	helmCommand := fmt.Sprintf("KUBECONFIG=%s HELM_HOME=%s %s", config.GetFullLocalAssetFilename(utils.ADMIN_KUBECONFIG), config.GetFullLocalAssetDirectory(utils.HELM_DATA_DIRECTORY), config.GetFullLocalAssetFilename(utils.HELM_BINARY))

	// Dependencies
	config.addCommand("swapoff", Labels{utils.NODE_WORKER}, "swapoff -a")
	config.addCommand("load-overlay", Labels{utils.NODE_WORKER}, "modprobe overlay")
	config.addCommand("load-btrfs", Labels{utils.NODE_WORKER}, "modprobe btrfs")
	config.addCommand("load-br_netfilter", Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, "modprobe br_netfilter")
	config.addCommand("enable-br_netfilter", Labels{utils.NODE_CONTROLLER, utils.NODE_WORKER}, "echo '1' > /proc/sys/net/bridge/bridge-nf-call-iptables")
	config.addCommand("flanneld-configuration", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s --ca-file=%s --cert-file=%s --key-file=%s --endpoints=%s set /coreos.com/network/config '{ \"Network\": \"%s\" }'", config.GetFullLocalAssetFilename(utils.ETCDCTL_BINARY), config.GetFullLocalAssetFilename(utils.CA_PEM), config.GetFullLocalAssetFilename(utils.KUBERNETES_PEM), config.GetFullLocalAssetFilename(utils.KUBERNETES_KEY_PEM), strings.Join(config.GetETCDClientEndpoints(), ","), config.Config.ClusterCIDR))
	config.addCommand("k8s-kubelet-setup", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s apply -f %s", kubectlCommand, config.GetFullLocalAssetFilename(utils.K8S_KUBELET_SETUP)))
	config.addCommand("k8s-admin-user-setup", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s apply -f %s", kubectlCommand, config.GetFullLocalAssetFilename(utils.K8S_ADMIN_USER_SETUP)))
	config.addCommand("k8s-kube-dns", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s apply -f https://storage.googleapis.com/kubernetes-the-hard-way/kube-dns.yaml", kubectlCommand))
	config.addCommand("k8s-helm-user-setup", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s apply -f %s", kubectlCommand, config.GetFullLocalAssetFilename(utils.K8S_HELM_USER_SETUP)))
	config.addCommand("helm-init", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s init --service-account %s --upgrade", helmCommand, utils.HELM_SERVICE_ACCOUNT))
	config.addCommand("helm-repo-update", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s repo update", helmCommand))
	config.addCommand("helm-kubernetes-dashboard", Labels{utils.NODE_BOOTSTRAPPER}, fmt.Sprintf("%s get svc kubernetes-dashboard -n kube-system || %s install stable/kubernetes-dashboard --name kubernetes-dashboard --set=service.type=NodePort,service.nodePort=32443 --namespace kube-system", kubectlCommand, helmCommand))
}

func (config *InternalConfig) Generate(deploymentDirectory string) {
	config.Config.APIServerPort = 6443
	config.Config.DeploymentDirectory = deploymentDirectory

	config.registerAssetDirectories()
	config.registerAssetFiles()
	config.registerCommands()
	config.registerServers()
}

func (config *InternalConfig) addServer(name string, labels []string, command string, arguments map[string]string) {
	// Do not add if already in the list
	for _, server := range config.Config.Servers {
		if server.Name == name {
			return
		}
	}

	config.Config.Servers = append(config.Config.Servers, ServerConfig{Name: name, Labels: labels, Command: command, Arguments: arguments, Logger: LoggerConfig{Enabled: true, Filename: path.Join(config.GetTemplateAssetDirectory(utils.LOGGING_DIRECTORY), name+".log")}})
}

func (config *InternalConfig) addCommand(name string, labels Labels, command string) {
	// Do not add if already in the list
	for _, command := range config.Config.Commands {
		if command.Name == name {
			return
		}
	}

	config.Config.Commands = append(config.Config.Commands, NewCommand(name, labels, command))
}

func (config *InternalConfig) addAssetFile(name string, labels Labels, _path string) {
	config.Config.Assets.Files[name] = NewAssetFile(labels, _path)
}

func (config *InternalConfig) addAssetDirectory(name string, labels Labels, directory string) {
	config.Config.Assets.Directories[name] = NewAssetDirectory(labels, directory)
}

func (config *InternalConfig) Dump() {
	log.WithFields(log.Fields{"base-directory": config.BaseDirectory}).Info("config")
	log.WithFields(log.Fields{"name": config.Name}).Info("config")

	if config.Node != nil {
		log.WithFields(log.Fields{"ip": config.Node.IP}).Info("config")
		log.WithFields(log.Fields{"labels": config.Node.Labels}).Info("config")
		log.WithFields(log.Fields{"index": config.Node.Index}).Info("config")
	}

	for name, assetFile := range config.Config.Assets.Files {
		log.WithFields(log.Fields{"name": name, "directory": assetFile.Directory, "labels": assetFile.Labels}).Info("config asset file")
	}

	for name, node := range config.Config.Nodes {
		log.WithFields(log.Fields{"name": name, "index": node.Index, "labels": node.Labels, "ip": node.IP}).Info("config node")
	}

	for name, command := range config.Config.Commands {
		log.WithFields(log.Fields{"name": name, "command": command.Command, "labels": command.Labels}).Info("config command")
	}

	for _, serverConfig := range config.Config.Servers {
		serverConfig.Dump()
	}
}

func (config *InternalConfig) getRelativeConfigDirectory() string {
	return path.Join(utils.CONFIG_SUBDIRECTORY, utils.K8S_TEW_SUBDIRECTORY)
}

func (config *InternalConfig) getConfigDirectory() string {
	return path.Join(config.BaseDirectory, config.getRelativeConfigDirectory())
}

func (config *InternalConfig) getConfigFilename() string {
	return path.Join(config.getConfigDirectory(), utils.CONFIG_FILENAME)
}

func (config *InternalConfig) Save() error {
	if error := utils.CreateDirectoryIfMissing(config.getConfigDirectory()); error != nil {
		return error
	}

	yamlOutput, error := yaml.Marshal(config.Config)
	if error != nil {
		return error
	}

	filename := config.getConfigFilename()

	if error := ioutil.WriteFile(filename, yamlOutput, 0644); error != nil {
		return error
	}

	log.WithFields(log.Fields{"filename": filename}).Info("saved config")

	return nil
}

func (config *InternalConfig) Load() error {
	var error error

	filename := config.getConfigFilename()

	// Check if config file exists
	if _, error := os.Stat(filename); os.IsNotExist(error) {
		return errors.New(fmt.Sprintf("config '%s' not found", filename))
	}

	yamlContent, error := ioutil.ReadFile(filename)

	if error != nil {
		return error
	}

	if error := yaml.Unmarshal(yamlContent, config.Config); error != nil {
		return error
	}

	if len(config.Name) == 0 {
		config.Name, error = os.Hostname()

		if error != nil {
			return error
		}
	}

	if config.Node == nil {
		for name, node := range config.Config.Nodes {
			if name != config.Name {
				continue
			}

			config.Node = node

			break
		}
	}

	return nil
}

func (config *InternalConfig) RemoveNode(name string) error {
	if _, ok := config.Config.Nodes[name]; !ok {
		return errors.New("node not found")
	}

	delete(config.Config.Nodes, name)

	return nil
}

func (config *InternalConfig) AddNode(name string, ip string, index uint, labels []string) (*Node, error) {
	name = strings.Trim(name, " \n")

	if len(name) == 0 {
		return nil, errors.New("empty node name")
	}

	if net.ParseIP(ip) == nil {
		return nil, errors.New("invalid or wrong ip format")
	}

	config.Config.Nodes[name] = NewNode(ip, index, labels)

	return config.Config.Nodes[name], nil
}

func (config *InternalConfig) GetETCDClientEndpoints() []string {
	result := []string{}

	for _, node := range config.Config.Nodes {
		if node.IsController() {
			result = append(result, fmt.Sprintf("https://%s:2379", node.IP))
		}
	}

	return result
}

func (config *InternalConfig) ApplyTemplate(label string, value string) (string, error) {
	var functions = template.FuncMap{
		"controllers_count": func() string {
			count := 0
			for _, node := range config.Config.Nodes {
				if node.IsController() {
					count += 1
				}
			}

			return fmt.Sprintf("%d", count)
		},
		"etcd_servers": func() string {
			result := ""

			for _, endpoint := range config.GetETCDClientEndpoints() {
				if len(result) > 0 {
					result += ","
				}

				result += endpoint
			}

			return result
		},
		"etcd_cluster": func() string {
			result := ""

			for name, node := range config.Config.Nodes {
				if !node.IsController() {
					continue
				}

				if len(result) > 0 {
					result += ","
				}

				result += fmt.Sprintf("%s=https://%s:2380", name, node.IP)
			}

			return result
		},
		"asset_file": func(name string) string {
			return config.GetFullTargetAssetFilename(name)
		},
		"asset_directory": func(name string) string {
			return config.GetFullTargetAssetDirectory(name)
		},
	}

	var newValue bytes.Buffer

	argumentTemplate, error := template.New(fmt.Sprintf(label)).Funcs(functions).Parse(value)

	if error != nil {
		return "", error
	}

	if error = argumentTemplate.Execute(&newValue, config); error != nil {
		return "", error
	}

	return newValue.String(), nil
}

func (config *InternalConfig) GetAPIServerIP() (string, error) {
	if len(config.Config.ControllerVirtualIP) > 0 {
		return config.Config.ControllerVirtualIP, nil
	}

	for _, node := range config.Config.Nodes {
		if node.IsController() {
			return node.IP, nil
		}
	}

	return "", errors.New("No API Server IP found")
}

func (config *InternalConfig) GetSortedNodeKeys() []string {
	result := []string{}

	for key := range config.Config.Nodes {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}

func (config *InternalConfig) GetKubeAPIServerAddresses() []string {
	result := []string{}

	for _, node := range config.Config.Nodes {
		if node.IsController() {
			result = append(result, fmt.Sprintf("%s:%d", node.IP, config.Config.APIServerPort))
		}
	}

	return result
}