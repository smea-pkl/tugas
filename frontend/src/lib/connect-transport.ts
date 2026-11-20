import { createConnectTransport } from "@connectrpc/connect-web";

export const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  interceptors: [
    function (next) {
      return function (req) {
        const token = localStorage.getItem("token");

        if (!token) {
          if (req.header.has("Authorization")) {
            req.header.delete("Authorization");
          }

          return next(req);
        }

        req.header.set("Authorization", `Bearer ${token}`);

        return next(req);
      };
    },
  ],
});
