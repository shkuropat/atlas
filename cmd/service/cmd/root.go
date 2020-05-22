// Copyright 2020 The Atlas Authors. All rights reserved.
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

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"

	"github.com/binarly-io/binarly-atlas/cmd/common"
	"github.com/binarly-io/binarly-atlas/pkg/config/service"
)

const (
	serviceAddressFlagName = "service-address"

	etcConfigFireDir     = "/etc/atlas"
	homedirConfigFileDir = ".atlas"
	defaultConfigFile    = "service.yaml"

	defaultServiceAddress = ":10000"
)

// CLI parameter variables
var (
	// serviceAddr specifies address of service to use
	serviceAddress string

	// tls specifies whether TLS be used or not
	tls bool
	// tlsCertFile specifies path to certificate file. To be used with TLS
	tlsCertFile string
	// tlsKeyFile specifies path to key file. To be used with TLS
	tlsKeyFile string

	// auth specifies whether to use OAuth2
	auth bool
	// jwtPublicKeyFile specifies path to RSA Public Key file to be used for JWT parsing
	jwtPublicKeyFile string

	// brokers specifies list of Kafka brokers
	brokers string

	// topic specifies Kafka topic work with
	topic string

	// rootCmd represents the base command when called without any sub-commands
	rootCmd = &cmd.Command{
		Use:   "atlas service [COMMAND]",
		Short: "Atlas service.",
		Long: heredoc.Docf(`
			For setting the address of the form HOST:PORT, you can
			- use the flag --%s=%s
			- or you can set '%s: %s' in config $HOME/%s
			`,
			serviceAddressFlagName,
			defaultServiceAddress,
			serviceAddressFlagName,
			defaultServiceAddress,
			defaultConfigFile,
		),
		PersistentPreRun: func(cmd *cmd.Command, args []string) {
			log.Debugf("using address: %s", conf.GetString(serviceAddressFlagName))
		},
	}
)

func init() {
	cmd.OnInitialize(func() {
		common.Init([]string{etcConfigFireDir}, []string{homedirConfigFileDir}, defaultConfigFile)
		config_service.ReadIn()
	})

	// Common section
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&common.ConfigFile, "config", "", fmt.Sprintf("config file (default: %s)", common.PrintConfigFilePaths([]string{etcConfigFireDir, "$HOME/" + homedirConfigFileDir}, defaultConfigFile)))

	// Service section
	rootCmd.PersistentFlags().StringVar(&serviceAddress, "service-address", defaultServiceAddress, fmt.Sprintf("The address of service to use in the format host:port, as %s", defaultServiceAddress))

	// TLS section
	rootCmd.PersistentFlags().BoolVar(&tls, "tls", false, "use TLS connection")
	rootCmd.PersistentFlags().StringVar(&tlsCertFile, "tls-cert-file", "", "The TLS cert file. To be used with TLS")
	rootCmd.PersistentFlags().StringVar(&tlsKeyFile, "tls-key-file", "", "The TLS key file. To be used with TLS")

	// OAuth section
	rootCmd.PersistentFlags().BoolVar(&auth, "oauth", false, "Whether to use OAuth2 for authentication")
	rootCmd.PersistentFlags().StringVar(&jwtPublicKeyFile, "jwt-public-key-file", "", "Public RSA key used for JWT parsing")

	// Kafka section
	rootCmd.PersistentFlags().StringVar(&brokers, "brokers", "", "List of Kafka brokers")
	rootCmd.PersistentFlags().StringVar(&topic, "topic", "", "Kafka topic to work with")

	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
