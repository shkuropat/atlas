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

package common

import (
	"strings"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	conf "github.com/spf13/viper"
)

var(
	// ConfigFile defines path to config file to be used
	ConfigFile string
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig(defaultConfigFile string) {

	if ConfigFile == "" {
		// Use config file from home directory
		homedir, err := homedir.Dir()
		if err != nil {
			log.Fatalf("unable to find homedir %v", err)
		}
		// Look for default config file in HOME dir
		conf.AddConfigPath(homedir)
		conf.SetConfigName(defaultConfigFile)
		conf.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		log.Infof("looking for config %s in %s", defaultConfigFile, homedir)
	} else {
		// Look for explicitly specified config file
		conf.SetConfigFile(ConfigFile)
		log.Infof("looking for config %s", ConfigFile)
	}

	// By default empty environment variables are considered unset and will fall back to the next configuration source.
	// To treat empty environment variables as set, use the AllowEmptyEnv method.
	conf.AllowEmptyEnv(false)
	// Check for an env var with a name matching the key uppercased and prefixed with the EnvPrefix
	// Prefix has "_" added automatically, so no need to say 'ATLAS_'
	conf.SetEnvPrefix("ATLAS")
	// SetEnvKeyReplacer allows you to use a strings.Replacer object to rewrite Env keys to an extent.
	// This is useful if you want to use - or something in your Get() calls, but want your environmental variables to use _ delimiters.
	conf.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	// Check ENV variables for all keys set in config, default & flags
	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err == nil {
		log.Debugf("config file used: %s", conf.ConfigFileUsed())
	} else if _, ok := err.(conf.ConfigFileNotFoundError); ok {
		// Config file not found
		log.Infof("no config file found")
	} else {
		// Config file was found but another error was produced
		log.Errorf("unable to read config file: %v", err)
	}
}
