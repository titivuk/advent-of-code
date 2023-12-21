let str = "this turns into 'this string'";

function split(str: string) {
  let arr: string[] = [];

  let left = 0;
  let right = 0;

  while (right < str.length) {
    if (str[left] === "'") {
      while (right < str.length && str[right] !== "'") right++;
      arr.push(str.substring(left, right + 1));
    } else {
      while (right < str.length && str[right] !== " ") right++;
      arr.push(str.substring(left, right));
    }

    left = right + 1;
    right = left + 1;
  }

  debugger;
}

split(str);
