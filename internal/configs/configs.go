package configs

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const envPrefix = "ova"

type JokeServiceConfig struct {
	GRPC    GRPCServerConfig
	Flusher FlusherConfig
	DB      DBConfig
	Metrics MetricsServerConfig
	Broker  BrokerConfig
}

type GRPCServerConfig struct {
	Addr string
}

type FlusherConfig struct {
	ChunkSize int
}

type DBConfig struct {
	Host string
	Port uint16
	Name string
	User string
	Pass string
}

func (d DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Pass, d.Name,
	)
}

type MetricsServerConfig struct {
	Addr string
}

type BrokerConfig struct {
	Addrs []string
}

func GetConfig() (*JokeServiceConfig, error) {
	if err := readConfig(); err != nil {
		return nil, err
	}

	initEnvs()

	if err := initFlags(); err != nil {
		return nil, err
	}

	config := &JokeServiceConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}

func readConfig() error {
	viper.SetConfigName("config")                  // name of config file (without extension)
	viper.SetConfigType("yml")                     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs/ova-joke-api/") // path to look for the config file in
	viper.AddConfigPath("$HOME/.ova-joke-api")     // call multiple times to add many search paths
	viper.AddConfigPath(".")                       // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok { //nolint:errorlint
			log.Warn().Msgf("configuration file not found: %v", err)
			return nil
		}
		return err
	}

	return nil
}

func initEnvs() {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initFlags() error {
	pflag.Parse()
	return viper.BindPFlags(pflag.CommandLine)
}
