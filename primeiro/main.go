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

type Produto struct {
	ID    int
	nome  string
	qtde  int
	preco float64
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

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS produtos (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(50) NOT NULL,
			qtde int NOT NULL,
			preco float NOT NULL
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

	// Criar produto
	produto := Produto{
		nome:  "sonho de valsa",
		qtde:  10,
		preco: 5,
	}
	err = createProduto(db, &produto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Produto criado com ID: %d\n", produto.ID)

	// Ler todos os usuários
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todos os usuários:")
	for _, u := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", u.ID, u.Username, u.Email)
	}

	// Ler todos os usuários
	produtos, err := getAllProdutos(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todos os produtos:")
	for _, u := range produtos {
		fmt.Printf("ID: %d, Nome: %s, Qtde: %d, Preço: %.2f\n", u.ID, u.nome, u.qtde, u.preco)
	}

	// Ler usuário por ID
	foundUser, err := getUserByID(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuário encontrado: ID: %d, Username: %s, Email: %s\n", foundUser.ID, foundUser.Username, foundUser.Email)

	foundProduto, err := getProdutoByID(db, produto.ID)
	if err != nil {
		fmt.Printf("entrou no erro")
		log.Fatal(err)
	}
	fmt.Printf("Produto encontrado: ID: %d, Nome: %s, Qtde: %d, Preço %.2f",
		foundProduto.ID, foundProduto.nome, foundProduto.qtde, foundProduto.preco)

	// Atualizar usuário
	user.Username = "john_smith"
	err = updateUser(db, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuário atualizado com sucesso!")

	produto.ID = 5
	produto.nome = "bolacha"
	produto.qtde = 10
	produto.preco = 15
	err = updateProduto(db, &produto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produto atualizado com sucesso!")

	// Deletar usuário
	err = deleteUser(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuário deletado com sucesso!")

	err = deleteProduto(db, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produto deletado com sucesso!")
}

// Função para criar um usuário
func createUser(db *sql.DB, user *User) error {
	err := db.QueryRow("INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id", user.Username, user.Email).Scan(&user.ID)
	return err
}

// Função para criar um produto
func createProduto(db *sql.DB, produto *Produto) error {
	err := db.QueryRow("INSERT INTO produtos (nome, qtde, preco) VALUES ($1, $2, $3) RETURNING id", produto.nome, produto.qtde, produto.preco).Scan(&produto.ID)
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

// Função para obter todos os usuários
func getAllProdutos(db *sql.DB) ([]Produto, error) {
	rows, err := db.Query("SELECT id, nome, qtde, preco FROM produtos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var produto Produto
		err := rows.Scan(&produto.ID, &produto.nome, &produto.qtde, &produto.preco)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}
	return produtos, nil
}

// Função para atualizar um usuário
func updateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", user.Username, user.Email, user.ID)
	return err
}

func updateProduto(db *sql.DB, produto *Produto) error {
	_, err := db.Exec("UPDATE produtos SET nome=$1, qtde=$2, preco=$3 WHERE id=$4", produto.nome, produto.qtde, produto.preco, produto.ID)
	return err
}

// Função para obter um usuário por ID
func getUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	return user, err
}

func getProdutoByID(db *sql.DB, id int) (Produto, error) {
	var produto Produto
	err := db.QueryRow("SELECT id, nome, qtde, preco FROM produtos WHERE id=$1", id).Scan(&produto.ID, &produto.nome, &produto.qtde, &produto.preco)
	return produto, err
}

// Função para deletar um usuário
func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func deleteProduto(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM produtos WHERE id=$1", id)
	return err
}
