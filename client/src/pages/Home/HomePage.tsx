import { useBooks } from "@/hooks/useBooks";
import { BookCard } from "@/components/book/bookCard";
import { Button } from "@/components/ui/button";
import HeroImage from "@/assets/HeroImage.png"
import { Link } from "react-router-dom";

const HomePage = () => {
  // const lat = 12.912; // placeholder
  // const lng = 77.642;

  const { data: books } = useBooks();

  return (
    <div className="container mx-auto bg-blue-100">
      // Hero section
      <div className="flex justify-center">
        <img src={HeroImage} alt="Hero Image" />
      </div>
    
    
      // Book grid 
      <div className="p-6 grid grid-cols-2 md:grid-cols-4 gap-4">
        {books?.map((b) => (
          <BookCard key={b.id} book={b} />
        ))}
      </div>
      <div className="flex justify-center p-6">
        <Button className="p-6">
          <Link to="/books">
          CLICK HERE TO ADD BOOK
          </Link>
        </Button>
      </div>
    </div>
  );
}

export default HomePage;
