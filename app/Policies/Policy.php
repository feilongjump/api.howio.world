<?php

namespace App\Policies;

use App\Models\User;

class Policy
{
    /**
     * Create a new policy instance.
     */
    public function __construct()
    {
        //
    }

    public function before(User $authUser): void
    {
        // This user is admin, then he will go anywhere at will
    }

    public function own(User $authUser, $association): bool
    {
        return $authUser->id == $association->user_id;
    }
}
