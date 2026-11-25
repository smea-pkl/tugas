import { useEffect, useRef } from "react";
import { Chessground } from "@lichess-org/chessground";
import type { Api } from "@lichess-org/chessground/api";
import type { Key } from "@lichess-org/chessground/types";
import { Chess, SQUARES } from "chess.js";

// Impor CSS bawaan Chessground
import "@lichess-org/chessground/assets/chessground.base.css";
import "@lichess-org/chessground/assets/chessground.brown.css";
import "@lichess-org/chessground/assets/chessground.cburnett.css";

// Tambahkan tipe eksplisit pada parameter (chess: Chess) dan return type (Map<Key, Key[]>)
function toDests(chess: Chess): Map<Key, Key[]> {
  const dests = new Map<Key, Key[]>();

  SQUARES.forEach((s) => {
    const ms = chess.moves({ square: s, verbose: true });
    if (ms.length) {
      dests.set(
        s as Key,
        ms.map((m) => m.to as Key),
      );
    }
  });

  return dests;
}

export function ChessBoard() {
  // Berikan tipe HTMLDivElement dan Api pada useRef
  const boardRef = useRef<HTMLDivElement | null>(null);
  const groundRef = useRef<Api | null>(null);
  const chessRef = useRef<Chess>(new Chess());

  useEffect(() => {
    if (!boardRef.current) return;

    const chess = chessRef.current;

    const ground = Chessground(boardRef.current, {
      fen: chess.fen(),
      orientation: "white",
      movable: {
        color: "white",
        free: false,
        dests: toDests(chess),
      },
      events: {
        move: (orig: Key, dest: Key) => {
          chess.move({ from: orig, to: dest, promotion: "q" });

          const turnColor = chess.turn() === "w" ? "white" : "black";

          ground.set({
            turnColor: turnColor,
            check: chess.isCheck(),
            movable: {
              color: turnColor,
              dests: toDests(chess),
            },
          });
        },
      },
    });

    groundRef.current = ground;

    return () => {
      ground.destroy();
    };
  }, []);

  const handleFlipBoard = () => {
    groundRef.current?.toggleOrientation();
  };

  const handleReset = () => {
    const chess = chessRef.current;
    chess.reset();
    groundRef.current?.set({
      fen: chess.fen(),
      turnColor: "white",
      check: false,
      lastMove: undefined,
      movable: {
        color: "white",
        dests: toDests(chess),
      },
    });
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        gap: "16px",
      }}
    >
      <div ref={boardRef} style={{ width: "480px", height: "480px" }} />

      <div style={{ display: "flex", gap: "8px" }}>
        <button onClick={handleFlipBoard}>Putar Papan</button>
        <button onClick={handleReset}>Reset Permainan</button>
      </div>
    </div>
  );
}
