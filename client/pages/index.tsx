// @ts-nocheck

import { useState, useEffect } from 'react';
import styles from '../styles/Home.module.css';

export default function Home() {
  const [gameState, setGameState] = useState(null);
  const [ws, setWs] = useState(null);
  const [character, setCharacter] = useState('');
  const [move, setMove] = useState('');

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/ws');
    socket.onmessage = (event) => {
      const state = JSON.parse(event.data);
      setGameState(state);
    };
    setWs(socket);

    return () => {
      socket.close();
    };
  }, []);

  const sendMove = () => {
    const player = gameState.turn === 'A' ? 'A' : 'B';
    ws.send(
      JSON.stringify({
        action: 'move',
        data: { player, character, move },
      })
    );
    setCharacter('');
    setMove('');
  };

  if (!gameState) return <div className={styles.loading}>Connecting...</div>;

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Turn-Based Chess-like Game</h1>
      <div className={styles.board}>
        {gameState.board?.map((row, x) =>
          row.map((cell, y) => (
            <div
              key={`${x}-${y}`}
              className={styles.cell}
              onClick={() => setCharacter(`${x},${y}`)}
              style={{ backgroundColor: cell ? '#ddd' : '#eee' }}
            >
              {cell ? <p>{cell}</p> : ''}
            </div>
          ))
        )}
      </div>
      <div className={styles.controls}>
        <input
          type="text"
          className={styles.input}
          value={character}
          onChange={(e) => setCharacter(e.target.value)}
          placeholder="Character (e.g., P1:0,1)"
        />
        <input
          type="text"
          className={styles.input}
          value={move}
          onChange={(e) => setMove(e.target.value)}
          placeholder="Move (e.g., R)"
        />
        <button className={styles.button} onClick={sendMove}>
          Send Move
        </button>
      </div>
      <div className={styles.turnIndicator}>
        {gameState.gameOver
          ? `Game Over! Winner: Player ${gameState.gameOver}`
          : `Player ${gameState.turn}'s turn`}
      </div>
    </div>
  );
}
