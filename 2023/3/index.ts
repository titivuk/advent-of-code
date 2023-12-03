import { readFile } from "fs/promises";

let directions: Array<[number, number]> = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
  [0, 1],
  [1, 1],
  [1, 0],
  [1, -1],
  [0, -1],
];

function isDigit(ch: string): boolean {
  return 48 <= ch.charCodeAt(0) && ch.charCodeAt(0) <= 57;
}

function getAdjacencyNumberPart1(
  matrix: string[][],
  i: number,
  j: number
): number {
  let adjacencyNumber = 0;

  let neighborI = 0,
    neighborJ = 0;
  for (let [iOffset, jOffset] of directions) {
    neighborI = i + iOffset;
    neighborJ = j + jOffset;

    if (
      neighborI >= 0 &&
      neighborI < matrix.length &&
      neighborJ >= 0 &&
      neighborJ < matrix[i].length &&
      matrix[neighborI][neighborJ] !== "." &&
      !isDigit(matrix[neighborI][neighborJ])
    ) {
      adjacencyNumber += 1;
    }
  }

  return adjacencyNumber;
}

function getAdjacencyNumberPart2(
  matrix: string[][],
  i: number,
  j: number,
  gears: Set<string>
): number {
  let adjacencyNumber = 0;

  let neighborI = 0,
    neighborJ = 0;
  for (let [iOffset, jOffset] of directions) {
    neighborI = i + iOffset;
    neighborJ = j + jOffset;
    if (
      neighborI >= 0 &&
      neighborI < matrix.length &&
      neighborJ >= 0 &&
      neighborJ < matrix[i].length &&
      // diff from part 1
      matrix[neighborI][neighborJ] === "*"
    ) {
      // remember gear that adjacent to given num
      gears.add(`${neighborI}_${neighborJ}`);

      adjacencyNumber += 1;
    }
  }

  return adjacencyNumber;
}

function part1(matrix: string[][]): number {
  let sum = 0;

  let num = 0,
    numAdjacency = 0;
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      // we found digit
      if (isDigit(matrix[i][j])) {
        // update num value
        num = num * 10 + Number.parseInt(matrix[i][j]);

        // update adjacency for a num
        numAdjacency += getAdjacencyNumberPart1(matrix, i, j);
      }
      // no digit found
      else {
        // if we found a number before and it is adjacent to valid symbols
        // add the number to the sum
        if (numAdjacency > 0) {
          sum += num;
        }

        // reset num information
        num = 0;
        numAdjacency = 0;
      }
    }
  }

  console.log("part1 sum:", sum);

  return sum;
}

/**
 * @description probably better approach is to iterate over gears("*")
 * and DO some kind of DFS to find numbers
 * but, again, I was lazy so made this piece of ... code with additional memory
 */
function part2(matrix: string[][]): number {
  let gearAdjacencyList = new Map<string, number[]>();

  let num = 0,
    numAdjacency = 0,
    // keep list of gears that are adjacent to num
    numGears = new Set<string>();
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      // we found digit
      if (isDigit(matrix[i][j])) {
        // update num value
        num = num * 10 + Number.parseInt(matrix[i][j]);

        // update adjacency for a num
        numAdjacency += getAdjacencyNumberPart2(matrix, i, j, numGears);
      }
      // no digit found
      else {
        // if we found a number before and it is adjacent to valid symbols
        if (numAdjacency > 0) {
          // for every gear that is adjacent to num
          // add the num to adjacency list for every gear
          for (const adjacentGear of numGears) {
            let adjacentsNums = gearAdjacencyList.get(adjacentGear) ?? [];

            adjacentsNums.push(num);
            gearAdjacencyList.set(adjacentGear, adjacentsNums);
          }
        }

        // reset num information
        num = 0;
        numAdjacency = 0;
        numGears.clear();
      }
    }
  }

  let sum = 0;
  for (const [gear, adjacentNums] of gearAdjacencyList) {
    // count only gears with 2 adjacent numbers
    if (adjacentNums.length === 2) {
      sum += adjacentNums[0] * adjacentNums[adjacentNums.length - 1];
    }
  }

  return sum;
}

async function main() {
  const fileContent = await readFile("./2023/3/input.txt", {
    encoding: "utf-8",
  });
  let fileRows = fileContent.split("\n");

  let matrix: string[][] = [];
  for (let i = 0; i < fileRows.length; i++) {
    matrix[i] = fileRows[i].split("");
  }

  let sumPart1 = part1(matrix);
  console.log("part 1 sum:", sumPart1);

  let sumPart2 = part2(matrix);
  console.log("part 2 sum:", sumPart2);
}

main();
