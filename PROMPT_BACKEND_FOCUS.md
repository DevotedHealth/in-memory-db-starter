# Technical Assessment: Backend-Focused In-Memory Database with UI

This exercise is designed to assess your ability to reason about state, handle user input, and build a clean, usable interface. We’re looking for correctness, clarity, and thoughtful design — both in your code and your UI.

You’ll build a small in-memory database with a user interface to interact with it. This project should take no more than 4 hours to complete.

---

## Requirements

### Core Commands

- `SET [name] [value]`  
  Sets the name in the database to the given value

- `GET [name]`  
  Displays the value for the given name in the UI. If the name is not in the database, display `NULL` or an appropriate message

- `DELETE [name]`  
  Deletes the name from the database

- `COUNT [value]`  
  Displays the number of names assigned to that value. If none, show `0`

### Transactions

- `BEGIN`  
  Begins a new transaction

- `ROLLBACK`  
  Rolls back the most recent transaction. If none exist, display an error

- `COMMIT`  
  Commits all open transactions

---

## UI Requirements

Your UI must include:

- An input field where the user can enter commands (for example: `SET a foo`)
- A button to submit the command
- A display area that shows the result of the most recent command (for example: `OK`, `foo`, `NULL`, or an error message)
- Useful error messages for invalid or unrecognized inputs

> You are not being evaluated on visual design. Basic, functional styling is more than sufficient!

---

## Technical Constraints

- Use only in-memory data structures — no database or database-like libraries
- Your implementation should not double memory usage for each transaction

---

## Deliverables

- Your working code
- A `README.md` file with instructions for running the frontend and backend (if applicable)

---

## Examples

### Example 1
```
>> GET a
NULL
>> SET a foo
>> SET b foo
>> COUNT foo
2
>> COUNT bar
0
>> DELETE a
>> COUNT foo
1
>> SET b baz
>> COUNT foo
0
>> GET b
baz
>> GET B
NULL
>> END
```

### Example 2
```
>> SET a foo
>> SET a foo
>> COUNT foo
1
>> GET a
foo
>> DELETE a
>> GET a
NULL
>> COUNT foo
0
>> END
```

### Example 3 (Transactions)
```
>> BEGIN 
>> SET a foo
>> GET a
foo
>> BEGIN 
>> SET a bar
>> GET a
bar
>> SET a baz
>> ROLLBACK 
>> GET a
foo
>> ROLLBACK 
>> GET a
NULL
>> END
```

### Example 4 (Nested Transactions - Optional)
```
>> SET a foo
>> SET b baz
>> BEGIN 
>> GET a
foo
>> SET a bar
>> COUNT bar
1
>> BEGIN 
>> COUNT bar
1
>> DELETE a
>> GET a
NULL
>> COUNT bar
0
>> ROLLBACK 
>> GET a
bar
>> COUNT bar
1
>> COMMIT 
>> GET a
bar
>> GET b
baz
>> END
```

---

## Follow-Up Conversation Topics

We may ask about:

- Your approach to nested transactions
- Ideas for extending the UI (for example: command history, current state)
- Testing strategies across layers

---

Let us know if you have any questions — we’re excited to see what you build!