package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"learn-go/models"
	"log"
	"strconv"
)

func GetHandler(ctx *fiber.Ctx, db *sql.DB) error {
	var accounts []models.Account
	rows, err := db.Query("SELECT * FROM accounts")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		ctx.JSON("Server error")
	}

	var res models.Account
	for rows.Next() {
		err := rows.Scan(&res.ID, &res.Username, &res.CreateAt)
		if err != nil {
			return err
		}
		accounts = append(accounts, res)
	}
	return ctx.Render("index", fiber.Map{
		"accounts": accounts,
	})
}

func PostHandler(ctx *fiber.Ctx, db *sql.DB) error {
	newAccount := models.Account{}
	if err := ctx.BodyParser(&newAccount); err != nil {
		log.Printf("Server error, %v", err)
		return ctx.SendString(err.Error())
	}
	if newAccount.Username != "" {
		_, err := db.Exec("INSERT INTO accounts(username) VALUES ($1)", newAccount.Username)
		if err != nil {
			log.Fatalf("Server error, %v", err)
		}
	} else {
		log.Printf("Invalid input %v", newAccount)
	}
	return ctx.Redirect("/accounts")
}

func PutHandler(ctx *fiber.Ctx, db *sql.DB) error {
	var data map[string]string
	if parseErr := ctx.BodyParser(&data); parseErr != nil {
		log.Fatalf("Server error, %v", parseErr)
	}
	_, err := db.Exec("UPDATE accounts SET username=$1 WHERE id=$2", data["username"], data["id"])
	log.Printf("%v", err)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}
	return ctx.Redirect("/accounts")
}

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	a := c.Query("accountId")
	accountId, _ := strconv.Atoi(a)
	_, err := db.Exec("DELETE from accounts WHERE ID=$1", accountId)
	if err != nil {
		return err
	}
	return c.SendString("deleted")
}
