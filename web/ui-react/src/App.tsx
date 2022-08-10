import React from 'react';
import './App.css';
import {Outlet, Route, Routes} from "react-router-dom";
import {Home} from "./Home";
import {Config} from "./Config";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {ReactQueryDevtools} from "@tanstack/react-query-devtools";
import Head from "./Head";

const queryClient = new QueryClient();

function Log() {
    return null;
}

function App() {

    return (
        <QueryClientProvider client={queryClient}>
            <div className="App">
                <Head/>
                <Routes>
                    <Route path="/" element={<Home/>}/>
                    <Route path="/config" element={<Config/>}/>
                    <Route path="/log" element={<Log/>}/>
                </Routes>
                <Outlet/>
            </div>
            <ReactQueryDevtools initialIsOpen={true}/>
        </QueryClientProvider>
    );
}

export default App;
