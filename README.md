# neo-cors
Enable CORS in your Neo application

## Example
```Go
package main

import (
	"github.com/ivpusic/cmap"
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo-cors"
)

func main() {
	app := neo.App()
	app.Use(cors.Init(cmap.C{}))
	app.Start()
}
```

### Settings
Default settings are:
```
Access-Control-Allow-Origin: "*"
Access-Control-Allow-Credentials: "true"
Access-Control-Allow-Methods: "*"
Access-Control-Allow-Headers: "*"
Access-Control-Max-Age: 600
Access-Control-Expose-Headers: ''
```

To override default settings, pass custom settings to ``Init`` function.
```Go
app.Use(cors.Init(cmap.C{
  "Access-Control-Allow-Credentials": "false",
  "Access-Control-Allow-Headers": "X-AUTH-HEADER",
  // etc
}))
```

# License
*MIT*
