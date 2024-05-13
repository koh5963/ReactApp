import React from 'react';
import logo from './logo.svg';
import './App.css';
import ReactDOM from 'react-dom/client';
import Back from './Back';
import root from './index';

function Test() {
  return (
    <div className="Test">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p><strong>これはApp.tsxではない</strong></p>
        <div>
            <button onClick={NewTest}>go</button>
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
export default Test;
