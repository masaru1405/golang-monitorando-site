package main
import "fmt"
import "os"
import "net/http"

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
	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibit Logs")
	fmt.Println("0-Sair do Programa")
}

func leComando() int{
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido: ", comando)
	return comando
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else{
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}
}



