package adapter

import "testing"

func Test(t *testing.T) {
	t.Run("adaptor: ", ConvertVolts)
}

func ConvertVolts(t *testing.T) {
	var adaptee Adaptee
	adaptee = &Volts220{}

	var target Target
	target = &Adapter{adaptee}
	target.CovertTo5V()


	// 100vēééåØ
	adaptee = &Volts100{}
	target = &Adapter{adaptee}
	target.CovertTo5V()
}