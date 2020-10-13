package main

func swap(s []int64, i, j int) {
	aux := s[i]

	s[i] = s[j]
	s[j] = aux
}

func Burbuja(s []int64) {

	for i, _ := range s {

		for j := len(s) - 1; j > i; j-- {

			if s[j-1] > s[j] {
				swap(s, j-1, j)
			}
		}

	}
}

func main() {

}
