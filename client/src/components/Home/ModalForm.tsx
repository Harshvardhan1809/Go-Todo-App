import React, { Fragment } from "react";
import { Dialog, TextField, styled, Button, Typography, Box } from "@mui/material";
import CreateIcon from '@mui/icons-material/Create';
import StarIcon from '@mui/icons-material/Star';
import { DatePicker } from "@mui/x-date-pickers";
import dayjs from 'dayjs';

// attempt to reduce the inline CSS
// can't use makeStyles since it is depreacted (@mui/styles)
// hence styled components

const StyledTextField = styled(TextField)(({}) => ({
    margin: "1rem",
    width: "350px",
}))

const StyledForm = styled('form')(({}) => ({
    display: "flex",
    flexDirection : "column",
    justifyContent: "center",
    alignItems: "center",
    padding: 2,
    margin: "25 px 0px 20px 0px"
}))

const StyledButtonDiv = styled('div')(({}) => ({
    width: "100%", 
    display: "flex",
    flexDirection:"row", 
    justifyContent: "space-evenly",
    margin: "10px auto 20px auto",
}))

const StyledImportantField = styled(Typography)(({}) => ({
    margin : "22.5px 15px 22.5px 15px", 
    color : "gray", 
    fontSize : "1rem"

}))

const ModalForm = (props: {open: boolean, handleClose: () => void}) => {

    const isOpen = props.open;
    const handleClose = props.handleClose;
    const value = dayjs();
    
    const handleSubmit = () =>{
        return 1
    }

    return (
        <Fragment>
            <Dialog open={isOpen} onClose={handleClose}>
                <Box sx={{"width": "450px"}}>
                    <StyledForm action="" onSubmit={handleSubmit}>
                        <Box sx={{"display": "flex", margin: "20px auto 5px auto"}}>
                            <CreateIcon sx={{"margin": "10px 10px 0px 0px"}}/>
                            <Typography variant="h4">
                                Create a new task
                            </Typography>
                        </Box>

                        <StyledTextField label="Title" />
                        <StyledTextField label="Description" />
                        <Box sx={{"display": "flex", "width": "100%", marginLeft: "75px"}}>
                            <StyledImportantField> Important :</StyledImportantField>
                            <Button>
                                <StarIcon sx={{margin: "14px 0px 18px 0px"}}/>
                            </Button>
                        </Box>
                        <Box sx={{"display": "flex", "width": "100%", marginLeft: "75px"}}>
                            <StyledImportantField> Date :</StyledImportantField>
                            <DatePicker value={value} />
                        </Box>
                        
                        <StyledButtonDiv>
                            <Button variant="outlined">Create task</Button>
                            <Button variant="outlined">Clear</Button>
                        </StyledButtonDiv>

                    </StyledForm>
                </Box>
            </Dialog>
        </Fragment>
    )
}

export default ModalForm;