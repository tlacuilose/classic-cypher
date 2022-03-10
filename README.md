# CMD Cipher

Command to encrypt and decrupt .txt files using classic methods: [caesar](https://mathworld.wolfram.com/CaesarsMethod.html) and [vigenere](https://pages.mtu.edu/~shene/NSF-4/Tutorial/VIG/Vig-Base.html)

## Usage

### Building and running

Build command with:

```bash
go build classic-cypher.go
```

Alternatively you can run the following commands with:

```bash
go run classic-cypher.go ...
```

instead of building and running:

```bash
./classic-cypher ...
```

### Command options

Encrypt a .txt file using caesar

```bash
./classic-cypher -encrypt -caesar [file-to-encrypt.txt] [target-for-ciphertext.txt]
```

Decrypt a .txt file using caesar

```bash
./classic-cypher -decrypt -caesar [ciphertext-to-decrypt.txt] [target-for-plaintext.txt]
```

Encrypt a .txt file using vigenere

```bash
./classic-cypher -encrypt -vigenere [file-to-encrypt.txt] [target-for-ciphertext.txt]
```

Decrypt a .txt file using vigenere

```bash
./classic-cypher -decrypt -vigenere [ciphertext-to-decrypt.txt] [target-for-plaintext.txt]
```

**REMEMBER TO USE THE SAME KEY**

## Testing

Test all unit and integration tests.

```bash
go test ./...
```

*Plaintext files for testing are in ./texts/*
