package filesize

// Byte is defined as uint64, which is equal 8 bit
type Byte uint64

const (
	// B is Byte
	B Byte = 1
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

	maxUint64 uint64 = (1 << 64) - 1
)

// Convert converts byte to other unit
func (b Byte) Convert(unit Byte) float64 {
	return float64(b) / float64(unit)
}
