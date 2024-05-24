import { useEffect, useState } from "react";
import { getTodos } from "./api/axios";

type Todo = {
  id: number;
  isDone: boolean;
  name: string;
  description: string;
};

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    loadTodos();
  }, []);

  const loadTodos = async () => {
    const response = await getTodos();
    setTodos(response);
  };

  console.log(todos);
  console.log(todos.map((todo) => todo.isDone));
  return (
    <main className="flex flex-col items-center text-center">
      <h1 className="text-xl font-semibold">Go Todos</h1>
      {todos.map((todo) => (
        <ul key={todo.id} className="p-4 m-2 bg-stone-700">
          <li>{todo.name}</li>
          <li>{todo.isDone ? "✅ Done" : "❌ Not Done"}</li>
          <li>{todo.description}</li>
        </ul>
      ))}
    </main>
  );
}
export default App;
