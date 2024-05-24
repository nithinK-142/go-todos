import axios from "axios";

const URL = import.meta.env.VITE_API_URL;

export async function getTodos() {
  try {
    const response = await axios.get(URL);
    return response.data;
  } catch (error) {
    throw new Error("failed to fetch!!");
  }
}

// export async function createTodo() {
//   try {
//     const response = await axios.post(URL);
//     return response.data;
//   } catch (error) {
//     throw new Error("failed to fetch!!");
//   }
// }

// export async function deleteTodo() {
//   try {
//     const response = await axios.delete(URL);
//     return response.data;
//   } catch (error) {
//     throw new Error("failed to fetch!!");
//   }
// }

// export async function editTodo() {
//   try {
//     const response = await axios.put(URL);
//     return response.data;
//   } catch (error) {
//     throw new Error("failed to fetch!!");
//   }
// }
