// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

//flags
var (
	StaticHost string
)

// frontCmd represents the front command
var frontCmd = &cobra.Command{
	Use:   "front",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start listen: ", StaticHost)
		dir, _ := os.Getwd()
		dir = dir + "/front"
		h := http.FileServer(http.Dir(dir))

		err := http.ListenAndServe(StaticHost, h)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(frontCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// frontCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// frontCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	frontCmd.Flags().StringVarP(&StaticHost, "host", "H", ":9090", "static file host")
}
