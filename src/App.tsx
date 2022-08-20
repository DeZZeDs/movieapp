import React from 'react';
import Layout from "./components/layout";
import HomePage from "./pages/home";

function App() {
  return (
    <div className="App">
        <Layout>
            <HomePage/>
        </Layout>
    </div>
  );
}

export default App;
