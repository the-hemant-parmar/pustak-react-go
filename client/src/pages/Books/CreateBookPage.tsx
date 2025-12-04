import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod"
import { bookSchema } from "@/services/validation/bookValidation";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { bookApi } from "@/services/api/bookApi";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

const CreateBookPage = () => {
    const form = useForm({
        resolver: zodResolver(bookSchema),
        defaultValues: {
            title: "",
            author: "",
            status: "not-available" as "not-available", // suboptimal fix for unnecessary typescript warning
        }
    }
    )

    return (
        <div className="container mx-auto max-w-2xl p-10">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(bookApi.createBook)} className="flex flex-col gap-6">
                    <FormField
                        control={form.control}
                        name="title"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Title</FormLabel>
                                <FormControl>
                                    <Input placeholder="Enter Books title" {...field} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="author"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Author</FormLabel>
                                <FormControl>
                                    <Input placeholder="Enter Author's Name" {...field} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="status"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Availablity Status</FormLabel>
                                <Select onValueChange={field.onChange} defaultValue={field.value}>
                                    <FormControl>
                                        <SelectTrigger>
                                            <SelectValue placeholder="Select availability" />
                                        </SelectTrigger>
                                    </FormControl>
                                    <SelectContent>
                                        <SelectItem value="not-available">Available</SelectItem>
                                        <SelectItem value="available">Not Available</SelectItem>
                                    </SelectContent>
                                </Select>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <Button className="p-4 cursor-pointer" type="submit">Register Book</Button>
                </form>
            </Form>
        </div>
    )
}

export default CreateBookPage;
