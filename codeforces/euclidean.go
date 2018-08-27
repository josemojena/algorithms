package codeforces


func Euclidean(a, b int) int{

  if a == 0 {
  	return b
  }

  return Euclidean(b%a, a)

}