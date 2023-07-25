package volume

import (
	"testing"
)

func TestPiramideVolume(t *testing.T) {
	areaBase := float32(10)
	altura := float32(5)
	expectedVolume := float32(16.666668)
	volume := PiramideVolume(areaBase, altura)

	if volume != expectedVolume {
		t.Errorf("Erro: o volume da pirâmide é %f, mas o valor esperado é %f", volume, expectedVolume)
	}

	// Teste com altura zero
	alturaZero := float32(0)
	expectedVolumeZero := float32(0)
	volumeZero := PiramideVolume(areaBase, alturaZero)

	if volumeZero != expectedVolumeZero {
		t.Errorf("Erro: o volume da pirâmide com altura zero é %f, mas o valor esperado é %f", volumeZero, expectedVolumeZero)
	}
}

func TestCuboVolume(t *testing.T) {
	aresta := float32(3)
	expectedVolume := float64(27)
	volume := CuboVolume(aresta)

	if volume != expectedVolume {
		t.Errorf("Erro: o volume do cubo é %f, mas o valor esperado é %f", volume, expectedVolume)
	}

	// Teste com aresta zero
	arestaZero := float32(0)
	expectedVolumeZero := float64(0)
	volumeZero := CuboVolume(arestaZero)

	if volumeZero != expectedVolumeZero {
		t.Errorf("Erro: o volume do cubo com aresta zero é %f, mas o valor esperado é %f", volumeZero, expectedVolumeZero)
	}
}

func TestParalelepipedoVolume(t *testing.T) {
	aresta1 := float32(2)
	aresta2 := float32(3)
	aresta3 := float32(4)
	expectedVolume := float32(24)
	volume := ParalelepipedoVolume(aresta1, aresta2, aresta3)

	if volume != expectedVolume {
		t.Errorf("Erro: o volume do paralelepípedo é %f, mas o valor esperado é %f", volume, expectedVolume)
	}

	// Teste com uma aresta zero
	arestaZero := float32(0)
	expectedVolumeZero := float32(0)
	volumeZero := ParalelepipedoVolume(aresta1, aresta2, arestaZero)

	if volumeZero != expectedVolumeZero {
		t.Errorf("Erro: o volume do paralelepípedo com uma aresta zero é %f, mas o valor esperado é %f", volumeZero, expectedVolumeZero)
	}
}

func TestEsferaVolume(t *testing.T) {
	raio := float32(2)
	expectedVolume := float32(33.510322)
	volume := EsferaVolume(raio)

	if volume != expectedVolume {
		t.Errorf("Erro: o volume da esfera é %f, mas o valor esperado é %f", volume, expectedVolume)
	}

	// Teste com raio zero
	raioZero := float32(0)
	expectedVolumeZero := float32(0)
	volumeZero := EsferaVolume(raioZero)

	if volumeZero != expectedVolumeZero {
		t.Errorf("Erro: o volume da esfera com raio zero é %f, mas o valor esperado é %f", volumeZero, expectedVolumeZero)
	}
}
