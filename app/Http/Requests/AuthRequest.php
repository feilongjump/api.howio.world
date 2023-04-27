<?php

namespace App\Http\Requests;

class AuthRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array|string>
     */
    public function rules(): array
    {
        return match ($this->getPathInfo()) {
            '/auth/sign-up' => [
                'name' => 'required|between:3,25|regex:/^[A-Za-z0-9\-\_]+$/|unique:users,name',
                'email' => 'required|email|unique:users,email',
                'password' => 'required|string|min:6',
                'verification_code' => 'required|string|size:6',
            ]
        };
    }
}
