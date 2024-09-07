export enum Meal {
    Breakfast = "breakfast",
    Lunch = "lunch",
    Dinner = "dinner",
    Snack = "snack"
}

export type Recipe = {
    id: string;
    name: string;
    cook_duration: string;
    instructions: string;
    image_url: string;
    meal: Meal;
}
