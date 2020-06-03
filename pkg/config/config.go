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

package config

import (
	hd "github.com/mitchellh/go-homedir"
	"strings"

	log "github.com/sirupsen/logrus"
	conf "github.com/spf13/viper"
)

var (
	// Check for an env var with a name matching the key upper-cased and prefixed with the EnvPrefix
	// Prefix has "_" added automatically, so no need to say 'ATLAS_'
	configEnvVarPrefix = "ATLAS"
)

var (
	// ConfigFile defines path to config file to be used
	ConfigFile string
)

func FullConfigFilePaths(paths []string, filename string) []string {
	res := []string{}
	for _, path := range paths {
		res = append(res, path+"/"+filename)
	}

	return res
}

func PrintConfigFilePaths(paths []string, filename string) string {
	return strings.Join(FullConfigFilePaths(paths, filename), " ")
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig(rootPaths, homeRelativePaths []string, defaultConfigFile string) {
	log.Info("InitConfig()")

	if ConfigFile == "" {
		// Use config file from home directory
		homedir, err := hd.Dir()
		if err != nil {
			log.Fatalf("InitConfig() - unable to find homedir %v", err)
		}
		// Look for default config file in root-based list of dirs, such as /etc, /opt/etc ...
		for _, path := range rootPaths {
			log.Infof("InitConfig() - add root path to look for config: %v", path)
			conf.AddConfigPath(path)
		}
		// Look for default config file in HOMEDIR-based list of dirs, such as $HOME/.atlas ...
		for _, path := range homeRelativePaths {
			homeRelativePath := homedir + "/" + path
			log.Infof("InitConfig() - add home relative path to look for config: %v : %v", path, homeRelativePath)
			conf.AddConfigPath(homeRelativePath)
		}
		log.Infof("InitConfig() - add config file name to look for: %v", defaultConfigFile)

		conf.SetConfigName(defaultConfigFile)
		conf.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	} else {
		// Look for explicitly specified config file
		conf.SetConfigFile(ConfigFile)
		log.Infof("InitConfig() - looking for explicitly specified config: %s", ConfigFile)
	}

	log.Infof("InitConfig() - set env prefix: %v_", configEnvVarPrefix)
	// By default empty environment variables are considered unset and will fall back to the next configuration source.
	// To treat empty environment variables as set, use the AllowEmptyEnv method.
	conf.AllowEmptyEnv(false)
	// Check for an env var with a name matching the key upper-cased and prefixed with the EnvPrefix
	// Prefix has "_" added automatically, so no need to say 'ATLAS_'
	conf.SetEnvPrefix(configEnvVarPrefix)
	// SetEnvKeyReplacer allows you to use a strings.Replacer object to rewrite Env keys to an extent.
	// This is useful if you want to use - or something in your Get() calls, but want your environmental variables to use _ delimiters.
	conf.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	// Check ENV variables for all keys set in config, default & flags
	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err == nil {
		log.Infof("InitConfig() - config file used: %s", conf.ConfigFileUsed())
	} else if _, ok := err.(conf.ConfigFileNotFoundError); ok {
		// Config file not found
		log.Infof("InitConfig() - no config file found")
	} else {
		// Config file was found but another error was produced
		log.Errorf("InitConfig() - unable to read config file: %v", err)
	}
}
