# Technical Assessment: In-Memory Database with UI

This exercise is designed to assess your ability to reason about state, handle user input, and build a clean, usable interface. We’re looking for correctness, clarity, and thoughtful design — both in your code and your UI.

You’ll build a small in-memory database with a user interface to interact with it. This project should take no more than 4 hours to complete.

---

## Requirements

### Core Commands

- `SET [name] [value]`
  - Sets the name in the database to the given value

- `GET [name]`
  - Prints the value for the given name. If the name is not in the database, prints `NULL`

- `DELETE [name]`
  - Deletes the name from the database

- `COUNT [value]`
  - Returns the number of names assigned to that value. If the value is not assigned, prints `0`

### Transactions

- `BEGIN`
  - Begins a new transaction

- `ROLLBACK`
  - Rolls back the most recent transaction. If there is no transaction to rollback, print an error message

- `COMMIT`
  - Commits all open transactions

### Additional Requirements

- Use only in-memory data structures (no databases or database-like libraries)
- Build a user interface — this can be a CLI-style input box or a simple form-based UI
- Provide useful error messages for invalid commands or inputs
- Ensure memory usage does not double for each new transaction

---

## Choose 2 of the Following to Implement

- Support nested transactions
- Display command history and results in the UI
- Add automated tests for your core logic

---

## Deliverables

- Your working code
- A `README.md` file that includes:
  - How to run your application
  - Which 2 bonus features you implemented

---

Let us know if you have any questions. We’re excited to see what you build!
