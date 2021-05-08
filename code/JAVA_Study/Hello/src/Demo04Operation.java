public class Demo04Operation {
    public static void main(String[] args){
        int x = 10,y = 3;
        int result = x / y;
        // "/",取整
        System.out.println(result);

        // "%"，取余/取模
        int result2 = x % y;
        System.out.println(result2);

        //注意事项
        //  一旦运算中有不同类型的数据，那么结果将会是数据范围大的那种
        //int + double
        double result3 = x + 2.5;
        System.out.println(result3);
    }
}
