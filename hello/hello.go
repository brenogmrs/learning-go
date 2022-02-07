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
	fmt.Scan(&comando)
	fmt.Println("o comando escolhido foi,", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 3: 
		fmt.Println("Saindo do programa...")
	default: 
		fmt.Println("Não conheço este comando")
	}

}