<?php

namespace Database\Factories;

use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Post>
 */
class PostFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'title' => fake()->sentence,
            'excerpt' => fake()->text,
            'published_at' => fake()->dateTimeInInterval('-1 years', '+2 months')
        ];
    }

    /**
     * @return $this
     */
    public function belongUser(): static
    {
        return $this->state(fn (array $attributes) => [
            'user_id' => User::all()->random(),
        ]);
    }
}
