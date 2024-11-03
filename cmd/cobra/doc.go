// Command gogen provides cryptographic key generation and password hashing functionality.
//
// Usage:
//
//	# Generate a 32-byte key
//	gogen key
//
//	# Generate a 64-byte key
//	gogen key -l 64
//
//	# Hash a password with default cost (12)
//	gogen hash mypassword
//
//	# Hash a password with custom cost
//	gogen hash -c 14 mypassword
//
//	# Run password hashing benchmark
//	gogen hash -b mypassword
package main
