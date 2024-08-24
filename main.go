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
	} else {
		fmt.Println(".m2ts files found:")
		for _, file := range m2tsFiles {
			fmt.Println(filepath.Join(currentDir, file))
		}
	}

	// 各.m2tsファイルに対してffmpegコマンドを実行
	for _, m2tsFile := range m2tsFiles {
		fmt.Println(m2tsFile)
		cmd := exec.Command("ffmpeg", "-i", m2tsFile, "-c", "copy", OutputDirName+"/"+m2tsFile)
		fmt.Println(cmd.String())
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", m2tsFile, err)
			fmt.Printf("Output:\n%s\n", string(out))
		} else {
			fmt.Printf("Successfully processed file: %s\n", m2tsFile)
		}
	}

	// 終了前にユーザーにキー入力を待つ
	fmt.Println("Enterキーを押してください...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
