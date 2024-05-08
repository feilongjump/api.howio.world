<?php

namespace App\Http\Controllers;

use App\Mail\VerificationCode;
use Illuminate\Support\Facades\Mail;
use Illuminate\Support\Facades\Cache;
use App\Http\Requests\VerificationCodeRequest;

class VerificationCodeController extends Controller
{
    public function verificationCode(VerificationCodeRequest $request, $medium)
    {
        return match ($medium) {
            'email' => $this->sendMailVerificationCode($request)
        };
    }

    private function sendMailVerificationCode(VerificationCodeRequest $request)
    {
        $email = $request->email;
        $key = 'mail_verification_code:' . $email;
        // 5 分钟内，只发一次验证码
        $code = Cache::get($key);
        if ($code) {
            abort(403, '请勿重复发送验证码。');
        }

        // 生成 6 位随机数，左侧补0
        $code = str_pad(random_int(1, 999999), 6, 0, STR_PAD_LEFT);

        // 异步发送验证码
        Mail::to($email)->send(new VerificationCode($code));

        // 缓存验证码 5 分钟过期。
        $expiredAt = now()->addMinutes(5);
        Cache::put($key, $code, $expiredAt);

        return response()->json([
            'expired_at' => $expiredAt->toDateTimeString(),
        ])->setStatusCode(201);
    }
}
