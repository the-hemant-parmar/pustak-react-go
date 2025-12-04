import { RouterProvider } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { routes } from "./routes";
import { GlobalNavigation } from "@/components/navigation/globalNavigation";

const queryClient = new QueryClient();

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <GlobalNavigation/>
      <RouterProvider router={routes} />
    </QueryClientProvider>
  );
}
