/*
使用变量注意事项
    1、如果创建多个变量，那么变量之间的名称不可重复
*/
public class Demo03VariableNotice {
    public static void main(String[] args){
        //1、如果创建多个变量，那么变量之间的名称不可重复
        int num1 = 10;
        //int num1 = 20;

        //对于float和long类型，后缀F和L不可丢掉
        float num2 = 2.3F;
        long num3 = 200000000000L;

        //没有进行赋值的变量，不能直接使用；一定要赋值之后使用
        int num4;
        //System.out.println(num4);


        //变量使用不能超出作用域的范围
            //作用域：从定义变量的一行开始，一直到所属的大括号结束为止。
        //System.out.println(num5);在创建变量之前，不能使用这个变量
        int num5 = 500;
        System.out.println(num5);
        {
            int num6 = 60;
            System.out.println(num6);
        }
        //int num6 = 40;可以重新定义
        //System.out.println(num6);已经超出大括号，超出了作用域，变量不能再使用了

        //可以通过一个语句创建多个变量，但是不推荐这样写。
        int a = 10,b = 20,c = 30;
        System.out.println(a);
        System.out.println(b);
        System.out.println(c);


    }
}
