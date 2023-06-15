<?php

namespace App\Models;

use Illuminate\Support\Str;
use Illuminate\Support\Stringable;
use Mews\Purifier\Facades\Purifier;
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

    /**
     * 校验短期内的发帖次数
     */
    public static function throttleCheck(User $user)
    {
        $lastPost = $user->posts()->latest()->first();

        if ($lastPost && $lastPost->created_at->gt(now()->subMinutes(config('throttle.thread.create')))) {
            \abort(403, '发贴太频繁');
        }
    }

    /**
     * 从 markdown 内容中获取摘录
     */
    public static function extractExcerptFromContent(string $markdown, int $limit = 450, string $end = ''): string
    {
        $splitStr = '~~---~~';

        $html = Str::of($markdown)
            ->whenContains($splitStr,
                function (Stringable $string) use ($splitStr, $limit, $end) {
                    return $string
                        ->beforeLast($splitStr)
                        ->markdown()
                        ->limit($limit, $end)
                        ->value();
                },
                fn () => ''
            );

        return Purifier::clean($html);
    }
}
