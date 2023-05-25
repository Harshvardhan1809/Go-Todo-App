import React  from "react"
import ReactDOM from 'react-dom/client';
import App from "./App"; 

const rootElement = document.getElementById('root');
if(!rootElement) throw new Error('Failed to find the root element');

const root: ReactDOM.Root = ReactDOM.createRoot(rootElement);
root.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>
)

// change the file extensions to .tsx for any files with component code or logic