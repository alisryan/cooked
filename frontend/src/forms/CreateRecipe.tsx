'use client'

import { createRecipe } from "@/api/recipe";
import Button from "@/components/Button/Button";
import { Meal, Recipe } from "@/types/recipe";
import { FormControl, InputLabel, MenuItem, Select, TextField, Box } from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import { useEffect } from "react";
import { useForm, Controller } from "react-hook-form";

type CreateRecipeProps = {
    onSuccess: () => void;
};

export default function CreateRecipe({ onSuccess }: CreateRecipeProps) {
    const {
        register,
        handleSubmit,
        formState: { errors },
        setError,
        control,
    } = useForm<Recipe>();

    const { mutate, isSuccess, error } = useMutation({
        mutationFn: createRecipe,
    });

    useEffect(() => {
        if (isSuccess) {
            onSuccess(); 
        }
    }, [isSuccess, onSuccess]);

    const parseDuration = (duration: string) => {
        const minutes = Number(duration);
        if (isNaN(minutes) || minutes < 0) {
            setError('cook_duration', { message: 'Duration must be a number greater than 0' });
            return null;
        }

        const hours = Math.floor(minutes / 60);
        const remainingMinutes = minutes % 60;
        const seconds = (minutes % 1) * 60;  

        return `${hours}h${remainingMinutes}m${Math.round(seconds)}s`;
    };

    const onSubmit = (data: Recipe) => {
        const parsedCookDuration = parseDuration(data.cook_duration);
        if (!parsedCookDuration) return;

        const parsedData: Recipe = {
            ...data,
            cook_duration: parsedCookDuration,
        };

        console.log("Parsed Data", parsedData);

        mutate(parsedData); 
    };

    return (
        <Box
            component="form"
            onSubmit={handleSubmit(onSubmit)}
            sx={{
                display: 'flex',
                flexDirection: 'column',
                width: '100%',
                gap: 2,
            }}
        >
            <TextField
                fullWidth
                label="Name"
                variant="outlined"
                {...register('name', { required: 'Name is required' })}
                error={!!errors.name}
                helperText={errors.name?.message}
            />
            
            <TextField
                fullWidth
                label="Time to Cook (in minutes)"
                variant="outlined"
                {...register('cook_duration', { required: 'Time to cook is required' })}
                error={!!errors.cook_duration}
                helperText={errors.cook_duration?.message}
            />
            
            <TextField
                fullWidth
                label="Instructions"
                variant="outlined"
                multiline
                rows={4}
                {...register('instructions', { required: 'Instructions are required' })}
                error={!!errors.instructions}
                helperText={errors.instructions?.message}
            />
            
            <TextField
                fullWidth
                label="Image URL"
                variant="outlined"
                {...register('image_url', { required: 'Image URL is required' })}
                error={!!errors.image_url}
                helperText={errors.image_url?.message}
            />

            <FormControl fullWidth error={!!errors.meal}>
                <InputLabel id="meal-label">Meal</InputLabel>
                <Controller
                    name="meal"
                    control={control}
                    defaultValue={Meal.Dinner}
                    render={({ field }) => (
                        <Select
                            labelId="meal-label"
                            {...field}
                            label="Meal"
                        >
                            <MenuItem value={Meal.Breakfast}>Breakfast</MenuItem>
                            <MenuItem value={Meal.Lunch}>Lunch</MenuItem>
                            <MenuItem value={Meal.Dinner}>Dinner</MenuItem>
                            <MenuItem value={Meal.Snack}>Snack</MenuItem>
                        </Select>
                    )}
                    rules={{ required: 'Meal selection is required' }}
                />
                {errors.meal && <Box sx={{ color: 'red' }}>{errors.meal.message}</Box>}
            </FormControl>

            {error && (
                <Box sx={{ color: 'red' }}>
                    {error.message}
                </Box>
            )}

            <Button type="submit" color="secondary">
                Create
            </Button>
        </Box>
    );
}
