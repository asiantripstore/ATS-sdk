package config

//ApiConfig from config.yml
type ApiConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
	}
	Gateway `yaml:"gateway"`
}

//Gateway ...
type Gateway struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	HostAPI     string `yaml:"hostApi"`
	RegisterKey string `yaml:"registerKey"`
}
