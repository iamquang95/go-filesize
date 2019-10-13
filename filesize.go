package filesize

import "errors"

// Unit is defined storage unit
type Unit uint64

const (
	// B is Byte
	B Unit = 1
	// KB is Kilobyte
	KB = B << 10
	// MB is Megabyte
	MB = KB << 10
	// GB is GigaByte
	GB = MB << 10
	// TB is Terabyte
	TB = GB << 10
	// PB is PetaByte
	PB = TB << 10
	// EB is Extrabyte
	EB = PB << 10
)

func (u Unit) isValid() bool {
	switch u {
	case B, KB, MB, GB, TB, PB, EB:
		return true
	}
	return false
}

// Byte is defined as uint64, which is equal 8 bit
type Byte uint64

// Convert converts byte to other unit, it will return err if input unit is invalid
func (b Byte) Convert(unit Unit) (float64, error) {
	if !unit.isValid() {
		return 0, errors.New("Invalid unit")
	}
	return float64(b) / float64(unit), nil
}
