
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ConfigMessage struct {
	LogLevel    string `json:"log_level"`
	GRPCPort    string `json:"grpc_port"`
	Environment string `json:"environment"`
	DBURL       string `json:"db_url"`
}

type DomainMessage struct {
	Endpoint string `json:"endpoint"`
	Domain   string `json:"domain"`
	LogLevel string `json:"log_level"`
}

type LoadCapabilityMessage struct {
	CPU    string `json:"cpus"`
	MEMORY  string `json:"memorys"`
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnvVariable(key string) string {
	value := os.Getenv(key)
	return value
}

func main() {
	r := gin.Default()

	r.GET("/config", func(c *gin.Context) {
		loadEnv()

		logLevel := getEnvVariable("LOG_LEVEL")
		grpcPort := getEnvVariable("GRPC_PORT")
		environment := getEnvVariable("ENVIRONMENT")
		dbURL := getEnvVariable("DB_URL")

		message := ConfigMessage{
			LogLevel:    logLevel,
			GRPCPort:    grpcPort,
			Environment: environment,
			DBURL:       dbURL,
		}

		c.JSON(200, message)
	})

	r.GET("/domain", func(c *gin.Context) {
		host := c.Request.Host
		fmt.Printf("Host: %s\n", host)

		logLevel := getEnvVariable("LOG_LEVEL")

		message := DomainMessage{
			Endpoint: "/domain",
			Domain:   host,
			LogLevel: logLevel,
		}

		c.JSON(200, message)
	})

	r.GET("/load-capability", func(c *gin.Context) {
		cpus := getEnvVariable("CPU")
		memorys := getEnvVariable("MEMORY")

		message := LoadCapabilityMessage{
			CPU:    cpus,
			MEMORY: memorys,
		}

		c.JSON(200, message)
	})

	port := ":8070"
	fmt.Printf("Server is listening on port %s...\n", port)
	r.Run(port)
}
