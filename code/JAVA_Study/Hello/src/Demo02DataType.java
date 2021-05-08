/*
强制类型转换
    1、特点：需要进行特殊处理
    2、格式：范围小的类型 范围小的变量名 = （范围小的类型）原本范围大的类型

注意事项
    1、强制类型转换一般不推荐使用，因为可能会发生精度丢失、数据溢出
    2、byte/short/char这三种类型都可以发生数学运算，例如"+"
    3、byte/short/char这三种类型在运算的时候，都会首先提升成为int类型，然后在计算
    4、boolean布尔类型不能发生数据类型转换
*/
public class Demo02DataType {
    public static void main(String[] args){
        //左边是int，右边是long
        //long --> int，不是从小到大
        //不能发生自动转换
        //格式：范围小的类型 范围小的变量名 = （范围小的类型）原本范围大的类型
        int num = (int) 100L;
        System.out.println(num);

        //long强制转化为int
        int num2 = (int) 6000000000L;//int最大表示21亿
        System.out.println(num2);//1705032704

        //double --> int，强制转换
        int num3 = (int) 3.99;
        System.out.println(num3);//不是四舍五入，小数都会舍弃掉

        char zifu1 = 'A';
        System.out.println(zifu1 + 1);//66，也就是大写字母A被当作65进行处理
        //计算机底层会用一个数字（二进制）来表示字符A，就是65
        //一旦char类型进行了数学运算，那么字符就会按照一定规则翻译成一个数字


        byte num4 = 40;//注意！右侧的数值大小不能超过左侧的类型范围
        byte num5 = 50;
        //byte + byte --> int + int --> int
        int result = num4 + num5;
        System.out.println(result);

        short num6 = 60;
        //byte + short --> int + int --> int
        //int强制转化为short，必须保证逻辑上的真实大小否则会发生数据溢出
        short result2 = (short) (num4 + num6);
        System.out.println(result2);

    }
}
