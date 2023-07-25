package perimetro

import (
	"math"
	"testing"
)

func TestPerimetroTriangulo(t *testing.T) {
	lado := float32(5)
	expected := float32(15)

	result := PerimetroTriangulo(lado)

	if result != expected {
		t.Errorf("Perímetro do triângulo incorreto. Esperado: %f, Obtido: %f", expected, result)
	}
}

func TestPerimetroRetangulo(t *testing.T) {
	base := float32(4)
	altura := float32(6)
	expected := float32(20)

	result := PerimetroRetangulo(base, altura)

	if result != expected {
		t.Errorf("Perímetro do retângulo incorreto. Esperado: %f, Obtido: %f", expected, result)
	}
}

func TestPerimetroQuadrado(t *testing.T) {
	lado := float32(3)
	expected := float32(12)

	result := PerimetroQuadrado(lado)

	if result != expected {
		t.Errorf("Perímetro do quadrado incorreto. Esperado: %f, Obtido: %f", expected, result)
	}
}

func TestPerimetroCirculo(t *testing.T) {
	raio := float32(2)
	expected := float32(2 * math.Pi * 2)

	result := PerimetroCirculo(raio)

	if result != expected {
		t.Errorf("Perímetro do círculo incorreto. Esperado: %f, Obtido: %f", expected, result)
	}
}
