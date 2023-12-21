import { readFile } from "fs/promises";

type Direction = [number, number];
type Point = [number, number];

type Beam = {
  position: Point;
  direction: Direction;
};

let topDirection: Direction = [-1, 0];
let leftDirection: Direction = [0, -1];
let bottomDirection: Direction = [1, 0];
let rightDirection: Direction = [0, 1];

function isValidDirection(i: number, j: number, m: number, n: number) {
  return 0 <= i && i < m && 0 <= j && j < n;
}

function startBeam(
  start: Point,
  grid: string[][],
  direction: Direction
): number {
  let sum = 1;

  let n = grid[0].length;

  let visited: Direction[][][] = [];
  for (let i = 0; i < n; i++) {
    visited[i] = [];
    for (let j = 0; j < n; j++) {
      visited[i][j] = [];
    }
  }
  visited[start[0]][start[1]].push(direction);

  let currentBeams: Beam[] = [{ position: start, direction }];
  let nextBeams: Beam[] = [];

  while (currentBeams.length > 0) {
    for (let beam of currentBeams) {
      let nextPosition: Point = [
        beam.position[0] + beam.direction[0],
        beam.position[1] + beam.direction[1],
      ];

      // TODO: refactor
      // very first attempt to reflect logic
      if (
        isValidDirection(nextPosition[0], nextPosition[1], n, n) &&
        visited[nextPosition[0]][nextPosition[1]].every(
          (dir) => beam.direction[0] !== dir[0] || beam.direction[1] !== dir[1]
        )
      ) {
        if (visited[nextPosition[0]][nextPosition[1]].length === 0) {
          sum += 1;
        }
        visited[nextPosition[0]][nextPosition[1]].push(beam.direction);

        // keep going in the same direction
        if (grid[nextPosition[0]][nextPosition[1]] === ".") {
          nextBeams.push({
            position: nextPosition,
            direction: [...beam.direction],
          });
        } else if (grid[nextPosition[0]][nextPosition[1]] === "-") {
          // pointy end of a splitter
          if (beam.direction[0] === 0) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...beam.direction],
            });
          }
          // split beam into two with opposite directions
          else {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...leftDirection],
            });
            nextBeams.push({
              position: [...nextPosition],
              direction: [...rightDirection],
            });
          }
        } else if (grid[nextPosition[0]][nextPosition[1]] === "|") {
          // pointy end of a splitter
          if (beam.direction[1] === 0) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...beam.direction],
            });
          }
          // split beam into two with opposite directions
          else {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...topDirection],
            });
            nextBeams.push({
              position: [...nextPosition],
              direction: [...bottomDirection],
            });
          }
        } else if (grid[nextPosition[0]][nextPosition[1]] === "/") {
          // right direction
          if (beam.direction[1] === 1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...topDirection],
            });
          }
          // left direction
          else if (beam.direction[1] === -1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...bottomDirection],
            });
          }
          // bottom direction
          else if (beam.direction[0] === 1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...leftDirection],
            });
          }
          // top direction
          else if (beam.direction[0] === -1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...rightDirection],
            });
          }
        } else if (grid[nextPosition[0]][nextPosition[1]] === "\\") {
          // right direction
          if (beam.direction[1] === 1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...bottomDirection],
            });
          }
          // left direction
          else if (beam.direction[1] === -1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...topDirection],
            });
          }
          // bottom direction
          else if (beam.direction[0] === 1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...rightDirection],
            });
          }
          // top direction
          else if (beam.direction[0] === -1) {
            nextBeams.push({
              position: [...nextPosition],
              direction: [...leftDirection],
            });
          }
        }
      }
    }

    currentBeams = nextBeams;
    nextBeams = [];
  }

  return sum;
}

async function part1() {
  let fileContent = await readFile("./2023/16/input.txt", "utf-8");

  let grid = fileContent.split("\n").map((row) => row.split(""));

  let sum = startBeam([0, 0], grid, rightDirection);

  debugger;
  console.log("part 1 sum", sum);
  return sum;
}

async function part2() {
  let fileContent = await readFile("./2023/16/input.txt", "utf-8");

  let grid = fileContent.split("\n").map((row) => row.split(""));
  let n = grid.length;

  let sum = 0;
  // top
  for (let j = 0; j < n; j++) {
    sum = Math.max(sum, startBeam([0, j], grid, bottomDirection));
  }
  // bottom
  for (let j = 0; j < n; j++) {
    sum = Math.max(sum, startBeam([n - 1, j], grid, topDirection));
  }
  // left
  for (let i = 0; i < n; i++) {
    sum = Math.max(sum, startBeam([i, 0], grid, rightDirection));
  }
  // right
  for (let i = 0; i < n; i++) {
    sum = Math.max(sum, startBeam([i, n - 1], grid, leftDirection));
  }

  debugger;
  console.log("part 2 sum", sum);
  return sum;
}

// part1();
part2();
