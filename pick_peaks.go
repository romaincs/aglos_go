type PosPeaks struct {
  Pos   []int
  Peaks []int
}

func PickPeaks(array []int) PosPeaks {
  var add bool
  var pk = PosPeaks {
    Pos: make([]int, 0),
    Peaks: make([]int, 0),
  }
  
  for i := 1; i < len(array) -1; i++ {
    add = false
    
    if array[i-1] < array[i] {
      if array[i] > array[i+1] {
        add = true  
      }
      
      if array[i] == array[i+1] {
        for j := i+1; j < len(array); j++ {
          if array[j] != array[i] && array[j] < array[i]  {
            add = true
          } else if array[j] > array[i] {
            break
          }
        } 
      }
      
      if add {
        pk.Pos = append(pk.Pos, i)
        pk.Peaks = append(pk.Peaks, array[i])
      }
    } 
  }
  return pk
}
