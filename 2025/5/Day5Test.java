public class Day5Test {

    public static void main(String[] args) throws Exception {
        Day5Test.test1("./test.txt", 3);
        Day5Test.test1("./input.txt", 761);

        Day5Test.test2("./test.txt", 14);
        Day5Test.test2("./input.txt", 345755049374932L);

        Day5Test.test2("./input-roma.txt", 344423158480189L);
    }

    private static void test1(String filePath, long expected) throws Exception {
        Day5 day5 = new Day5(filePath);
        int actual = day5.part1();

        assert actual == expected : String.format(
            "Test failed for %s input. Expected %d, actual %d",
            filePath,
            expected,
            actual
        );
    }

    private static void test2(String filePath, long expected) throws Exception {
        Day5 day5 = new Day5(filePath);
        long actual = day5.part2();

        assert actual == expected : String.format(
            "Test failed for %s input. Expected %d, actual %d",
            filePath,
            expected,
            actual
        );
    }
}
