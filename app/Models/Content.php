<?php

namespace App\Models;

use Illuminate\Support\Str;
use Mews\Purifier\Facades\Purifier;
use Illuminate\Database\Eloquent\Casts\Attribute;

class Content extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array<int, string>
     */
    protected $fillable = [
        'contentable_type', 'contentable_id', 'body', 'markdown'
    ];

    /**
     * The attributes that should be hidden for serialization.
     *
     * @var array<int, string>
     */
    protected $hidden = [
        'created_at', 'updated_at', 'deleted_at',
    ];

    protected function body(): Attribute
    {
        return Attribute::make(
            set: fn (string $value) => Purifier::clean($value),
        );
    }

    public static function toHTML(string $markdown): string
    {
        return Str::markdown($markdown);
    }
}
