package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringTimes = 3
const monitoringDelay = 5

func main() {

	showIntro()
	getSitesFromFile()
	for {
		showMenu()	
	
		command := readCommand()
	
		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0: 
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default: 
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
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
	fmt.Println("")
	
	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := getSitesFromFile()

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			verifySite(site)
		}

		time.Sleep(monitoringDelay * time.Second)
		fmt.Println("")
	}
	
	fmt.Println("")
	
}

func verifySite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("site:", site, "carregado com sucesso!")
		registerLog(site, true)
	}else {
		fmt.Println("site:", site, "está com problemas. Status Code:", response.StatusCode)
		registerLog(site, true)
	}
}

func getSitesFromFile() []string {

	var sites []string

	file, err := os.Open("sitesToMonitor.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
		
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {

	permissions := os.O_CREATE | os.O_RDWR | os.O_APPEND

	file, err := os.OpenFile("siteStatusLog.txt", permissions, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	logString := "[" + time.Now().Format("02/01/2006 - 15:04:05") + "] - " + site + " - online: " + strconv.FormatBool(status) + "\n"

	file.WriteString(logString)

	file.Close()
}

func showLogs(){
	fmt.Println("Exibindo logs...")

	file, err := ioutil.ReadFile("siteStatusLog.txt")

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(file))
	
}
