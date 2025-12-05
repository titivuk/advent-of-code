import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

class Range {

    public long start;
    public long end;

    public Range(long start, long end) {
        this.start = start;
        this.end = end;
    }

    @Override
    public String toString() {
        return String.format("[%d, %d]", start, end);
    }
}

public class Day5 {

    private final String inputPath;

    public Day5(String inputPath) {
        this.inputPath = inputPath;
    }

    public int part1() throws Exception {
        String input = this.readInput(inputPath);
        String[] lines = input.lines().toArray(String[]::new);

        List<Range> freshIngredients = new ArrayList<>();
        int i = 0;
        while (i < lines.length && lines[i] != "") {
            String[] parts = lines[i].split("-");
            Range r = new Range(
                Long.parseLong(parts[0]),
                Long.parseLong(parts[1])
            );

            freshIngredients.add(r);

            i++;
        }
        // sort interval by start, so we can use binary search
        freshIngredients.sort(
            Comparator.comparing((Range r) -> r.start).thenComparingLong(
                (Range r) -> r.end
            )
        );
        freshIngredients = mergeRanges(freshIngredients);

        // extract availableIngredients
        List<Long> availableIngredients = new ArrayList<>();
        // advance pointer
        i++;
        while (i < lines.length) {
            availableIngredients.add(Long.parseLong(lines[i]));
            i++;
        }

        int freshCounter = 0;
        for (long ai : availableIngredients) {
            int l = 0,
                r = freshIngredients.size() - 1;

            while (l <= r) {
                // possible overflow, but should be fine
                int mid = (l + r) / 2;

                if (
                    freshIngredients.get(mid).start <= ai &&
                    ai <= freshIngredients.get(mid).end
                ) {
                    freshCounter++;
                    break;
                }

                if (ai < freshIngredients.get(mid).start) {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            }
        }

        return freshCounter;
    }

    public long part2() throws Exception {
        String input = this.readInput(inputPath);

        List<Range> freshIngredients = input
            .lines()
            .takeWhile(l -> !l.isEmpty())
            .map(l -> {
                String[] parts = l.split("-");
                return new Range(
                    Long.parseLong(parts[0]),
                    Long.parseLong(parts[1])
                );
            })
            .sorted(
                Comparator.comparingLong((Range r) ->
                    r.start
                ).thenComparingLong((Range r) -> r.end)
            )
            .collect(Collectors.toList());
        freshIngredients = mergeRanges(freshIngredients);

        long sum = 0;
        for (Range r : freshIngredients) {
            sum += (r.end - r.start + 1);
        }

        return sum;
    }

    private ArrayList<Range> mergeRanges(List<Range> list) {
        ArrayList<Range> merged = new ArrayList<>();

        int i = 0;
        while (i < list.size()) {
            Range accum = new Range(list.get(i).start, list.get(i).end);

            int j = i + 1;
            while (j < list.size() && list.get(j).start <= accum.end) {
                accum.end = Math.max(list.get(j).end, accum.end);
                j++;
            }

            merged.add(accum);
            i = j;
        }

        return merged;
    }

    private String readInput(String inputPath) throws Exception {
        String input = Files.readString(Path.of(inputPath));
        return input;
    }
}
