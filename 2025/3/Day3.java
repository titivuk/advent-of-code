import java.nio.file.Files;
import java.nio.file.Path;
import java.util.PriorityQueue;

record IndexedValue(int val, int idx) {
    @Override
    public String toString() {
        return "IndexedValue(val=" + val + ", idx=" + idx + ")";
    }
}

public class Day3 {

    private final String inputPath;

    public Day3(String inputPath) {
        this.inputPath = inputPath;
    }

    public long part1() throws Exception {
        String input = this.readInput(inputPath);

        String[] lines = input.lines().toArray(String[]::new);
        long sum = 0;
        for (String line : lines) {
            char first = '0';
            int firstIdx = 0;
            for (int i = 0; i < line.length() - 1; i++) {
                if (first < line.charAt(i)) {
                    first = line.charAt(i);
                    firstIdx = i;
                }
            }

            char second = '0';
            for (int i = firstIdx + 1; i < line.length(); i++) {
                if (second < line.charAt(i)) {
                    second = line.charAt(i);
                }
            }

            sum += (first - 48) * 10 + (second - 48);
        }

        return sum;
    }

    public long part2() throws Exception {
        String input = this.readInput(inputPath);

        String[] lines = input.lines().toArray(String[]::new);
        long sum = 0;
        for (String line : lines) {
            PriorityQueue<IndexedValue> maxQueue = new PriorityQueue<>(
                (a, b) -> {
                    if (a.val() == b.val()) {
                        return a.idx() - b.idx();
                    }

                    return b.val() - a.val();
                }
            );

            // add all possible elements for the 1st digit
            // except the rightmost (we will add it later)
            for (int i = 0; i < line.length() - 12; i++) {
                maxQueue.add(new IndexedValue(line.charAt(i) - 48, i));
            }

            long joltage = 0;
            int curIdx = 0;
            for (int i = 11; i >= 0; i--) {
                // for every next digit new element becomes available as potential max value
                maxQueue.add(
                    new IndexedValue(
                        line.charAt(line.length() - i - 1) - 48,
                        line.length() - i - 1
                    )
                );

                // get max value
                IndexedValue maxValue = maxQueue.poll();
                if (maxValue == null) {
                    throw new Exception("Unexpected empty queue");
                }

                // increase joltage
                joltage += (long) maxValue.val() * (long) Math.pow(10, i);
                curIdx = maxValue.idx();

                // remove elements from PQ which are on the left side of maxValue
                while (
                    maxQueue.peek() != null && maxQueue.peek().idx() <= curIdx
                ) maxQueue.poll();
            }

            sum += joltage;
        }

        return sum;
    }

    private String readInput(String inputPath) throws Exception {
        String input = Files.readString(Path.of(inputPath));
        return input;
    }
}
