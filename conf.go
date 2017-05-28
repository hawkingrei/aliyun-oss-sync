package main

type Config struct {
	Day            int
	Producer       int
	Consumer       int
	constomer      int
	DateList       []string
	Bucket         string
	ACCESS_ID      string
	ACCESS_SEC_KEY string
	Endpoint       string
	PrefixPath     string
}

func NewConfig() *Config {
	return &Config{
		Day:            365 * 4 - 10,
		Producer:       4,
		Consumer:       4 * 4,
		DateList:       []string{},
		Bucket:         "",
		ACCESS_ID:      "",
		ACCESS_SEC_KEY: "",
		Endpoint:       "",
		PrefixPath:     "",
	}
}
