import { readFile } from "fs/promises";

type Coordinate = [number, number];

function findDistance(
  a: Coordinate,
  b: Coordinate,
  rowsExpansion: number[],
  colsExpansion: number[]
) {
  let distance = 0;

  let startRow = Math.min(a[0], b[0]);
  let endRow = Math.max(a[0], b[0]);
  for (let i = startRow + 1; i <= endRow; i++) {
    distance += rowsExpansion[i];
  }

  let startCol = Math.min(a[1], b[1]);
  let endCol = Math.max(a[1], b[1]);
  for (let i = startCol + 1; i <= endCol; i++) {
    distance += colsExpansion[i];
  }

  return distance;
}

async function main(expansionFactor: number) {
  let fileContent = await readFile("./2023/11/input.txt", "utf-8");
  let fileRows = fileContent.split("\n");

  let n = fileRows.length;

  let galaxies: Coordinate[] = [];
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (fileRows[i][j] === "#") {
        galaxies.push([i, j]);
      }
    }
  }

  let rowsExpansion: number[] = new Array(n).fill(1);
  let colsExpansion: number[] = new Array(n).fill(1);
  for (let d = 0; d < n; d++) {
    if (!fileRows[d].includes("#")) rowsExpansion[d] = expansionFactor;
    if (fileRows.every((row) => row[d] === "."))
      colsExpansion[d] = expansionFactor;
  }

  let sum = 0;

  for (let i = 0; i < galaxies.length - 1; i++) {
    for (let j = i + 1; j < galaxies.length; j++) {
      sum += findDistance(
        galaxies[i],
        galaxies[j],
        rowsExpansion,
        colsExpansion
      );
    }
  }

  debugger;
  return sum;
}

main(1000000);
