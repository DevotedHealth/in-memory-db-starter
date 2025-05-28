type Command = 'SET' | 'GET' | 'DELETE' | 'COUNT' | 'BEGIN' | 'ROLLBACK' | 'COMMIT';

export const handleCommand = async (command: Command): Promise<string> => {
    return `Executed: ${command}`;
};
