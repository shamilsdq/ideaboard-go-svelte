import { w3cwebsocket as WebSocket } from "websocket";

const API_ROOT = import.meta.env.VITE_SERVER_URL + "/api/boards";

const create = async (payload) => {
  const response = await fetch(API_ROOT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  const responseJson = await response.json();
  if (!response.ok) {
    throw new Error(responseJson);
  } else {
    return responseJson;
  }
};

const connectSocket = async (boardId) => {
  const socketUrl = `${API_ROOT.replace(/^https?:\/\//i, "ws://")}/${boardId}`;
  const ws = new WebSocket(socketUrl);
  ws.onopen = () => console.log(`Opened WebSocket connection at ${socketUrl}`);
  ws.onclose = () => console.log(`Closed WebSocket connection`);
  return ws;
};

const boardsService = { create, connectSocket };
export default boardsService;
