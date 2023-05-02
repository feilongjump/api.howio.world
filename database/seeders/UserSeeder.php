<?php

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use Illuminate\Support\Str;

class UserSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        User::factory()
            ->unverified()
            ->count(10)
            ->create();

        $user = User::first();
        $user->name = 'super';
        $user->email = 'super@example.com';
        $user->email_verified_at = now();
        $user->remember_token = Str::random();
        $user->save();
    }
}
