/*
Copyright 2022 Red Hat
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

package helpers

import (
	"github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateNamespace creates a Kubernetes Namespace resource.
//
// Example usage:
//
//	th.CreateNamespace("test-namespace")
//
// Note: the namespace should be unique and not be already present in the cluster, otherwise,
// the function will fail.
func (tc *TestHelper) CreateNamespace(name string) *corev1.Namespace {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	gomega.Expect(tc.K8sClient.Create(tc.Ctx, ns)).Should(gomega.Succeed())
	return ns
}

// DeleteNamespace deletes a Kubernetes Namespace resource.
//
// Example usage:
//
//	th.DeleteNamespace("test-namespace")
//
// or
//
//	DeferCleanup(th.DeleteNamespace, namespace)
//
// Note: the namespace should exist in the cluster, otherwise, the function will fail.
func (tc *TestHelper) DeleteNamespace(name string) {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	gomega.Expect(tc.K8sClient.Delete(tc.Ctx, ns)).Should(gomega.Succeed())
}
