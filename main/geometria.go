package main

import (
	"area"
	"fmt"
	"os"
	"perimetro"
	"volume"
)

func printer() {
	fmt.Println("Calculadora de Geometria Plana e Espacial")
	fmt.Println("(1) Triângulo equilátero")
	fmt.Println("(2) Retângulo")
	fmt.Println("(3) Quadrado")
	fmt.Println("(4) Círculo")
	fmt.Println("(5) Pirâmide com base quadrangular")
	fmt.Println("(6) Cubo")
	fmt.Println("(7) Paralelepípedo")
	fmt.Println("(8) Esfera")
	fmt.Println("(0) Sair")
	fmt.Printf("Digite sua opção: ")
}

func lerEntradaFloat(prompt string) (float32, error) {
	var input float32
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	return input, err
}

func main() {
	var escolha int

	for {
		printer()
		_, err := fmt.Scanln(&escolha)
		if err != nil {
			fmt.Println("Erro ao ler a entrada do usuário.")
			os.Exit(1)
		}

		switch escolha {
		case 1:
			base, err := lerEntradaFloat("Digite a base do triângulo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			altura, err := lerEntradaFloat("Digite a altura do triângulo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			lado, err := lerEntradaFloat("Digite o lado do triângulo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaTri := area.AreaTriangulo(base, altura)
			perimetroTri := perimetro.PerimetroTriangulo(lado)
			fmt.Println("Área do Triângulo: ", areaTri)
			fmt.Println("Perímetro do Triângulo: ", perimetroTri)

		case 2:
			base, err := lerEntradaFloat("Digite a base do retângulo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			altura, err := lerEntradaFloat("Digite a altura do retângulo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaRet := area.AreaRetangulo(base, altura)
			perimetroRet := perimetro.PerimetroRetangulo(base, altura)
			fmt.Println("Área do Retângulo: ", areaRet)
			fmt.Println("Perímetro do Retângulo: ", perimetroRet)

		case 3:
			lado, err := lerEntradaFloat("Digite o lado do quadrado: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaQua := area.AreaQuadrado(lado)
			perimetroQua := perimetro.PerimetroQuadrado(lado)
			fmt.Println("Área do Quadrado: ", areaQua)
			fmt.Println("Perímetro do Quadrado: ", perimetroQua)

		case 4:
			raio, err := lerEntradaFloat("Digite o raio da circunferência: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaCirc := area.AreaCirculo(raio)
			circunferencia := perimetro.PerimetroCirculo(raio)
			fmt.Println("Área da Circunferência: ", areaCirc)
			fmt.Println("Circunferência: ", circunferencia)

		case 5:
			base, err := lerEntradaFloat("Area da base da pirâmide: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			lateral, err := lerEntradaFloat("Digite a area lateral da pirâmide: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			altura, err := lerEntradaFloat("Digite a altura da pirâmide: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaPir := area.AreaPiramide(base, lateral)
			volumePir := volume.PiramideVolume(base, altura)
			fmt.Println("Área da Pirâmide: ", areaPir)
			fmt.Println("Volume da Pirâmide: ", volumePir)

		case 6:
			aresta, err := lerEntradaFloat("Digite a aresta do cubo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaCubo := area.AreaCubo(aresta)
			volumeCubo := volume.CuboVolume(aresta)
			fmt.Println("Área do Cubo: ", areaCubo)
			fmt.Println("Volume do Cubo: ", volumeCubo)

		case 7:
			aresta1, err := lerEntradaFloat("Digite a primeira aresta do paralelepípedo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			aresta2, err := lerEntradaFloat("Digite a segunda aresta do paralelepípedo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			aresta3, err := lerEntradaFloat("Digite a terceira aresta do paralelepípedo: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaPar := area.AreaParalelepipedo(aresta1, aresta2, aresta3)
			volumePar := volume.ParalelepipedoVolume(aresta1, aresta2, aresta3)
			fmt.Println("Área do Paralelepípedo: ", areaPar)
			fmt.Println("Volume do Paralelepípedo: ", volumePar)

		case 8:
			raio, err := lerEntradaFloat("Digite o raio da esfera: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			areaEsfera := area.AreaEsfera(raio)
			volumeEsfera := volume.EsferaVolume(raio)
			fmt.Println("Área da Esfera: ", areaEsfera)
			fmt.Println("Volume da Esfera: ", volumeEsfera)

		case 0:
			fmt.Println("Encerrando o Programa...")
			os.Exit(0)

		default:
			fmt.Println("Número inválido!")
		}

		fmt.Println()
	}
}
