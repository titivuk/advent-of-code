import { readFile } from "fs/promises";

function getHash(str: string): number {
  let currentValue = 0;
  for (let i = 0; i < str.length; i++) {
    currentValue += str.charCodeAt(i);
    currentValue = (currentValue * 17) % 256;
  }

  return currentValue;
}

async function part1() {
  let sequence = await readFile("./2023/15/input.txt", "utf-8");

  let steps = sequence.split(",");

  let sum = 0;
  for (let step of steps) {
    sum += getHash(step);
  }

  console.log("part 1 sum", sum);
  return sum;
}

async function part2() {
  let sequence = await readFile("./2023/15/input.txt", "utf-8");

  let steps = sequence.split(",");

  let boxes: Set<string>[] = [];
  for (let i = 0; i < 256; i++) {
    boxes.push(new Set<string>());
  }

  let lenses = new Map<string, number>();

  let sum = 0;
  for (let step of steps) {
    let i = 0;
    while (step[i] !== "=" && step[i] !== "-") i++;
    let label = step.substring(0, i);
    let operation = step[i];
    let box = getHash(label);

    if (operation === "=") {
      let focalLength = Number.parseInt(step.substring(i + 1), 10);

      lenses.set(label, focalLength);

      if (!boxes[box].has(label)) {
        boxes[box].add(label);
      }
    } else if (operation === "-") {
      boxes[box].delete(label);
    }
  }

  for (let i = 0; i < boxes.length; i++) {
    let currIdx = 1;
    for (const label of boxes[i]) {
      let focusingPower = (i + 1) * currIdx * lenses.get(label)!;
      sum += focusingPower;
      currIdx += 1;
    }
  }

  console.log("part 2 sum", sum);
  return sum;
}

// part1();
part2();
