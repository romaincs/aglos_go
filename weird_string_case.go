func toWeirdCase(str string) (out string) {
  var i int
  for _, c := range str {
    s := string(c)
    
    if i%2 == 0 {
      out += strings.ToUpper(s)
    } else {
      out += strings.ToLower(s)
    }
    
    if s == " " {
      i = 0
    } else {
      i ++
    }
  }
  
  return
}
