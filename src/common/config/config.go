package config

import (
	"context"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	awsSSM "github.com/aws/aws-sdk-go-v2/service/ssm"
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
		client *awsSSM.Client
	}
}

var config *Config

func Init() {
	decrypt := true

	ssmConfig, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	ssmClient := awsSSM.NewFromConfig(ssmConfig)

	config = &Config{
		ParamStore: struct {
			client *awsSSM.Client
		}{client: ssmClient},
	}

	config.LoadParams(decrypt)
}

func (config *Config) LoadParams(decrypt bool) {
	config.Params.Steam.Key = config.ParameterGet("steam/SteamApiKey", decrypt)
	config.Params.Steam.Endpoints.GetUsers = config.ParameterGet("steam/endpoints/GetUsers", decrypt)
	config.Params.DataBase.Url = config.ParameterGet("database/URL", decrypt)
}

func (config *Config) ParameterGet(key string, decrypt bool) string {
	param, err := config.ParamStore.client.GetParameter(context.Background(), &awsSSM.GetParameterInput{
		Name:           aws.String("/slanj/api/dev/" + key),
		WithDecryption: aws.Bool(decrypt),
	})
	if err != nil {
		panic(err)
	}

	return *param.Parameter.Value
}

func GetConfig() *Config {
	return config
}
