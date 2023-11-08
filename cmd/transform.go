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

package cmd

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/christopherfriedrich/tf-archi/internal/extraction"
	"github.com/christopherfriedrich/tf-archi/internal/transformer"
	"github.com/christopherfriedrich/tf-archi/internal/util"
	"github.com/spf13/cobra"
)

func init() {
	transformCmd.Flags().StringVarP(&inputFilePath, "input", "i", "state.json", "Terraform State JSON file to transform.")
	transformCmd.Flags().StringVarP(&outputFilePath, "output", "o", "model.xml", "ArchiMate XML file to write the transformed model to.")
	transformCmd.Flags().StringVarP(&modelName, "model-name", "m", "Terraform State", "Name of the model to create.")
	transformCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(transformCmd)
}

// inputFilePath is the path to the Terraform State JSON file to transform.
var inputFilePath string

// inputFilePath is the path to the place where the ouput should be placed.
var outputFilePath string

// modelName is the name of the model to create.
var modelName string

var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform a Terraform state file into an ArchiMate model.",
	Long:  `Transform a Terraform state file into an ArchiMate model.`,
	Run: func(cmd *cobra.Command, args []string) {
		transformCommandExec(cmd, args)
	},
}

func transformCommandExec(cmd *cobra.Command, _ []string) {
	infrastructureState, err := extraction.ExtractFromTerraformState(inputFilePath)
	cobra.CheckErr(err)
	model, errors := transformer.TransformState(infrastructureState, modelName)
	if len(errors) > 0 {
		fmt.Println("Transformation finished with errors:")
		for _, err := range errors {
			fmt.Printf(" - %sError during transformation: %s%s\n", util.ColorRed, err, util.ColorNone)
		}
	}
	modelXML, err := xml.MarshalIndent(model, "", "  ")
	cobra.CheckErr(err)
	// Write XML to file
	os.WriteFile(outputFilePath, []byte(xml.Header+string(modelXML)), 0644)
}
