import React from 'react';
import {CommandInput} from './components/CommandInput';
import {OutputDisplay} from './components/OutputDisplay';

export const App = () => {
    return (
        <div style={{ padding: '1rem', fontFamily: 'sans-serif' }}>
            <h1>In-Memory DB</h1>
            <CommandInput />
            <OutputDisplay />
        </div>
    );
};
