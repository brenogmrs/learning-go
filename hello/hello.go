package main

import "fmt"

func main() {

	var nome = "Breno"
	
	vesao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", vesao)

	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - exibir logs")
	fmt.Println("0 - sair do programa")

	var comando int
	// fmt.Scan() -> infere o tipo
	fmt.Scan(&comando)

	// fmt.Scanf() -> não infere o tipo e requer um modificador
	fmt.Scanf("%d", &comando)

	fmt.Println("o endereço da variavel comando é", &comando)
	fmt.Println("o comando escolhido foi,", comando)
}