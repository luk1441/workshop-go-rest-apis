# Exercise 6 – Add a Method to Clear User Name

**Goal:**  
Extend the `User` struct by adding a new method `ClearName` that sets the user's `Name` field to an empty string (`""`).  

- You should **only** implement this new method.  
- Do **not** modify existing methods (`New` or `Display`).

---

## Hints / Guidance

- Method receivers can be **pointer receivers** if you want to modify the original struct:

```go
func (u *User) ClearName() {
    // TODO: set u.Name to ""
}
