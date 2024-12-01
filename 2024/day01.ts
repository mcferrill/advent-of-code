
import fs from 'node:fs';

function sortIntsAscending(a, b: int): int {
	if (a < b) return -1
	else if (a > b) return 1
	else return 0
}

fs.readFile('day01.txt', (err, data) => {
	let list1: number[] = []
	let list2: number[] = []
	data.toString().split('\n').map(line => {
		const parts = line.split('   ')
		if (parts.length < 2) return
		list1.push(parseInt(parts[0]))
		list2.push(parseInt(parts[1]))
	})
	list1.sort(sortIntsAscending)
	list2.sort(sortIntsAscending)
	let total = 0
	for (const i in list1) {
		total += Math.max(list1[i], list2[i]) - Math.min(list1[i], list2[i])
	}
	console.log(total)
})
