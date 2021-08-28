package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"parking/model"
)

func readInstructionFile(FileName string) ([]model.Instruction, error) {
	file, err := os.Open(FileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open the file %s\n", FileName))
	}
	defer file.Close()

	var ins [][]byte
	var insParam []string
	insSet := make([]model.Instruction, 0)

	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	for err == nil {
		ins = bytes.Split(line, []byte(" "))
		for _, v := range ins {
			insParam = append(insParam, string(v))

		}
		i := model.Instruction{
			Command:   insParam[0],
			Arguments: insParam[1:],
		}
		insSet = append(insSet, i)
		insParam = nil
		line, _, err = reader.ReadLine()
	}
	return insSet, nil
}
