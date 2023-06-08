package main
import "fmt"
import "os"
import "net/http"
import "time"

const MONITORAMENTOS = 3
const DELAY = 5

func main(){
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao(){
	nome := "Kaio"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu(){
	fmt.Println("***********************")
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibit Logs")
	fmt.Println("0-Sair do Programa")
	fmt.Println("***********************")
}

func leComando() int{
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido: ", comando)
	return comando
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	//sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br/"}

	sites := leSitesDoArquivo()

	for i:= 0; i < MONITORAMENTOS; i++{
		for j, site := range sites{
			fmt.Println("Posição", j, "- Site:", site)
			testaSite(site)
		}
		time.Sleep(DELAY * time.Second) //espera 5 segundos
	}
}

func testaSite(site string){
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else{
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string{
	arquivo, _ := os.Open("sites.txt")
	fmt.Println(arquivo)

	var sites []string
	return sites
}



