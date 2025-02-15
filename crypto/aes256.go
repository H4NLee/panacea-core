package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

// DeriveSharedKey derives a shared key (which can be used for asymmetric encryption)
// using a specified KDF (Key Derivation Function)
// from a shared secret generated by Diffie-Hellman key exchange (ECDH).
func DeriveSharedKey(priv *btcec.PrivateKey, pub *btcec.PublicKey, kdf func([]byte) []byte) []byte {
	sharedSecret := btcec.GenerateSharedSecret(priv, pub)
	return kdf(sharedSecret)
}

// KDFSHA256 is a key derivation function which uses SHA256.
func KDFSHA256(in []byte) []byte {
	out := sha256.Sum256(in)
	return out[:]
}

// Encrypt encrypts data using a AES256 cryptography.
func Encrypt(secretKey, nonce, data []byte) ([]byte, error) {
	if len(secretKey) != 32 {
		return nil, fmt.Errorf("secret key is not for AES-256: total %d bits", 8*len(secretKey))
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(nonce) != aesgcm.NonceSize() {
		return nil, fmt.Errorf("nonce length must be %v", aesgcm.NonceSize())
	}

	cipherText := aesgcm.Seal(nil, nonce, data, nil)

	return cipherText, nil
}

// Decrypt decrypts data using a AES256 cryptography.
func Decrypt(secretKey, nonce, ciphertext []byte) ([]byte, error) {
	if len(secretKey) != 32 {
		return nil, fmt.Errorf("secret key is not for AES-256: total %d bits", 8*len(secretKey))
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
