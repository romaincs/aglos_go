import "strconv"

func RangeExtraction(list []int) (out string) {
  var cpt int 
  for i, v := range list {
    
    var skip bool
    if i+1 != len(list) {      
      skip = list[i+1] == v+1       
    } 
    
    if !skip || cpt == 0 {      
      if i != 0  {        
        if cpt < 2 {
          out += ","  
        } else {
          out += "-"
        }        
      } 
      
      out += strconv.Itoa(v)
    }
    
    if skip  { 
      cpt++     
    } else {
      cpt = 0
    }
  }
  return 
}
