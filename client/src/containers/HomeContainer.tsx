import React, { Fragment, useState } from "react";
import TodoList from "../components/Home/Todolist"
import ModalForm from "../components/Home/ModalForm"
import { Container, Typography } from "@mui/material";
import Navbar from "../components/general/Navbar"; 

const Home = () => {

    const [modalOpen, setModalOpen] =  useState(false);

    const handleOpen = () => {
        setModalOpen(true);
    }

    const handleClose = () => {
        setModalOpen(false);
    }

    return (
        <main>
            <div>
                <Navbar/>
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
    )
}

export default Home
