package config

import (
	awsParamStore "github.com/PaddleHQ/go-aws-ssm"
	"github.com/aws/aws-sdk-go/aws"
)

type DatabaseConfig struct {
	Url string
}

type SteamConfig struct {
	Key       string
	Endpoints struct {
		GetUsers string
	}
}

type Config struct {
	Params struct {
		DataBase DatabaseConfig
		Steam    SteamConfig
	}
	ParamStore struct {
		client *awsParamStore.ParameterStore
	}
}

var config *Config

func Init() {
	decrypt := true
	region := aws.String("eu-central-1")

	paramStore, err := awsParamStore.NewParameterStore(&aws.Config{Region: region})
	if err != nil {
		panic(err)
	}

	config = &Config{
		ParamStore: struct {
			client *awsParamStore.ParameterStore
		}{client: paramStore},
	}

	config.LoadParams(decrypt)
}

func (config *Config) LoadParams(decrypt bool) {
	config.Params.Steam.Key = config.ParameterGet("steam/SteamApiKey", decrypt)
	config.Params.Steam.Endpoints.GetUsers = config.ParameterGet("steam/endpoints/GetUsers", decrypt)
	config.Params.DataBase.Url = config.ParameterGet("database/URL", decrypt)
}

func (config *Config) ParameterGet(key string, decrypt bool) string {
	param, err := config.ParamStore.client.GetParameter("/slanj/api/dev/"+key, decrypt)
	if err != nil {
		panic("No value for key:" + "/slanj/api/dev/" + key)
	}

	return param.GetValue()
}

func GetConfig() *Config {
	return config
}
