package middlewares

import (
	"github.com/brettmerri/GoBlog/backend/db"
	"github.com/gin-gonic/gin"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
