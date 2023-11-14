package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvVar() {

	if fh, isExists := os.LookupEnv("CONFIG_FILE"); isExists {
		readFile, err := os.Open(fh)

		if err != nil {
			log.Fatal("Unable to read config file", err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			env := strings.Split(fileScanner.Text(), ":")
			// setting env var, if not set
			if _, isExists := os.LookupEnv(env[0]); !isExists {
				err := os.Setenv(env[0], strings.Trim(env[1], " "))
				if err != nil {
					log.Fatal("Unable to set the env var", err)
				}
			}
		}
	} else {
		if _, isExists := os.LookupEnv("ENV_TYPE"); !isExists {
			log.Print("Defaulting ENV_TYPE to PROD")
			err := os.Setenv("ENV_TYPE", "PROD")
			if err != nil {
				log.Fatal("Unable to set the env var ENV_TYPE", err)
			}
		}
		if _, isExists := os.LookupEnv("PORT"); !isExists {
			log.Fatal("PORT must be set")
		}
		if _, isExists := os.LookupEnv("KEY"); !isExists {
			log.Fatal("KEY must be set")
		}
		if _, isExists := os.LookupEnv("PUB_KEY"); !isExists {
			log.Fatal("PUB_KEY must be set")
		}
	}
}
