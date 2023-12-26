import { readFile } from "fs/promises";

type Part = Record<string, number>;

type Rule = {
  category: string;
  value: number;
  comparator: typeof less;
  destination: string;
};

type Workflow = {
  name: string;
  rules: Rule[];
  fallbackDestination: string;
};

function less(part: number, rule: number) {
  return part < rule;
}

function greater(part: number, rule: number) {
  return part > rule;
}

function getWorkflows(workflowData: string) {
  let workflowsMap = new Map<string, Workflow>();

  workflowData.split("\n").map((line) => {
    let [workflow, rest] = line.split("{");
    let rules = rest.substring(0, rest.length - 1).split(",");

    let parsedRules: Rule[] = [];

    // skip last rule
    for (let i = 0; i < rules.length - 1; i++) {
      let rule = rules[i];

      let [expression, destination] = rule.split(":");

      let j = 0;
      while (expression[j] !== "<" && expression[j] !== ">") j += 1;

      let comparator = expression[j] === "<" ? less : greater;

      let category = expression.substring(0, j);
      let value = Number.parseInt(expression.substring(j + 1), 10);

      parsedRules.push({
        category,
        value,
        comparator,
        destination,
      });
    }

    let fallbackDestination = rules[rules.length - 1];

    workflowsMap.set(workflow, {
      name: workflow,
      rules: parsedRules,
      fallbackDestination,
    });
  });

  return workflowsMap;
}

async function part1() {
  let fileContent = await readFile("./2023/19/input.txt", "utf-8");

  let [workflowData, ratingLines] = fileContent.split("\n\n");
  let workflowsMap = getWorkflows(workflowData);

  let sum = 0;
  ratingLines.split("\n").map((line) => {
    let part: Part = {};
    line
      .substring(1, line.length - 1)
      .split(",")
      .forEach((rating) => {
        let [name, value] = rating.split("=");
        part[name] = Number.parseInt(value, 10);
      });

    let currWorkflow = "in";
    let nextWorkflow = "";
    while (currWorkflow !== "A" && currWorkflow !== "R") {
      let workflow = workflowsMap.get(currWorkflow)!;

      for (const rule of workflow.rules) {
        if (rule.comparator(part[rule.category], rule.value)) {
          nextWorkflow = rule.destination;
          break;
        }
      }

      if (nextWorkflow === "") {
        nextWorkflow = workflow.fallbackDestination;
      }

      currWorkflow = nextWorkflow;
      nextWorkflow = "";
    }

    if (currWorkflow === "A") {
      sum += Object.values(part).reduce((total, value) => total + value, 0);
    }
  });

  return sum;
}

async function part2() {
  let fileContent = await readFile("./2023/19/input.txt", "utf-8");

  let [workflowData] = fileContent.split("\n\n");
  let workflowsMap = getWorkflows(workflowData);

  let sum = 0;

  function dfs(currWorkflow: string, part: Record<string, [number, number]>) {
    if (currWorkflow === "R") {
      return;
    }

    // if part is accepted
    // calculate all combinations possible with given part state
    if (currWorkflow === "A") {
      sum += Object.values(part).reduce(
        (accum, value) => accum * (value[1] - value[0] + 1),
        1
      );

      return;
    }

    let workflow = workflowsMap.get(currWorkflow)!;

    part = structuredClone(part);
    for (let rule of workflow.rules) {
      // less
      if (rule.comparator === less) {
        if (part[rule.category][0] < rule.value) {
          let newPart = structuredClone(part);
          newPart[rule.category][1] = rule.value - 1;

          dfs(rule.destination, newPart);

          // after current rule is handled update part state to make it falsy for the current rule
          // because the next rule is only applicable when all previous rules are falsy
          part[rule.category][0] = rule.value;
        }
      }
      // greater
      else {
        if (part[rule.category][1] > rule.value) {
          let newPart = structuredClone(part);
          newPart[rule.category][0] = rule.value + 1;

          dfs(rule.destination, newPart);

          // after current rule is handled update part state to make it falsy for the current rule
          // because the next rule is only applicable when all previous rules are falsy
          part[rule.category][1] = rule.value;
        }
      }
    }

    dfs(workflow.fallbackDestination, part);
  }

  dfs("in", {
    x: [1, 4000],
    m: [1, 4000],
    a: [1, 4000],
    s: [1, 4000],
  });

  debugger;
}

// part1();
part2();
