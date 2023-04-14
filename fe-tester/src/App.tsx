import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [users, setUsers] = useState(["1", "2", "3", "4"]);

  const [requestUser, setRequestUser] = useState(["1", "2", "3", "4"]);
  const queryClient = useQueryClient();

  const { data } = useQuery(
    ["users-status"],
    async () => {
      return fetch("http://localhost:1324/api/users/status?users=1,2,3,4").then(
        (res) => res.json()
      );
    },
    {
      refetchInterval: 1000,
    }
  );

  const { mutate } = useMutation(["users"], (user: string) => {
    return fetch("http://localhost:1324/api/users/hb?user=" + user, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    }).then((res) => res.json());
  });

  useEffect(() => {
    const interval = setInterval(() => {
      requestUser.forEach((user) => {
        mutate(user);
      });
    }, 3000);

    return () => clearInterval(interval);
  }, [requestUser]);

  return users.map((user, index) => {
    return (
      <div key={user}>
        <button onClick={() => mutate(user)}>{user}</button>
        {data?.[index] ? "Online" : "Offline"}
        <button
          onClick={() => {
            setRequestUser((prev) => {
              const newUsers = prev.filter((u) => u !== user);
              return newUsers;
            });
          }}
        >
          {" "}
          Make it offline{" "}
        </button>
      </div>
    );
  });
}

export default App;
