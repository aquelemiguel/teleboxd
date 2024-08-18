package locales

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type LocaleData struct {
	selected string
	locales map[string](map[string]string)
}

var (
	instance *LocaleData
	once sync.Once
)

func LoadLocales() error {
	var localePaths []string

	filepath.WalkDir("src/locales", func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".json") {
			localePaths = append(localePaths, path)
		}
		return nil
	})

	for _, localePath := range localePaths {
		content, err := parseLocale(localePath)
		if err != nil {
			return err
		}
		instance := getInstance()
		instance.locales[getLocaleKey(localePath)] = content
	}
	return nil
}

func SetLocale(locale string) {
	getInstance().selected = locale
}

func Translate(key string) (string, bool) {
	instance := getInstance()
	message, ok := instance.locales[instance.selected][key]
	if !ok {
		message, ok = instance.locales["en-US"][key]
		if !ok {
			log.Printf("cannot translate key %s", key)
			return "", false
		}
	}
	return message, ok
}

func getInstance() *LocaleData {
	once.Do(func() {
		instance = &LocaleData{
			locales: map[string](map[string]string){},
		}
	})
	return instance
}

func parseLocale(localePath string) (map[string]string, error) {
	content, err := os.ReadFile(localePath)
	if err != nil {
		return nil, err
	}
	
	var data map[string]string
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getLocaleKey(path string) string {
	filename := filepath.Base(path)
    return strings.TrimSuffix(filename, filepath.Ext(filename))
}