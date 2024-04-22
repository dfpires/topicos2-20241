package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Struct para representar um usuário
type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	// Substitua com suas próprias credenciais
	connStr := "user=postgres dbname=golang password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Testar conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão com o banco de dados estabelecida!")

	// Criar tabela se não existir
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	// Exemplo de CRUD
	// Criar usuário
	user := User{
		Username: "joao",
		Email:    "joao.valeriano.silva@gmail.com",
	}
	err = createUser(db, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário criado com ID: %d\n", user.ID)

	/*
		// Ler todos os usuários
		users, err := getAllUsers(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Todos os usuários:")
		for _, u := range users {
			fmt.Printf("ID: %d, Username: %s, Email: %s\n", u.ID, u.Username, u.Email)
		}

		// Atualizar usuário
		user.Username = "john_smith"
		err = updateUser(db, &user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Usuário atualizado com sucesso!")

		// Ler usuário por ID
		foundUser, err := getUserByID(db, user.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Usuário encontrado: ID: %d, Username: %s, Email: %s\n", foundUser.ID, foundUser.Username, foundUser.Email)

		// Deletar usuário
		err = deleteUser(db, user.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Usuário deletado com sucesso!") */
}

// Função para criar um usuário
func createUser(db *sql.DB, user *User) error {
	err := db.QueryRow("INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id", user.Username, user.Email).Scan(&user.ID)
	return err
}

// Função para obter todos os usuários
func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Função para atualizar um usuário
func updateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", user.Username, user.Email, user.ID)
	return err
}

// Função para obter um usuário por ID
func getUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	return user, err
}

// Função para deletar um usuário
func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
