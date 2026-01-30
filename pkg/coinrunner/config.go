package coinrunner

import (
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func (c Config) GetGameState(key string) GameState {
	return GameState(c.GetInt(key))
}

func GetConfig() Config {

	v := viper.New()

	v.SetDefault("startroom", StartPage)

	return Config{v}
}
