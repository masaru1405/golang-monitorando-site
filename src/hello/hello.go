package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	//"io/ioutil"
	"bufio"
	"io"
	"strconv"
)

const MONITORAMENTOS = 3
const DELAY = 5

func main() {
	exibeIntroducao()

	leSitesDoArquivo()

	//for infinito
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Kaio"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("***********************")
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibit Logs")
	fmt.Println("0-Sair do Programa")
	fmt.Println("***********************")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido: ", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br/"}

	sites := leSitesDoArquivo()

	for i := 0; i < MONITORAMENTOS; i++ {
		for j, site := range sites {
			fmt.Println("Posição", j, "- Site:", site)
			testaSite(site)
		}
		fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
		time.Sleep(DELAY * time.Second) //espera 5 segundos
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	//fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	var sites []string
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		//fmt.Println(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	//fmt.Println(sites)

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	//ioutil.ReadFile já fecha o arquivo automaticamente
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
