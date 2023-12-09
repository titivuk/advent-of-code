import { readFile } from "fs/promises";

// larger is better
enum HandType {
  FiveOfAKind = 7,
  FourOfAKind = 6,
  FullHouse = 5,
  ThreeOfAKind = 4,
  TwoPair = 3,
  OnePair = 2,
  HighCard = 1,
}

type Hand = {
  hand: string;
  bid: number;
  priority: number;
};

let cardStrengthPart1: Record<string, number> = {
  A: 13,
  K: 12,
  Q: 11,
  J: 10,
  T: 9,
  "9": 8,
  "8": 7,
  "7": 6,
  "6": 5,
  "5": 4,
  "4": 3,
  "3": 2,
  "2": 1,
};

let cardStrengthPart2: Record<string, number> = {
  A: 13,
  K: 12,
  Q: 11,
  T: 9,
  "9": 8,
  "8": 7,
  "7": 6,
  "6": 5,
  "5": 4,
  "4": 3,
  "3": 2,
  "2": 1,
  J: 0,
};

let cardCounter = new Map<string, number>();
function calculateHandPriorityPart1(hand: string): number {
  cardCounter.clear();

  let jokerCounter = 0;
  for (let card of hand) {
    if (card === "J") {
      jokerCounter += 1;
    }

    cardCounter.set(card, (cardCounter.get(card) ?? 0) + 1);
  }

  let distinctCards = cardCounter.size;
  let handType = 0;
  let cardHandType = 0;
  for (let [card, counter] of cardCounter) {
    if (counter === 5) {
      cardHandType = HandType.FiveOfAKind;
    } else if (counter === 4) {
      cardHandType = HandType.FourOfAKind;
    } else if (counter === 3 && distinctCards === 2) {
      cardHandType = HandType.FullHouse;
    } else if (counter === 3 && distinctCards === 3) {
      cardHandType = HandType.ThreeOfAKind;
    } else if (counter === 2 && distinctCards === 3) {
      cardHandType = HandType.TwoPair;
    } else if (counter === 2 && distinctCards === 4) {
      cardHandType = HandType.OnePair;
    } else if (distinctCards === 5) {
      cardHandType = HandType.HighCard;
    }

    handType = Math.max(handType, cardHandType);
  }

  cardCounter.clear();
  return handType;
}

function calculateHandPriorityPart2(hand: string): number {
  cardCounter.clear();

  let jokerCounter = 0;
  for (let card of hand) {
    if (card === "J") {
      jokerCounter += 1;
    }

    cardCounter.set(card, (cardCounter.get(card) ?? 0) + 1);
  }

  let distinctCards = cardCounter.size;
  let handType = 0;
  let cardHandType = 0;
  for (let [card, counter] of cardCounter) {
    // reset
    distinctCards = cardCounter.size;

    // if we can use joker
    if (card !== "J" && jokerCounter > 0) {
      // increase card counter by number of jokers
      counter += jokerCounter;
      // reduce distinct cards number
      // because we converted all jokers to current card
      distinctCards -= 1;
    }

    if (counter === 5) {
      cardHandType = HandType.FiveOfAKind;
    } else if (counter === 4) {
      cardHandType = HandType.FourOfAKind;
    } else if (counter === 3 && distinctCards === 2) {
      cardHandType = HandType.FullHouse;
    } else if (counter === 3 && distinctCards === 3) {
      cardHandType = HandType.ThreeOfAKind;
    } else if (counter === 2 && distinctCards === 3) {
      cardHandType = HandType.TwoPair;
    } else if (counter === 2 && distinctCards === 4) {
      cardHandType = HandType.OnePair;
    } else if (distinctCards === 5) {
      cardHandType = HandType.HighCard;
    }

    handType = Math.max(handType, cardHandType);
  }

  cardCounter.clear();
  return handType;
}

function getHandsComparator(cardStrength: Record<string, number>) {
  return function handsComparator(a: Hand, b: Hand) {
    if (a.priority === b.priority) {
      for (let i = 0; i < a.hand.length; i++) {
        if (a.hand[i] !== b.hand[i]) {
          return cardStrength[a.hand[i]] - cardStrength[b.hand[i]];
        }
      }
    }

    return a.priority - b.priority;
  };
}

async function part1() {
  const fileContent = await readFile("./2023/7/input.txt", "utf-8");

  const hands: Hand[] = fileContent.split("\n").map((row) => {
    let [hand, bidStr] = row.split(" ");

    return {
      hand,
      bid: Number.parseInt(bidStr, 10),
      priority: calculateHandPriorityPart1(hand),
    };
  });

  let handsComparator = getHandsComparator(cardStrengthPart1);
  hands.sort(handsComparator);

  let totalWinnings = 0;
  for (let i = 0; i < hands.length; i++) {
    totalWinnings += hands[i].bid * (i + 1);
  }

  console.log(totalWinnings);
  return totalWinnings;
}

async function part2() {
  const fileContent = await readFile("./2023/7/input.txt", "utf-8");

  const hands: Hand[] = fileContent.split("\n").map((row) => {
    let [hand, bidStr] = row.split(" ");

    return {
      hand,
      bid: Number.parseInt(bidStr, 10),
      priority: calculateHandPriorityPart2(hand),
    };
  });

  let handsComparator = getHandsComparator(cardStrengthPart2);
  hands.sort(handsComparator);

  let totalWinnings = 0;
  for (let i = 0; i < hands.length; i++) {
    totalWinnings += hands[i].bid * (i + 1);
  }

  console.log(totalWinnings);
  return totalWinnings;
}

// part1();
// part2();
