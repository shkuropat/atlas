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
)

const (
	serviceAddressFlagName = "service-address"

	defaultConfigFileName = ".atlas-client.yaml"
	defaultServiceAddress = "localhost:10000"
)

// CLI parameter variables
var (
	// verbose specifies whether app should be verbose
	verbose bool

	// configFile defines path to config file to be used
	configFile string

	// serviceAddr specifies address of service to use
	serviceAddress string

	tls                bool
	caFile             string
	serverHostOverride string

	auth         bool
	clientID     string
	clientSecret string
	tokenURL     string

	// rootCmd represents the base command when called without any sub-commands
	rootCmd = &cmd.Command{
		Use:   "atlas client [COMMAND]",
		Short: "Atlas client.",
		Long: heredoc.Docf(`
			For setting the address of the form HOST:PORT, you can
			- use the flag --%s=%s
			- or you can set '%s: %s' in config $HOME/%s
			`,
			serviceAddressFlagName,
			defaultServiceAddress,
			serviceAddressFlagName,
			defaultServiceAddress,
			defaultConfigFileName,
		),
		PersistentPreRun: func(cmd *cmd.Command, args []string) {
			log.Debugf("using address: %s", conf.GetString(serviceAddressFlagName))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf("config file (default is $HOME/%s)", defaultConfigFileName))

	rootCmd.PersistentFlags().StringVar(&serviceAddress, "service-address", defaultServiceAddress, fmt.Sprintf("The address of service to use in the format host:port, as %s", defaultServiceAddress))

	// TLS section
	rootCmd.PersistentFlags().BoolVar(&tls, "tls", false, "use TLS connection")
	rootCmd.PersistentFlags().StringVar(&caFile, "ca-file", "", "CA root cert file")
	rootCmd.PersistentFlags().StringVar(&serverHostOverride, "server-host-override", "x.test.youtube.com", "server name use to verify the hostname returned by TLS handshake")

	// OAuth section
	rootCmd.PersistentFlags().BoolVar(&auth, "oauth", false, "Whether to use OAuth2 for authentication")
	rootCmd.PersistentFlags().StringVar(&clientID, "client-id", "", "ClientID used for Identity server access")
	rootCmd.PersistentFlags().StringVar(&clientSecret, "client-secret", "", "ClientSecret used for Identity server access")
	rootCmd.PersistentFlags().StringVar(&tokenURL, "token-url", "", "URL of Identity server's token service")

	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	common.InitLog(verbose)
	common.InitConfig(configFile, defaultConfigFileName)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}