import assert from "assert";
import { readFile } from "fs/promises";

async function part1() {
  const fileContent = await readFile("./2023/6/input.txt", "utf-8");

  const [times, distances] = fileContent.split("\n").map((row) =>
    row
      .split(":")[1]
      .trim()
      .split(/\s+/g)
      .map((value) => +value)
  );

  assert(
    times.length === distances.length,
    "time.length and distance.length not equal"
  );

  let n = times.length;
  let result = 1;

  let d = 0;
  let t = 0;
  for (let i = 0; i < n; i++) {
    t = times[i];
    d = distances[i];

    let ways = 0;
    for (let s = 1; s < t; s++) {
      if (s * (t - s) > d) ways += 1;
      else if (ways > 0) break;
    }

    result *= ways;
  }

  console.log("part 1", result);
  return result;
}

/**
 * some kind of binary search to optimize?
 */
async function part2() {
  const fileContent = await readFile("./2023/6/input.txt", "utf-8");

  const [time, distance] = fileContent
    .split("\n")
    .map((row) => +row.split(":")[1].trim().replaceAll(" ", ""));

  let minSpeed = 1;
  while (minSpeed * (time - minSpeed) <= distance) {
    minSpeed += 1;
  }

  let maxSpeed = time - 1;
  while (maxSpeed * (time - maxSpeed) <= distance) {
    maxSpeed -= 1;
  }

  console.log("part 2", maxSpeed - minSpeed + 1);
  return maxSpeed - minSpeed + 1;
}

part1();
part2();
