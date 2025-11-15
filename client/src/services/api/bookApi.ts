import { http } from "./http";
import { type Book } from "../../types/book";

export const bookApi = {
  async listNearby(lat: number, lng: number, radius: number) {
    const res = await http.get<Book[]>("/books/nearby", {
      params: { lat, lng, radius },
    });
    return res.data;
  },

  async get(id: string) {
    const res = await http.get<Book>(`/books/${id}`);
    return res.data;
  }
};
