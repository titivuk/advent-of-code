import { readFile } from "fs/promises";

function tiltNorth(dish: string[][]) {
  let m = dish.length;
  let n = dish[0].length;

  let sum = 0;

  // for every column
  for (let c = 0; c < n; c++) {
    let lastOccupiedRow = -1;

    // iterate over column from top to bottom
    for (let r = 0; r < m; r++) {
      let cell = dish[r][c];

      // update lastOccupiedRow
      if (cell === "#") {
        lastOccupiedRow = r;
      }
      // roll rounded rock to the north
      else if (cell === "O") {
        dish[r][c] = ".";
        dish[lastOccupiedRow + 1][c] = "O";

        lastOccupiedRow += 1;

        sum += m - lastOccupiedRow;
      }
    }
  }

  return sum;
}

function tiltWest(dish: string[][]) {
  let m = dish.length;
  let n = dish[0].length;

  // for every row
  for (let r = 0; r < m; r++) {
    let lastOccupiedColumn = -1;

    // iterate over row from left to right
    for (let c = 0; c < n; c++) {
      let cell = dish[r][c];

      // update lastOccupiedColumn
      if (cell === "#") {
        lastOccupiedColumn = c;
      }
      // roll rounded rock to the west
      else if (cell === "O") {
        dish[r][c] = ".";
        dish[r][lastOccupiedColumn + 1] = "O";

        lastOccupiedColumn += 1;
      }
    }
  }
}

function tiltSouth(dish: string[][]) {
  let m = dish.length;
  let n = dish[0].length;

  // for every column
  for (let c = 0; c < n; c++) {
    let lastOccupiedRow = m;

    // iterate over column from bottom to top
    for (let r = m - 1; r >= 0; r--) {
      let cell = dish[r][c];

      // update lastOccupiedRow
      if (cell === "#") {
        lastOccupiedRow = r;
      }
      // roll rounded rock to the south
      else if (cell === "O") {
        dish[r][c] = ".";
        dish[lastOccupiedRow - 1][c] = "O";

        lastOccupiedRow -= 1;
      }
    }
  }
}

function tiltEast(dish: string[][]) {
  let m = dish.length;
  let n = dish[0].length;

  // for every row
  for (let r = 0; r < m; r++) {
    let lastOccupiedColumn = n;

    // iterate over row from right to left
    for (let c = n - 1; c >= 0; c--) {
      let cell = dish[r][c];

      // update lastOccupiedColumn
      if (cell === "#") {
        lastOccupiedColumn = c;
      }
      // roll rounded rock to the east
      else if (cell === "O") {
        dish[r][c] = ".";
        dish[r][lastOccupiedColumn - 1] = "O";

        lastOccupiedColumn -= 1;
      }
    }
  }
}

async function part1() {
  let fileContent = await readFile("./2023/14/input.txt", "utf-8");
  let dish = fileContent.split("\n").map((row) => row.split(""));

  let sum = tiltNorth(dish);

  console.log("part 1 sum", sum);
  return sum;
}

function getDishKey(dish: string[][]) {
  return dish.map((row) => row.join("")).join("");
}

function calculateNorthLoad(dish: string[][]) {
  let m = dish.length;
  let n = dish[0].length;

  let sum = 0;

  for (let i = 0; i < n; i++) {
    for (let j = 0; j < m; j++) {
      if (dish[j][i] === "O") {
        sum += m - j;
      }
    }
  }

  return sum;
}

async function part2() {
  let fileContent = await readFile("./2023/14/input.txt", "utf-8");
  let dish = fileContent.split("\n").map((row) => row.split(""));

  let cycles = 1_000_000_000;

  let cycle2Dish = new Map<number, string[][]>();
  let dish2Cycle = new Map<string, number>();

  for (let i = 1; i <= cycles; i++) {
    tiltNorth(dish);
    tiltWest(dish);
    tiltSouth(dish);
    tiltEast(dish);

    let dishKey = getDishKey(dish);

    if (dish2Cycle.has(dishKey)) {
      // iteration (counts from 1) where dish matrix starts being repeated with some period
      let cycleStart = dish2Cycle.get(dishKey)!;
      // repeat period
      let period = i - cycleStart;
      // number of times the cycle repeats until 1e9 cycle
      let numberOfCycles = Math.floor((cycles - cycleStart) / period);
      // number of the cycle which has the same state as 1e9 cycle
      let it = cycles - numberOfCycles * period;

      let sum = calculateNorthLoad(cycle2Dish.get(it)!);

      return sum;
    }

    dish2Cycle.set(dishKey, i);
    cycle2Dish.set(i, structuredClone(dish));
  }
}

// part1();
part2();
