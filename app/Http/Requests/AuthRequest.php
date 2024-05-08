<?php

namespace App\Http\Requests;

class AuthRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return match ($this->getPathInfo()) {
            '/auth/sign-up' => [
                'name' => 'required|between:5,32|unique:users,name',
                'email' => 'required|email|unique:users,email',
                'password' => 'required|string|min:6',
                'verification_code' => 'required|size:6',
            ],
            '/auth/sign-in' => [
                'username' => 'required|min:5',
                'password' => 'required|string|min:6',
            ]
        };
    }
}
