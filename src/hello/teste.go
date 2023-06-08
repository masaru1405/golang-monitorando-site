package main
import "fmt"
import "os"

func main(){
	leSitesDoArquivo()
}

func leSitesDoArquivo() []string{
	arquivo, _ := os.Open("sites.txt")
	fmt.Println(arquivo)

	var sites []string
	return sites
}