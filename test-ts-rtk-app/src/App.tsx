import React from 'react';
import logo from './logo.svg';
import './App.css';
import MainLayout from './components/layouts/MainLayout';
import { BrowserRouter } from 'react-router-dom';
import AppRouter from './Router';
import { configureStore } from '@reduxjs/toolkit';
import { Provider } from 'react-redux';
import rootReducer from './common/indexStore';
import { loadState, saveState } from './common/localstorage';


function App() {
    return (
        <BrowserRouter>
            <AppRouter />
        </BrowserRouter>
    );
}

export default App;
