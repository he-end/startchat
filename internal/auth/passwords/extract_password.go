package authpassword

func ExtractPassword(store_hash string) (salt, pwd string) {
	if store_hash == "" {
		return "", ""
	}
	split := splitOne(store_hash, "$")
	if split[0] == "" && split[1] == "" {
		return "", ""
	}
	return string(split[0]), string(split[1])
}

func splitOne(s, sep string) []string {
	idx := -1
	for i := range s {
		if string(s[i]) == sep {
			idx = i
			break
		}
	}

	if idx == -1 {
		return []string{s}
	}
	return []string{s[:idx], s[idx+1:]}
}
