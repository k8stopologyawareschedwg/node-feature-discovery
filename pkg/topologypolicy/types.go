package topologypolicy

// K8sTopologyPolicies are resource allocation policies constants
type K8sTopologyManagerPolicies string

const (
	singleNumaNode K8sTopologyManagerPolicies = "single-numa-node"
	restricted     K8sTopologyManagerPolicies = "restricted"
	bestEffort     K8sTopologyManagerPolicies = "best-effort"
	none           K8sTopologyManagerPolicies = "none"
)

// K8sTopologyScopes are constants which defines the granularity
// at which you would like resource alignment to be performed.
type K8sTopologyManagerScopes string

const (
	pod       K8sTopologyManagerScopes = "pod"
	container K8sTopologyManagerScopes = "container"
)
