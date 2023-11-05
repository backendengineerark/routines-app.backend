package configs

import "github.com/spf13/viper"

type Config struct {
	WebServerPort              string `mapstructure:"SERVER_PORT"`
	DBDriver                   string `mapstructure:"DB_DRIVER"`
	DBHost                     string `mapstructure:"DB_HOST"`
	DBPort                     string `mapstructure:"DB_PORT"`
	DBUser                     string `mapstructure:"DB_USER"`
	DBPassword                 string `mapstructure:"DB_PASSWORD"`
	DBName                     string `mapstructure:"DB_NAME"`
	CreateTodayRoutineTaskCron string `mapstructure:"CREATE_ROUTINES_TASK_CRON"`
}

func LoadConfig(path string) *Config {
	var cfg *Config

	viper.SetConfigName("app_config")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
