import React, {Fragment} from "react";
import {Card, CardContent, CardHeader, Typography, Box, Button, Tooltip} from "@mui/material"
import StarIcon from '@mui/icons-material/Star';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import RadioButtonUncheckedIcon from '@mui/icons-material/RadioButtonUnchecked';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import ButtonBase from '@mui/material/ButtonBase';
import { DatePicker } from "@mui/x-date-pickers";
import * as dayjs from 'dayjs';

const TodoList = (props: { modalOpen : () => void }) => {
    
    const modalOpen = props.modalOpen;
    const value = dayjs();

    return (
        <Fragment>
            <Card sx={{width: 800, "margin": "auto", "margin-top": "15px", "padding-top": "30px"}}>

                <Box sx={{"display": "flex", "justify-content": "center"}}>
                    {/* <CardHeader title="Tasks for today, 29/05/2023" sx={{"text-align": "center"}}/>
                    <Button>
                        <CalendarMonthIcon sx={{"margin": "0px 0px 0px 5px"}}/>
                    </Button> */}
                    <CardHeader title="Tasks for " sx={{"text-align": "center"}}/>
                    <DatePicker/>
                </Box>

                <Box sx={{"display": "flex"}}>
                    <Typography sx={{"margin": "2.5px 10px 0px 75px", "color": "red"}} variant="h6">INCOMPLETE</Typography>
                    <Tooltip title="Add new task">
                        <Button onClick={modalOpen}>
                            <AddCircleIcon sx={{"margin": "0px 0px 0px 0px"}} />
                        </Button>
                    </Tooltip>
                </Box>
                <CardContent>

                    <Box sx={{"margin": "0px 50px 10px 50px", "display": "flex", "border": "solid", "border-radius": "5px","border-width": "1px", "border-color": "gray", "padding": "7.5px 10px 7.5px 20px"}}>
                        <Tooltip title="Done">
                            <RadioButtonUncheckedIcon sx={{"margin": "5px 20px auto 0px"}}/>
                        </Tooltip>
                        <Box>
                            <Typography variant="h6" sx={{"margin": "auto 0 auto 0"}}>Go to gym and workout</Typography>
                            <Typography>Did arms workout yesterday so do back workout today. Tommorrow do legs workout</Typography>
                        </Box>
                        <Box sx={{"display": "flex", "margin-left": "auto"}}>
                            <Tooltip title="Star as important">
                                <StarIcon sx={{"margin": "auto 15px auto 0px"}}/>
                            </Tooltip>
                            <Button variant="contained">DONE</Button>
                            <Button variant="contained" color="error" sx={{"margin": "0px 0px 0px 10px"}}>DELETE</Button>
                        </Box>
                    </Box>

                    <Box sx={{"margin": "0px 50px 10px 50px", "display": "flex", "border": "solid", "border-radius": "5px","border-width": "1px", "border-color": "gray", "padding": "7.5px 10px 7.5px 20px"}}>
                        <Tooltip title="Done">
                            <RadioButtonUncheckedIcon sx={{"margin": "5px 20px auto 0px"}}/>
                        </Tooltip>                        
                        <Box>
                            <Typography variant="h6" sx={{"margin": "auto 0 auto 0"}}>Do schoolwork</Typography>
                            <Typography>Write code and report for experiments, create the video for TOEFL lecture, prepare for midterms</Typography>
                        </Box>
                        <Box sx={{"display": "flex", "margin-left": "auto"}}>
                            <Tooltip title="Unstar">
                                <StarIcon sx={{"margin": "auto 15px auto 0px", "color": "gold"}}/>
                            </Tooltip>
                            <Button variant="contained">DONE</Button>
                            <Button variant="contained" color="error" sx={{"margin": "0px 0px 0px 10px"}}>DELETE</Button>                      
                        </Box>
                    </Box>

                </CardContent>
                
                <Typography sx={{"margin": "0px 50px 0px 75px", "color": "green"}} variant="h6">COMPLETED</Typography>
                <CardContent>

                    <Box sx={{"margin": "0px 50px 10px 50px", "display": "flex", "border": "solid", "border-radius": "5px","border-width": "1px", "border-color": "gray", "padding": "7.5px 10px 7.5px 20px"}}>
                        <Tooltip title="Undo">
                            <CheckCircleIcon sx={{"margin": "5px 20px auto 0px"}}/>
                        </Tooltip>
                        <Box>
                            <Typography variant="h6" sx={{"margin": "auto 0 auto 0"}}>Completed Integrated Circuit assignment</Typography>
                            <Typography>Completed assignment 13</Typography>
                        </Box>
                        <Box sx={{"display": "flex", "margin-left": "auto"}}>
                            <StarIcon sx={{"margin": "auto 15px auto 0px"}}/>
                        </Box>
                    </Box>
                    <Typography sx={{"color": "black"}}>displays completed tasks from today, can't clear but automatically disappears in 24 hours</Typography>

                </CardContent>

                <Typography sx={{"margin": "0px 50px 0px 75px", "color": "gray"}} variant="h6">PREVIOUS TASKS</Typography>
                <CardContent>

                    <Box sx={{"margin": "0px 50px 10px 50px", "display": "flex", "border": "solid", "border-radius": "5px","border-width": "1px", "border-color": "gray", "padding": "7.5px 10px 7.5px 20px"}}>
                        <Tooltip title="Done">
                            <RadioButtonUncheckedIcon sx={{"margin": "5px 20px auto 0px"}}/>
                        </Tooltip>
                        <Box>
                            <Typography variant="h6" sx={{"margin": "auto 0 auto 0"}}>Pay the water bill for April and May</Typography>
                            <Typography>2 months unpaid lol</Typography>
                            <Typography>22nd April 2023</Typography>
                        </Box>
                        <Box sx={{"display": "flex", "margin-left": "auto"}}>
                            <Tooltip title="Star as important">
                                <StarIcon sx={{"margin": "auto 15px auto 0px"}}/>
                            </Tooltip>
                            <Button variant="contained">DONE</Button>
                            <Button variant="contained" color="error" sx={{"margin": "0px 0px 0px 10px"}}>DELETE</Button>                 
                        </Box>
                    </Box>
                    <Typography sx={{"color": "black"}}>view more (fetches 5 more incomplete tasks, currently seeing tasks from yesterday and the day before yesterday)</Typography>
                </CardContent>

            </Card>
        </Fragment>
    )
}

export default TodoList;