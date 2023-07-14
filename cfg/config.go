package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	cfgSingleton sync.Once
	Instance     *Config
)

type Config struct {
	GitHubAccessToken string `mapstructure:"GITHUB_ACCESS_TOKEN"`
	GitHubOwnerName   string `mapstructure:"GITHUB_OWNER_NAME"`
	GitHubRepoName    string `mapstructure:"GITHUB_REPO_NAME"`
	GitHubIssueNumber int    `mapstructure:"GITHUB_ISSUE_NUMBER"`
}

func LoadEnv() {
	// notify that app try to load config file
	log.Println("Load configuration file . . . .")
	cfgSingleton.Do(func() {
		// find environment file
		viper.AutomaticEnv()
		// error handling for specific case
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				panic(".env file not found!, please copy .env.example and paste as .env")
			}
			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
		// notify that config file is ready
		log.Println("configuration file: ready")
		// extract config to struct
		if err := viper.Unmarshal(&Instance); err != nil {
			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
	})
}
