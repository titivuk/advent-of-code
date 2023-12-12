import { readFile } from "fs/promises";

type Direction = [number, number];
type Point = [number, number];

let northDirection: Direction = [-1, 0];
let eastDirection: Direction = [0, 1];
let southDirection: Direction = [1, 0];
let westDirection: Direction = [0, -1];

function isValidDirection(i: number, j: number, m: number, n: number) {
  return 0 <= i && i < m && 0 <= j && j < n;
}

function isConnectedPipe(
  i: number,
  j: number,
  matrix: string[][],
  sourceDirection: Direction
) {
  return (
    isValidDirection(i, j, matrix.length, matrix[i].length) &&
    matrix[i][j] !== "." &&
    pipeToDirectionsMapping[matrix[i][j]].some(
      (dir) => sourceDirection[0] === -dir[0] && sourceDirection[1] === -dir[1]
    )
  );
}

let pipeToDirectionsMapping: Record<string, Direction[]> = {
  "|": [northDirection, southDirection],
  "-": [westDirection, eastDirection],
  L: [northDirection, eastDirection],
  J: [northDirection, westDirection],
  "7": [southDirection, westDirection],
  F: [southDirection, eastDirection],
};

function findStartPoint(matrix: string[][]): Point {
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (matrix[i][j] === "S") {
        return [i, j];
      }
    }
  }

  throw new Error("Stat position not found");
}

function arePointsEqual(p1: Point, p2: Point): boolean {
  return p1[0] === p2[0] && p1[1] === p2[1];
}

function getNextPoint(
  prevPoint: Point,
  currPoint: Point,
  matrix: string[][]
): Point {
  let pipeSymbol = matrix[currPoint[0]][currPoint[1]];

  let [possibleDirection1, possibleDirection2] =
    pipeToDirectionsMapping[pipeSymbol];

  // there are only 2 possibles directions
  // one of the are prev point where we don't need to go
  // pick the one of the 2 possible directions that are not prevPoint
  let nextDirection = arePointsEqual(prevPoint, [
    currPoint[0] + possibleDirection1[0],
    currPoint[1] + possibleDirection1[1],
  ])
    ? possibleDirection2
    : possibleDirection1;

  return [currPoint[0] + nextDirection[0], currPoint[1] + nextDirection[1]];
}

async function part1() {
  const fileContent = await readFile("./2023/10/input.txt", "utf-8");

  let matrix = fileContent.split("\n").map((row) => row.split(""));
  let start = findStartPoint(matrix);

  let answer = 0;

  for (let [direction1, direction2] of Object.values(pipeToDirectionsMapping)) {
    let currP1: Point = [start[0] + direction1[0], start[1] + direction1[1]];
    let prevP1: Point = [...start];

    let currP2: Point = [start[0] + direction2[0], start[1] + direction2[1]];
    let prevP2: Point = [...start];

    // pipe should be connected to 2 neighbors
    if (
      !isConnectedPipe(currP1[0], currP1[1], matrix, direction1) ||
      !isConnectedPipe(currP2[0], currP2[1], matrix, direction2)
    ) {
      continue;
    }

    let steps = 1;
    // traverse from both sides in parallel
    // until pointers meet
    while (!arePointsEqual(currP1, currP2)) {
      let newCurr = getNextPoint(prevP1, currP1, matrix);
      prevP1 = currP1;
      currP1 = newCurr;

      newCurr = getNextPoint(prevP2, currP2, matrix);
      prevP2 = currP2;
      currP2 = newCurr;

      steps += 1;
    }

    answer = Math.max(answer, steps);
  }

  console.log("part 1 steps", answer);
  return answer;
}

part1();
