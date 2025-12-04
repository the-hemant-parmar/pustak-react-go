import { type Book } from "@/types/book";
import Placeholder from "@/assets/PlaceholderImage.png";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";

export const BookCard = ({ book }: { book: Book }) => {
  return (
    <Card>
      <CardHeader>
        <img src={book.coverUrl ?? Placeholder} className="h-64 w-full object-cover rounded-md" />
      </CardHeader>
      <CardContent>
        <CardTitle>
          {book.title}
        </CardTitle>
        <CardDescription>
          {book.author}
        </CardDescription>
      </CardContent>
      <CardFooter>
        <Button>Learn More</Button>
      </CardFooter>

    </Card>
  );
}

