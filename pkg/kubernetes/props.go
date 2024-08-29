package kubernetes

func MapToAnnotations(m map[string]string) *map[string]*string {
	antsMap := make(map[string]*string, len(m))

	for k, v := range m {
		antsMap[k] = &v
	}
	return &antsMap
}
