import java.nio.file.Files;
import java.nio.file.Path;

public class Day1 {
    public static int solve1(String inputFile) throws Exception {
        String input = Day1.readInput(inputFile);

        /*
        locals[0] - dial
        locals[1] - answer
         */
        int[] locals = {50, 0};
        input.lines().forEach(line -> {
            char direction = line.charAt(0);
            int distance = Integer.parseInt(line.substring(1));

            if (direction == 'L') {
                // update dial position
                locals[0] = locals[0] - (distance % 100);
                // left overflow
                if (locals[0] < 0) {
                    locals[0] = 100 + locals[0];
                }
            } else {
                locals[0] = (locals[0] + distance) % 100;
            }

            if (locals[0] == 0) {
                locals[1]++;
            }
        });

        return locals[1];
    }

    public static int solve2(String inputFile) throws Exception {
        String input = Day1.readInput(inputFile);

        final int DIAL_LEN = 100;
        final int START_DIAL = 50;

        /*
        locals[0] - dial
        locals[1] - answer
         */
        int[] locals = {START_DIAL, 0};
        input.lines().forEach(line -> {
            char direction = line.charAt(0);
            int distance = Integer.parseInt(line.substring(1));

            if (direction == 'L') {
                // we are on the border or cross it
                if (locals[0] - distance <= 0) {
                    // count how many times we were on the border during rotation
                    locals[1] += (distance - locals[0]) / DIAL_LEN;
                    if (locals[0] != 0) {
                        locals[1]++;
                    }
                }

                // update dial position
                locals[0] = locals[0] - (distance % DIAL_LEN);
                // left overflow
                if (locals[0] < 0) {
                    locals[0] = DIAL_LEN + locals[0];
                }
            } else {
                locals[1] += (locals[0] + distance) / DIAL_LEN;
                locals[0] = (locals[0] + distance) % DIAL_LEN;
            }
        });

        return locals[1];
    }

    private static String readInput(String inputPath) throws Exception {
        String input = Files.readString(Path.of(inputPath));
        return input;
    }
}
