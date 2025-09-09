
# models (clean)

Drop these files into your `models/` folder. Remove any older copies that had
string-concatenated tags (e.g. `+ "json:..." +`). Keep ONLY this `validation.go` (delete any `validations.go`).

After copying:
  go get github.com/go-playground/validator/v10
  go mod tidy
  go build ./...
