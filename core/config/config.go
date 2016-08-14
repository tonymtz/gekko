package config

// Config ...
type Config struct {
	AppName string `yaml:"app_name"`
	Port    int    `yaml:"http_port"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`

	GoogleID       string `yaml:"google_id"`
	GoogleSecret   string `yaml:"google_secret"`
	GoogleCallback string `yaml:"google_callback"`

	DropboxID       string `yaml:"dropbox_id"`
	DropboxSecret   string `yaml:"dropbox_secret"`
	DropboxCallback string `yaml:"dropbox_callback"`
}
