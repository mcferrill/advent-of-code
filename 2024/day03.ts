
import fs from 'node:fs';

function part1(instructions: string[]): number {
	let sum = 0
	for (const inst of instructions) {
		const left = parseInt(inst[1])
		const right = parseInt(inst[2])
		sum += (left * right)
	}
	return sum
}

fs.readFile('day03.txt', (_, data) => {
	const instructions = [...data.toString().matchAll(/mul\((\d+)\,(\d+)\)/g)]

	console.log(part1(instructions))
})
