# GoSnowflakeApi

GoSnowflakeApi is an api wrapper for the snowflake api written in golang for golang developers.

## Example

```go
package main

import (
    "fmt"
    "github.com/DevSnowflake/gosnowflakeapi"
)

func main() {
    client := gosnowflakeapi.Client{"YOUR_API_TOKEN"}
    fmt.Println(client.ChatWith("Hello!", "Snowflake", "male", 1));
}

```

## Todo

- Write documentation for undocumented fields in the structs.
- Make support for /api/ytsearch endpoints.