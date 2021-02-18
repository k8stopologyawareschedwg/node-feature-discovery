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
