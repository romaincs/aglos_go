import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
  var ret string
  
  for i:= 0; i<len(numbers); i++ {
    if i == 0 {
      ret += "("
    }
    
    if i == 3 {
      ret += ") "
    }
    
    if i == 6 {
      ret += "-"
    }
    
    ret += fmt.Sprintf("%v", numbers[i])
  }
  
  return ret
}
