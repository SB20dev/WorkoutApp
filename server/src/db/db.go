package db

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Datasource string `yaml:"datasource"`
}

type Configs map[string]Config

func (c *Configs) Open(env string) (*gorm.DB, error) {
	dsn := (*c)[env].Datasource
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ReadConfigs(filePath string) (Configs, error) {
	// ファイル読み込み
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	// 設定情報をConfigsに束縛
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func ConnectDB(filePath string, env string) *gorm.DB {
	configs, err := ReadConfigs(filePath)
	if err != nil {
		panic(err)
	}
	db, err := configs.Open(env)
	if err != nil {
		panic(err)
	}
	return db
}
