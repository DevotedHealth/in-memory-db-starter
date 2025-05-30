/**
 * This function receives a full user-entered command string (such as "SET a foo")
 * and returns a string result to display in the UI (such as "OK", "NULL", or an error message).
 *
 * You can choose how to implement the logic
 * Option 1: Parse and handle the command entirely in TypeScript
 * Option 2: Send the command string to a backend server (via `fetch`, for example)
 */

export const handleCommand = async (command: string): Promise<string> => {
  const [operation, ...args] = command.trim().split(/\s+/);

  switch (operation.toUpperCase()) {
    case 'SET':
      // TODO: implement SET [name] [value]
      return 'OK';

    case 'GET':
      // TODO: implement GET [name]
      return 'NULL';

    case 'DELETE':
      // TODO: implement DELETE [name]
      return 'OK';

    case 'COUNT':
      // TODO: implement COUNT [value]
      return '0';

    case 'BEGIN':
      // TODO: start a new transaction
      return 'OK';

    case 'ROLLBACK':
      // TODO: rollback the most recent transaction
      return 'TRANSACTION NOT FOUND';

    case 'COMMIT':
      // TODO: commit all open transactions
      return 'OK';

    default:
      return 'ERROR: Unknown command';
  }
};
