import { createReadStream } from "fs";
import { createInterface } from "readline/promises";

type GameRequiredColors = {
  gameId: number;
  red: number;
  green: number;
  blue: number;
  [key: string]: number;
};

function getNumberOfRequiredColorsForGame(
  gameLine: string
): GameRequiredColors {
  const [header, payload] = gameLine.split(": ");

  let [, id] = header.split(" ");
  let sets = payload.split("; ");

  let results: GameRequiredColors = {
    gameId: Number.parseInt(id),
    red: 0,
    green: 0,
    blue: 0,
  };

  for (let set of sets) {
    let setItems = set.split(", ");

    for (const item of setItems) {
      let [num, color] = item.split(" ");
      results[color] = Math.max(results[color], Number.parseInt(num, 10));
    }
  }

  return results;
}

async function main() {
  const inputStream = createReadStream("./2023/2/input.txt", {
    encoding: "utf-8",
  });
  const rl = createInterface({ input: inputStream });

  let part1Sum = 0,
    part2Sum = 0;

  const RED_MAX = 12,
    GREEN_MAX = 13,
    BLUE_MAX = 14;

  for await (let line of rl) {
    let results = getNumberOfRequiredColorsForGame(line);

    // part 1
    if (
      results.red <= RED_MAX &&
      results.green <= GREEN_MAX &&
      results.blue <= BLUE_MAX
    ) {
      part1Sum += results.gameId;
    }

    // part2
    part2Sum += results.blue * results.green * results.red;
  }

  console.log("part1 sum:", part1Sum);
  console.log("part2 sum:", part2Sum);

  return part1Sum;
}

main();
