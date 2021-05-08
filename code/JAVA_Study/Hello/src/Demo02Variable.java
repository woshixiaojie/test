/*
变量：程序运行期间，内容可以发生改变的量。

创建一个变量并使用的格式；

数据类型 变量名称；//创建了一个变量
变量名称 = 数据值 //赋值，将右边的数据值，赋值给左边的变量

一步到位格式：
    数据类型 变量名称 = 数据值 //在创建一个变量的同时，立刻放入指定的数据值

*/

public class Demo02Variable {
    public static void main(String[] args){
        //创建变量
        int num1;
        //给变量赋值
        num1 = 10;
        //打印出来显示的变量内容
        System.out.println(num1);

        //改变变量中的数字
        num1 = 24;
        System.out.println(num1);

        //使用一步到位的格式
        int num2 = 30;
        System.out.println(num2);

        num2 = 40;
        System.out.println(num2);

        byte num3 = 30;//注意：右侧数值范围不能超过左侧数据类型
        System.out.println(num3);

        //num3 = 400; 右侧超出了byte范围，错误

        short num4 = 50;
        System.out.println(num4);

        long num5 = 3000000L;
        System.out.println(num5);

        float num6 = 2.5F;
        System.out.println(num6);

        double num7 = 1.2;
        System.out.println(num7);

        char zifu1 = 'a';
        System.out.println(zifu1);

        zifu1 = '中';
        System.out.println(zifu1);

        boolean var1 = true;
        System.out.println(var1);

        var1 = false;
        System.out.println(var1);

        //将一个变量的数据内容，赋值给另外一个变量
        boolean var2 = var1;
        System.out.println(var2);

    }
}
