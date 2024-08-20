import React from 'react';
import HomePage from "./components/HomePage";
import './App.css'; // Подключаем стили

const App = () => {
    return (
        <div className="container">
            <h1>Link Shortener</h1>
            <HomePage />
        </div>
    );
};

export default App;
