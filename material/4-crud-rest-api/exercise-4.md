# Exercise 4: Implement Update User Endpoint

## Goal
Add the ability to update an existing user via a PUT request to `/user/{userId}`.

## Task

Add a new case `http.MethodPut:` to the switch statement in `UserHandler` function in `handlers/handlers.go`.

## Requirements

- Decode the request body into a `user.UserInput` struct
- Call `user.UpdateUser(u)` (already implemented)
- Return the updated user as JSON with status `200 OK`
- Handle errors:
  - Invalid JSON → `400 Bad Request` with message "Invalid User JSON"
  - User not found → `404 Not Found` with error message

## Example Request

```bash
curl -X PUT http://localhost:8080/user/1 \
  -H "Content-Type: application/json" \
  -d '{"id": 1, "name": "Updated Name", "age": 30}'
```

## Hints

- Use `json.NewDecoder(r.Body).Decode(&u)` to decode
- Use `user.UpdateUser(u)` to update
- Use `json.NewEncoder(w).Encode(updatedUser)` to respond

Check `handlers/handlers.go` for the solution!