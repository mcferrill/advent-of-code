
import fs from 'node:fs';

function sortIntsAscending(a: number, b: number): number {
	if (a < b) return -1
	else if (a > b) return 1
	else return 0
}

function part1(list1: number[], list2: number[]): number {
	list1.sort(sortIntsAscending)
	list2.sort(sortIntsAscending)
	let total = 0
	for (const i in list1) {
		total += Math.max(list1[i], list2[i]) - Math.min(list1[i], list2[i])
	}
	return total;
}

function part2(list1: number[], list2: number[]): number {
	let total = 0;
	for (const i in list1) {
		const count = list2.filter(n => n == list1[i]).length
		list1[i] = list1[i] * count
		total += list1[i]
	}
	return total
}

fs.readFile('day01.txt', (_, data) => {
	let list1: number[] = []
	let list2: number[] = []
	data.toString().split('\n').map(line => {
		const parts = line.split('   ')
		if (parts.length < 2) return
		list1.push(parseInt(parts[0]))
		list2.push(parseInt(parts[1]))
	})
	console.log(part1(list1, list2))
	console.log(part2(list1, list2))
})
