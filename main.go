package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const OutputDirName = "output"

func main() {
	// カレントディレクトリを取得
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// カレントディレクトリ内の全ファイルを取得
	files, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	// .m2tsファイルのみをリストする
	var m2tsFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".m2ts") {
			m2tsFiles = append(m2tsFiles, file.Name())
		}
	}

	// .m2tsファイルが見つかったか確認
	if len(m2tsFiles) == 0 {
		fmt.Println("No .m2ts files found in the current directory.")
		return
	} else {
		fmt.Println(".m2ts files found:")
		for _, file := range m2tsFiles {
			fmt.Println(filepath.Join(currentDir, file))
		}
	}

	// outputディレクトリが存在するかチェックし、存在しない場合は作成する
	outputDirPath := filepath.Join(currentDir, OutputDirName)
	if _, err := os.Stat(outputDirPath); os.IsNotExist(err) {
		fmt.Printf("Creating output directory: %s\n", outputDirPath)
		err = os.Mkdir(outputDirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	// 各.m2tsファイルに対してffmpegコマンドを実行
	for _, m2tsFile := range m2tsFiles {
		outputFilePath := filepath.Join(outputDirPath, m2tsFile)
		cmd := exec.Command("ffmpeg", "-i", m2tsFile, "-c", "copy", outputFilePath)

		// ffmpegコマンドを実行して出力を取得
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", m2tsFile, err)
			fmt.Printf("Output:\n%s\n", string(out))
		} else {
			fmt.Printf("Successfully processed file: %s -> %s\n", m2tsFile, outputFilePath)
		}
	}

	// 終了前にユーザーにキー入力を待つ
	fmt.Println("Enterキーを押してください...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
