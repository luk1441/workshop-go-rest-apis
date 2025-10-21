# Exercise 3 – Decode JSON Request

**Goal:**  
Implement the `decodeHandler` function that:

1. Reads a JSON request body from the client.
2. Parses it into the `RequestBody` struct.
3. Prints the parsed struct to the console.
4. Returns nothing to the client (status 200 OK by default).

---

## Hints / Guidance

- Use `json.NewDecoder` to read JSON from `r.Body`.
- Always handle errors in case the client sends invalid JSON.
- The `RequestBody` struct is already defined as:

```go
type RequestBody struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
