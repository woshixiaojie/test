/*
当数据类型不一样时，将会发生数据类型转换

自动类型转换（隐式）
    1、代码不需要进行特殊处理，自动完成
    2、规则：数据范围从小到大。与字节数不一定相关
*/
public class Demo01DataType {
    public static void main(String[] args){
        System.out.println(1024);//整数，默认int
        System.out.println(3.14);//浮点数，默认double

        //左边是long类型，右边是int类型，左右不一样
        //一个等号代表赋值，将右侧的int常量，交给左侧的long进行存储
        //int --> long，符合了数据范围从小到大的要求
        //这一行代码发生了自动类型的转换
        long num1 = 100;
        System.out.println(num1);

        //左边是double类型，右边是float类型，左右不一样
        //float --> double，符合从小到大的规则
        //也发生了自动类型的转换
        double num2 = 2.5F;
        System.out.println(num2);

        //左边是float类型，右边是long类型
        //long --> float,范围是float更大一些，符合从小到大的规则
        //也发生了自动转换
        float num3 = 30L;
        System.out.println(num3);//30.0

    }
}
