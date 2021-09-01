/*
Copyright 2014 The Kubernetes Authors.

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

package main

import (
	"context"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/openziti-incubator/kubectl/pkg/cmd"
	"github.com/openziti-incubator/kubectl/pkg/cmd/plugin"
	"github.com/spf13/cobra"

	"github.com/openziti/sdk-golang/ziti"
	"github.com/openziti/sdk-golang/ziti/config"

	"github.com/openziti-incubator/kubectl/pkg/util/logs"
	"github.com/sirupsen/logrus"

	// Import to initialize client auth plugins.
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

var configFilePath string
var serviceName string

type ZitiFlags struct {
	ZConfig string
	Service string
}

var zFlags = ZitiFlags{}

func main() {
	rand.Seed(time.Now().UnixNano())
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.WrapConfigFn = wrapConfigFn

	command := cmd.NewDefaultKubectlCommandWithArgsAndConfigFlags(cmd.NewDefaultPluginHandler(plugin.ValidPluginFilenamePrefixes), os.Args, os.Stdin, os.Stdout, os.Stderr, kubeConfigFlags)
	command = setZitiFlags(command)
	command.PersistentFlags().Parse(os.Args)

	configFilePath = command.Flag("ZConfig").Value.String()
	serviceName = command.Flag("Service").Value.String()

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// cliflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func dialFunc(ctx context.Context, network, address string) (net.Conn, error) {
	service := serviceName
	configFile, err := config.NewFromFile(configFilePath)

	if err != nil {
		logrus.WithError(err).Error("Error loading config file")
		os.Exit(1)
	}

	context := ziti.NewContextWithConfig(configFile)
	return context.Dial(service)
}

func wrapConfigFn(restConfig *rest.Config) *rest.Config {

	restConfig.Dial = dialFunc
	return restConfig
}

func setZitiFlags(command *cobra.Command) *cobra.Command {

	command.PersistentFlags().StringVarP(&zFlags.ZConfig, "ZConfig", "c", "", "Path to ziti config file")
	command.PersistentFlags().StringVarP(&zFlags.Service, "Service", "S", "", "Service name")

	return command
}
