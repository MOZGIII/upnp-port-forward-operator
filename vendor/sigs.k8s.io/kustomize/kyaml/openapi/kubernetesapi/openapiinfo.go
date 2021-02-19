// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by ./scripts/makeOpenApiInfoDotGo.sh; DO NOT EDIT.

package kubernetesapi

import (
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1184"
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1186"
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1188"
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1190"
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1191"
)

const Info = "{title:Kubernetes,version:v1.18.4}\n{title:Kubernetes,version:v1.18.6}\n{title:Kubernetes,version:v1.18.8}\n{title:Kubernetes,version:v1.19.0}\n{title:Kubernetes,version:v1.19.1}"

var OpenAPIMustAsset = map[string]func(string) []byte{
	"v1184": v1184.MustAsset,
	"v1186": v1186.MustAsset,
	"v1188": v1188.MustAsset,
	"v1190": v1190.MustAsset,
	"v1191": v1191.MustAsset,
}

const DefaultOpenAPI = "v1191"
