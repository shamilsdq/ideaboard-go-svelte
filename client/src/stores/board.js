import { writable } from "svelte/store";

const handle = (code, content, board) =>
  board.update((data) => {
    switch (code) {
      case "INITIAL":
        return content;
      case "MEMBER_JOINED":
      case "MEMBER_EXITED":
        return { ...data, ...content };
      case "POST_CREATED":
        return { ...data, posts: [...data.posts, content] };
      case "POST_UPDATED":
        return {
          ...data,
          posts: data.posts.map((p) => (p.id === content.id ? content : p)),
        };
      case "POST_DELETED":
        return {
          ...data,
          posts: [...data.posts.filter((p) => p.id !== content.id)],
        };
      default:
        console.log("Unknown action:", { code, content });
        return data;
    }
  });

const createBoard = () => {
  const board = writable(null);
  return {
    subscribe: board.subscribe,
    handle: (code, content) => handle(code, content, board),
    clear: () => board.set(null),
  };
};

export const boardStore = createBoard();
