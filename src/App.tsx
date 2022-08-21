import React from 'react';
import {Routes,Route} from "react-router-dom";
import Layout from "./components/layout";
import HomePage from "./pages/home";

function App() {
  return (
    <div className="App">
        <Layout>
            <Routes>
                <Route path="/" element={<HomePage/>}/>
                <Route path="*" element={<HomePage/>}/>
            </Routes>
        </Layout>
    </div>
  );
}

export default App;
