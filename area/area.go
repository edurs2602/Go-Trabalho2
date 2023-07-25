package area

import (
	"math"
)

// AreaTriangulo calcula a área de um triângulo com base e altura fornecidas pelo usuário.
func AreaTriangulo(base, altura float32) float32 {
	areaTri := (base * altura) / 2
	return areaTri
}

// AreaRetangulo calcula a área de um retângulo com base e altura fornecidas pelo usuário.
func AreaRetangulo(base, altura float32) float32 {
	areaRet := base * altura
	return areaRet
}

// AreaQuadrado calcula a área de um quadrado com lado fornecido pelo usuário.
func AreaQuadrado(lado float32) float32 {
	areaQua := lado * lado
	return areaQua
}

// AreaCirculo calcula a área de um círculo com raio fornecido pelo usuário.
func AreaCirculo(raio float32) float32 {
	areaCirc := math.Pi * raio * raio
	return areaCirc
}

// AreaPiramide calcula a área de uma pirâmide com área da base e área lateral fornecidas pelo usuário.
func AreaPiramide(areaBase, areaLateral float32) float32 {
	areaPir := areaBase + areaLateral
	return areaPir
}

// AreaCubo calcula a área de um cubo com aresta fornecida pelo usuário.
func AreaCubo(aresta float32) float32 {
	areaCub := 6 * (aresta * aresta)
	return areaCub
}

// AreaParalelepipedo calcula a área de um paralelepípedo com as três arestas fornecidas pelo usuário.
func AreaParalelepipedo(aresta1, aresta2, aresta3 float32) float32 {
	areaPara := (2 * aresta1 * aresta2) + (2 * aresta1 * aresta3) + (2 * aresta2 * aresta3)
	return areaPara
}

// AreaEsfera calcula a área de uma esfera com raio fornecido pelo usuário.
func AreaEsfera(raio float32) float32 {
	areaEsf := 4 * math.Pi * (raio * raio)
	return areaEsf
}
