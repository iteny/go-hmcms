package ini

import (
	"log"
	"time"
)

func Value(section string, key string) string {
	val := cfg.Section(section).Key(key).Value()
	return val
}
func Int(section string, key string) (int, error) {
	val, err := cfg.Section(section).Key(key).Int()
	if err != nil {
		log.Fatal(err)
	}
	return val, err
}
func Int64(section string, key string) (int64, error) {
	val, err := cfg.Section(section).Key(key).Int64()
	if err != nil {
		log.Fatal(err)
	}
	return val, err
}

func Duration(section string, key string) (time.Duration, error) {
	val, err := cfg.Section(section).Key(key).Duration()
	if err != nil {
		log.Fatal(err)
	}
	return val, err
}
func MustDuration(section string, key string) time.Duration {
	val := cfg.Section(section).Key(key).MustDuration()
	return val
}
