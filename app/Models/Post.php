<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;

class Post extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array<int, string>
     */
    protected $fillable = [
        'user_id', 'title', 'excerpt', 'published_at'
    ];

    /**
     * The attributes that should be hidden for serialization.
     *
     * @var array<int, string>
     */
    protected $hidden = [
        'deleted_at',
    ];

    public function user(): \Illuminate\Database\Eloquent\Relations\BelongsTo
    {
        return $this->belongsTo(User::class);
    }

    public function content(): \Illuminate\Database\Eloquent\Relations\morphOne
    {
        return $this->morphOne(Content::class, 'contentable');
    }

    public function scopePublished(Builder $query): void
    {
        $query->where('published_at', '<=', now());
    }

    public function scopeKeywords(Builder $query, $keywords)
    {
        $query->where(function (Builder $query) use ($keywords) {
            $query->where('title', 'like', "%{$keywords}%")
                ->orWhere('excerpt', 'like', "%{$keywords}%");
        });
    }
}
