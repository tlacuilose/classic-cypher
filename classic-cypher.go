package main

import (
	"log"

	"github.com/tlacuilose/classic-cypher/caesar"
	"github.com/tlacuilose/classic-cypher/cyphercmd"
	"github.com/tlacuilose/classic-cypher/vigenere"
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
			log.Println(err)
		}
	case cmd.Action == cyphercmd.CaesarDecrypt:
		err := caesar.Decrypt(cmd.SourceFile, cmd.TargetFile)
		if err != nil {
			log.Println(err)
		}
	case cmd.Action == cyphercmd.VigenereEncrypt:
		err := vigenere.Encrypt(cmd.SourceFile, cmd.TargetFile)
		if err != nil {
			log.Println(err)
		}
	case cmd.Action == cyphercmd.VigenereDecrypt:
		err := vigenere.Decrypt(cmd.SourceFile, cmd.TargetFile)
		if err != nil {
			log.Println(err)
		}
	}

}
