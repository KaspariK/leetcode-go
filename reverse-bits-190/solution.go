package reverse_bits_190

func reverseBits(num uint32) uint32 {
	reverse := uint32(0)
	shift := uint32(31)

	for num != 0 {
		// AND 1 gets last digit, arithmetic left shift places the digit in its correct spot
		reverse += (num & 1) << shift
		shift -= 1

		// Arithmetic right shift drops the least significant digit and copies the most significant digit
		num = num >> 1
	}

	return reverse
}
