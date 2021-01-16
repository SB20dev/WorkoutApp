package db

import (
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Datasource string `yaml:"datasource"`
}

type Configs map[string]Config

func (c *Configs) Open(env string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", (*c)[env].Datasource)
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
