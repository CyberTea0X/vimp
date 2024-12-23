package helpers

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Fill заполняет область на изображении цветом newColor с учетом допустимости tolerancy.
func Fill(img *ebiten.Image, newColor color.Color, x, y int, tolerancy int) {
	// Получаем исходный цвет точки
	initialColor := img.At(x, y)

	// Если новый цвет совпадает с исходным, ничего не делаем
	if !isColorDifferent(initialColor, newColor, tolerancy) {
		return
	}

	// Создаем очередь для BFS
	queue := []imagePoint{{x, y}}
	visited := make(map[imagePoint]bool)

	// Основной цикл BFS
	for len(queue) > 0 {
		// Берем текущую точку
		current := queue[0]
		queue = queue[1:]

		// Проверяем границы и избегаем повторных посещений
		if current.x < 0 || current.y < 0 || current.x >= img.Bounds().Dx() || current.y >= img.Bounds().Dy() {
			continue
		}
		if visited[current] {
			continue
		}
		visited[current] = true

		// Проверяем цвет текущей точки
		currentColor := img.At(current.x, current.y)
		if isColorDifferent(initialColor, currentColor, tolerancy) {
			continue
		}

		// Устанавливаем новый цвет
		img.Set(current.x, current.y, newColor)

		// Добавляем соседние точки
		queue = append(queue, imagePoint{current.x + 1, current.y})
		queue = append(queue, imagePoint{current.x - 1, current.y})
		queue = append(queue, imagePoint{current.x, current.y + 1})
		queue = append(queue, imagePoint{current.x, current.y - 1})
	}
}

// imagePoint представляет точку на изображении
type imagePoint struct {
	x, y int
}

// isColorDifferent проверяет, отличается ли цвет от базового цвета в пределах tolerancy
func isColorDifferent(c1, c2 color.Color, tolerancy int) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return absDiff(int(r1>>8), int(r2>>8)) > tolerancy ||
		absDiff(int(g1>>8), int(g2>>8)) > tolerancy ||
		absDiff(int(b1>>8), int(b2>>8)) > tolerancy ||
		absDiff(int(a1>>8), int(a2>>8)) > tolerancy
}

// absDiff вычисляет абсолютную разницу между двумя числами
func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
