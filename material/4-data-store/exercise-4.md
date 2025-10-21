# Exercise 4 – CRUD Operations on Users

**Goal:**  
Extend the Go program that manages a slice of `User` structs by implementing **only the Delete and Update functions**.  

> **Important:** You should **not** modify or rewrite any other part of the program. Focus **only** on implementing `DeleteUser` and `UpdateUser`.

The functions should allow:

1. **Update** an existing user's name by ID.
2. **Delete** a user by ID.

---

## Hints / Guidance

- Users are stored in a slice:

```go
var users = []User{
    {Id: 0, Name: "Max"},
    {Id: 1, Name: "Anton"},
}
