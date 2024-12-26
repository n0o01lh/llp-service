package domain

type ByOrder []Resource

// Len() devuelve el número de elementos en el array
func (a ByOrder) Len() int {
	return len(a)
}

// Swap() intercambia dos elementos
func (a ByOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less() determina si el elemento i debe estar antes que el elemento j
func (a ByOrder) Less(i, j int) bool {
	return a[i].ExtraFields["order"].(int) < a[j].ExtraFields["order"].(int)
}
