<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\VerificationCodeController;

Route::get('/', function () {
    return response()->json([
        'msg' => 'Hello World'
    ]);
});

// Auth
Route::post('auth/sign-up', [AuthController::class, 'signUp']);
Route::post('auth/sign-in', [AuthController::class, 'signIn']);

// User
Route::post('user/{medium}/verification-code', [VerificationCodeController::class, 'verificationCode'])
    ->where('medium', 'email');

Route::middleware('auth')->group(function () {
    Route::get('me', [UserController::class, 'me']);
});
