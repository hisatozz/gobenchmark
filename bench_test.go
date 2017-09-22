package gobenchmark

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"testing"
)

func BenchmarkBinhexEncode(b *testing.B) {
	dst := make([]byte, hex.EncodedLen(len(bin1024)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hex.Encode(dst, bin1024)
	}
}

func BenchmarkBinhexDecode(b *testing.B) {
	src := []byte(hex1024)
	dst := make([]byte, hex.DecodedLen(len(src)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hex.Decode(dst, src)
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	src := []byte(str1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString(src)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.DecodeString(strbase64)
	}
}

func BenchmarkHMAC(b *testing.B) {
	src := []byte(str1024)
	key := []byte(key256bit)

	mac := hmac.New(sha256.New, key)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mac.Reset()
		mac.Write(src)
		expectedMAC := mac.Sum(nil)
		// log.Print(expectedMAC)
		_ = expectedMAC
	}
}

func BenchmarkAES256Enc(b *testing.B) {
	src := []byte(str1024)
	key := []byte(key256bit)

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nonce := make([]byte, 12)
		io.ReadFull(rand.Reader, nonce)
		ciphertext := aesgcm.Seal(nil, nonce, src, nil)
		_ = ciphertext
		// log.Printf("%x", nonce)
		// log.Printf("%x", ciphertext)
	}
}

func BenchmarkAES256Dec(b *testing.B) {
	src, _ := hex.DecodeString(ciphertextAES256)
	key := []byte(key256bit)
	nonce, _ := hex.DecodeString(nonceGCM)

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		plaintext, _ := aesgcm.Open(nil, nonce, src, nil)
		// log.Printf("%s", plaintext)
		_ = plaintext
	}
}

func nop() {
	log.Print("never called")
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	var animals []Animal

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(jsonBlob, &animals)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		_ = err
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	animal := Animal{
		Name:  "Quoll",
		Order: "Dasyuromorphia",
	}
	var animals [22]Animal
	for i := 0; i < len(animals); i++ {
		animals[i] = animal
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(animals)
	}
}
