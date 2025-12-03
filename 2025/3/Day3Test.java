public class Day3Test {

    public static void main(String[] args) throws Exception {
        Day3Test.test1("./test.txt", 357L);
        Day3Test.test1("./input.txt", 17301L);

        Day3Test.test2("./test.txt", 3121910778619L);
        Day3Test.test2("./input.txt", 172162399742349L);
    }

    private static void test1(String filePath, long expected) throws Exception {
        Day3 day3 = new Day3(filePath);
        long actual = day3.part1();

        assert actual == expected : String.format(
            "Test failed for %s input. Expected %d, actual %d",
            filePath,
            expected,
            actual
        );
    }

    private static void test2(String filePath, long expected) throws Exception {
        Day3 day3 = new Day3(filePath);
        long actual = day3.part2();

        assert actual == expected : String.format(
            "Test failed for %s input. Expected %d, actual %d",
            filePath,
            expected,
            actual
        );
    }
}
