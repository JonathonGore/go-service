package config

type Config interace {
	GetStrings(keys []string) []string
	GetString(key string) string
	GetInts(keys []string) []string
	GetString(key string) string
}
