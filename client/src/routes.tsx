import { createBrowserRouter } from "react-router-dom";
import HomePage from "./pages/Home/HomePage";
// import LoginPage from "./pages/Auth/LoginPage";
// import RegisterPage from "./pages/Auth/RegisterPage";
// import BookDetailsPage from "./pages/Books/BookDetailsPage";

export const routes = createBrowserRouter([
  { path: "/", element: <HomePage /> },
  // { path: "/login", element: <LoginPage /> },
  // { path: "/register", element: <RegisterPage /> },
  // { path: "/books/:id", element: <BookDetailsPage /> },
]);
