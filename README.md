i18naddress can tell you all the names of cities, states, provinces of all nations around the world.

```go
package main

import (
  "github.com/site-name/i18naddress"
  "log"
)

func main() {
  
  params := &i18naddress.Params{
    CountryCode: "vn",
  }

  rules, err := i18naddress.GetValidationRules(params)
  if err != nil {
    log.Fatalln(err)
  }

  fmt.Println(rules)
}

```
