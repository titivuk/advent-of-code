class Day1Test {
    public static void main(String[] args) throws Exception {
        testSolve1("./test.txt", 3);
        testSolve1("./input.txt", 1177);

        testSolve2("./test.txt", 6);
        testSolve2("./input.txt", 6768);
    }

    private static void testSolve1(String filePath, int expected) throws Exception {
            int actual = Day1.solve1(filePath);
            assert actual == expected : String.format("Test failed for %s input. Expected %d, actual - %d", filePath, expected, actual);
    }

    private static void testSolve2(String filePath, int expected) throws Exception {
        int actual = Day1.solve2(filePath);
        assert actual == expected : String.format("Test failed for %s input. Expected %d, actual - %d", filePath, expected, actual);
    }
}