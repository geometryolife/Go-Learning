package tempconv

// CToF 把摄氏温度转换为华氏温度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC 把华氏温度转换为摄氏温度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC 把开尔文温度转换为摄氏温度
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CToK 把摄氏温度转换为开尔文温度
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToF 把卡尔文温度转换为华氏温度
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

// FToK 把华氏温度转换为开尔文温度
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9) + 273.15 }
