package pull

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LocaleFiles []*LocaleFile

type LocaleFile struct {
	Path, Name, ID, Code, Tag, FileFormat string
	ExistsRemote                          bool
}

func (localeFile *LocaleFile) RelPath() string {
	callerPath, _ := os.Getwd()
	relativePath, _ := filepath.Rel(callerPath, localeFile.Path)
	return relativePath
}

// Locale to Path mapping
func (localeFile *LocaleFile) Message() string {
	str := ""
	if Config.Debug {
		if localeFile.Name != "" {
			str = fmt.Sprintf("%s Name: %s", str, localeFile.Name)
		}
		if localeFile.ID != "" {
			str = fmt.Sprintf("%s Id: %s", str, localeFile.ID)
		}
		if localeFile.Code != "" {
			str = fmt.Sprintf("%s Code: %s", str, localeFile.Code)
		}
		if localeFile.Tag != "" {
			str = fmt.Sprintf("%s Tag: %s", str, localeFile.Tag)
		}
		if localeFile.FileFormat != "" {
			str = fmt.Sprintf("%s Format: %s", str, localeFile.FileFormat)
		}
	} else {
		str = fmt.Sprintf("%s", localeFile.Name)
	}
	return strings.TrimSpace(str)
}
