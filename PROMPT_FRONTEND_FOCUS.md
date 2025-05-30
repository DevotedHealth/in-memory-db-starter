# Technical Assessment: Frontend-Focused In-Memory Database with UI

This exercise is designed to assess your ability to manage frontend state, handle user input, and build a clear and responsive interface. We’re looking for correctness, clarity, and thoughtful design — both in your code and your UI.

You’ll build a small in-memory database with a user interface to interact with it. This project should take no more than 4 hours to complete.

---

## Requirements

### Core Commands

- `SET [name] [value]`  
  Sets the name in the database to the given value

- `GET [name]`  
  Displays the value for the given name in the UI. If the name is not in the database, display a meaningful message

- `DELETE [name]`  
  Deletes the name from the database

- `COUNT [value]`  
  Displays the number of names assigned to that value. If none, show `0`

---

### UI Requirements

Your UI must include:

- An input field where the user can enter commands (for example: `SET a foo`)
- A button to submit the command
- A display area that shows the result of the most recent command (for example: `foo` if the input is `GET a`)
- A display area showing the full history of entered commands and results
  - This history should be editable and update the DB state when changed
- A display area showing the current state of the database
  - This should live-update whenever a new valid command is entered or the history is edited
- Useful error messages for invalid inputs

> You are not being evaluated on visual design. Basic, functional styling is more than sufficient!

---

### Technical Constraints

- Use only in-memory data structures — no database or database-like libraries

---

## Deliverables

- Your working code
- A `README.md` file with instructions for running your application

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

### Example 3 (Editable History and Live DB View)
```
>> SET a foo
>> SET b bar
>> DELETE a
>> COUNT foo
0
```

At this point, the UI might display:

**Command History**
```
1. SET a foo
2. SET b bar
3. DELETE a
4. COUNT foo → 0
```

**Database State**
```
b = bar
```

---

Now, imagine the user edits line 3 in the command history, changing:
```
DELETE a
```
to:
```
SET a baz
```

The app should automatically re-evaluate and update the state accordingly:

**Updated Command History**
```
1. SET a foo
2. SET b bar
3. SET a baz
4. COUNT foo → 0
```

**Updated Database State**
```
a = baz
b = bar
```

Then, if the user edits the final line to:
```
COUNT baz
```
the command history should update to:
```
1. SET a foo
2. SET b bar
3. SET a baz
4. COUNT baz → 1
```

### Example 4 (Editable History and Live DB View)

```
>> SET x alpha
>> SET y beta
>> DELETE x
>> SET z gamma
>> COUNT alpha
0
```

Initial UI state:

**Command History**
```
1. SET x alpha
2. SET y beta
3. DELETE x
4. SET z gamma
5. COUNT alpha → 0
```

**Database State**
```
y = beta
z = gamma
```

---

Now, the user edits line 3 to:
```
SET x delta
```

**Updated Command History**
```
1. SET x alpha
2. SET y beta
3. SET x delta
4. SET z gamma
5. COUNT alpha → 0
```

**Updated Database State**
```
x = delta
y = beta
z = gamma
```

Next, the user reorders the commands:
```
1. SET z gamma
2. SET x alpha
3. DELETE x
4. SET y beta
5. COUNT alpha → 0
```

**Updated Database State**
```
y = beta
z = gamma
```

Then, the user deletes line 3 entirely. Now the state becomes:
```
1. SET z gamma
2. SET x alpha
3. SET y beta
4. COUNT alpha → 1
```

**Updated Database State**
```
x = alpha
y = beta
z = gamma
```

---

## Follow-Up Conversation Topics

We may ask about:
- How you might support transactions in the future (see description in the [backend-focused prompt](https://github.com/DevotedHealth/in-memory-db-starter/blob/master/PROMPT_BACKEND_FOCUS.md#transactions))
- Approaches to state management
- Testing strategies across layers
- How you might enhance usability (for example: syntax highlighting in the input box)

---

Let us know if you have any questions — we’re excited to see what you build!
