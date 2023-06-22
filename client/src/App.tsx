import React, { Fragment, useState } from "react";
import { BrowserRouter as Router, Routes, Route, Link, Navigate} from 'react-router-dom';
import Login from "./scenes/Login";
import Signup from "./scenes/Signup";
import Home from "./scenes/Home";
import AuthWrapper from "./components/AuthWrapper";

const App = () => {

    return (
            <div className="app">
                <Router>
                    <Routes>
                        <Route path="/login" element={<Login />} />
                        <Route path="/signup" element={<Signup />} />
                        <Route path="" element={
                            <AuthWrapper>
                                <Home/>
                            </AuthWrapper>
                        }/>
                    </Routes>
                </Router>
            </div>
    )
}

export default App;

// Navbar, hero area (with a background design and the todo section) 
