// package adapter 适配器，以手机充电为例，适配器将220v电压转为5v电压
package adapter

import "fmt"

// ####################################################### 1 源接口

// Adaptee 源接口
type Adaptee interface {
	// OutputPower 输出电压
	OutputPower()
}

// Volts220 源结构体，电压220v
type Volts220 struct{
}

// OutputPower 输出电压，实现了Adaptee接口
func (v Volts220) OutputPower() {
	fmt.Println("电源输出了220V电压")
}


// Volts100 源结构体，电压100v
type Volts100 struct{
}

// OutputPower 输出电压，实现了Adaptee接口
func (v Volts100) OutputPower() {
	fmt.Println("电源输出了100V电压")
}


// ####################################################### 2 目的接口



// Target 目的接口 --- 这个目的接口存在的意义貌似就是统一适配器的 - 输出接口
// 比如不管如何实现适配器，这个适配器都有一个固定的输出
// 其实就是指定了适配器的实现
type Target interface {
	// CovertTo5V 转换到5V电压
	CovertTo5V()
}


// ####################################################### 3 适配器


// Adapter 适配器结构体，嵌套Adaptee接口，实现Target接口
type Adapter struct  {
	Adaptee
}

// CovertTo5V 转换到5V电压
func (a Adapter) CovertTo5V() {
	//这里实现自己的转换逻辑，会调用源结构体的方法
	a.OutputPower()
	fmt.Println("通过手机电源适配器，转成了5V电压，可供手机充电")
}