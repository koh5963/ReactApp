import logo from './logo.svg';
import './App.css';
import Back from './Back';
import root from './index';
import React from 'react';
import axios from 'axios';

function Test() {
  return (
    <div className="Test">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p><strong>これはApp.tsxではない</strong></p>
        <div>
            <button onClick={NewTest}>go</button>
            <button onClick={App}>test</button>
        </div>
      </header>
    </div>
  );
}

function NewTest() {
    root.render (
        <div className="Test">
            <header className="App-header">
                <p><strong>ボタン押下の結果がこれだよ</strong></p>
                <button onClick={() => Back(<Test />)}>back</button>
            </header>
        </div>
    )
}

function App() {
  const handleClick = async () => {
    try {
      await axios.post('http://localhost:8080/write', { message: 'おした！' });
      console.log('Successfully wrote to file.');
    } catch (error) {
      console.error('Error writing to file:', error);
    }
  };

  return (
    <div>
      <h1>React App</h1>
      <button onClick={handleClick}>ボタン</button>
    </div>
  );
}

export default Test;
