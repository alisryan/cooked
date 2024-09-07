import { API_BASE_URL } from "./const";
import { Recipe } from "@/types/recipe";

const RECIPE_BASE_URL = API_BASE_URL + '/recipes';

export async function fetchAllRecipes(): Promise<Recipe[]> {
    const response = await fetch(RECIPE_BASE_URL);

    if (response.status > 299) {
        console.error("recipe service error: ", response.status);
        throw new Error("Failed fetching recipes")
    }

    return response.json();
}

export async function createRecipe(recipe: Recipe): Promise<void> {
        const response = await fetch(RECIPE_BASE_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(recipe)
        });

        if (response.status > 299) {
            console.error("recipe service error: ", response.status);
            throw new Error("Failed creating recipe")
        }
}
