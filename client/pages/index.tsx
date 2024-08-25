// @ts-nocheck

import { useState, useEffect } from 'react';

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
    ws.send(JSON.stringify({
      action: 'move',
      data: { player, character, move }
    }));
    setCharacter('');
    setMove('');
  };

  if (!gameState) return <div>Connecting...</div>;

  return (
    <div>
      <h1>Turn-Based Chess-like Game</h1>
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(5, 50px)', gridGap: '1px', width: '256px', margin: '0 auto' }}>
        {gameState.board?.map((row, x) =>
          row.map((cell, y) => (
            <div key={`${x}-${y}`} style={{ width: '50px', height: '50px', backgroundColor: cell ? '#ddd' : '#eee', display: 'flex', alignItems: 'center', justifyContent: 'center', cursor: 'pointer' }} onClick={() => setCharacter(`${x},${y}`)}>
              {cell ? <p>{cell}</p> : ''}
            </div>
          ))
        )}
      </div>
      <div style={{ margin: '20px', textAlign: 'center' }}>
        <input type="text" value={character} onChange={(e) => setCharacter(e.target.value)} placeholder="Character (e.g., P1:0,1)" />
        <input type="text" value={move} onChange={(e) => setMove(e.target.value)} placeholder="Move (e.g., R)" />
        <button onClick={sendMove}>Send Move</button>
      </div>
      <div id="status" style={{ marginTop: '20px', textAlign: 'center' }}>
        {gameState.gameOver ? `Game Over! Winner: Player ${gameState.gameOver}` : `Player ${gameState.turn}'s turn`}
      </div>
    </div>
  );
}
