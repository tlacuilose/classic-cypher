# CMD Cipher

Command to encrypt and decrupt .txt files using classic methods: [caesar](https://mathworld.wolfram.com/CaesarsMethod.html) and [vigenere](https://pages.mtu.edu/~shene/NSF-4/Tutorial/VIG/Vig-Base.html)

## Usage

Build command with:

```bash
go build classic-cypher.go
```

Encrypt a .txt file using caesar

```bash
./classic-cypher -encrypt -caesar [file-to-encrypt.txt] [target-for-ciphertext.txt]
```

Decrypt a .txt file using caesar

```bash
./classic-cypher -decrypt -caesar [ciphertext-to-decrypt.txt] [target-for-plaintext.txt]
```

## Testing

Test all unit and integration tests

```bash
go test ./...
```
