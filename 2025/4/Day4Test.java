public class Day4Test {

    public static void main(String[] args) throws Exception {
        Day4Test.test1("./test.txt", 13);
        Day4Test.test1("./input.txt", 1419);

        Day4Test.test2("./test.txt", 43);
        Day4Test.test2("./input.txt", 8739);
    }

    private static void test1(String filePath, long expected) throws Exception {
        Day4 day4 = new Day4(filePath);
        int actual = day4.part1();

        assert actual == expected : String.format(
            "Test failed for %s input. Expected %d, actual %d",
            filePath,
            expected,
            actual
        );
    }

   private static void test2(String filePath, long expected) throws Exception {
       Day4 day4 = new Day4(filePath);
       long actual = day4.part2();

       assert actual == expected : String.format(
           "Test failed for %s input. Expected %d, actual %d",
           filePath,
           expected,
           actual
       );
   }
}
