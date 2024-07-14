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

const connectSocket = (boardId, config) => {
  const { loadingSetter, instanceSetter, messageHandler } = config;
  const socketUrl = `${API_ROOT.replace(/^https?:\/\//i, "ws://")}/${boardId}`;
  const socket = new WebSocket(socketUrl);

  socket.onopen = () => {
    loadingSetter?.(false);
    instanceSetter?.(socket);
    console.log("Socket opened");
  };

  socket.onclose = () => {
    instanceSetter?.(undefined);
    console.log("Socket closed");
  };

  socket.onerror = () => {
    console.error("Socket error");
  };

  socket.onmessage = (message) => {
    const { errors, ...parsedData } = JSON.parse(message.data.toString());
    if (errors) console.error(errors);
    else messageHandler?.(parsedData);
  };
};

const boardsService = { create, connectSocket };
export default boardsService;
