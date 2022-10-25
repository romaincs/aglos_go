func alphanumeric(str string) bool {
  if len(str) == 0 {
    return false
  }
  
  for _, r := range str {
    isDigit := r > 47 && r < 58
    isUpperLetter := r > 64 && r < 91
    isLowerLetter := r > 96 && r < 123
    
    if !isDigit && !isUpperLetter && !isLowerLetter {
      return false
    }
  }
  return true
}
