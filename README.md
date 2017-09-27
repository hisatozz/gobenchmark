# Encode/Decode benchmark for GO
From go language standard library, select the function frequently used around WEB API, wrote a benchmark test.

## The target functions
 - binhex
   - "encoding/hex"
 - base64
    - "encoding/base64"
 - HMAC-SHA256
	 - "crypto/sha256"
	 - "crypto/hmac"
 - AES256-GCM
	 - "crypto/cipher"
	 - "crypto/aes"
 - JSON marshal/unmarshal
	 - "encoding/json"


## Results

### AWS t2 micro

 ```
BenchmarkBinhexEncode            2000000               878 ns/op
BenchmarkBinhexDecode             500000              2475 ns/op
BenchmarkBase64Encode            1000000              2156 ns/op
BenchmarkBase64Decode             200000              6735 ns/op
BenchmarkHMAC_SHA256              300000              4760 ns/op
BenchmarkAES256_GCM_Enc          1000000              1898 ns/op
BenchmarkAES256_GCM_Dec          2000000               951 ns/op
BenchmarkJSONUnmarshal            200000              6269 ns/op
BenchmarkJSONMarshal              100000             13694 ns/op
```

### Intel Core i3 Clarkdale

CPU without AES acceleration.

```
BenchmarkBinhexEncode-4          2000000               858 ns/op
BenchmarkBinhexDecode-4           500000              3030 ns/op
BenchmarkBase64Encode-4          1000000              2152 ns/op
BenchmarkBase64Decode-4           200000              8299 ns/op
BenchmarkHMAC_SHA256-4            200000              6697 ns/op
BenchmarkAES256_GCM_Enc-4         100000             17681 ns/op
BenchmarkAES256_GCM_Dec-4         100000             17350 ns/op
BenchmarkJSONUnmarshal-4          200000              7675 ns/op
BenchmarkJSONMarshal-4            100000             13769 ns/op
```

