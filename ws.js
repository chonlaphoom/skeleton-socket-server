import { WebSocketServer } from "ws";

const server = new WebSocketServer({
  port: 8081,
});

server.on("connection", (socket) => {
  console.log("Client connected");

  let interval = null;

  socket.on("message", (message) => {
    console.log(`Received: ${message}`);

    if (interval != null) {
      clearInterval(interval);
    }

    interval = setInterval(() => {
      socket.send(`Server: ${message}`);
    }, 3000);
  });

  socket.on("close", () => {
    console.log("Client disconnected");
  });
});

console.log("WebSocket server is running on ws://localhost:8081");
