import java.nio.file.Files;
import java.nio.file.Path;

public class Day4 {

    private final String inputPath;

    public Day4(String inputPath) {
        this.inputPath = inputPath;
    }

    public int part1() throws Exception {
        String input = this.readInput(inputPath);

        char[][] grid = input
            .lines()
            .map(String::toCharArray)
            .toArray(char[][]::new);
        int m = grid.length;
        int n = grid[0].length;

        int rollsCounter = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '@' && isRollAccessible(i, j, m, n, grid)) {
                    rollsCounter++;
                }
            }
        }

        return rollsCounter;
    }

    public int part2() throws Exception {
        String input = this.readInput(inputPath);

        char[][] grid = input
            .lines()
            .map(String::toCharArray)
            .toArray(char[][]::new);
        int m = grid.length;
        int n = grid[0].length;

        int localRollsCounter;
        int rollsCounter = 0;
        do {
            localRollsCounter = 0;

            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    if (
                        grid[i][j] == '@' && isRollAccessible(i, j, m, n, grid)
                    ) {
                        localRollsCounter++;
                        grid[i][j] = '.'; // remove roll
                    }
                }
            }

            rollsCounter += localRollsCounter;
        } while (localRollsCounter > 0);

        return rollsCounter;
    }

    private boolean isRollAccessible(
        int i,
        int j,
        int m,
        int n,
        char[][] grid
    ) {
        int[][] adjacentCells = {
            { i - 1, j - 1 },
            { i - 1, j },
            { i - 1, j + 1 },
            { i, j - 1 },
            { i, j + 1 },
            { i + 1, j - 1 },
            { i + 1, j },
            { i + 1, j + 1 },
        };

        int adjacentRolls = 0;
        for (int[] adj : adjacentCells) {
            int ai = adj[0];
            int aj = adj[1];

            if (inRange(ai, aj, m, n) && grid[ai][aj] == '@') {
                adjacentRolls += 1;
            }
        }

        return adjacentRolls < 4;
    }

    private boolean inRange(int i, int j, int m, int n) {
        return 0 <= i && i < m && 0 <= j && j < n;
    }

    private String readInput(String inputPath) throws Exception {
        String input = Files.readString(Path.of(inputPath));
        return input;
    }
}
