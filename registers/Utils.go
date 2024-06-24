package registers

import (
	"fmt"
	"math"
)

// appendFloat32 appends the 32-bit floating point value to the given registers.
func appendFloat32(registers *[]uint16, value float32) {
	fmt.Printf("Value = %f\n", value)
	bits := math.Float32bits(value)
	*registers = append(*registers, uint16(bits>>16), uint16(bits&0xffff))
}
