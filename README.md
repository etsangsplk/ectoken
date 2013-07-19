Go implementation to generate authentication tokens, to access protected content via Edgecast CDN.

Requires a valid key registered on edgecast and takes an optional list of extra security parameters (such as expiry time, etc.) and generates a token.

To find out more about the params to use, refer to the Edgecast documentation at <https://my.edgecast.com/support/docs/> (Look for a zip file named _EdgeCast Token Authentication & Binaries_ and refer to the PDF document inside)

## Usage

```
import (
    "github.com/soundcloud/ectoken"
    "fmt"
)

func main() {
    token, _ := ectoken.Generate("secret", "ec_expire=1333238400&ec_url_allow=test.mp3j")
    fmt.Println(token)
}
```

