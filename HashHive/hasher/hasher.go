package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// List of file extensions to hash
var validExtensions = []string{".exe", ".dmg", ".sh", ".bin", ".zip", ".dll", ".iso", ".rar"}

// Function to check if a file has a valid extension
func hasValidExtension(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// Function to compute MD5 hash of a file
func hashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Function to walk the filesystem and hash files with valid extensions
func WalkAndHash(rootDir, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	fmt.Printf("Starting file hashing from: %s\n", rootDir)

	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil // Continue even if there's an error
		}

		// Only hash files with valid extensions, not directories
		if !d.IsDir() && hasValidExtension(path) {
			hash, err := hashFile(path)
			if err == nil {
				outFile.WriteString(fmt.Sprintf("%s,%s\n", path, hash))
			}
		}
		return nil
	})

	return err
}
