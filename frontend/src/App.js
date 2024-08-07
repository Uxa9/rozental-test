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
                    <Route path='' element={<>Welcome page</>} />
                    <Route path='personal_accounts' element={<PersonalAccounts/>} />
                    <Route path='*' element={<>Page Not Found</>} />
                </Routes>
            </MainLayout>
        </BrowserRouter>
    )
}

export default App;
