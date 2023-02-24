package main

import "fmt"

//A função imprime números múltiplos de 3
//de sequência iniciada em 1 e terminada pelo argumento recebido
func listarMultiplosDe3ApartirDe1(numeroLimite int) {
	for i := 1; i <= numeroLimite; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}		
	}
}

//Dada uma sequência numérica iniciada em 1 e terminada em um argumento recebido
//exibe 'Pin' para múltiplos de '3';
//'Pan' para multiplos de 5;
//'PinPan' para múltiplos de 15;
//valores numéricos para os demais casos
func pinPan(numeroLimite int) {
	for i := 1; i <= numeroLimite; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("PinPan")
		} else if i % 3 == 0 {
			fmt.Println("Pin")
		} else if i % 5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	} 
}

func main() {
	numeroLimitrofe := 100
	listarMultiplosDe3ApartirDe1(numeroLimitrofe)
	pinPan(numeroLimitrofe)
}