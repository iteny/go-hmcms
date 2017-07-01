package ini

func Value(section string, key string) string {
	val := cfg.Section(section).Key(key).Value()
	return val
}
