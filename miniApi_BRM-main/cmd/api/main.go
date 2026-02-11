package main

import (
	"log"
	"net/http"
	"os"

	"miniApi_BRM/internal/db"
	httpHandler "miniApi_BRM/internal/http"
	"miniApi_BRM/internal/repository"
	"miniApi_BRM/internal/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env, usando variables de entorno del sistema")
	}

	dbConfig := db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		Database: getEnv("DB_NAME", "defaultdb"),
		TLS:      getEnv("DB_TLS", "verify"),
		SSLCA:    getEnv("DB_SSL_CA", "/ca.pem"),
	}

	log.Printf("Conectando a: %s:%s@%s:%s/%s", dbConfig.User, "****", dbConfig.Host, dbConfig.Port, dbConfig.Database)

	database, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer database.Close()

	userRepo := repository.NewMySQLUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := httpHandler.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	port := getEnv("SERVER_PORT", "8080")
	addr := "0.0.0.0:" + port
	log.Printf("Servidor iniciado en http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
