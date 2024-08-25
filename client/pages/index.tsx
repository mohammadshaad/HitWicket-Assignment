// pages/index.js
import { useEffect, useState } from "react";
import styles from "../styles/Home.module.css";

const HomePage = () => {
  const [messages, setMessages] = useState<any[]>([]);
  const [input, setInput] = useState("");
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws");
    setWs(socket);

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, message]);
    };

    socket.onclose = () => console.log("WebSocket connection closed");

    return () => socket.close();
  }, []);

  const sendMessage = () => {
    if (ws) {
      ws.send(JSON.stringify({ player: "Player1", move: input }));
      setInput("");
    }
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Turn-based Chess-like Game</h1>
      <div className={styles.board}>
        {messages.map((msg, index) => (
          <p key={index} className={styles.message}>
            {JSON.stringify(msg)}
          </p>
        ))}
      </div>
      <div className={styles.controls}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Enter move (e.g., P1:L)"
          className={styles.input}
        />
        <button onClick={sendMessage} className={styles.button}>
          Send Move
        </button>
      </div>
    </div>
  );
};

export default HomePage;
