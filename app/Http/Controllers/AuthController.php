<?php

namespace App\Http\Controllers;

use Carbon\Carbon;
use App\Models\User;
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

        if (! $token = auth('api')->login($user)) {
            abort(401, '未经授权');
        }

        return $this->respondWithToken($token);
    }

    public function signIn(AuthRequest $request)
    {
        $username = $request->username;

        $credentials['password'] = $request->password;
        filter_var($username, FILTER_VALIDATE_EMAIL) ?
            $credentials['email'] = $username :
            $credentials['name'] = $username;

        if (! $token = auth('api')->attempt($credentials)) {
            abort(401, '未经授权');
        }

        return $this->respondWithToken($token);
    }

    protected function respondWithToken($token)
    {
        return response()->json([
            'access_token' => $token,
            'token_type' => 'bearer',
            'expired_at' => Carbon::now()
                ->addMinutes(auth('api')->factory()->getTTL())
                ->toDatetimeString()
        ]);
    }
}
