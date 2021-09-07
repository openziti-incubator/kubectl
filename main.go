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
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/openziti-incubator/kubectl/pkg/cmd"
	"github.com/openziti-incubator/kubectl/pkg/cmd/plugin"
	"github.com/spf13/cobra"

	"github.com/openziti/sdk-golang/ziti"
	"github.com/openziti/sdk-golang/ziti/config"

	"github.com/openziti-incubator/kubectl/pkg/util/logs"
	"github.com/sirupsen/logrus"

	// Import to initialize client auth plugins.
	"github.com/go-yaml/yaml"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var configFilePath string
var serviceName string

type ZitiFlags struct {
	zConfig string
	service string
}

type Context struct {
	ZConfig string `yaml:"zConfig"`
	Service string `yaml:"service"`
}

type MinKubeConfig struct {
	Contexts []struct {
		Context Context `yaml:"context"`
		Name    string  `yaml:"name"`
	} `yaml:"contexts"`
}

var zFlags = ZitiFlags{}

func main() {
	rand.Seed(time.Now().UnixNano())
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.WrapConfigFn = wrapConfigFn

	command := cmd.NewDefaultKubectlCommandWithArgsAndConfigFlags(cmd.NewDefaultPluginHandler(plugin.ValidPluginFilenamePrefixes), os.Args, os.Stdin, os.Stdout, os.Stderr, kubeConfigFlags)
	command = setZitiFlags(command)
	command.PersistentFlags().Parse(os.Args)

	configFilePath = command.Flag("zConfig").Value.String()
	serviceName = command.Flag("service").Value.String()

	//readConfig(getKubeconfigPath(command))

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
		logrus.WithError(err).Error("Error loading ziti config file")
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

	command.PersistentFlags().StringVarP(&zFlags.zConfig, "zConfig", "c", "", "Path to ziti config file")
	command.PersistentFlags().StringVarP(&zFlags.service, "service", "S", "", "Service name")

	return command
}

func getKubeconfigPath(command *cobra.Command) string {
	kubeconfig := command.Flag("kubeconfig").Value.String()

	home := homedir.HomeDir()
	if kubeconfig == "" && home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else if home == "" {
		logrus.Error("Could not find kubeconfig file in the default location")
		os.Exit(1)
	}

	return kubeconfig
}

func readConfig(kubeconfig string) {

	logrus.Infof("kubeconfig: ", kubeconfig)

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)

	currentContext := config.CurrentContext

	filename, _ := filepath.Abs(kubeconfig)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var minKubeConfig MinKubeConfig

	err = yaml.Unmarshal(yamlFile, &minKubeConfig)
	if err != nil {
		panic(err)
	}

	var context Context
	for _, ctx := range minKubeConfig.Contexts {

		if ctx.Name == currentContext {
			context = ctx.Context
		}
	}

	if configFilePath == "" {
		configFilePath = context.ZConfig
	}

	if serviceName == "" {
		serviceName = context.Service
	}
}
