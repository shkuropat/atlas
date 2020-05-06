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
	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"
	"golang.org/x/oauth2/dcrp"
)

var (
	registerClientURL  string
	initialAccessToken string
)

var registerCmd = &cmd.Command{
	Use:   "register",
	Short: "Register new user",
	Long: heredoc.Docf(`
			Register new user
			`,
	),
	Args: func(cmd *cmd.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("requires an filename as argument")
		//}
		return nil
	},
	Run: func(cmd *cmd.Command, args []string) {
		if registerClientURL == "" {
			log.Fatalf("Need to specify --register-url=URL and possibly --initial-access-token=TOKEN\n")
		}

		config := dcrp.Config{
			InitialAccessToken:            initialAccessToken,
			ClientRegistrationEndpointURL: registerClientURL,
			Metadata: dcrp.Metadata{
				ClientName:              "new fluffy ControlPlaneClient",
				TokenEndpointAuthMethod: "client_secret_basic",
				GrantTypes:              []string{"client_credentials"},
				SoftwareID:              "atlas",
				SoftwareVersion:         "0.0.1",
			},
		}
		if cl, err := config.Register(); err != nil {
			log.Errorf("Error: %s\n", err.Error())
		} else {
			log.Infof("Registered:\nclient_id:%s\nclient_secret:%s\n", cl.ClientID, cl.ClientSecret)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&registerClientURL, "register-url", "", "Register client URL")
	rootCmd.PersistentFlags().StringVar(&initialAccessToken, "initial-access-token", "", "Initial access token")

	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(sendCmd)
}
