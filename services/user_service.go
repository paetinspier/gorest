package services

import (
	"psn/gorest/database"
	"psn/gorest/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /users [get]
func GetAllUsers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Uid, &u.Active, &u.Verified, &u.Dob)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
		}

		users = append(users, u)
	}

	return c.Status(200).JSON(users)
}

type CreateUserDTO struct {
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Dob       string `json:"dob" db:"dob"`
	Uid       string `json:"uid" db:"uid"`
	Active    bool   `json:"active" db:"active"`
	Verified  bool   `json:"verified" db:"verified"`
}

func CreateUser(c *fiber.Ctx) error {
	newUser := new(CreateUserDTO)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	// Prepare the SQL statement with placeholders for the values to be inserted
	sqlStatement := `
        INSERT INTO users (first_name, last_name, email, dob, uid, active, verified)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`

	// Execute the SQL statement with the user's data and retrieve the inserted ID
	var id int
	err := database.DB.QueryRow(
		sqlStatement,
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		newUser.Dob,
		newUser.Uid,
		newUser.Active,
		newUser.Verified,
	).Scan(&id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// Return the ID of the inserted user
	return c.Status(200).JSON(fiber.Map{"id": id})
}
