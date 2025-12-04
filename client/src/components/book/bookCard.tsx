import { type Book } from "../../types/book";
import Placeholder from  "../../assets/Gemini_Generated_Image_hwyne4hwyne4hwyn-Cpd7Tx4m.png";

export const BookCard = ({ book }: { book: Book }) => {
  return (
    <div className="p-4 border rounded-xl shadow-sm hover:shadow-md transition">
      <img src={book.coverUrl ?? Placeholder} className="h-48 w-full object-cover rounded-md" />
      <h3 className="font-semibold text-lg mt-2">{book.title}</h3>
      <p className="text-sm text-gray-600">{book.author}</p>
    </div>
  );
}
