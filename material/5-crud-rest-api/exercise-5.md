# Exercise 5 – Filter Users by Name (Query Parameter)

**Goal:**  
Extend the `UsersHandler` function so that it can **filter users by name** using a query parameter.  

- If a query parameter `name` is provided, return only users whose names **contain** the given string (case-insensitive).  
- If no query parameter is provided, return all users.  

> Example:  
`GET /users?name=an` should return users like `"Anton"` because `"an"` is part of the name.

---

## Hints / Guidance

- Retrieve the query parameter using:

```go
nameFilter := r.URL.Query().Get("name")
