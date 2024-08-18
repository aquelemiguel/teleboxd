package locales

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	instance *LocaleData
)

type LocaleData struct {
	primary map[string]string
	secondary map[string]string
}

func LoadLocale(locale string) (bool, error) {
	if instance == nil {
		content, err := parseLocale("en-US")
		if err != nil {
			return false, err
		}
		instance = &LocaleData{secondary: content}
	}
	content, err := parseLocale(locale)
	if err != nil {
		return false, err
	}
	instance.primary = content
	return true, nil
}

func Translate(key string) (string, bool) {
	if instance == nil {
		log.Printf("cannot translate without a loaded locale")
		return "", false
	}
	message, ok := instance.primary[key]
	if !ok {
		message, ok = instance.secondary[key]
		if !ok {
			log.Printf("cannot translate key %s", key)
			return "", false
		}
	}
	return message, ok
}

func parseLocale(locale string) (map[string]string, error) {
	localePath := filepath.Join("src/locales", fmt.Sprintf(("%s.json"), locale))
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