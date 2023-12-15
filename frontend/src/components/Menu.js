import React from 'react';
import { Link } from 'react-router-dom';

function Menu() {
    return (
        <div className="menu">
            <Link to="/">Home</Link>
            <Link to="/app1">App1</Link>
            <Link to="/app2">App2</Link>
            <Link to="/app3">App3</Link>
        </div>
    );
}

export default Menu;