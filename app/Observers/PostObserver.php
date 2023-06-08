<?php

namespace App\Observers;

use App\Models\Post;
use Illuminate\Support\Arr;

class PostObserver
{
    public function saved(Post $post)
    {
        $contentData = Arr::only(request()->input('content', []), 'markdown');

        $post->content()->updateOrCreate(['contentable_id' => $post->id], $contentData);

        $post->loadMissing('content');
    }
}
