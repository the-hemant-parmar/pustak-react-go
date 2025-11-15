import { type Book } from "../../types/book";

export default function BookCard({ book }: { book: Book }) {
  return (
    <div className="p-4 border rounded-xl shadow-sm hover:shadow-md transition">
      <img src={book.coverUrl} className="h-48 w-full object-cover rounded-md" />
      <h3 className="font-semibold text-lg mt-2">{book.title}</h3>
      <p className="text-sm text-gray-600">{book.author}</p>
    </div>
  );
}
