import java.nio.file.Files;
import java.nio.file.Path;

record Day2Result(Long part1, Long part2) {}

public class Day2 {

    private final String inputPath;

    public Day2(String inputPath) {
        this.inputPath = inputPath;
    }

    public Day2Result solve() throws Exception {
        String input = this.readInput(inputPath);

        String[] ranges = input.split(",");
        long sum1 = 0;
        long sum2 = 0;
        for (String range : ranges) {
            String[] borders = range.split("-");
            long start = Long.parseLong(borders[0]);
            long end = Long.parseLong(borders[1]);

            for (long id = start; id <= end; id++) {
                if (!isValidP1(id)) {
                    sum1 += id;
                }

                if (!isValidP2(id)) {
                    sum2 += id;
                }
            }
        }

        return new Day2Result(sum1, sum2);
    }

    private boolean isValidP1(long id) {
        long digits = (long) (Math.log10(id)) + 1;

        // if sequence is repeated twice, number of digits has to be even
        if (digits % 2 == 1) {
            return true;
        }

        long first = id / (long) Math.pow(10, digits / 2);
        long second = id % (long) Math.pow(10, digits / 2);

        return first != second;
    }

    private boolean isValidP2(long id) {
        String str = Long.toString(id);

        StringBuilder seq = new StringBuilder();
        for (int i = 0; i < str.length() / 2; i++) {
            seq.append(str.charAt(i));

            if (str.split(seq.toString()).length == 0) {
                return false;
            }
        }

        return true;
    }

    private String readInput(String inputPath) throws Exception {
        String input = Files.readString(Path.of(inputPath));
        return input;
    }
}
