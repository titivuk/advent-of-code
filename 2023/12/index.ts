import { strictEqual } from "assert";
import { readFile } from "fs/promises";

/**
 * check that every subsequence of '#' satisfies corresponding group length
 */
function isValidSpring(spring: string[], groups: number[]) {
  let currGroupIndex = 0;
  let currGroupLength = 0;

  let i = 0;
  while (i < spring.length) {
    if (spring[i] === "#") {
      while (spring[i] === "#") {
        currGroupLength += 1;
        i += 1;
      }

      if (currGroupLength !== groups[currGroupIndex]) {
        return false;
      }

      currGroupLength = 0;
      currGroupIndex += 1;
    } else if (spring[i] === "?") {
      return false;
    } else {
      i += 1;
    }
  }

  // all groups checked
  return currGroupIndex === groups.length;
}

async function part1() {
  let fileContent = await readFile("./2023/12/input.txt", "utf-8");
  let fileRows = fileContent.split("\n");

  let sum = 0;

  function backtrack(
    spring: string[],
    startIndex: number,
    groups: number[]
  ): number {
    if (spring[startIndex] === "?") {
      spring[startIndex] = "#";
      let r1 = backtrack(spring, startIndex + 1, groups);
      spring[startIndex] = "?";

      spring[startIndex] = ".";
      let r2 = backtrack(spring, startIndex + 1, groups);
      spring[startIndex] = "?";

      return r1 + r2;
    }

    return backtrack(spring, startIndex + 1, groups);
  }

  for (const row of fileRows) {
    let [springStr, groupsStr] = row.split(" ");
    let spring = springStr.split("");
    let groups = groupsStr.split(",").map((g) => +g);

    sum += backtrack(spring, 0, groups);
  }
}

/**
 * for part 2
 */

// let [springStr, groupsStr] = row.split(" ");
// let spring = new Array(5).fill(springStr).join("?").split("");
// let groups = new Array(5)
//   .fill(groupsStr)
//   .join(",")
//   .split(",")
//   .map((g) => +g);

part1();

strictEqual(isValidSpring("#.#.###".split(""), [1, 1, 3]), true);
strictEqual(isValidSpring("#.#.###".split(""), [1, 1, 2]), false);
strictEqual(isValidSpring("#.#.##".split(""), [1, 1, 3]), false);
strictEqual(isValidSpring("#.#".split(""), [1, 1, 3]), false);
strictEqual(isValidSpring("#.#.".split(""), [1, 1, 3]), false);
strictEqual(isValidSpring(".??..??...?##.".split(""), [1, 1, 3]), false);

strictEqual(isValidSpring(".###.##.#...".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###.##..#..".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###.##...#.".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###.##....#".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###..##.#..".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###..##..#.".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###..##...#".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###...##.#.".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###...##..#".split(""), [3, 2, 1]), true);
strictEqual(isValidSpring(".###....##.#".split(""), [3, 2, 1]), true);
