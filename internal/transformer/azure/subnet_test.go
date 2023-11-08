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

package azure_test

import (
	"strings"
	"testing"

	"github.com/christopherfriedrich/tf-archi/internal/transformer/azure"
)

// Write a test for the AzureSubnetTransformer.Type() method.
func TestAzureSubnetTransformer_Type(t *testing.T) {
	testee := azure.NewSubnetTransformer()
	want := "azurerm_subnet"
	got := testee.Type()
	if strings.Compare(want, got) != 0 {
		t.Errorf("AzureSubnetTransformer.Type() = %q, want %q", got, want)
	}
}
