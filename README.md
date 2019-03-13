# Harvest

Harvest is Polar Squad internal application/bot/etc for helping empyees to use Harvest. Easing the hour entries, requesting overtime hour calculations, notifications, prefilled time entries, etc...

The structure of the repo is so that the root has the main application code as a main package, and the sub folder _harvest_ has it's own package that can be used more genericly on other Golang projects also, just by importin it.
But first clone it in your folder since it's a private repo.  
```golang
package main

import "./harvest"`

func main() {
    configFile := "config.json"
    c := config.LoadConfig(configFile)
    
    h := harvest.Init(configFile)
}
```
