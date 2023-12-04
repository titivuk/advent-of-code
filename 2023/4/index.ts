import { readFile } from "fs/promises";

function isDigit(ch: string): boolean {
  return 48 <= ch.charCodeAt(0) && ch.charCodeAt(0) <= 57;
}

async function part1() {
  const fileContent = await readFile("./2023/4/input-test.txt", {
    encoding: "utf-8",
  });
  let fileRows = fileContent.split("\n");

  let sum = 0;
  let winningNumbers = new Set<string>();
  let matchingNumbers = 0;
  for (const row of fileRows) {
    matchingNumbers = 0;
    winningNumbers.clear();

    /**
     * Find winning number and put them into winningNumbers set
     */
    let j = row.length - 1;
    let numEnd = -1;
    while (row[j] !== "|") {
      if (isDigit(row[j])) {
        if (numEnd === -1) numEnd = j;
      } else if (numEnd !== -1) {
        winningNumbers.add(row.substring(j + 1, numEnd + 1));
        numEnd = -1;
      }

      j -= 1;
    }

    /**
     * Find elf numbers and calculate how many of them are winning
     */
    numEnd = -1;
    j -= 1;
    while (row[j] !== ":") {
      if (isDigit(row[j])) {
        if (numEnd === -1) numEnd = j;
      } else if (numEnd !== -1) {
        if (winningNumbers.has(row.substring(j + 1, numEnd + 1))) {
          matchingNumbers += 1;
        }

        numEnd = -1;
      }

      j -= 1;
    }

    if (matchingNumbers > 0) {
      sum += 2 ** (matchingNumbers - 1);
    }
  }

  console.log("part 1 sum:", sum);
}

async function part2() {
  const fileContent = await readFile("./2023/4/input.txt", {
    encoding: "utf-8",
  });
  let fileRows = fileContent.split("\n");

  let cardCounter = new Array<number>(fileRows.length).fill(1);
  let totalNumberOfCards = 0;
  let winningNumbers = new Set<string>();
  let matchingNumbers = 0;
  let row = "";
  for (let i = 0; i < fileRows.length; i++) {
    row = fileRows[i];
    matchingNumbers = 0;
    winningNumbers.clear();

    /**
     * Find winning number and put them into winningNumbers set
     */
    let j = row.length - 1;
    let numEnd = -1;
    while (row[j] !== "|") {
      if (isDigit(row[j])) {
        if (numEnd === -1) numEnd = j;
      } else if (numEnd !== -1) {
        winningNumbers.add(row.substring(j + 1, numEnd + 1));
        numEnd = -1;
      }

      j -= 1;
    }

    /**
     * Find elf numbers and calculate how many of them are winning
     */
    numEnd = -1;
    j -= 1;
    while (row[j] !== ":") {
      if (isDigit(row[j])) {
        if (numEnd === -1) numEnd = j;
      } else if (numEnd !== -1) {
        if (winningNumbers.has(row.substring(j + 1, numEnd + 1))) {
          matchingNumbers += 1;
        }

        numEnd = -1;
      }

      j -= 1;
    }

    if (matchingNumbers > 0) {
      // don't really like it, may be it's possible to optimize it
      let startCardIndex = i + 1;
      for (
        let cardIndex = startCardIndex;
        cardIndex < startCardIndex + matchingNumbers;
        cardIndex++
      ) {
        cardCounter[cardIndex] += cardCounter[i];
      }
    }
  }

  for (let counter of cardCounter) {
    totalNumberOfCards += counter;
  }

  console.log("total number of cards:", totalNumberOfCards);
}

part2();
