package config

type Theme struct {
}

type ThemeType struct {
	Name      string
	Overrides Theme
}

type ThemeOverrides struct {
}

type Config struct {
	Theme ThemeType `toml:"Theme"`
}
