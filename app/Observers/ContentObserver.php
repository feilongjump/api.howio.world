<?php

namespace App\Observers;

use App\Models\Content;

class ContentObserver
{
    public function saving(Content $content): void
    {
        $content->body = Content::toHTML($content->markdown);
    }
}
