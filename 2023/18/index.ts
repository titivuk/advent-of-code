import { readFile } from "fs/promises";

type Direction = [number, number];
type Point = [number, number];

let topDirection: Direction = [0, 1];
let leftDirection: Direction = [-1, 0];
let downDirection: Direction = [0, -1];
let rightDirection: Direction = [1, 0];

let digit2Direction: Record<string, Direction> = {
  "0": rightDirection,
  "1": downDirection,
  "2": leftDirection,
  "3": topDirection,
};

let symbol2Direction: Record<string, Direction> = {
  U: topDirection,
  R: rightDirection,
  D: downDirection,
  L: leftDirection,
};

async function part1() {
  let fileContent = await readFile("./2023/18/input.txt", "utf-8");

  let points: Point[] = [[0, 0]];

  let P = 0;
  fileContent.split("\n").forEach((row) => {
    let [dirStr, lengthStr] = row.split(" ");

    let direction = symbol2Direction[dirStr];
    let length = Number.parseInt(lengthStr, 10);

    P += length;

    points.push([
      points[points.length - 1][0] + direction[0] * length,
      points[points.length - 1][1] + direction[1] * length,
    ]);
  });

  let S = calculateS(points, P);

  return S;
}

async function part2() {
  let fileContent = await readFile("./2023/18/input.txt", "utf-8");

  let points: Point[] = [[0, 0]];

  let P = 0;
  fileContent.split("\n").forEach((row) => {
    let [, , color] = row.split(" ");

    let length = Number.parseInt(color.substring(2, color.length - 2), 16);
    let direction = digit2Direction[color[color.length - 2]];

    P += length;

    points.push([
      points[points.length - 1][0] + direction[0] * length,
      points[points.length - 1][1] + direction[1] * length,
    ]);
  });

  let S = calculateS(points, P);

  return S;
}

function calculateS(points: Point[], P: number): number {
  /**
   * Area calculation using shoelace algorithm
   * https://www.math.tolaso.com.gr/?p=1451
   */

  // reverse points excluding the first and the last element
  // note: firstEl = lastEl = start
  let left = 1;
  let right = points.length - 2;
  while (left < right) {
    [points[left], points[right]] = [points[right], points[left]];

    left += 1;
    right -= 1;
  }

  let A = 0;
  for (let i = 0; i < points.length - 1; i++) {
    A += (points[i + 1][0] + points[i][0]) * (points[i + 1][1] - points[i][1]);
  }
  A /= 2;

  // https://en.wikipedia.org/wiki/Pick%27s_theorem
  // A = i + b/2 - 1
  // where i - interior points (unknown)
  //       b - boundary points - in our case P (perimeter)
  //       A - area calculated by shoelace algorithm
  // we need to know total number of points S, i.e. S = (i + b)
  // let's derive it from the formula
  // i + b/2 = A + 1                 =>
  // 2i + b/2 = 2A + 2               =>
  // i + b = 2A + 2 - i              => derive i from the inital formula and replace it in the right side
  // i + b = 2A + 2 - (A - b/2 + 1)  =>
  // i + b = A + b/2 + 1             => replace symbols with the variables we use
  // S = i + b = A + P/2 + 1
  let S = A + P / 2 + 1;
  return S;
}

// part1();
part2();
