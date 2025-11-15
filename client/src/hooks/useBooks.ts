import { useQuery } from "@tanstack/react-query";
import { bookApi } from "../services/api/bookApi";

export function useNearbyBooks(lat: number, lng: number) {
  return useQuery({
    queryKey: ["books", lat, lng],
    queryFn: () => bookApi.listNearby(lat, lng, 3000),
    enabled: !!lat && !!lng,
  });
}
