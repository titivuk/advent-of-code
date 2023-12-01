import { createReadStream } from "fs";
import { createInterface } from "readline/promises";

let digitsMapping = new Map<string, number>([
  ["one", 1],
  ["two", 2],
  ["three", 3],
  ["four", 4],
  ["five", 5],
  ["six", 6],
  ["seven", 7],
  ["eight", 8],
  ["nine", 9],
  ["1", 1],
  ["2", 2],
  ["3", 3],
  ["4", 4],
  ["5", 5],
  ["6", 6],
  ["7", 7],
  ["8", 8],
  ["9", 9],
]);

type LineParser = (line: string) => number;

function part1LineParser(line: string): number {
  let num = 0;

  let digits: number[] = [];
  let charCode = 0;
  for (let i = 0; i < line.length; i++) {
    charCode = line.charCodeAt(i);
    if (48 <= charCode && charCode <= 57) {
      digits.push(Number.parseInt(line[i], 10));
    }
  }
  num = digits[0] * 10 + digits[digits.length - 1];

  return num;
}

function part2LineParser(line: string): number {
  let num = 0;

  let digits: number[] = [];
  for (let i = 0; i < line.length; i++) {
    for (const [key, value] of digitsMapping) {
      if (line.startsWith(key, i)) {
        digits.push(value);
        // initially I moved pointer to key.length - 1
        // -1 because I use for-loop and too lazy to rewrite
        // so it will be incremented by for i++ statement
        // then I found that "eighthree" => 83 (strings can intersect)
        // maximum possible intersection is 1
        // as a result i += key.length - 1 - intersection = key.length - 2
        // but it does not work when key.length = 1
        // so Math.max comes to rescue
        i += Math.max(0, key.length - 2);
        break;
      }
    }
  }
  num = digits[0] * 10 + digits[digits.length - 1];

  return num;
}

async function main(lineParser: LineParser) {
  const inputStream = createReadStream("./2023/1/input.txt", {
    encoding: "utf-8",
  });

  const rl = createInterface({ input: inputStream });

  let sum = 0;

  for await (const line of rl) {
    sum += lineParser(line);
  }

  return sum;
}

main(part2LineParser);
