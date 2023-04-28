<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use App\Http\Requests\AuthRequest;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\Hash;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Validation\ValidationException;

class AuthController extends Controller
{
    public function signUp(AuthRequest $request)
    {
        $key = 'mail_verification_code:' . $request->email;
        $verificationCode = Cache::get($key);
        if (! $verificationCode) {
            abort(403, '验证码已失效');
        }
        if (! hash_equals($verificationCode, $request->verification_code)) {
            abort(422, '验证码错误');
        }

        // 清除验证码缓存
        Cache::forget($key);

        $request->merge(['email_verified_at' => now()]);
        $user = User::create($request->all());

        return $user->createDeviceToken($request->device_name);
    }

    /**
     * @throws ValidationException|AuthenticationException
     */
    public function signIn(AuthRequest $request)
    {
        $user = User::where('name', $request->username)
            ->orWhere('email', $request->username)
            ->first();
        if (!$user || !Hash::check($request->password, $user->password)) {
            throw ValidationException::withMessages([
                'username' => [__('auth.failed')],
                'password' => [__('auth.failed')],
            ]);
        }
        if (! $user->email_verified_at) {
            throw new AuthenticationException(__('auth.frozen'));
        }

        return $user->createDeviceToken($request->device_name);
    }
}
