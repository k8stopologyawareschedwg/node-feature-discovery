package topologypolicy

import (
	v1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
)

// DetectTopologyPolicy returns TopologyManagerPolicy type which present
// both Topology manager policy and scope
func DetectTopologyPolicy(policy string, scope string) v1alpha1.TopologyManagerPolicy {
	k8sTmPolicy := K8sTopologyManagerPolicies(policy)
	k8sTmScope := K8sTopologyManagerScopes(scope)
	switch k8sTmPolicy {
	case singleNumaNode:
		if k8sTmScope == pod {
			return v1alpha1.SingleNUMANodePodLevel
		}
		// default scope for single-numa-node
		return v1alpha1.SingleNUMANodeContainerLevel
	case restricted:
		return v1alpha1.Restricted
	case bestEffort:
		return v1alpha1.BestEffort
	default:
		return v1alpha1.None
	}
}
