
import fs from 'node:fs';
import { sortIntsAscending } from './util';

function isReportSafe(report: number[]): boolean {
	const sorted = report.toSorted(sortIntsAscending)
	const ascending = report.toString() == sorted.toString()
	const descending = report.toString() == sorted.toReversed().toString()
	if (!(ascending || descending)) return false
	let isSafe = true
	for (let idx in report) {
		const i = parseInt(idx)
		if (i == 0) continue
		const diff = Math.max(report[i], report[i - 1]) - Math.min(report[i], report[i - 1])
		if ((diff < 1) || (diff > 3)) {
			isSafe = false
			break
		}
	}
	return isSafe
}

function part1(reports: number[][]): number {
	let safe = 0;
	for (const report of reports) {
		if (isReportSafe(report)) safe++
	}
	return safe
}

function part2(reports: number[][]): number {
	let safe = 0;
	for (const report of reports) {
		if (isReportSafe(report)) {
			safe++
			continue
		}

		for (const i in report) {
			let tmp = [...report]
			tmp.splice(parseInt(i), 1)
			if (isReportSafe(tmp)) {
				safe++
				break
			}
		}
	}
	return safe
}

fs.readFile('day02.txt', (_, data) => {
	let reports: number[][] = []
	data.toString().split('\n').map(line => {
		const parts = line.split(' ')
		if (parts.length == 1) return
		let report: number[] = [];
		for (let char of parts) {
			report.push(parseInt(char))
		}
		reports.push(report)
	})
	console.log(part1(reports))
	console.log(part2(reports))
})
