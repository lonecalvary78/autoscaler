/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package system

import (
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/drainability"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/framework"
	"k8s.io/autoscaler/cluster-autoscaler/utils/drain"
)

// KubeSystemNamespace is the namespase includes system pods
const KubeSystemNamespace = "kube-system"

// Rule is a drainability rule on how to handle system pods.
type Rule struct {
	BspDisruptionTimeout time.Duration
}

// New creates a new Rule.
func New(bspDisruptionTimeout time.Duration) *Rule {
	return &Rule{BspDisruptionTimeout: bspDisruptionTimeout}
}

// Name returns the name of the rule.
func (r *Rule) Name() string {
	return "System"
}

// Drainable decides what to do with system pods on node drain.
func (r *Rule) Drainable(drainCtx *drainability.DrainContext, pod *apiv1.Pod, _ *framework.NodeInfo) drainability.Status {
	if isBlockingSystemPod(drainCtx, pod) {
		if r.isBspPassedDisruptionTimeout(pod, drainCtx.Timestamp) {
			return drainability.NewDrainableStatus()
		}
		return drainability.NewBlockedStatus(drain.UnmovableKubeSystemPod, fmt.Errorf("non-daemonset, non-mirrored, non-pdb-assigned kube-system pod present: %s", pod.Name))
	}
	return drainability.NewUndefinedStatus()
}

func isBlockingSystemPod(drainCtx *drainability.DrainContext, pod *apiv1.Pod) bool {
	return pod.Namespace == KubeSystemNamespace && len(drainCtx.RemainingPdbTracker.MatchingPdbs(pod)) == 0
}

func (r *Rule) isBspPassedDisruptionTimeout(pod *apiv1.Pod, drainTime time.Time) bool {
	return !pod.ObjectMeta.CreationTimestamp.IsZero() &&
		drainTime.After(pod.ObjectMeta.CreationTimestamp.Add(r.BspDisruptionTimeout))
}
