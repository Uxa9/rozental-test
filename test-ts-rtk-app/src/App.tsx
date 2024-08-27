import React from 'react';
import logo from './logo.svg';
import './App.css';
import MainLayout from './components/layouts/MainLayout';
import { BrowserRouter } from 'react-router-dom';
import AppRouter from './Router';
import { legacy_createStore as createStore } from "redux";
import { Provider } from 'react-redux';
import rootReducer from './common/indexStore';
import { loadState, saveState } from './common/localstorage';

const localStorage = loadState()

const store = createStore(rootReducer, localStorage)

store.subscribe(() => {
    saveState(store.getState())
})

function App() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <AppRouter />
            </BrowserRouter>
        </Provider>
    );
}

export default App;
