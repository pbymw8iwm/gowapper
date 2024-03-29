package utilswapper

import (
	"crypto/cipher"
	"fmt"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
	"golang.org/x/crypto/blowfish"
)

func BlowfishEcbEncrypt(pt, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt) // pad last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

func BlowfishEcbDecrypt(ct, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		panic(err.Error())
	}
	return pt
}

func BlowfishCbcEncrypt(pt, key, iv []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt) // pad last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

//输入密文，得到明文
func BlowfishCbcDecrypt(ct, key, iv []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		panic(err.Error())
	}
	return pt
}
func Blowfish_example() {
	pt := []byte("xft2QfpeNdN3CjCYSXpmB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq26rVja0aRvBgd3rXoDC3amB3etegMLdqYHd616Awt2o5H68TOzR3fGB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqcQ3La_nIgHBgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3apdFaG0AC2NUYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pLqtKE0owDn2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqwwaD1j0RM8Jgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3alLt08Si3iL2YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rjzjlRKaNxI2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdq78JzaibYFkQhvh3PRy6MT2B3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdql0VobQALY1Rgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3akuq0oTSjAOfYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rDBoPWPREzwmB3etegMLdqYHd616Awt2pgd3rXoDC3amB3etegMLdqUu3TxKLeIvZgd3rXoDC3amB3etegMLdqYHd616Awt2pgd3rXoDC3auPOOVEpo3EjYHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2rbqtWNrRpG8GB3etegMLdqYHd616Awt2pgd3rXoDC3ajkfrxM7NHd8YHd616Awt2pgd3rXoDC3amB3etegMLdqYHd616Awt2pxDctr")
	key := []byte("#^0^$no.1!@@dh")

	ct := BlowfishEcbEncrypt(pt, key)
	fmt.Printf("Ciphertext: %x\n", ct)

	recoveredPt := BlowfishEcbDecrypt(ct, key)
	fmt.Printf("Recovered plaintext: %s\n", recoveredPt)
}
