import { store } from '../db/store'; // If using Option 1 - Frontend logic in TypeScript

/**
 * This function receives a full user-entered command string (for example: "SET a foo")
 * and returns a string result to display in the UI (for example:  "OK", "NULL", or an error message).
 *
 * You can implement the command logic in one of two ways:
 *
 * Option 1 — Frontend logic in TypeScript: parse and execute the command directly
 * Option 2 — Backend logic via HTTP: send the command to a server that returns the result
 */

export const handleCommand = async (command: string): Promise<string> => {
  /* // Option 1 - Frontend logic in TypeScript
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
        // TODO: begin transaction
        return 'OK';
      case 'ROLLBACK':
        return 'TRANSACTION NOT FOUND';
      case 'COMMIT':
        return 'OK';
      default:
        return 'ERROR: Unknown command';
    }
  */

  /* // Option 2 — Backend logic via HTTP
    try {
      const response = await fetch('/api/command', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({command}),
      });

      if (!response.ok) {
        return 'ERROR: Failed to reach backend';
      }

      const data = await response.json();
      return data.output ?? 'ERROR: No output from backend';
    } catch (err) {
      return 'ERROR: Network or server issue';
    }
  */
};
