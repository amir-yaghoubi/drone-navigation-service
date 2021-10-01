package config

type Config struct {
	IsProduction bool `env:"PRODUCTION"`
	Port         int  `env:"PORT" envDefault:"8080"`
	SectionID    int  `env:"SECTION_ID" envDefault:"1"`
}
