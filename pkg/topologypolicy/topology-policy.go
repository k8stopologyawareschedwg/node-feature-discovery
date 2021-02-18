package topologypolicy

// DetectTopologyPolicy returns TopologyManagerPolicy type which present
// both Topology manager policy and scope
func DetectTopologyPolicy(policy string, scope string) TopologyManagerPolicy {
	switch policy {
	case "single-numa-node":
		if scope == "pod" {
			return SingleNumaPodScope
		}
		// default scope for single-numa-node
		return SingleNumaContainerScope
	case "Restricted":
		return Restricted
	case "BestEffort":
		return BestEffort
	default:
		return None
	}
}
