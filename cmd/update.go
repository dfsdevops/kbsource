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

func readAndUpdateManifests() error {
	reader := io.Reader(os.Stdin)
	decoder := yaml.NewDecoder(reader)
	for {
		var node yaml.Node
		err := decoder.Decode(node)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		content, err := yaml.Marshal(&node)
		if err != nil {
			panic(err)
		}

		fmt.Println("Content", content)
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
