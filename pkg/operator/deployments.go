/*
Copyright The Voyager Authors.

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

package operator

import (
	"github.com/appscode/go/log"
	"github.com/golang/glog"
	"k8s.io/client-go/tools/cache"
	"kmodules.xyz/client-go/tools/queue"
	wpi "kmodules.xyz/webhook-runtime/apis/workload/v1"
)

func (op *Operator) initDeploymentWatcher() {
	op.dpInformer = op.kubeInformerFactory.Apps().V1().Deployments().Informer()
	op.dpQueue = queue.New("Deployment", op.MaxNumRequeues, op.NumThreads, op.reconcileDeployment)
	op.dpInformer.AddEventHandler(queue.NewDeleteHandler(op.dpQueue.GetQueue()))
	op.dpLister = op.kubeInformerFactory.Apps().V1().Deployments().Lister()
}

func (op *Operator) reconcileDeployment(key string) error {
	_, exists, err := op.dpInformer.GetIndexer().GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}
	if !exists {
		glog.Warningf("Deployment %s does not exist anymore\n", key)
		if ns, name, err := cache.SplitMetaNamespaceKey(key); err != nil {
			return err
		} else {
			return op.restoreDeployment(name, ns)
		}
	}
	return nil
}

// requeue ingress if user deletes haproxy-deployment
func (op *Operator) restoreDeployment(name, ns string) error {
	items, err := op.listIngresses()
	if err != nil {
		return err
	}
	for i := range items {
		ing := &items[i]
		if ing.DeletionTimestamp == nil &&
			ing.ShouldHandleIngress(op.IngressClass) &&
			ing.Namespace == ns &&
			ing.WorkloadKind() == wpi.KindDeployment &&
			ing.OffshootName() == name {
			if key, err := cache.MetaNamespaceKeyFunc(ing); err != nil {
				return err
			} else {
				op.getIngressQueue(ing.APISchema()).Add(key)
				log.Infof("Add/Delete/Update of haproxy deployment %s/%s, Ingress %s re-queued for update", ns, name, key)
				break
			}
		}
	}
	return nil
}
