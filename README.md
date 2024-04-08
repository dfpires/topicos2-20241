Passo 1: Instalar o GO disponível em https://go.dev/dl/
Passo 2: Abrir o VSCode
Passo 3: Vá em Extensões, digite GO e instale a primeira extensão que aparece (Go Team at Google)
Passo 4: Vá em Ajuda, depois em Mostrar todos os comandos
Passo 5: Digite GO:Install/Updates. Marque todas as opções e clique em OK
Passo 6: Fecha e abre o VSCOde novamente
Passo 7: Abra uma pasta com o nome go/primeiro
Passo 8: Abra um terminal de comando
Passo 9: Digite o comando para criar um projeto GO: go mod init teste
Passo 10: Observe que foi criado um arquivo chamado go.mod
Passo 11: criar um arquivo chamado main.go contendo o conteúdo a seguir:
package main
func main() {
println("Hello, World!")
}
Passo 12: vamos compilar e executar o programa. Execute o comando
go run main.go
Passo 13: Pode-se gerar também o build da aplicação, que é um executável
go build
Passo 14: Perceba que foi gerado o arquivo .exe do programa
Passo 15: Entrar no site https://go-tour-br.appspot.com/list e vamos seguir os exemplos para aprender
Passo 16: Developing a RESTful API with Go and Gin - https://go.dev/doc/tutorial/web-service-gin
