import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { connect, sendMsg } from "./api";
import React from 'react';

// function App() {
  

//   return (
//     <>
      
//     </>
//   )
// }

// export default App

class App extends React.Component {
  constructor(props: {}) {
    super(props);
    connect();
  }

  send() {
    console.log("Sending message");
    sendMsg("Hello");
  }

  render(){
    return (
      <>
        <div className='App'>
          <button onClick={this.send}>Hit this</button>
        </div>
      </>
    )
  }
}

export default App;
