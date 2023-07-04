<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\PostController;
use App\Http\Controllers\VerificationCodeController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

// Auth
Route::post('auth/sign-up', [AuthController::class, 'signUp']);
Route::post('auth/sign-in', [AuthController::class, 'signIn']);

// User
Route::post('user/{medium}/verification-code', [VerificationCodeController::class, 'verificationCode'])
    ->where('medium', 'email');

Route::middleware('auth')->group(function () {
    Route::get('me', [UserController::class, 'me']);
});

\LaravelUploader::routes([
    'as' => 'files.upload',
    'middleware' => ['auth'],
]);

Route::resources([
    'posts' => PostController::class
]);
