<?php

namespace App\Models;

use Illuminate\Support\Str;
use Illuminate\Support\Carbon;
use Illuminate\Support\Stringable;
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
     * 从 Markdown 内容中截取摘录信息
     *
     * @param string $markdown
     * @param int $maxLength
     * @param string $end
     * @return string
     */
    public static function extractExcerptFromMarkdown(string $markdown, int $maxLength = 450, string $end = ''): string
    {
        $str = '---';

        return Str::of($markdown)
            ->whenContainsAll([$str, $str],
                function (Stringable $stringable) use ($str, $maxLength, $end) {
                    // 最大只获取前 550 以内的内容进行转换
                    // 因为可能存在发布时间以及标签文字，所以得进行限制字符长度
                    // TODO: 如何更改截取方式，会更加好
                    return $stringable
                        ->limit(550)
                        ->betweenFirst($str, $str)
                        ->markdown()
                        ->limit($maxLength, $end)
                        ->value();
                },
                fn () => ''
            );
    }

    /**
     * 从 Markdown 内容中截取发布时间
     *
     * @param string $markdown
     * @param Post $post
     * @return string|null
     */
    public static function extractPublishedAtFromMarkdown(string $markdown, Post $post): string | null
    {
        $publishedAt = $post->getOriginal('published_at');
        $publishedAt = $publishedAt === null ? null : Carbon::parse($publishedAt);
        // 已发布的无法更改发布时间
        if ($publishedAt && $publishedAt->lte(now())) return $publishedAt;

        $str = '`published_at:';
        $endStr = '`';
        $updatedPublishedAt = Str::of($markdown)
            ->whenContainsAll([$str, $endStr],
                function (Stringable $stringable) use ($str, $endStr) {
                    $value = $stringable->betweenFirst($str, $endStr)->trim()->value();
                    if (empty($value) || $value === 'null') return '';

                    return new Carbon($value);
                },
                fn () => ''
            );

        // 更改为：待发布
        if (empty($updatedPublishedAt)) return null;

        // 严禁更改为过往时间
        if ($updatedPublishedAt->lte(now())) abort(422, '禁止发布过往时间的帖子');

        // 更改为：更新发布时间
        return $updatedPublishedAt;
    }
}
