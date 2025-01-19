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
    try {
      const sendMsg = { message: inputValue };
      // const backendUrl = process.env.REACT_APP_BACKEND_URL;
      const backendUrl = "http://localhost:8080";
      const post = await axios.post(`${backendUrl}/write`, sendMsg);
      console.log('Successfully wrote to file.');
      alert('Successfully wrote to file.');
    } catch (error) {
      if (error instanceof Error) {
        alert(`エラーが発生しました: ${error.message}`);
      } else {
        alert('未知のエラーが発生しました');
      }
      console.error('Error writing to file:', error);
    }
  };

  return (
    <div className="Test">
      <header className="App-header">
      </header>
      <body className='App-body'>
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
          <button onClick={TableTest}>table</button>
        </div>
      </body>
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

function TableTest() {
  root.render (
    <div className='Test'>
      <div className='App-header'>
        <table className='Table-test'>
          <thead>
            <tr>
              <th scope="col">header1</th>
              <th scope="col">header2</th>
              <th scope="col">header3</th>
              <th scope="col">header4</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
            <tr>
              <td>データ１</td>
              <td>データ２</td>
              <td>データ３</td>
              <td>データ４</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Test;
