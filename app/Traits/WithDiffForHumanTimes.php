<?php

namespace App\Traits;

use Illuminate\Support\Str;

trait WithDiffForHumanTimes
{
    public function __call($method, $arguments)
    {
        if (Str::endsWith($method, 'TimeAgoAttribute')) {
            $date = $this->{Str::snake(substr($method, 3, -16))};

            return $date ? $date->diffForHumans() : null;
        }

        return parent::__call($method, $arguments);
    }
}
