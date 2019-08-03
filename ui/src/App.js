import React from 'react';
import logo from './logo.svg';
import './App.css';
import 'antd/dist/antd.css';
import Home from './containers/home';

function App() {
  return (
    <div className="App">
      <header className="App-header">
         <Home/>
      </header>
    </div>
  );
}

export default App;
