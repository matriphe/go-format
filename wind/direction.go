package wind

var directions = []string{
	"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
}

// Direction converts wind direction in degree to wind direction. Ref: https://stackoverflow.com/a/7490772/3460840.
func Direction(deg float32) string {
	// Divide the angle by 22.5 because 360deg/16 directions = 22.5deg/direction change.
	angle := deg / 22.5

	// Add .5 so that when you truncate the value you can break the 'tie' between the change threshold.
	angle += .5

	// Truncate the value using integer division (so there is no rounding).
	index := int(angle) % 16

	// Directly index into the array and return the value (mod 16).
	return directions[index]
}
