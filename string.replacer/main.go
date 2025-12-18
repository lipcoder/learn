package main

import (
	"fmt"
	"strings"
)

func main(){
	broken :="G# r#cks"
	replacer := strings.NewReplacer("#","c","s","f")
	fixed := replacer.Replace(broken)
	// replacer是值，Replace是方法
	fmt.Println(fixed)
}

// 面向对象里面
// 如果使用一个类似于函数的东西，当然这里面叫方法，是需要有一个类，这个类里面是有这个方法的定义
// 调用时先new一个类的实例，然后这个实例是有这个类的所有属性，比如变量和类里面的方法
// 现在就可以直接使用.来调用实例里面的方法

/*
class Person {
	方法1:输入1输出123
	方法2:输入2输出234
	实例变量/成员变量:name
}
p1 = Person.new() 创建了一个对象
name = p1.nmae
shuzi = p1.方法1(1)
*/

/*Ruby的写法
class Replacer
  def initialize(old, new)
    @old = old
    @new = new
  end

  def replace(str)
    str.gsub(@old, @new)
  end
end

replacer = Replacer.new("#", "c")
fixed = replacer.replace("G# r#cks")
*/

// 上面这是面向对象语言的说法，不代表go的实际，只不过写法是真的像