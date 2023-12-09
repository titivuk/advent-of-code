import { readFile } from "fs/promises";

let instructionToIndex: Record<string, number> = {
  L: 0,
  R: 1,
};
function getInstruction(instructions: string, step: number) {
  return instructionToIndex[instructions[step % instructions.length]];
}

function gcd(a: number, b: number) {
  if (b === 0) {
    return a;
  }

  return gcd(b, a % b);
}

// Returns LCM of array elements
function findLcm(nums: number[]) {
  // Initialize result
  let answer = nums[0];

  // ans contains LCM of nums[0], ..nums[i]
  // after i'th iteration,
  for (let i = 1; i < nums.length; i++)
    answer = (nums[i] * answer) / gcd(nums[i], answer);

  return answer;
}

async function part1() {
  const fileContent = await readFile("./2023/8/input.txt", "utf-8");

  let fileRows = fileContent.split("\n");

  let instructions = fileRows[0];
  let tree: Map<string, [string, string]> = new Map();
  let node = "";
  let left = "";
  let right = "";
  for (let i = 2; i < fileRows.length; i++) {
    node = fileRows[i].substring(0, 3);
    left = fileRows[i].substring(7, 10);
    right = fileRows[i].substring(12, 15);

    if (!tree.has(node)) {
      tree.set(node, [left, right]);
    }
  }

  let step = 0;
  let currentNode = "AAA";
  let nextNode = "";
  while (currentNode !== "ZZZ") {
    nextNode = tree.get(currentNode)!!![getInstruction(instructions, step)];
    step += 1;

    currentNode = nextNode;
  }

  console.log(step);
  return step;
}

/**
 * I tried to stop while loop when I will be in the situation
 * where all nextNodes are Z nodes
 * but it turned out it's trickier the step is so big
 * that we need to calculate step for every node separately
 * and then find LCM for array of the found steps
 */
async function part2() {
  const fileContent = await readFile("./2023/8/input.txt", "utf-8");

  let fileRows = fileContent.split("\n");

  let instructions = fileRows[0];
  let tree: Map<string, [string, string]> = new Map();
  let currentNodes: string[] = [];
  let node = "";
  let left = "";
  let right = "";
  for (let i = 2; i < fileRows.length; i++) {
    node = fileRows[i].substring(0, 3);
    left = fileRows[i].substring(7, 10);
    right = fileRows[i].substring(12, 15);

    if (!tree.has(node)) {
      tree.set(node, [left, right]);
    }

    if (node.at(-1) === "A") {
      currentNodes.push(node);
    }
  }

  let step = 0;
  let steps: number[] = [];
  let nextNodes: string[] = [];
  while (currentNodes.length > 0) {
    let direction = getInstruction(instructions, step);

    step += 1;

    for (let currNode of currentNodes) {
      let nextNode = tree.get(currNode)![direction];

      // if we found Z from currNode
      // stop that branch and remember how many steps it took
      if (nextNode.at(-1) === "Z") {
        steps.push(step);
      }
      // if Z not found -> keep searching
      else {
        nextNodes.push(nextNode);
      }
    }

    currentNodes = nextNodes;
    nextNodes = [];
  }

  // after we found how many steps required for every Z node
  // find LCM (The Least Common Multiple)
  let answer = findLcm(steps);

  console.log(answer);
  return answer;
}

// part1();
// part2();
