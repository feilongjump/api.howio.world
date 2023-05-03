<?php

namespace App\Observers;

use App\Models\Thread;
use Illuminate\Support\Arr;

class ThreadObserver
{
    public function creating(Thread $thread)
    {
        $thread->user_id = request()->user()->id;
    }

    public function saved(Thread $thread)
    {
        $contentData = Arr::only(request()->input('content', []), 'markdown');

        $thread->content()->updateOrCreate(['contentable_id' => $thread->id], $contentData);

        $thread->loadMissing('content');
    }
}
