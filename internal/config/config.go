package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// //////////////////////////////////////////////////////////////////////
// Config structer
// //////////////////////////////////////////////////////////////////////
type Config struct {
	AppEnv    string
	AppPort   string
	PostgeSQL string
	DBDriver  string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
}

// //////////////////////////////////////////////////////////////////////
// LoadConfig
// //////////////////////////////////////////////////////////////////////
func LoadConfig() *Config {
	root := findProjectRoot()

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}

	envPath := filepath.Join(root, "env", appEnv+".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: Cannot load from local %s", envPath)
	}

	cfg := &Config{
		AppEnv:    appEnv,
		AppPort:   getEnv("APP_PORT", "3000"),
		PostgeSQL: getEnv("PostgeSQL_URL", ""),
		DBDriver:  getEnv("DB_DRIVER", "postgres"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASSWORD", "root1234"),
		DBName:    getEnv("DB_NAME", "mydb"),
	}

	log.Println("Config loaded.")
	return cfg
}

// //////////////////////////////////////////////////////////////////////
// Get ENV
// //////////////////////////////////////////////////////////////////////
func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

// ======================================================
//
//	หาตำแหน่ง project root อัตโนมัติ (สำคัญที่สุด)
//
// ======================================================
func findProjectRoot() string {
	dir, _ := os.Getwd()

	for {
		// ถ้าข้างในมีโฟลเดอร์ env → ถือว่าเจอ project root
		if exists(filepath.Join(dir, "env")) {
			return dir
		}

		// ถ้าถึง root filesystem แล้ว → ใช้ CWD เดิม
		parent := filepath.Dir(dir)
		if parent == dir {
			return dir
		}

		dir = parent
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
