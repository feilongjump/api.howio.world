<?php

namespace App\Policies;

use App\Models\Post;
use App\Models\User;

class PostPolicy extends Policy
{
    /**
     * Determine whether the user can view the model.
     */
    public function view(?User $user, Post $post): bool
    {
        return $user?->id === $post->user_id || $post->published_at;
    }

    /**
     * Determine whether the user can create models.
     */
    public function create(User $user): bool
    {
        return true;
    }
}
