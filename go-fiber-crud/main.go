package main

import (
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{}

// GET /users - list all users
func getUsers(c *fiber.Ctx) error {
    return c.JSON(users)
}

// GET /users/:id - get single user
func getUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).SendString("Invalid ID")
    }

    for _, user := range users {
        if user.ID == id {
            return c.JSON(user)
        }
    }
    return c.Status(404).SendString("User not found")
}

// POST /users - create new user
func createUser(c *fiber.Ctx) error {
    newUser := new(User)
    if err := c.BodyParser(newUser); err != nil {
        return c.Status(400).SendString("Bad request")
    }

    newUser.ID = len(users) + 1
    users = append(users, *newUser)
    return c.JSON(newUser)
}

// PUT /users/:id - update user
func updateUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    update := new(User)
    if err := c.BodyParser(update); err != nil {
        return c.Status(400).SendString("Bad request")
    }

    for i, u := range users {
        if u.ID == id {
            users[i].Name = update.Name
            return c.JSON(users[i])
        }
    }
    return c.Status(404).SendString("User not found")
}

// DELETE /users/:id - delete user
func deleteUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))

    for i, u := range users {
        if u.ID == id {
            users = append(users[:i], users[i+1:]...)
            return c.SendString("Deleted successfully")
        }
    }
    return c.Status(404).SendString("User not found")
}

func main() {
    app := fiber.New()

    // Routes
    app.Get("/users", getUsers)
    app.Get("/users/:id", getUser)
    app.Post("/users", createUser)
    app.Put("/users/:id", updateUser)
    app.Delete("/users/:id", deleteUser)

    app.Listen(":3000")
}
