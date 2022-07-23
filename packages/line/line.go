package main

import "math"

// Primeira letra maiúscula: pública
// Do contrário: privado
// Tudo isso dentro do pacote
// Não existe nada que defina visibilidade dentro do Go
// É apenas convensão
// Ainda, pode-se usar o padrão javascript para privado: _private

// Representa uma coordenada no plano cartesiano
type Point struct {
	x, y float64
}

func cathetus(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

// Responsável por calcular a distância linear entre dois pontos
func Distance(a, b Point) (d float64) {
	cx, cy := cathetus(a, b)

	d = math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))

	return
}
