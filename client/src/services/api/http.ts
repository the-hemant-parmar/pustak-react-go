import axios from "axios";
import { env } from "../../config/env";

export const http = axios.create({
  baseURL: env.API_URL,
  withCredentials: false,
});

http.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) config.headers.Authorization = `Bearer ${token}`;
  return config;
});
