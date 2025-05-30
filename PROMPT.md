# Technical Assessment: In-Memory Database with UI

This exercise is designed to assess your ability to reason about state, handle user input, and build a clean, usable interface. We’re looking for correctness, clarity, and thoughtful design — both in your code and your UI.

You’ll build a small in-memory database with a user interface to interact with it. This project should take no more than 4 hours to complete.

---

## Requirements

### Core Commands

- `SET [name] [value]`
  - Sets the name in the database to the given value

- `GET [name]`
  - Displays the value for the given name in the UI. If the name is not in the database, displays `NULL` or a meaningful error message

- `DELETE [name]`
  - Deletes the name from the database

- `COUNT [value]`
  - Returns the number of names assigned to that value. If the value is not assigned, displays `0`

### Transactions

- `BEGIN`
  - Begins a new transaction

- `ROLLBACK`
  - Rolls back the most recent transaction. If there is no transaction to rollback, display an error message

- `COMMIT`
  - Commits all open transactions

### UI Requirements

Your submission should include a simple user interface that allows someone to interact with the in-memory database. Specifically, your UI must include:

- An input field where the user can enter commands (for example, `SET a foo`)
- A button to submit the command
- A display area that shows the result of the most recent command (for example, `OK`, `foo`, `NULL` or an error message)
- Useful error messages for invalid commands or inputs

You are free to style your UI however you like, but it should be clear, usable, and responsive to user input.

**Note:** we are not evaluating visual design - basic, functional styling is totally fine!

### Additional Requirements

- Use only in-memory data structures (no databases or database-like libraries)
- Ensure memory usage does not double for each new transaction

---

## Bonus Features (Optional)

- Support nested transactions
- Display command history and/or results in the UI
- Add automated tests for your core logic

---

## Deliverables

- Your working code
- A `README.md` file that explains how to run the frontend and backend (if applicable)

---

## Examples of Use

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

### Example 3
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

Let us know if you have any questions. We’re excited to see what you build!
