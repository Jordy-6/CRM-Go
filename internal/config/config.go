package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"database"`

	Storer struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"storer"`
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Println("Fichier de configuration 'config.yaml' non trouvé, utilisation des valeurs par défaut ou des variables d'environnement.")
		}
	} else {
		log.Println("Fichier de configuration 'config.yaml' chargé avec succès.")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Erreur lors du déchargement de la configuration: %v", err)
	}

	log.Printf("Configuration chargée: DBName=%s, StorerType=%s\n", AppConfig.Database.Name, AppConfig.Storer.Type)
}
