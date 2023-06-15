<?php

namespace App\Observers;

use App\Models\Post;
use Illuminate\Support\Arr;

class PostObserver
{
    public function creating(Post $post): void
    {
        // 非命令行下执行
        if (! app()->runningInConsole()) {
            $post->user_id = auth()->id() ?? 1;

            // 非开发环境，限制一段时间内的发帖次数
            app()->isLocal() || Post::throttleCheck($post->user);

            // 获取内容的摘录
            $post->excerpt = Post::extractExcerptFromContent(request()->input('content.markdown'));
        }
    }

    public function saving(Post $post): void
    {
        // 获取内容的摘录
        $post->excerpt = Post::extractExcerptFromContent(request()->input('content.markdown'));
    }

    public function saved(Post $post): void
    {
        if (! app()->runningInConsole()) {
            $contentData = Arr::only(request()->input('content', []), 'markdown');

            $post->content()->updateOrCreate(['contentable_id' => $post->id], $contentData);

            $post->loadMissing('content');
        }
    }
}
