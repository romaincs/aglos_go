import(
  "sort"
)

func TwoOldestAges(ages []int) [2]int {
  sort.Ints(ages)
  var arr [2]int
  copy(arr[:], ages[len(ages)-2:len(ages)])
  return arr
}
