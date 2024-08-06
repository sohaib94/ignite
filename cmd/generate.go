/*
Copyright © 2024 Sohaib Ashraf <sohaib.ashraf94@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/sohaib94/ignite/ignite"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("generate called")

		if !(strings.HasSuffix(generateFilePath, ".yml") || strings.HasSuffix(generateFilePath, ".yaml")) {
			log.Printf("Only yaml format supported - please provide a yaml file path ending in .yaml or .yml")
			return
		}

		i := ignite.Ignite{OutputParentPath: outputParentPath}
		f := &ignite.IgniteFile{
			Path:    generateFilePath,
			Reader:  os.ReadFile,
		}

		err := i.Handle(f)
		if err != nil {
			log.Println("Failed to ignite project.")
			return
		}
	},
}

var (
	generateFilePath, outputParentPath string
)

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.PersistentFlags().StringVarP(&generateFilePath, "file", "f", "./ignite.yml", "Yaml file containing ignition details. Default is ./ignite.yaml")
	generateCmd.PersistentFlags().StringVarP(&outputParentPath, "output", "o", ".", "Path to local directory where to create repo. Default to create within directory `generate` is used.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
