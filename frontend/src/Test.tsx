import logo from './logo.svg';
import './App.css';
import Back from './Back';
import root from './index';
import React, { useState } from 'react';
import axios from 'axios';

function Test() {
  const [inputValue, setInputValue] = useState('');
  const handleChange = (event: { target: { value: React.SetStateAction<string>; }; }) => {
    setInputValue(event.target.value); // 入力値を更新
  };

  const handleClick = async () => {
    alert("リクエストを送信しています...");
    try {
      alert(inputValue);
      const sendMsg = { message: inputValue }; // 正しい形に変更
      console.log(sendMsg);
      await axios.post('http://localhost:8080/write', sendMsg);
      console.log('Successfully wrote to file.');
    } catch (error) {
      if (error instanceof Error) {
        alert(`エラーが発生しました: ${error.message}`); // ユーザーにエラーメッセージを表示
      } else {
        alert('未知のエラーが発生しました');
      }
      console.error('Error writing to file:', error);
    }
  };

  return (
    <div className="Test">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p><strong>これはApp.tsxではない</strong></p>
        <div>
          <input 
            type="text" 
            value={inputValue} 
            onChange={handleChange}
          />
        </div>
        <div>
            <button onClick={NewTest}>go</button>
            <button onClick={handleClick}>test</button>
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
