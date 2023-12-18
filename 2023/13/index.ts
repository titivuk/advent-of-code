import { readFile } from "fs/promises";

async function part1() {
  let fileContent = await readFile("./2023/13/input.txt", "utf-8");
  let patterns = fileContent.split("\n\n");

  let sum = 0;

  for (let p of patterns) {
    let rows = p.split("\n");

    let m = rows.length;
    let n = rows[0].length;

    let left = 0;
    let right = 1;

    // find vertical reflection
    while (right < n) {
      let isReflection = true;
      let numberOfPairs = Math.min(left, n - 1 - right) + 1;
      for (let offset = 0; offset < numberOfPairs; offset++) {
        for (let i = 0; i < m; i++) {
          if (rows[i][left - offset] !== rows[i][right + offset]) {
            isReflection = false;
            break;
          }
        }

        if (!isReflection) break;
      }

      // vertical reflection found
      if (isReflection) {
        sum += left + 1;
        break;
      }

      left += 1;
      right += 1;
    }

    // vertical reflection not found
    // find horizontal reflection
    if (right >= n) {
      let top = 0;
      let bottom = 1;
      while (bottom < m) {
        let isReflection = true;
        let numberOfPairs = Math.min(top, m - 1 - bottom) + 1;
        for (let offset = 0; offset < numberOfPairs; offset++) {
          if (rows[top - offset] !== rows[bottom + offset]) {
            isReflection = false;
            break;
          }
        }

        if (isReflection) {
          sum += (top + 1) * 100;
          break;
        }

        top += 1;
        bottom += 1;
      }
    }
  }

  debugger;
  console.log("part 1 sum", sum);
}

async function part2() {
  let fileContent = await readFile("./2023/13/input.txt", "utf-8");
  let patterns = fileContent.split("\n\n");

  let sum = 0;

  for (let p of patterns) {
    let rows = p.split("\n");

    let m = rows.length;
    let n = rows[0].length;

    let left = 0;
    let right = 1;

    // find vertical reflection
    while (right < n) {
      let isReflection = true;
      let diffChars = 0;

      let numberOfPairs = Math.min(left, n - 1 - right) + 1;
      for (let offset = 0; offset < numberOfPairs; offset++) {
        for (let i = 0; i < m; i++) {
          if (rows[i][left - offset] !== rows[i][right + offset]) {
            diffChars += 1;
          }

          if (diffChars > 1) {
            isReflection = false;
            break;
          }
        }

        if (!isReflection) break;
      }

      // vertical reflection found
      if (isReflection && diffChars === 1) {
        sum += left + 1;
        break;
      }

      left += 1;
      right += 1;
    }

    // vertical reflection not found
    // find horizontal reflection
    if (right >= n) {
      let top = 0;
      let bottom = 1;
      while (bottom < m) {
        let isReflection = true;
        let diffChars = 0;

        let numberOfPairs = Math.min(top, m - 1 - bottom) + 1;
        for (let offset = 0; offset < numberOfPairs; offset++) {
          for (let i = 0; i < n; i++) {
            if (rows[top - offset][i] !== rows[bottom + offset][i]) {
              diffChars += 1;
            }

            if (diffChars > 1) {
              isReflection = false;
              break;
            }
          }
        }

        // horizontal reflection found
        if (isReflection && diffChars === 1) {
          sum += (top + 1) * 100;
          break;
        }

        top += 1;
        bottom += 1;
      }
    }
  }

  debugger;
  console.log("part 1 sum", sum);
}

// part1();
part2();
