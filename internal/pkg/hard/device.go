package hard

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

func AuthUkeyInfo(serialNumber, pubkey, signData []byte) error {
	sm2pub, err := X509Sm2PublicKeyFromBytes(pubkey)
	if err != nil {
		log.Printf("AuthUkeyInfo: invalid sm2 public key data\n")
		return err
	}
	if len(signData) != 64 {
		log.Printf("AuthUkeyInfo: invalid sm2 signature length: %d\n", len(signData))
		return fmt.Errorf("AuthUkeyInfo: invalid sm2 signature length: %d", len(signData))
	}
	valid := sm2.Sm2Verify(sm2pub, serialNumber, []byte("1234567812345678"), new(big.Int).SetBytes(signData[:32]), new(big.Int).SetBytes(signData[32:]))
	if !valid {
		return errors.New("invalid sm2 signature")
	}
	return nil
}

func X509Sm2PublicKeyFromBytes(pubKeyData []byte) (*sm2.PublicKey, error) {
	if len(pubKeyData) != 64 && !(len(pubKeyData) == 65 && pubKeyData[0] == 0x04) {
		return nil, fmt.Errorf("invalid sm2 public key: bytes length error: %d", len(pubKeyData))
	}
	pub := &sm2.PublicKey{Curve: sm2.P256Sm2()}
	offset := len(pubKeyData) - 64
	X := new(big.Int).SetBytes(pubKeyData[offset:(offset + 32)])
	Y := new(big.Int).SetBytes(pubKeyData[(offset + 32):])
	if !sm2.P256Sm2().IsOnCurve(X, Y) {
		return nil, fmt.Errorf("invalid sm2 public key: not on curve")
	}
	pub.X, pub.Y = X, Y
	return pub, nil
}
