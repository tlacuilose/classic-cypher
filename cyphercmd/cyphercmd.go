package cyphercmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CmdBody struct {
	action     ActionType
	sourceFile string
	targetFile string
}

type Command struct {
	Action     ActionType
	SourceFile *os.File
	TargetFile *os.File
}

func (c *Command) Terminate() {
	c.SourceFile.Close()
	c.TargetFile.Close()
}

type ActionType int64

const (
	CaesarEncrypt ActionType = iota
	CaesarDecrypt
	VigenereEncrypt
	VigenereDecrypt
)

func parseFlagsAndArgs() (CmdBody, error) {

	// Parse cmd flags.
	encryptFlag := flag.Bool("encrypt", false, "Encrypt from txt file.")
	decryptFlag := flag.Bool("decrypt", false, "Decrypt from txt file.")
	caesarFlag := flag.Bool("caesar", false, "Use Caesar method.")
	vigenereFlag := flag.Bool("vigenere", false, "Use Vigenere method.")
	flag.Parse()

	// Parse cmd args
	args := flag.Args()

	// Expect two filename, source, target
	if len(args) != 2 {
		return CmdBody{}, errors.New("Expected two filename, $ classic-cypher [-caesar or -vigenere] [-encrypt or -decrypt] [source-filename.txt] [target-filename.txt")
	}

	// Expect a .txt file for source and target files.
	fileNames := args[0:1]
	for _, f := range fileNames {
		if !strings.HasSuffix(f, ".txt") {
			return CmdBody{}, errors.New("Expected a .txt file for both the source and the target filename.")

		}
	}

	// Get the cyphercmd correct action type
	isCaesar := *caesarFlag && !*vigenereFlag
	isVigenere := *vigenereFlag && !*caesarFlag
	isEncrypt := *encryptFlag && !*decryptFlag
	isDecrypt := *decryptFlag && !*encryptFlag

	// Return action type, args and error
	switch {
	case (isCaesar && isEncrypt):
		return CmdBody{
			action:     CaesarEncrypt,
			sourceFile: args[0],
			targetFile: args[1],
		}, nil
	case (isCaesar && isDecrypt):
		return CmdBody{
			action:     CaesarDecrypt,
			sourceFile: args[0],
			targetFile: args[1],
		}, nil
	case (isVigenere && isEncrypt):
		return CmdBody{
			action:     VigenereEncrypt,
			sourceFile: args[0],
			targetFile: args[1],
		}, nil
	case (isVigenere && isDecrypt):
		return CmdBody{
			action:     VigenereDecrypt,
			sourceFile: args[0],
			targetFile: args[1],
		}, nil
	}

	return CmdBody{}, errors.New("Invalid command body. $ classic-cypher [-caesar or -vigenere] [-encrypt or -decrypt] [source-filename.txt] [target-filename.txt] See -h for help")
}

func openSourceFile(filename string) (*os.File, error) {

	// Expect a valid source file
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error opening source file %s", filename))
	}

	return file, nil
}

func createTargetFile(filename string) (*os.File, error) {

	// Expect a valid target file
	file, err := os.Create(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error opening file %s", filename))
	}

	return file, nil
}

func NewCmdFromBash() (Command, error) {

	cmdBody, err := parseFlagsAndArgs()

	// Open source file.
	sourceFile, err := openSourceFile(cmdBody.sourceFile)
	if err != nil {
		return Command{}, err
	}

	// Open target file.
	targetFile, err := createTargetFile(cmdBody.targetFile)
	if err != nil {
		return Command{}, err

	}

	// Return new command.
	return Command{
		Action:     cmdBody.action,
		SourceFile: sourceFile,
		TargetFile: targetFile,
	}, nil

}
