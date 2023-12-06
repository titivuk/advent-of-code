import { readFile } from "fs/promises";

type AlmanacMapEntry = {
  from: number;
  to: number;
  length: number;
};

type AlmanacMap = {
  name: string;
  entries: AlmanacMapEntry[];
};

function parseFilePart1(fileContent: string) {
  let sections = fileContent.split("\n\n");

  let seeds = sections[0]
    .split(": ")[1]
    .split(" ")
    .map((seed) => Number.parseInt(seed, 10));

  let mapSequence: AlmanacMap[] = parseMaps(sections);

  return { mapSequence, seeds };
}

function parseFilePart2(fileContent: string) {
  let sections = fileContent.split("\n\n");

  // cannot create seedRanges of seeds because its too large to keep in memory (O_O)
  let seedNumbers = sections[0]
    .split(": ")[1]
    .split(" ")
    .map((seed) => Number.parseInt(seed, 10));

  let mapSequence: AlmanacMap[] = parseMaps(sections);

  return { mapSequence, seedNumbers };
}

function parseMaps(sections: string[]) {
  let mapSequence: AlmanacMap[] = [];

  let section: string[] = [];
  let from = "",
    to = "",
    length = "";
  for (let i = 1; i < sections.length; i++) {
    section = sections[i].split("\n");

    let map: AlmanacMap = {
      name: section[0],
      entries: [],
    };
    for (let j = 1; j < section.length; j++) {
      [to, from, length] = section[j].split(" ");

      map.entries.push({
        from: Number.parseInt(from, 10),
        to: Number.parseInt(to, 10),
        length: Number.parseInt(length, 10),
      });
    }

    mapSequence.push(map);
  }

  return mapSequence;
}

async function part1() {
  let fileContent = await readFile("./2023/5/input.txt", {
    encoding: "utf-8",
  });

  let { mapSequence, seeds } = parseFilePart1(fileContent);

  let minLocation = Number.POSITIVE_INFINITY;
  let currentNumber = 0;
  for (let seed of seeds) {
    // start from seed
    currentNumber = seed;

    for (let map of mapSequence) {
      // try to find range that currentNumber fit
      for (let { from, to, length } of map.entries) {
        if (from <= currentNumber && currentNumber < from + length) {
          currentNumber += to - from;
          break;
        }
      }
    }

    minLocation = Math.min(minLocation, currentNumber);
  }

  console.log("part 1 min location:", minLocation);
}

/**
 * brute force
 */
async function part2() {
  let fileContent = await readFile("./2023/5/input.txt", {
    encoding: "utf-8",
  });

  let { mapSequence, seedNumbers } = parseFilePart2(fileContent);

  let minLocation = Number.POSITIVE_INFINITY;
  let currentNumber = 0;

  let start = 0;
  let end = 0;
  let seed = 0;
  for (let i = 0; i < seedNumbers.length - 1; i += 2) {
    start = seedNumbers[i];
    end = seedNumbers[i] + seedNumbers[i + 1] - 1;

    console.log(
      `seed sequence from ${start} to ${end}. Number of seeds: ${end - start}`
    );

    seed = start;
    while (seed <= end) {
      // start from seed
      currentNumber = seed;

      for (let map of mapSequence) {
        // try to find range that currentNumber fit
        for (let { from, to, length } of map.entries) {
          if (from <= currentNumber && currentNumber < from + length) {
            currentNumber += to - from;
            break;
          }
        }
      }

      minLocation = Math.min(minLocation, currentNumber);

      seed += 1;
    }
  }

  console.log("part 2 min location:", minLocation);
}

part2();
