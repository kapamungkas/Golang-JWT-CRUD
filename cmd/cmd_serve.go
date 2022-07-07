package main

import (
	"betest/internal/databases"
	"betest/internal/routes"
)

func main() {

	// os.Setenv("BASE_URL", "http://localhost:8080")

	// os.Setenv("CONFIG_SMTP_HOST", "mail.mediaceria.com")
	// os.Setenv("CONFIG_SMTP_PORT", "587")
	// os.Setenv("CONFIG_SENDER_NAME", "Backend Test Email <krisnaanggapamungkas@gmail.com>")
	// os.Setenv("CONFIG_AUTH_EMAIL", "no-reply@mediaceria.com")
	// os.Setenv("CONFIG_AUTH_PASSWORD", "macantutul123")

	// os.Setenv("MYSQL_HOST", "localhost")
	// os.Setenv("MYSQL_PORT", "3306")
	// os.Setenv("MYSQL_DB_USER", "root")
	// os.Setenv("MYSQL_DB_NAME", "gotest")
	// os.Setenv("MYSQL_DB_PASSWORD", "password")

	// os.Setenv("JWT_SIGNATURE_KEY", "wndjsa67sd6a5sd78asvahjsda")
	// os.Setenv("REFRESH_JWT_SIGNATURE_KEY", "asjdnajsdhyausd62355326j")

	// os.Setenv("STORAGE_PATH", "./storages/")

	db, _ := databases.Connect()

	routes.InitRoute(db)
}
