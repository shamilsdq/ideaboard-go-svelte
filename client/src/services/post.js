const addPost = (socket, payload) =>
  socket.send(JSON.stringify({ code: "POST_CREATE", content: payload }));

const postsService = { addPost };
export default postsService;
