/*
数字与字符的对照关系表
    ASCII码表：美国信息交换标准代码
    Unicode：万国表。开头0-127与ASCII完全一样，但是128开始包含更多字符

    48 - '0'
    65 - 'A'
    97 - 'a'
*/

public class Demo03DataType {
    public static void main(String[] args){
        char zifu1 = '1';
        System.out.println(zifu1);

        char zifu3 = 'c';

        /*
        左侧是int，右侧是char
        char --> int，从小到大
        发生了自动转换
        */
        int num = zifu3;
        System.out.println(num);

        char zifu4 = '中';
        System.out.println(zifu4 + 0);
    }
}
