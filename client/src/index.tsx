import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { CssBaseline } from '@mui/material';
import { QueryClientProvider, } from 'react-query';
import { queryClient } from './utils/queryClient';
import { ThemeProvider } from '@mui/material/styles';
import { theme } from './utils/theme';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider theme={theme}>
          <CssBaseline/>
          <App />
        </ThemeProvider>
      </QueryClientProvider>
    </LocalizationProvider> 
  </React.StrictMode>,
)
