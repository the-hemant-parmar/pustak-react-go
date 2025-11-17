import { useBooks } from "../../hooks/useBooks";
import BookCard from "../../components/Book/BookCard";

export default function HomePage() {
  // const lat = 12.912; // placeholder
  // const lng = 77.642;

  const { data: books } = useBooks();

  return (
    <div className="p-6 grid grid-cols-2 md:grid-cols-4 gap-4">
      {books?.map((b) => (
        <BookCard key={b.id} book={b} />
      ))}
    </div>
  );
}
