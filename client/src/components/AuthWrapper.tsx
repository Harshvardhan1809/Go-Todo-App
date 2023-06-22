import React, {Fragment, useState, useEffect} from 'react'
import { Typography, Card, CardHeader, CardContent, CardActions, Box, TextField, Button, CssBaseline } from '@mui/material';
import {useTheme} from '@mui/material';
import axios from "axios";
import { useQuery, useMutation } from "react-query";
import { queryClient } from '../index';

export type AuthWrapperProps = {
    children: React.ReactNode
}

const AuthWrapper = ({children}: AuthWrapperProps) => {

    // here we make a query to the /auth/validate to check if we are logged in

    return (
        <>
            <h1>auth wrapper</h1>
            <div>
                {children}
            </div>
        </>
    )
}

export default AuthWrapper
