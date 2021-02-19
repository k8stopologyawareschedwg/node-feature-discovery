package topologypolicy

// TopologyManagerPolicy constants which represent the current configuration
// for Topology manager policy and Topology manager scope in Kubelet config
type TopologyManagerPolicy string

const (
	SingleNumaContainerScope = "SingleNUMANodeContainerLevel"
	SingleNumaPodScope       = "SingleNUMANodePodLevel"
	Restricted               = "Restricted"
	BestEffort               = "BestEffort"
	None                     = "None"
)

// K8sTopologyPolicies are resource allocation policies constants
type K8sTopologyManagerPolicies string

const (
	singleNumaNode = "single-numa-node"
	restricted     = "restricted"
	bestEffort     = "best-effort"
	none           = "none"
)

// K8sTopologyScopes are constants which defines the granularity
// at which you would like resource alignment to be performed.
type K8sTopologyManagerScopes string

const (
	pod       = "pod"
	container = "container"
)
