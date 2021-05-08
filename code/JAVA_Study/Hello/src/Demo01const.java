public class Demo01const {
    public static void main(String[] args){
        //字符串常量
        System.out.println("asd");
        System.out.println("");

        //整数常量
        System.out.println(100);
        System.out.println(0);
        System.out.println(-100);

        //浮点数常量
        System.out.println(3.14);
        System.out.println(-2.3);
        System.out.println(0.00);

        //字符常量
        System.out.println('a');
        System.out.println('6');
      /*  两个单引号有切仅有一个字符
        System.out.println('');
        System.out.println('11');*/

        //布尔常量
        System.out.println(true);
        System.out.println(false);

        //空常量不能直接打印输出
        //System.out.println(null);
        /*
        基本数据类型
            整数形 byte short int long
            浮点型 float double
            字符型 char
            布尔型 boolean

        引用数据类型
            字符串、数组、类、接口lamdba

        注意事项
            1、字符串不是基本类型，而是引用类型
            2、浮点型可能只是一个近视值，并非精确值
            3、数据范围与字节数不一定相关，例如float数据范围比long更加广泛，但是float是4字节，long是8字节
            4、浮点数当中默认类型是double。如果一定要使用float类型，需要加上一个后缀F。
                如果是整数，默认是int类型，如果一定要使用long类型，需要加上一个后缀L。都是推荐使用大写
                System.out.println(100L);
        */
    }
}
