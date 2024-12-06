
import fs from 'node:fs';


fs.readFile('day05.txt', (_, data) => {
	let [chunk1, chunk2] = data.toString().split('\n\n')
		.map(chunk => chunk.split('\n'))

	// Convert rules to regex
	const rules = chunk1.map(rule => {
		const [before, after] = rule.split('|')
		return new RegExp(`${after}.*(?=${before})`, 'g')
	})

	const validUpdates = chunk2.filter(update => {
		for (const rule of rules) {
			if (rule.exec(update)) {
				return false
			}
		}
		return true
	})

	let sum = 0
	validUpdates.map(update => {
		const parts = update.split(',')
		const value = parseInt(parts[(parts.length - 1) / 2])
		if (!isNaN(value)) sum += value
	})

	console.log(sum)
})
