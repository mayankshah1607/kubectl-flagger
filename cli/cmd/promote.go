/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/mayankshah1607/kubectl-flagger/pkg/flagger"
	"github.com/spf13/cobra"
)

// promoteCmd represents the promote command
var promoteCmd = &cobra.Command{
	Use:     "promote [canary name] [canary namespace]",
	Short:   "Promote a canary deployment to primary",
	Long:    "Promote a canary deployment to primary",
	Example: "kubectl flagger promote podinfo test -n flagger",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatalf("Invalid usage")
		}

		opts := parseArgs(args)
		err := flagger.Promote(opts.name, opts.namespace, config.loadTesterNs)
		if err != nil {
			log.Fatalf("failed to promote:\n%s", err)
		}
		log.Printf("Successfully promoted canary/%s in namespace/%s", opts.name, opts.namespace)
	},
}

func init() {
	rootCmd.AddCommand(promoteCmd)
}

func parseArgs(args []string) *cmdArgs {
	return &cmdArgs{
		name:      args[0],
		namespace: args[1],
	}
}
