// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}