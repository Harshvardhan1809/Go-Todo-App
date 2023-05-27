import React, { Fragment } from "react";
import { BrowserRouter } from 'react-router-dom'
import Navbar  from "./components/Navbar";
import { Typography } from "@mui/material";

const App = () => {
    return (
            <div className="app">
                <BrowserRouter>
                    <Fragment>
                        <Navbar/>
                        <Typography variant="h2">Hello World</Typography>
                    </Fragment>
                </BrowserRouter>
            </div>
    )
}

export default App;

// Navbar, hero area (with a background design and the todo section) 
