package volume

import (
	"math"
)

// PiramideVolume calcula o volume de uma pirâmide usando a função AreaPiramide da package "area".
func PiramideVolume(areaBase, altura float32) float32 {
	piramideVol := (1.0 / 3.0) * areaBase * altura
	return piramideVol
}

// CuboVolume calcula o volume de um cubo usando a função AreaCubo da package "area".
func CuboVolume(aresta float32) float64 {
	cuboVol := math.Pow(float64(aresta), 3)
	return cuboVol
}

// ParalelepipedoVolume calcula o volume de um paralelepípedo usando a função AreaParalelepipedo da package "area".
func ParalelepipedoVolume(aresta1, aresta2, aresta3 float32) float32 {
	paralelepipedoVol := aresta1 * aresta2 * aresta3
	return paralelepipedoVol
}

// EsferaVolume calcula o volume de uma esfera usando a função AreaEsfera da package "area".
func EsferaVolume(raio float32) float32 {
	esferaVol := (4.0 / 3.0) * math.Pi * float32(math.Pow(float64(raio), 3))
	return esferaVol
}
