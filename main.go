package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	// Min pulse length out of 4096
	servoMin = 150
	// Max pulse length out of 4096
	servoMax = 700
	// Limiting the max this servo can rotate (in deg)
	maxDegree = 180
	// Number of degrees to increase per call
	degIncrease = 10
	yawDeg      = 90
)

func main() {

	adaptor := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(adaptor, "12")

	work := func() {
		/*
			gobot.Every(1*time.Second, func() {
				servo.Move(10)
			})
		*/
		servo.Move(10)
	}

	robot := gobot.NewRobot("superBot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()

}
