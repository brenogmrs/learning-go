package main

import "fmt"
import "os"

func main() {

	showIntro()
	showMenu()	

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0: 
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default: 
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func showIntro() {
	name := "Breno"
	version := 1.1
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - iniciar monitoramento")
	fmt.Println("2 - exibir logs")
	fmt.Println("0 - sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("o comando escolhido foi,", command)

	return command
}
