/*
Copyright 2021 Red Hat

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

package service

import (
	"time"

	corev1 "k8s.io/api/core/v1"
)

// Service -
type Service struct {
	service         *corev1.Service
	timeout         time.Duration
	clusterIPs      []string
	externalIPs     []string
	ipFamilies      []corev1.IPFamily
	serviceHostname string
}

// GenericServiceDetails -
type GenericServiceDetails struct {
	Name      string
	Namespace string
	Labels    map[string]string
	Selector  map[string]string
	Port      GenericServicePort
	ClusterIP string
}

// GenericServicePort -
type GenericServicePort struct {
	Name     string
	Port     int32
	Protocol corev1.Protocol // corev1.ProtocolTCP/ corev1.ProtocolUDP/ corev1.ProtocolSCTP - https://pkg.go.dev/k8s.io/api@v0.23.6/core/v1#Protocol
}

// MetalLBServiceDetails -
type MetalLBServiceDetails struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	Labels      map[string]string
	Selector    map[string]string
	Port        GenericServicePort
}

const (
	// MetalLBAddressPoolAnnotation -
	MetalLBAddressPoolAnnotation = "metallb.universe.tf/address-pool"
	// MetalLBAllowSharedIPAnnotation -
	MetalLBAllowSharedIPAnnotation = "metallb.universe.tf/allow-shared-ip"
	// MetalLBLoadBalancerIPs -
	MetalLBLoadBalancerIPs = "metallb.universe.tf/loadBalancerIPs"
)
