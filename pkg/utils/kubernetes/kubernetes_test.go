// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"testing"
)

func TestKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Suite")
}

var _ = Describe("kubernetes", func() {
	Describe("#CreateTwoWayMergePatch", func() {
		It("should fail for two different object types", func() {
			_, err := CreateTwoWayMergePatch(&corev1.ConfigMap{}, &corev1.Secret{})
			Expect(err).To(HaveOccurred())
		})

		It("Should correctly create a patch", func() {
			patch, err := CreateTwoWayMergePatch(
				&corev1.ConfigMap{Data: map[string]string{"foo": "bar"}},
				&corev1.ConfigMap{Data: map[string]string{"foo": "baz"}})

			Expect(err).NotTo(HaveOccurred())
			Expect(patch).To(Equal([]byte(`{"data":{"foo":"baz"}}`)))
		})
	})

	DescribeTable("#IsEmptyPatch",
		func(patch string, expected bool) {
			Expect(IsEmptyPatch([]byte(patch))).To(Equal(expected))
		},
		Entry("non-empty-patch", `{"foo": "bar"}`, false),
		Entry("non-json-patch", `random input`, false),
		Entry("empty string", ``, true),
		Entry("empty string with spaces", `  `, true),
		Entry("empty json object", `{}`, true),
		Entry("empty json object with spaces", ` { } `, true),
	)
})
