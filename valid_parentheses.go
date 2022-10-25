func ValidParentheses(parens string) bool {
    var nbp int
    for _, c := range parens {
      if c == 40 {
        nbp++
      } else if c == 41 {
        nbp--
      }
      
      if nbp < 0 {
        return false
      }
    }
  
    return nbp == 0
}
