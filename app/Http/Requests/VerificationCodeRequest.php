<?php

namespace App\Http\Requests;

class VerificationCodeRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array|string>
     */
    public function rules(): array
    {
        return match ($this->getPathInfo()) {
            '/user/email/verification-code' => [
                'email' => 'required|email',
            ]
        };
    }
}
