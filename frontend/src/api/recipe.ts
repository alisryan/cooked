import { API_BASE_URL } from "./const";
import { Recipe } from "@/types/recipe";

const RECIPE_BASE_URL = API_BASE_URL + '/recipes';

export async function fetchAllRecipes(): Promise<Recipe[]> {
    try {
        const response = await fetch(RECIPE_BASE_URL);
        console.log("recipe service fetching data - status: ", response.status)
        return response.json();
    }
    catch (error) {
        console.error("recipe service error: ", error);
        throw new Error("Failed fetching recipes")
    }
}

export async function createRecipe(recipe: Recipe): Promise<void> {
    try {
        const response = await fetch(RECIPE_BASE_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(recipe)
        });
        console.log("recipe service creating data - status: ", response.status)
    }
    catch (error) {
        console.error("recipe service error: ", error);
        throw new Error("Failed creating recipe")
    }
}
