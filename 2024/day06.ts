
import fs from 'node:fs';

const directions = {
	'^': [0, -1],
	'>': [1, 0],
	'v': [0, 1],
	'<': [-1, 0],
}

function value(grid: string[][], x: number, y: number): string {
	return grid[y][x];
}

function inBounds(grid: string[][], x: number, y: number): boolean {
	if (x < 0 || x > grid[0].length - 1) return false
	if (y < 0 || y > grid.length - 1) return false
	return true
}

fs.readFile('day06.txt', (_, data) => {
	const grid = data.toString().split('\n')
		.filter(r => !!r)
		.map(r => r.split(''))

	// Find starting position and direction
	let x = -1
	let y = -1
	let direction: string | undefined = undefined
	for (const i in grid) {
		if (x > -1 && y > -1) break
		for (const dir in directions) {
			if (grid[i].includes(dir)) {
				direction = dir
				x = grid[i].indexOf(dir)
				y = parseInt(i)
			}
		}
	}
	grid[y][x] = '.'

	const arrows = Object.keys(directions)
	let steps = new Set<string>()
	while (true) {
		steps.add(`${x} ${y}`)
		const nextX = x + directions[direction][0]
		const nextY = y + directions[direction][1]
		if (!inBounds(grid, nextX, nextY)) break
		if (grid[nextY][nextX] == '#') {
			direction = arrows[arrows.indexOf(direction) + 1]
			if (direction == undefined) direction = '^'
			continue
		}
		x = nextX
		y = nextY
	}
	console.log(steps.size)
})
