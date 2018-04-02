package utils

import (
	"github.com/spf13/viper"
)

// init viper config
func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		Log.Errorf(ErrTemplate(), GetFunctionName(InitConfig), err)
	}

	// database default
	viper.SetDefault("username", "postgres")
	viper.SetDefault("password", "")
	viper.SetDefault("database", "postgres")

	// gRPC default
	viper.SetDefault("grpc.network", "tcp")
	viper.SetDefault("grpc.host", "")
	viper.SetDefault("grpc.port", "50051")
}
