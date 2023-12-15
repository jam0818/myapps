import React from 'react';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import {Helmet} from 'react-helmet';
import Menu from './components/Menu';
import Home from './components/Home';
import App1 from './components/App1';
import App2 from './components/App2';
import App3 from './components/App3';
import './style/App.css';

function App() {
    return (<Router>
            <div className="App">
                <Helmet>
                    <title>my app</title>
                </Helmet>
                <Menu/>
                <div className="content">
                    <Routes>
                        <Route path="/" element={<Home/>}/>
                        <Route path="/app1" element={<App1/>}/>
                        <Route path="/app2" element={<App2/>}/>
                        <Route path="/app3" element={<App3/>}/>
                    </Routes>
                </div>
            </div>
        </Router>);
}

export default App;