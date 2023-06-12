import React, { Fragment, useState } from "react";
import { BrowserRouter } from 'react-router-dom'
import Navbar from "./components/Navbar";
import TodoList from "./components/Todolist"
import ModalForm from "./components/ModalForm"
import { Container, Typography } from "@mui/material";

const App = () => {

    const [modalOpen, setModalOpen] =  useState(false);

    const handleOpen = () => {
        setModalOpen(true);
    }

    const handleClose = () => {
        setModalOpen(false);
    }

    return (
            <div className="app">
                <BrowserRouter>
                    <Fragment>
                        <Navbar/>
                        <main>
                            <div>
                                <Container maxWidth="sm" sx={{"text-align": "center", "margin-top": "25px"}}>
                                    <Typography variant="h4" gutterBottom>TODO App</Typography>
                                    <Typography variant="h6" gutterBottom>A todo app to list down and manage your daily tasks!</Typography>
                                </Container>
                                <Container>
                                    <TodoList modalOpen={handleOpen}/>
                                    <ModalForm open={modalOpen} handleClose={handleClose}/>
                                </Container>
                            </div>
                        </main>
                    </Fragment>
                </BrowserRouter>
            </div>
    )
}

export default App;

// Navbar, hero area (with a background design and the todo section) 
