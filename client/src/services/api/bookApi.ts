import { http } from "./http";
import { type Book } from "../../types/book";

export const bookApi = {
  // temporarily using a degraded version
  async listNearby(lat: number, lng: number, radius: number) {
    const res = await http.get<Book[]>("/books/nearby", {
      params: { lat, lng, radius },
    });
    return res.data;
  },

  async listBooks() {
    const res = await http.get<Book[]>("/books");
    return res.data;
  },

  async get(id: string) {
    const res = await http.get<Book>(`/books/${id}`);
    return res.data;
  }
};
