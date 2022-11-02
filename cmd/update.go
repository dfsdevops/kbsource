/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update /path/to/output/directory",
	Short: "Update files in specified directory",
	Long: `Read from stdin and update commands in specified directory.

By default output filename format is [kind].[name].yaml`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Outputting to directory", args[0])
		readAndUpdateManifests()
	},
}

type K8sobject struct {
	Kind       string `yaml:"kind"`
	Name       string `yaml:"metadata.name"`
	ApiVersion string `yaml:"apiVersion"`
}

func readAndUpdateManifests() error {
	reader := io.Reader(os.Stdin)
	decoder := yaml.NewDecoder(reader)
	for {
		// var node yaml.Node
		data := make(map[string]interface{})
		// var k8sobj K8sobject
		err := decoder.Decode(data)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		if len(data) == 0 {
			break
		}
		fmt.Println(data)

		metadata := data["metadata"].(map[string]interface{})

		// fmt.Println(data)
		fmt.Println("\napiVersion: ", data["apiVersion"])
		fmt.Println("kind: ", data["kind"])
		fmt.Println("name: ", metadata["name"])
		// yaml.Unmarshal(data, &k8sobj)
		// fmt.Printf("name: %s\n", k8sobj.Name)

	}

	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
