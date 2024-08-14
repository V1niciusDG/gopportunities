package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createExampleEnv() error {
	inputFilePath := ".env"
	outputFilePath := "example.env"

	// Abrir o arquivo .env para leitura
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// Criar e/ou limpar o arquivo example.env
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// Processar cada linha do arquivo .env
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("failed to write to output file: %v", err)
			}
		} else {
			// Dividir a linha em chave e valor
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]
			_, err := outputFile.WriteString(key + "=\n")
			if err != nil {
				return fmt.Errorf("failed to write to output file: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	return nil
}