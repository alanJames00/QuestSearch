// global config setup
package configs

import "os"

type Config struct {
	Addr string 
	MongoURI string
	DBName string
}

func LoadConfig() *Config {
	// get the params from current env
	return &Config{
		Addr: getEnv("PORT", ":8000"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName: getEnv("DB_NAME", "test"),
	}
}


func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
