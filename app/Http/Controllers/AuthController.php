<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use App\Http\Requests\AuthRequest;
use Illuminate\Support\Facades\Cache;

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
}
