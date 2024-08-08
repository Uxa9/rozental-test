import React from 'react';
import MainLayout from './Layouts/mainLayout';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import PersonalAccounts from './PersonalAccounts';
import styles from './App.css';

const App = (props) => {

    return (
        <BrowserRouter>
            <MainLayout>
                <Routes>
                    <Route path='*' element={<PersonalAccounts/>} />
                </Routes>
            </MainLayout>
        </BrowserRouter>
    )
}

export default App;
