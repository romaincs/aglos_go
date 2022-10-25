import (
  "regexp"
  "strconv"
)

func NbDig(n int, d int) (ret int) {  
    reg := regexp.MustCompile(strconv.Itoa(d))
    for i:=0;i<=n;i++{
      ret += len(reg.FindAllString(strconv.Itoa(i*i), -1))
    }
    return
    
}
