/*
Copyright 2023 The tf-archi Authors.

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

package extraction_test

import (
	"testing"

	"github.com/christopherfriedrich/tf-archi/internal/extraction"
)

func TestConvertTerraformResourceIdToGenericId(t *testing.T) {
	resourceID := "/subscriptions/58b1e9asda-eqwwtt-4c11-342fd3a-da34dd4fr2342/resourceGroups/example-rg"
	want := "subscriptions-58b1e9asda-eqwwtt-4c11-342fd3a-da34dd4fr2342-resourceGroups-example-rg"
	testee := extraction.ConvertTerraformResourceIDToGenericID(resourceID)
	if !(want == testee) {
		t.Fatalf("convertTerraformResourceIdToGenericId() = '%v', want '%v'", testee, want)
	}
}
