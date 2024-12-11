
import fs from 'node:fs';

enum Operator {
	ADD,
	MULTIPLY,
}

function getCombinations(places: number): Operator[][] {
	if (places == 1) return [[Operator.ADD], [Operator.MULTIPLY]]
	let combinations: Operator[][] = []
	for (const prefix of getCombinations(1)) {
		for (const sub of getCombinations(places - 1)) {
			combinations.push([...prefix, ...sub])
		}
	}
	return combinations
}

fs.readFile('day07.txt', (_, data) => {
	const rows = data.toString().split('\n')
		.filter(r => !!r)

	let total = 0
	for (let row of rows) {
		const [testValue, right] = row.split(': ')
		let numbers = right.split(' ').map(number => parseInt(number))
		for (const combo of getCombinations(numbers.length - 1)) {
			let sum = 0
			for (const key in numbers) {
				const i = parseInt(key)
				const number = numbers[i]
				const operator = combo[i - 1]

				if (i == 0 || operator == Operator.ADD) sum += number
				else if (operator == Operator.MULTIPLY) sum *= number
			}
			if (sum == parseInt(testValue)) {
				total += sum
				break
			}
		}
	}

	console.log(total)
})
