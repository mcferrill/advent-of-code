
import fs from 'node:fs';

function part1(instructions: RegExpMatchArray[]): number {
	let sum = 0
	for (const inst of instructions) {
		sum += (parseInt(inst[1]) * parseInt(inst[2]))
	}
	return sum
}

function part2(instructions: RegExpMatchArray[]): number {
	let sum = 0
	let enabled = true
	for (const inst of instructions) {
		if (inst[0] == 'do()') enabled = true
		else if (inst[0] == "don't()") enabled = false
		else if (enabled) sum += (parseInt(inst[2]) * parseInt(inst[3]))
	}
	return sum
}

fs.readFile('day03.txt', (_, data) => {
	const muls = [...data.toString().matchAll(/mul\((\d+)\,(\d+)\)/g)]
	const mulsAndConditions = [...data.toString().matchAll(/(mul\((\d+)\,(\d+)\))|don\'t\(\)|do\(\)/g)]

	console.log(part1(muls))
	console.log(part2(mulsAndConditions))
})
