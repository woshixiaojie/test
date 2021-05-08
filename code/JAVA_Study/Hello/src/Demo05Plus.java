/*
四则运算中"+"常用三种方法
    1、数值 - 加法
    2、char - char提升成int - 然后计算
        char int - 对照表 - ASCII Unicode
    3、字符串 String - 首字母大写，并不是关键字 - 表示连接
*/

public class Demo05Plus {
    public static void main(String[] args){
        //字符串 - 基本使用
        //数据类型 变量名 = 数据值
        String str1 = "Hello";
        System.out.println(str1);

        System.out.println("Hello" + "World");//HelloWorld

        //任何数据类型和字符串连接的时候，结果都会变成字符串
        String str2 = "java";
        // String + int --> String
        System.out.println(str2 + 20);//java20

        //优先级问题
        // String + int + int
        // String       + int
        // String
        System.out.println(str2 + 20 + 30);//java2030
        //小括号中的优先级是最先的
        System.out.println(str2 + (20 + 30));//java50

    }
}
