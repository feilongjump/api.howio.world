<?php

namespace App\Http\Requests;

class UpdateThreadRequest extends FormRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array|string>
     */
    public function rules(): array
    {
        return [
            'title' => 'required|min:2|max:64',
            'published_at' => 'nullable|date|after_or_equal:now',
            'content.markdown' => 'required',
        ];
    }
}
