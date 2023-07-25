package perimetro

import (
	"math"
)

// PerimetroTriangulo calcula o perímetro de um triângulo com um lado fornecido pelo usuário.
func PerimetroTriangulo(lado1 float32) float32 {
	perimetroTri := lado1 * 3
	return perimetroTri
}

// PerimetroRetangulo calcula o perímetro de um retângulo usando a função AreaRetangulo da package "area".
func PerimetroRetangulo(base, altura float32) float32 {
	perimetroRet := 2 * (base + altura)
	return perimetroRet
}

// PerimetroQuadrado calcula o perímetro de um quadrado usando a função AreaQuadrado da package "area".
func PerimetroQuadrado(lado float32) float32 {
	perimetroQua := 4 * lado
	return perimetroQua
}

// PerimetroCirculo calcula o perímetro de um círculo usando a função AreaCirculo da package "area".
func PerimetroCirculo(raio float32) float32 {
	perimetroCirc := 2 * math.Pi * raio
	return perimetroCirc
}
