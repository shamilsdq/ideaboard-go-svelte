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

const boardsService = { create };
export default boardsService;
