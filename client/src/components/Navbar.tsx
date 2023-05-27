import { Fragment } from 'react';
import { AppBar, CssBaseline, Typography, Button } from '@mui/material';
import { Box } from '@mui/system';
import ChecklistIcon from '@mui/icons-material/Checklist';

// Using Box component from MUI System as a wrapper to give styling to MUI Material components

const Navbar = () => {

    return (
        <Fragment>
            <CssBaseline />
                <AppBar position="static" >
                    <Box sx={{ 'display': 'flex' }}>
                        <Box sx={{ 'padding': '15px 10px 10px 15px',  }}>
                            <ChecklistIcon />
                        </Box>
                        <Box sx={{ 'padding-top': '10px', 'padding-left': '15px' }}>
                            <Typography variant="h6">TODO</Typography>
                        </Box>

                        <Box sx={{"display":"flex", "padding-top": "10px", "margin-left": "auto", "margin-right": "10px"}}>
                            <Box sx={{ "padding-left": "10px", "padding-top": "2.5px"}}>
                                <Typography variant="h6">Harsh</Typography>
                            </Box>
                            <Box sx={{"padding-left": "10px"}}>
                                <Button variant="contained" color="success">Logout</Button>
                            </Box>
                        </Box>
                    </Box>
                </AppBar>

        </Fragment>
    )

}

export default Navbar;