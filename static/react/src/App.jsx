import { useRef, useState } from 'react'
import './App.css'

function App() {
  const [messages, setMessages] = useState([]);
  const [connectionStatus, setConnectionStatus] = useState(false);
  const inputRef = useRef();

  const socketRef = useRef();

  const connect = () => {
    socketRef.current = new WebSocket("ws://localhost:8080/ws?name=react-basic");

    socketRef.current.onopen = () => {
      console.log("Successfully Connected");
      setConnectionStatus(true);
    };

    socketRef.current.onmessage = msg => {
      console.log(msg);
      setMessages(prevMessages => [...prevMessages, msg.data]);
    };

    socketRef.current.onclose = event => {
      console.log("Socket Closed Connection: ", event);
      setConnectionStatus(false);
    };

    socketRef.current.onerror = error => {
      console.log("Socket Error: ", error);
      setConnectionStatus(false);
    };
  }


  const disconnect = () => {
    socketRef.current?.close();
  }


  const send = () => {
    if (!inputRef.current) {
      return;
    }
    const msg = inputRef.current.value;
    socketRef.current.send(msg);
    inputRef.current.value = null;
  }

  return (
    <>
      <h1>React websockets demo</h1>
      <h2>Websocket status: {connectionStatus ? "Connected" : "Disconnected"}</h2>
      {connectionStatus ? <button onClick={disconnect}>Disconnect</button> :
        <button onClick={connect}>Connect</button>}
      <input id="message" name="message" type="text" ref={inputRef} />
      <button onClick={send}>Send message</button>
      <ul>
        {messages.map((m, i) => <li key={`${m}_${i}`}> {m}</li>)}
      </ul >
    </>)
}

export default App
