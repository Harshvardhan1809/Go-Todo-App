import React from "react";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginContainer from "./containers/LoginContainer";
import SignupContainer from "./containers/SignupContainer";
import Home from "./containers/HomeContainer";
import AuthWrapper from "./components/general/AuthWrapper";

const App: React.FC = React.memo(() => {

    return (
            <div className="app">
                <Router>
                    <Routes>
                        <Route path="/login" element={<LoginContainer />} />
                        <Route path="/signup" element={<SignupContainer />} />
                        <Route path="" element={
                                <Home/>
                        }/>
                    </Routes>
                </Router>
            </div>
    )
})

export default App;
//<AuthWrapper>                            </AuthWrapper>


// Navbar, hero area (with a background design and the todo section) 
