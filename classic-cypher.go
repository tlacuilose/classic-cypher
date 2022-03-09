package main

import (
	"fmt"
	"log"

	"github.com/tlacuilose/classic-cypher/caesar"
	"github.com/tlacuilose/classic-cypher/cyphercmd"
)

func main() {

	// Parse command body
	cmd, err := cyphercmd.NewCmdFromBash()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Close files from cyphercmd
	defer cmd.Terminate()

	switch {
	case cmd.Action == cyphercmd.CaesarEncrypt:
		err := caesar.Encrypt(cmd.SourceFile, cmd.TargetFile)
		if err != nil {
			fmt.Println(err)
		}
	case cmd.Action == cyphercmd.CaesarDecrypt:
		err := caesar.Decrypt(cmd.SourceFile, cmd.TargetFile)
		if err != nil {
			fmt.Println(err)
		}
	case cmd.Action == cyphercmd.VigenereEncrypt:
		log.Printf("Encrypting %s with vigenere.", cmd.SourceFile.Name())
	case cmd.Action == cyphercmd.VigenereDecrypt:
		log.Printf("Decrypting %s with vigenere.", cmd.SourceFile.Name())
	}

}
