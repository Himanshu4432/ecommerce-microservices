// internal/models/user.go\npackage models\n\n// Define your models here\n
// services/user-service/internal/models/user.go
package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
