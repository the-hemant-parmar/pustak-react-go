import * as z from "zod";

export const bookSchema = z.object({
    title: z.string().min(2, {
        message: "Title must be at least 2 characters.",
    }),
    author: z.string().min(1, {
        message: "Author name must be at least one character.",
    }),
    status: z.enum(["available", "not-available"])
});

export type bookSchemaValues = z.infer<typeof bookSchema>;
