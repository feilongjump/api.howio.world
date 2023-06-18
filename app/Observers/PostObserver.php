<?php

namespace App\Observers;

use App\Models\Post;
use Illuminate\Support\Arr;

class PostObserver
{
    public function saving(Post $post)
    {
        if (! app()->runningInConsole()) {
            $post->user_id = auth()->id() ?? 0;

            $post->excerpt = Post::extractExcerptFromMarkdown(request()->input('content.markdown'));
            $post->published_at = Post::extractPublishedAtFromMarkdown(request()->input('content.markdown'), $post);
        }
    }

    public function saved(Post $post)
    {
        if (! app()->runningInConsole()) {
            $contentData = Arr::only(request()->input('content', []), 'markdown');

            $post->content()->updateOrCreate(['contentable_id' => $post->id], $contentData);

            $post->loadMissing('content');
        }
    }
}
