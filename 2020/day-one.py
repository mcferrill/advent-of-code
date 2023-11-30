#!/usr/bin/env python

INPUT = open('day1.txt').read()


def find_matching_rows(rows):
    for row1 in rows:
        for row2 in rows:
            if row1 == row2:
                continue
            for row3 in rows:
                if row3 in (row1, row2):
                    continue
                if row1 + row2 + row3 == 2020:
                    return row1, row2, row3
    raise ValueError('Could not find values summing to 2020 in', rows)

def main():
    rows = [int(n) for n in INPUT.splitlines() if n]
    row1, row2, row3 = find_matching_rows(rows)
    print(row1 * row2 * row3)


if __name__ == "__main__":
    main()
