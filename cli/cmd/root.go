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
	"github.com/spf13/cobra"
)

type cmd struct {
	loadTesterNs string
	openDuration int32
}

type cmdArgs struct {
	name      string
	namespace string
}

// Config holds the values from the flags
var config = &cmd{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flagger",
	Short: "A kubectl plugin for manually gating Flagger based Canary deployments",
	Long:  "A kubectl plugin for manually gating Flagger based Canary deployments",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&config.loadTesterNs,
		"namespace", "n", "ci", "Namespace where flagger-loadtester is installed")
	rootCmd.PersistentFlags().Int32VarP(&config.openDuration,
		"open-duration", "t", 3, "Duration (in seconds) after which gate must be closed automatically")
}
