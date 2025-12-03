public class Day2Test {

    public static void main(String[] args) throws Exception {
        Day2Test.test("./test.txt", new Day2Result(1227775554l, 4174379265l));
        Day2Test.test(
            "./input.txt",
            new Day2Result(24157613387l, 33832678380l)
        );
    }

    private static void test(String filePath, Day2Result expected)
        throws Exception {
        Day2 day2 = new Day2(filePath);
        Day2Result actual = day2.solve();

        assert actual.equals(expected) : String.format(
            "Test failed for %s input. Expected part1: %d, part2: %d actual part1: %d, part2: %d",
            filePath,
            expected.part1(),
            expected.part2(),
            actual.part1(),
            actual.part2()
        );
    }
}
