package area

import (
	"math"
	"testing"
)

func TestAreaTriangulo(t *testing.T) {
	base := float32(5)
	altura := float32(6)
	expected := float32(15)

	result := AreaTriangulo(base, altura)

	if result != expected {
		t.Errorf("Área do triângulo incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaRetangulo(t *testing.T) {
	base := float32(4)
	altura := float32(6)
	expected := float32(24)

	result := AreaRetangulo(base, altura)

	if result != expected {
		t.Errorf("Área do retângulo incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaQuadrado(t *testing.T) {
	lado := float32(3)
	expected := float32(9)

	result := AreaQuadrado(lado)

	if result != expected {
		t.Errorf("Área do quadrado incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaCirculo(t *testing.T) {
	raio := float32(2)
	expected := float32(math.Pi * 4)

	result := AreaCirculo(raio)

	if result != expected {
		t.Errorf("Área do círculo incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaPiramide(t *testing.T) {
	areaBase := float32(10)
	areaLateral := float32(30)
	expected := float32(40)

	result := AreaPiramide(areaBase, areaLateral)

	if result != expected {
		t.Errorf("Área da pirâmide incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaCubo(t *testing.T) {
	aresta := float32(3)
	expected := float32(54)

	result := AreaCubo(aresta)

	if result != expected {
		t.Errorf("Área do cubo incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaParalelepipedo(t *testing.T) {
	aresta1 := float32(2)
	aresta2 := float32(3)
	aresta3 := float32(4)
	expected := float32(52)

	result := AreaParalelepipedo(aresta1, aresta2, aresta3)

	if result != expected {
		t.Errorf("Área do paralelepípedo incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}

func TestAreaEsfera(t *testing.T) {
	raio := float32(2)
	expected := float32(4 * math.Pi * 4)

	result := AreaEsfera(raio)

	if result != expected {
		t.Errorf("Área da esfera incorreta. Esperada: %f, Obtida: %f", expected, result)
	}
}
