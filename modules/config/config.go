package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type Config struct {
	cfg *config.Config
}

type Yandex struct {
	CloudId  string `example:"cloud_id"`
	FolderId string `example:"folder_id"`
}

type ServiceAccount struct {
	Id          pulumi.StringOutput
	Name        string
	Description string
	Roles       []string
}

func NewConfig(ctx *pulumi.Context) *Config {

	cf := config.New(ctx, "")
	config := &Config{
		cfg: cf,
	}

	return config
}

func (c *Config) GetYandexData() *Yandex {
	var ya Yandex
	c.cfg.RequireObject("yandex", &ya)
	return &ya
}
func (c *Config) GetServiceAccount(data string) *ServiceAccount {
	var sa ServiceAccount
	c.cfg.GetObject(data, &sa)
	return &sa
}
