// @ts-nocheck

import { useEffect, useState } from "react";
import styles from "../styles/Home.module.css";

const HomePage = () => {
  const [gameState, setGameState] = useState({
    grid: Array(5).fill(null).map(() => Array(5).fill("")),
    turn: "Player1",
  });
  const [input, setInput] = useState("");
  const [ws, setWs] = useState(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws");
    setWs(socket);

    socket.onmessage = (event) => {
      const state = JSON.parse(event.data);
      setGameState(state);
    };

    socket.onclose = () => console.log("WebSocket connection closed");

    return () => socket.close();
  }, []);

  const sendMove = () => {
    if (ws && input) {
      const move = { player: gameState.turn, move: input };
      ws.send(JSON.stringify(move));
      setInput("");
    }
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Turn-based Chess-like Game</h1>
      <div className={styles.board}>
        {gameState.grid?.map((row, rowIndex) => (
          row.map((cell, colIndex) => (
            <div key={`${rowIndex}-${colIndex}`} className={styles.cell}>
              {cell}
            </div>
          ))
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
        <button onClick={sendMove} className={styles.button}>
          Send Move
        </button>
      </div>
      <div className={styles.turnIndicator}>
        Current Turn: {gameState.turn}
      </div>
    </div>
  );
};

export default HomePage;
