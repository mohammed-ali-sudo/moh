# MoH Registry Models (Go)

- One file per model under `models/`
- Simple, general `FirstError()` returning "`Field` is invalid"
- All IDs are `string` with `validate:"uuid4"` (except BIGSERIAL IDs for Outbox/Audit)
- Uses: `github.com/go-playground/validator/v10`

## Usage

```go
import "path/to/models"

drug := models.Drug{
    // ...
}
if msg, ok := drug.Validate(); !ok {
    // handle msg
}
```
