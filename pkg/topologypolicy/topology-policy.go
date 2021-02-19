package topologypolicy

// DetectTopologyPolicy returns TopologyManagerPolicy type which present
// both Topology manager policy and scope
func DetectTopologyPolicy(policy string, scope string) TopologyManagerPolicy {
	switch policy {
	case singleNumaNode:
		if scope == pod {
			return SingleNumaPodScope
		}
		// default scope for single-numa-node
		return SingleNumaContainerScope
	case restricted:
		return Restricted
	case bestEffort:
		return BestEffort
	default:
		return None
	}
}
