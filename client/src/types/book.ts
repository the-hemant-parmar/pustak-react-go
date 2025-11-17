export interface Book {
  id: string;
  title: string;
  author: string;
  coverUrl?: string;
  price?: number;
  lend: boolean;
  location: {
    lat: number;
    lng: number;
  };
}
