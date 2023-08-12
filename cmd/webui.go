// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"github.com/kppotato/coredns/controller"
	"github.com/kppotato/coredns/dao/etcd"
	"github.com/kppotato/coredns/g"
	"github.com/kppotato/coredns/httpGin"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// webuiCmd represents the webui command
var webuiCmd = &cobra.Command{
	Use:   "webui",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if g.Etcd_url == nil {
			os.Exit(1)
		}
		if len(g.Etcd_url) == 0 {
			os.Exit(1)
		}
		if g.Etcd_path == "" {
			g.Etcd_path = "/skydns"
		} else {
			if !strings.HasPrefix(g.Etcd_path, "/") {
				g.Etcd_path = "/" + g.Etcd_path
			}
		}
		//初始化检测etcd链接情况
		etcd.OninitCheck()
	},
	Run: func(cmd *cobra.Command, args []string) {
		controller.Oninit()
		httpGin.StartHttp()
	},
}

func init() {
	RootCmd.AddCommand(webuiCmd)
	webuiCmd.Flags().StringSliceVar(&g.Etcd_url, "etcdurl", nil, "etcd url not empty")
	webuiCmd.Flags().StringVar(&g.Etcd_path, "etcdpath", "", "etcd url not empty")
}
