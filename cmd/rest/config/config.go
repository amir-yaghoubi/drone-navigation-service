package config

type Config struct {
	IsProduction bool `env:"PRODUCTION"`
	Port         int  `env:"PORT" envDefault:"8080"`
	SectorID     int  `env:"SECTOR_ID" envDefault:"1"`
}
