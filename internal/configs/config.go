package configs

import "github.com/spf13/viper"

var config *Config

type option struct {
	configFile   string
	configType   string
	configFolder []string
}

func Init(opts ...Option) error {
	opt := &option{
		configFile:   getDefaultConfigFile(),
		configType:   getDefaultConfigType(),
		configFolder: getDefaultConfigFolder(),
	}

	for _, o := range opts {
		o(opt) // akan mengubah nilai dari opt dari main function
	}

	for _, folder := range opt.configFolder {
		viper.AddConfigPath(folder)
	}

	// viper.SetConfigFile(opt.configFile)
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	config = new(Config)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&config)
}

type Option func(*option)

func getDefaultConfigFolder() []string {
	return []string{"./configs"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolder []string) Option {
	return func(o *option) {
		// o.configFolder = append(o.configFolder, configFolder)
		o.configFolder = configFolder
	}
}

func WithConfigFile(configFile string) Option {
	return func(o *option) {
		o.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(o *option) {
		o.configType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
