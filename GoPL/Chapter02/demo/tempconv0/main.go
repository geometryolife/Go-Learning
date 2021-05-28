// 为了说明类型声明，我们把不同的温度标度转换为不同类型:
// 包 tempconv 进行摄氏温度和华氏温度的转换计算
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF、FToC用来在两种温度计量单位之间转换，返回不同的数值
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// 声明 String() 方法，关联到 Celsius 类型，返回c变量的数字值，
// 后面跟着摄氏温度的符号°C
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// 1. 这个包定义了两个类型--Celsius(摄氏温度)和Fahrenheit(华氏温度)，
// 它们分别对应两种温度计量单位。即使使用相同的底层类型float64，它们
// 也不是相同的类型，所以它们不能使用算术表达式进行比较和合并。
// 2. 区分这些类型可以防止无意间合并不同计量单位的温度值；从float64
// 转换为Celsius(t)或Fahrenheit(t)需要显式类型转换。
// 3. Celsius(t)和Fahrenheit(t)是类型转换，而不是函数调用。它们不会
// 改变值和表达式，但改变了显式意义。
