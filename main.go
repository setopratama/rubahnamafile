package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Tanya prefix
	fmt.Print("Masukkan prefix untuk nama file (contoh: FT): ")
	prefix, _ := reader.ReadString('\n')
	prefix = strings.TrimSpace(prefix)

	// Tanya path folder dengan contoh
	fmt.Print("Masukkan path folder yang ingin diubah (contoh: C:\\Users\\NamaAnda\\Documents\\FolderUntukDiubah atau /home/namaanda/documents/folderuntukdiubah): ")
	folder, _ := reader.ReadString('\n')
	folder = strings.TrimSpace(folder)

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println("Error membaca folder:", err)
		return
	}

	// Urutkan file berdasarkan nama
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Ubah nama file
	for i, file := range files {
		if file.IsDir() {
			continue // Lewati folder
		}

		ext := filepath.Ext(file.Name())
		newName := fmt.Sprintf("%s%04d%s", prefix, i+1, ext)
		oldPath := filepath.Join(folder, file.Name())
		newPath := filepath.Join(folder, newName)

		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error mengubah nama file %s: %v\n", file.Name(), err)
		} else {
			fmt.Printf("Berhasil mengubah nama: %s -> %s\n", file.Name(), newName)
		}
	}
}
