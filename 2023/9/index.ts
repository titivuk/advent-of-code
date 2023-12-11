import { readFile } from "fs/promises";

async function part1() {
  const fileContent = await readFile("./2023/9/input.txt", "utf-8");

  let rows = fileContent
    .split("\n")
    .map((row) => row.split(" ").map((val) => Number.parseInt(val, 10)));

  let sum = 0;

  for (let row of rows) {
    let currSum = -1;
    let currSeq: number[] = [];
    let nextSeq = row;
    let lastElements: number[] = [];
    let someNonZeroes = true;

    while (someNonZeroes) {
      currSum = 0;
      currSeq = nextSeq;
      nextSeq = [];
      someNonZeroes = false;

      for (let i = 1; i < currSeq.length; i++) {
        let pairDiff = currSeq[i] - currSeq[i - 1];

        if (pairDiff !== 0) someNonZeroes = true;

        nextSeq.push(pairDiff);
        currSum += pairDiff;
      }

      lastElements.push(currSeq.at(-1)!);
    }

    let extrapolation = 0;
    for (let i = lastElements.length - 1; i >= 0; i--) {
      extrapolation += lastElements[i];
    }

    sum += extrapolation;
  }

  console.log("part 1 sum", sum);
  return sum;
}

async function part2() {
  const fileContent = await readFile("./2023/9/input.txt", "utf-8");

  let rows = fileContent
    .split("\n")
    .map((row) => row.split(" ").map((val) => Number.parseInt(val, 10)));

  let sum = 0;

  for (let row of rows) {
    let currSum = -1;
    let currSeq: number[] = [];
    let nextSeq = row;
    let firstElements: number[] = [];
    let someNonZeroes = true;

    while (someNonZeroes) {
      currSum = 0;
      currSeq = nextSeq;
      nextSeq = [];
      someNonZeroes = false;

      for (let i = 1; i < currSeq.length; i++) {
        let pairDiff = currSeq[i] - currSeq[i - 1];

        if (pairDiff !== 0) someNonZeroes = true;

        nextSeq.push(pairDiff);
        currSum += pairDiff;
      }

      firstElements.push(currSeq[0]);
    }

    let prevExtrapolation = 0;
    let extrapolation = 0;
    for (let i = firstElements.length - 1; i >= 0; i--) {
      extrapolation = firstElements[i] - prevExtrapolation;
      prevExtrapolation = extrapolation;
    }

    sum += extrapolation;
  }

  console.log("part 1 sum", sum);
  return sum;
}

// part1();
part2();
