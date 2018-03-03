package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts Kelvin temperature to Celsius.
func KToC(k Kelvin) Kelvin { return Kelvin(k - Kelvin(AbsoluteZeroC)) }
