package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:Zxcvbnm123@62.72.23.250:5432/fortest")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "вах вах",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		rows, err := dbpool.Query(context.Background(), "SELECT id, name, email FROM users")
		if err != nil {
			log.Println("Query error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database query error"})
			return
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name, email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				log.Println("Scan error:", err)
				continue
			}
			users = append(users, gin.H{
				"id":    id,
				"name":  name,
				"email": email,
			})
		}

		c.JSON(http.StatusOK, users)
	})
	r.Run(":8889")
}
