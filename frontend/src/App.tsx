import './App.css'
import { connect, sendMsg } from "./api";
import React from 'react';
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';


function App() {

  // const [messages, setMessages] = React.useState([]);

  const handleClick = () => {
    console.log("sending message");
    sendMsg("Am now communicating with the server");
  }
  

  return (
    <>
      <div className='App'>
        <Header />
        <button onClick={handleClick}>Hit this</button>
        <ChatHistory chatHistory={[ {data: "This is a message"}]} />
      </div>
    </>
  )
}


// class App extends React.Component {
//   constructor(props: {}) {
//     super(props);
//     connect();
//   }

//   send() {
//     console.log("Sending message");
//     sendMsg("Hello");
//   }

//   render(){
//     return (
//       <>
//         <div className='App'>
//           <button onClick={this.send}>Hit this</button>
//         </div>
//       </>
//     )
//   }
// }

export default App;
